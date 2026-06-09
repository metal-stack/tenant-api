package client

import (
	"log/slog"
	"testing"

	"connectrpc.com/connect"
	v1 "github.com/metal-stack/tenant-api/go/api/v1"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	const (
		namespace = "a"
	)

	var (
		log = slog.Default()
	)

	t.Run("check namespace interceptor sets missing namespace", func(t *testing.T) {
		client, err := New(&DialConfig{
			BaseURL:   "http://localhost",
			Log:       log,
			Namespace: namespace,
			Interceptors: []connect.Interceptor{
				NewTestInterceptor(t, []ClientCall{
					{
						WantRequest: &v1.ProjectMemberServiceCreateRequest{
							ProjectMember: &v1.ProjectMember{
								ProjectId: "project-a",
								TenantId:  "tenant-a",
								Namespace: namespace,
							},
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.ProjectMemberServiceCreateResponse{})
						},
					},
					{
						WantRequest: &v1.ProjectMemberServiceListRequest{
							ProjectId: new("project-a"),
							TenantId:  new("tenant-a"),
							Namespace: namespace,
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.ProjectMemberServiceListResponse{})
						},
					},
					{
						WantRequest: &v1.TenantMemberServiceCreateRequest{
							TenantMember: &v1.TenantMember{
								TenantId:  "tenant-a",
								Namespace: namespace,
							},
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantMemberServiceCreateResponse{})
						},
					},
					{
						WantRequest: &v1.TenantMemberServiceListRequest{
							TenantId:  new("tenant-a"),
							Namespace: namespace,
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantMemberServiceListResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceFindParticipatingProjectsRequest{
							TenantId:  "tenant-a",
							Namespace: namespace,
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceFindParticipatingProjectsResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceFindParticipatingTenantsRequest{
							TenantId:  "tenant-a",
							Namespace: namespace,
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceFindParticipatingTenantsResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceListTenantMembersRequest{
							TenantId:  "tenant-a",
							Namespace: namespace,
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceListTenantMembersResponse{})
						},
					},
				}),
			},
		})

		_, err = client.Apiv1().ProjectMember().Create(t.Context(), &v1.ProjectMemberServiceCreateRequest{
			ProjectMember: &v1.ProjectMember{
				ProjectId: "project-a",
				TenantId:  "tenant-a",
			},
		})
		require.NoError(t, err)

		_, err = client.Apiv1().ProjectMember().List(t.Context(), &v1.ProjectMemberServiceListRequest{
			ProjectId: new("project-a"),
			TenantId:  new("tenant-a"),
		})
		require.NoError(t, err)

		_, err = client.Apiv1().TenantMember().Create(t.Context(), &v1.TenantMemberServiceCreateRequest{
			TenantMember: &v1.TenantMember{
				TenantId: "tenant-a",
			},
		})
		require.NoError(t, err)

		_, err = client.Apiv1().TenantMember().List(t.Context(), &v1.TenantMemberServiceListRequest{
			TenantId: new("tenant-a"),
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().FindParticipatingProjects(t.Context(), &v1.TenantServiceFindParticipatingProjectsRequest{
			TenantId: "tenant-a",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().FindParticipatingTenants(t.Context(), &v1.TenantServiceFindParticipatingTenantsRequest{
			TenantId: "tenant-a",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().ListTenantMembers(t.Context(), &v1.TenantServiceListTenantMembersRequest{
			TenantId: "tenant-a",
		})
		require.NoError(t, err)
	})

	t.Run("check explicit namespace can be set anyway", func(t *testing.T) {
		client, err := New(&DialConfig{
			BaseURL:   "http://localhost",
			Log:       log,
			Namespace: namespace,
			Interceptors: []connect.Interceptor{
				NewTestInterceptor(t, []ClientCall{
					{
						WantRequest: &v1.ProjectMemberServiceCreateRequest{
							ProjectMember: &v1.ProjectMember{
								ProjectId: "project-a",
								TenantId:  "tenant-a",
								Namespace: "b",
							},
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.ProjectMemberServiceCreateResponse{})
						},
					},
					{
						WantRequest: &v1.ProjectMemberServiceListRequest{
							ProjectId: new("project-a"),
							TenantId:  new("tenant-a"),
							Namespace: "b",
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.ProjectMemberServiceListResponse{})
						},
					},
					{
						WantRequest: &v1.TenantMemberServiceCreateRequest{
							TenantMember: &v1.TenantMember{
								TenantId:  "tenant-a",
								Namespace: "b",
							},
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantMemberServiceCreateResponse{})
						},
					},
					{
						WantRequest: &v1.TenantMemberServiceListRequest{
							TenantId:  new("tenant-a"),
							Namespace: "b",
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantMemberServiceListResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceFindParticipatingProjectsRequest{
							TenantId:  "tenant-a",
							Namespace: "b",
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceFindParticipatingProjectsResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceFindParticipatingTenantsRequest{
							TenantId:  "tenant-a",
							Namespace: "b",
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceFindParticipatingTenantsResponse{})
						},
					},
					{
						WantRequest: &v1.TenantServiceListTenantMembersRequest{
							TenantId:  "tenant-a",
							Namespace: "b",
						},
						WantResponse: func() connect.AnyResponse {
							return connect.NewResponse(&v1.TenantServiceListTenantMembersResponse{})
						},
					},
				}),
			},
		})

		_, err = client.Apiv1().ProjectMember().Create(t.Context(), &v1.ProjectMemberServiceCreateRequest{
			ProjectMember: &v1.ProjectMember{
				ProjectId: "project-a",
				TenantId:  "tenant-a",
				Namespace: "b",
			},
		})
		require.NoError(t, err)

		_, err = client.Apiv1().ProjectMember().List(t.Context(), &v1.ProjectMemberServiceListRequest{
			ProjectId: new("project-a"),
			TenantId:  new("tenant-a"),
			Namespace: "b",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().TenantMember().Create(t.Context(), &v1.TenantMemberServiceCreateRequest{
			TenantMember: &v1.TenantMember{
				TenantId:  "tenant-a",
				Namespace: "b",
			},
		})
		require.NoError(t, err)

		_, err = client.Apiv1().TenantMember().List(t.Context(), &v1.TenantMemberServiceListRequest{
			TenantId:  new("tenant-a"),
			Namespace: "b",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().FindParticipatingProjects(t.Context(), &v1.TenantServiceFindParticipatingProjectsRequest{
			TenantId:  "tenant-a",
			Namespace: "b",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().FindParticipatingTenants(t.Context(), &v1.TenantServiceFindParticipatingTenantsRequest{
			TenantId:  "tenant-a",
			Namespace: "b",
		})
		require.NoError(t, err)

		_, err = client.Apiv1().Tenant().ListTenantMembers(t.Context(), &v1.TenantServiceListTenantMembersRequest{
			TenantId:  "tenant-a",
			Namespace: "b",
		})
		require.NoError(t, err)
	})
}
