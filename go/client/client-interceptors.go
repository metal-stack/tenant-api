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
			stream, err := next(ctx, spec)
			if err != nil {
				return nil, err
			}
			return &loggingClientStream{ClientStream: stream, log: log, procedure: spec.Procedure}, nil
		}
	}
}

type loggingClientStream struct {
	connect.ClientStream
	log       *slog.Logger
	procedure string
}

func (s *loggingClientStream) Send(msg any) error {
	s.log.Debug("request", "procedure", s.procedure, "payload", msg)
	return s.ClientStream.Send(msg)
}

func (s *loggingClientStream) Receive(msg any) error {
	err := s.ClientStream.Receive(msg)
	s.log.Debug("response", "procedure", s.procedure, "payload", msg, "error", err)
	return err
}

func NamespaceInterceptor(namespace string) connect.ClientInterceptor {
	return func(next connect.ClientFunc) connect.ClientFunc {
		return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
			stream, err := next(ctx, spec)
			if err != nil {
				return nil, err
			}
			return &namespaceClientStream{ClientStream: stream, namespace: namespace}, nil
		}
	}
}

type namespaceClientStream struct {
	connect.ClientStream
	namespace string
}

func (s *namespaceClientStream) Send(msg any) error {
	switch r := msg.(type) {
	case interface {
		GetNamespace() string
	}:
		if r.GetNamespace() == "" {
			reflect.Indirect(reflect.ValueOf(r)).FieldByName("Namespace").Set(reflect.ValueOf(s.namespace))
		}
	case interface {
		GetProjectMember() *v1.ProjectMember
	}:
		if r.GetProjectMember().Namespace == "" {
			r.GetProjectMember().Namespace = s.namespace
		}
	case interface {
		GetTenantMember() *v1.TenantMember
	}:
		if r.GetTenantMember().Namespace == "" {
			r.GetTenantMember().Namespace = s.namespace
		}
	}

	return s.ClientStream.Send(msg)
}
