# Findings

Inconsistencies, bugs, and logical problems discovered in the tenant-api codebase.

---

## F4: `ProjectMember.tenant_id` has different semantics than `Project.tenant_id` — `proto/api/v1/project_member.proto:27-28`, `proto/api/v1/project.proto:33-34`

| Message         | Field       | Meaning                                         |
|-----------------|-------------|-------------------------------------------------|
| `Project`       | `tenant_id` | The tenant that **owns** this project           |
| `ProjectMember` | `tenant_id` | The tenant that is a **member** of this project |

The same field name and comment pattern (`// TenantId of this project member`) are used for a member tenant in `ProjectMember`, which is semantically different from `Project.tenant_id` (the owner tenant). A consumer reading `ProjectMember.tenant_id` would likely assume it refers to the owning tenant of the project, when in fact it refers to the member tenant.

Similarly, `TenantMember` uses:

- `tenant_id` = parent (owner) tenant
- `member_id` = member tenant

But `ProjectMember` uses:

- `project_id` = the project
- `tenant_id` = member tenant

The naming is inconsistent: `TenantMember` distinguishes between parent and member with two different field names, while `ProjectMember` uses only `tenant_id` for the member, leaving the project's owning tenant ID absent.

---

## F5: `NamespaceInterceptor` mutates request in place — `go/client/client-interceptors.go:70-105`

The interceptor mutates received request pointers directly:

```go
case *v1.TenantMemberServiceCreateRequest:
    if r.TenantMember.Namespace == "" {
        r.TenantMember.Namespace = namespace
    }
```

**Problems:**

1. **Race condition in concurrent reuse**: Connect may reuse request objects for retries. Mutating the request means a retry could send a namespace that was already set, or the mutation could be observed by caller-side code.

2. **`TenantServiceListRequest` and `ProjectServiceListRequest` are not covered** — the interceptor does not set `namespace` on these list requests, but it does cover their counterpart `TenantServiceListTenantMembersRequest` and member-specific list requests. This asymmetry may be intentional but is undocumented.

---

## F8: `Schema()` uses backtick-wrapped SQL with template-embedded backtick — `generate/genscanvaluer.go:114`

```go
info["schema"] = fmt.Sprintf("`%s`", renderedBytesSchema.String())
```

The SQL template string is wrapped in Go backticks for the generated code's raw string literal. If the schema template itself ever contains backtick characters (which it currently doesn't), this would break. Currently safe, but fragile.

---

## F11: `op char NOT NULL` has no length in migration — `*_scanvaluer.go`

All generated schema files use `op char NOT NULL` without a length specifier. In PostgreSQL, `char` (unary) is equivalent to `character(1)`, which happens to be correct for single-character operation codes. However, `char` without parentheses can be confusing — explicitly using `char(1)` would be clearer and more portable across databases.

---

## F12: No `TenantServiceListRequest` namespace handling — `go/client/client-interceptors.go`

The `NamespaceInterceptor` handles namespace for member requests and finding-participating requests, but **not** for `TenantServiceListRequest` or `ProjectServiceListRequest`. If list operations should also default to the namespace, this is missing coverage. If intentional, it should be documented.

---
