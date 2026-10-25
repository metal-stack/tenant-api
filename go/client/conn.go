package client

import (
	"log/slog"
	"net/http"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connectgzip"
	"connectrpc.com/connect/v2/connecthttp"
)

type (
	// DialConfig is the configuration to create an tenant-apiserver connection
	DialConfig struct {
		BaseURL string
		Token   string

		Namespace string

		// Optional client Interceptors
		Interceptors []connect.ClientInterceptor

		UserAgent string

		Transport http.RoundTripper

		Log *slog.Logger
	}
	client struct {
		config *DialConfig

		interceptors []connect.ClientInterceptor

		httpClient *connect.Client
	}
)

func New(config *DialConfig) (Client, error) {
	c := &client{
		config:       config,
		interceptors: []connect.ClientInterceptor{},
	}

	if config.Token != "" {
		authInterceptor := newAuthInterceptor(config.Token)
		c.interceptors = append(c.interceptors, authInterceptor)
	}
	if config.Log != nil {
		loggingInterceptor := newLoggingInterceptor(config.Log)
		c.interceptors = append(c.interceptors, loggingInterceptor)
	}
	if config.Namespace != "" {
		c.interceptors = append(c.interceptors, NamespaceInterceptor(config.Namespace))
	}
	if config.UserAgent != "" {
		c.interceptors = append(c.interceptors, newUserAgentInterceptor(config.UserAgent))
	}
	c.interceptors = append(c.interceptors, config.Interceptors...)

	gzipCompressor := connectgzip.New()
	options := []connecthttp.Option{
		connecthttp.WithCompressors(gzipCompressor),
		connecthttp.WithSendCompression(gzipCompressor.Name()),
	}
	c.httpClient = connect.NewClient(
		connecthttp.NewTransport(c.config.HttpClient(), c.config.BaseURL, options...),
		c.interceptors...,
	)

	return c, nil
}

func (d *DialConfig) HttpClient() *http.Client {
	transport := http.DefaultTransport
	if d.Transport != nil {
		transport = d.Transport
	}

	return &http.Client{
		Transport: transport,
	}
}
