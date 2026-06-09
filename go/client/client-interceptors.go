package client

import (
	"context"
	"reflect"

	"connectrpc.com/connect"
	v1 "github.com/metal-stack/tenant-api/go/api/v1"
)

// authinterceptor adds the required auth headers
type authInterceptor struct {
	config *DialConfig
}

func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		request.Header().Add("Authorization", "Bearer "+i.config.Token)
		return next(ctx, request)
	})
}

func (i *authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return &streamingInterceptorConn{
			StreamingClientConn: next(ctx, spec),
			token:               i.config.Token,
		}
	}
}

func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

type streamingInterceptorConn struct {
	connect.StreamingClientConn
	token string
}

func (conn *streamingInterceptorConn) Send(m any) error {
	conn.RequestHeader().Add("Authorization", "Bearer "+conn.token)
	return conn.StreamingClientConn.Send(m)
}

type loggingInterceptor struct {
	config *DialConfig
}

func (i *loggingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		i.config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "body", request.Any())
		response, err := next(ctx, request)
		if err != nil {
			return nil, err
		}
		i.config.Log.Debug("intercept", "request procedure", request.Spec().Procedure, "response", response.Any())
		return response, err
	})
}

func (i *loggingInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	// TODO also log here
	return next
}

func (i *loggingInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return next
}

func NamespaceInterceptor(namespace string) connect.UnaryInterceptorFunc {
	return func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			switch r := ar.Any().(type) {
			case interface {
				GetNamespace() string
			}:
				if r.GetNamespace() == "" {
					reflect.Indirect(reflect.ValueOf(r)).FieldByName("Namespace").Set(reflect.ValueOf(namespace))
				}
			case interface {
				GetProjectMember() *v1.ProjectMember
			}:
				if r.GetProjectMember().Namespace == "" {
					r.GetProjectMember().Namespace = namespace
				}
			case interface {
				GetTenantMember() *v1.TenantMember
			}:
				if r.GetTenantMember().Namespace == "" {
					r.GetTenantMember().Namespace = namespace
				}
			}

			return uf(ctx, ar)
		}
	}
}

func userAgentInterceptor(userAgent string) connect.UnaryInterceptorFunc {
	return func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			ar.Header().Add("User-Agent", userAgent)
			return uf(ctx, ar)
		}
	}
}
