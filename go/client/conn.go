package client

import (
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
)

type (
	// DialConfig is the configuration to create an tenant-apiserver connection
	DialConfig struct {
		BaseURL string
		Token   string

		Namespace string

		// Optional client Interceptors
		Interceptors []connect.Interceptor

		UserAgent string

		Transport http.RoundTripper

		Log *slog.Logger
	}
)

func (d *DialConfig) HttpClient() *http.Client {
	transport := http.DefaultTransport
	if d.Transport != nil {
		transport = d.Transport
	}

	return &http.Client{
		Transport: transport,
	}
}
