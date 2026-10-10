package client

import (
	"context"
	"log/slog"
	"reflect"

	"connectrpc.com/connect/v2"
	v1 "github.com/metal-stack/tenant-api/go/tenant/api/v1"
)

// authinterceptor adds the required auth headers
func newAuthInterceptor(token string) connect.ClientInterceptor {
	return func(next connect.ClientFunc) connect.ClientFunc {
		return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
			if info, ok := connect.CallInfoForClientContext(ctx); ok {
				info.RequestHeader().Add("Authorization", "Bearer "+token)
			}
			return next(ctx, spec)
		}
	}
}

func newUserAgentInterceptor(userAgent string) connect.ClientInterceptor {
	return func(next connect.ClientFunc) connect.ClientFunc {
		return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
			if info, ok := connect.CallInfoForClientContext(ctx); ok {
				info.RequestHeader().Add("User-Agent", userAgent)
			}
			return next(ctx, spec)
		}
	}
}

func newLoggingInterceptor(log *slog.Logger) connect.ClientInterceptor {
	return func(next connect.ClientFunc) connect.ClientFunc {
		return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
			log.Debug("intercept", "request procedure", spec.Procedure)
			return next(ctx, spec)
		}
	}
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
