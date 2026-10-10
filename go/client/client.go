// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"github.com/metal-stack/tenant-api/go/tenant/api/v1/apiv1connect"
)

type (
	Client interface {
		Apiv1() Apiv1
	}

	Apiv1 interface {
		Health() apiv1connect.HealthServiceClient
		Project() apiv1connect.ProjectServiceClient
		ProjectMember() apiv1connect.ProjectMemberServiceClient
		Tenant() apiv1connect.TenantServiceClient
		TenantMember() apiv1connect.TenantMemberServiceClient
		Version() apiv1connect.VersionServiceClient
	}

	apiv1 struct {
		healthservice        apiv1connect.HealthServiceClient
		projectservice       apiv1connect.ProjectServiceClient
		projectmemberservice apiv1connect.ProjectMemberServiceClient
		tenantservice        apiv1connect.TenantServiceClient
		tenantmemberservice  apiv1connect.TenantMemberServiceClient
		versionservice       apiv1connect.VersionServiceClient
	}
)

func (c *client) Apiv1() Apiv1 {
	a := &apiv1{
		healthservice:        apiv1connect.NewHealthServiceClient(c.httpClient),
		projectservice:       apiv1connect.NewProjectServiceClient(c.httpClient),
		projectmemberservice: apiv1connect.NewProjectMemberServiceClient(c.httpClient),
		tenantservice:        apiv1connect.NewTenantServiceClient(c.httpClient),
		tenantmemberservice:  apiv1connect.NewTenantMemberServiceClient(c.httpClient),
		versionservice:       apiv1connect.NewVersionServiceClient(c.httpClient),
	}
	return a
}

func (c *apiv1) Health() apiv1connect.HealthServiceClient {
	return c.healthservice
}
func (c *apiv1) Project() apiv1connect.ProjectServiceClient {
	return c.projectservice
}
func (c *apiv1) ProjectMember() apiv1connect.ProjectMemberServiceClient {
	return c.projectmemberservice
}
func (c *apiv1) Tenant() apiv1connect.TenantServiceClient {
	return c.tenantservice
}
func (c *apiv1) TenantMember() apiv1connect.TenantMemberServiceClient {
	return c.tenantmemberservice
}
func (c *apiv1) Version() apiv1connect.VersionServiceClient {
	return c.versionservice
}
