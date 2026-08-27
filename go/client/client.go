// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
	"connectrpc.com/connect"
	compress "github.com/klauspost/connect-compress/v2"

	"github.com/metal-stack/tenant-api/go/tenant/api/v1/apiv1connect"
)

type (
	Client interface {
		Apiv1() Apiv1
	}

	client struct {
		config *DialConfig

		interceptors []connect.Interceptor
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

func New(config *DialConfig) (Client, error) {
	c := &client{
		config:       config,
		interceptors: []connect.Interceptor{},
	}

	if config.Token != "" {
		authInterceptor := &authInterceptor{config: config}
		c.interceptors = append(c.interceptors, authInterceptor)
	}
	if config.Log != nil {
		loggingInterceptor := &loggingInterceptor{config: config}
		c.interceptors = append(c.interceptors, loggingInterceptor)
	}
	if config.Namespace != "" {
		c.interceptors = append(c.interceptors, NamespaceInterceptor(config.Namespace))
	}
	if config.UserAgent != "" {
		c.interceptors = append(c.interceptors, userAgentInterceptor(config.UserAgent))
	}
	c.interceptors = append(c.interceptors, config.Interceptors...)

	return c, nil
}

func (c *client) Apiv1() Apiv1 {
	a := &apiv1{
		healthservice: apiv1connect.NewHealthServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		projectservice: apiv1connect.NewProjectServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		projectmemberservice: apiv1connect.NewProjectMemberServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		tenantservice: apiv1connect.NewTenantServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		tenantmemberservice: apiv1connect.NewTenantMemberServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
		versionservice: apiv1connect.NewVersionServiceClient(
			c.config.HttpClient(),
			c.config.BaseURL,
			connect.WithInterceptors(c.interceptors...),
			compress.WithAll(compress.LevelBalanced),
		),
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
