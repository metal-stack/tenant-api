// Code generated generate_clients.go. DO NOT EDIT.
package client

import (
{{ range $name, $api := . -}}
	"github.com/metal-stack/tenant-api/go{{ $api.Path }}/{{ $api.Name }}connect"
{{ end }}
)

type (
	Client interface {
{{ range $name, $api := . -}}
	{{ $name | title }}() {{ $name | title }}
{{ end }}

	}

{{ range $name, $api := . -}}
	{{ $name | title }} interface {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | trimSuffix "Service" }}() {{ $name }}connect.{{ $svc.Name }}Client
{{ end }}
	}

    {{ $name }} struct {
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }} {{ $name }}connect.{{ $svc.Name }}Client
{{ end }}
    }

{{ end }}
)

{{ range $name, $api := . -}}
func (c *client) {{ $name | title }}() {{ $name | title }} {
	a := &{{ $name }}{
{{ range $svc := $api.Services -}}
	{{ $svc.Name | lower }}:  {{ $name }}connect.New{{ $svc.Name }}Client(c.httpClient),
{{ end }}
	}
	return a
}

{{ range $svc := $api.Services -}}
func (c  *{{ $name }} ) {{ $svc.Name | trimSuffix "Service" }}() {{ $name }}connect.{{ $svc.Name }}Client {
	return c.{{ $svc.Name | lower }}
}
{{ end }}

{{ end }}
