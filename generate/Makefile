.PHONY: generate
generate:
	# we remove files explicitly to ensure files really get generated
	rm -f ../go/permissions/servicepermissions.go
	rm -f ../go/client/client.go

	go run ./generate.go

	go run ./genscanvaluer.go -package apiv1 -type Project -output ../go/tenant/api/v1/project_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type ProjectMember -output ../go/tenant/api/v1/project_member_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type Tenant -output ../go/tenant/api/v1/tenant_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type TenantMember -output ../go/tenant/api/v1/tenant_member_scanvaluer.go

