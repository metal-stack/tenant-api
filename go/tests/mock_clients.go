// Code generated generate_clients.go. DO NOT EDIT.
package apitests

import (
	"testing"

	apiclient "github.com/metal-stack/tenant-api/go/client"
	"github.com/metal-stack/tenant-api/go/tenant/api/v1/apiv1connect"
	apiv1mocks "github.com/metal-stack/tenant-api/go/tests/mocks/tenant/api/v1/apiv1connect"

	"github.com/stretchr/testify/mock"
)

type (
	client struct {
		apiv1service *apiv1
	}

	ClientMockFns struct {
		Apiv1Mocks *Apiv1MockFns
	}

	wrapper struct {
		t *testing.T
	}
	apiv1 struct {
		healthservice        *apiv1mocks.HealthServiceClient
		projectservice       *apiv1mocks.ProjectServiceClient
		projectmemberservice *apiv1mocks.ProjectMemberServiceClient
		tenantservice        *apiv1mocks.TenantServiceClient
		tenantmemberservice  *apiv1mocks.TenantMemberServiceClient
		versionservice       *apiv1mocks.VersionServiceClient
	}

	Apiv1MockFns struct {
		Health        func(m *mock.Mock)
		Project       func(m *mock.Mock)
		ProjectMember func(m *mock.Mock)
		Tenant        func(m *mock.Mock)
		TenantMember  func(m *mock.Mock)
		Version       func(m *mock.Mock)
	}
)

func New(t *testing.T) *wrapper {
	return &wrapper{t: t}
}

func (w wrapper) Client(fns *ClientMockFns) *client {
	return &client{
		apiv1service: w.Apiv1(fns.Apiv1Mocks),
	}
}

func (c *client) Apiv1() apiclient.Apiv1 {
	return c.apiv1service
}

func (w wrapper) Apiv1(fns *Apiv1MockFns) *apiv1 {
	return newapiv1(w.t, fns)
}

func newapiv1(t *testing.T, fns *Apiv1MockFns) *apiv1 {
	a := &apiv1{
		healthservice:        apiv1mocks.NewHealthServiceClient(t),
		projectservice:       apiv1mocks.NewProjectServiceClient(t),
		projectmemberservice: apiv1mocks.NewProjectMemberServiceClient(t),
		tenantservice:        apiv1mocks.NewTenantServiceClient(t),
		tenantmemberservice:  apiv1mocks.NewTenantMemberServiceClient(t),
		versionservice:       apiv1mocks.NewVersionServiceClient(t),
	}

	if fns != nil {
		if fns.Health != nil {
			fns.Health(&a.healthservice.Mock)
		}
		if fns.Project != nil {
			fns.Project(&a.projectservice.Mock)
		}
		if fns.ProjectMember != nil {
			fns.ProjectMember(&a.projectmemberservice.Mock)
		}
		if fns.Tenant != nil {
			fns.Tenant(&a.tenantservice.Mock)
		}
		if fns.TenantMember != nil {
			fns.TenantMember(&a.tenantmemberservice.Mock)
		}
		if fns.Version != nil {
			fns.Version(&a.versionservice.Mock)
		}

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
