package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"connectrpc.com/connect/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/testing/protocmp"
)

type TestClientInterceptor struct {
	t     *testing.T
	calls []ClientCall
	count int
}

type ClientCall struct {
	WantRequest  proto.Message
	WantResponse func() proto.Message
	WantError    *connect.Error
}

func NewTestInterceptor(t *testing.T, calls []ClientCall) connect.ClientInterceptor {
	tci := &TestClientInterceptor{
		t:     t,
		calls: calls,
	}
	return tci.intercept
}

func (t *TestClientInterceptor) intercept(connect.ClientFunc) connect.ClientFunc {
	return func(ctx context.Context, spec connect.Spec) (connect.ClientStream, error) {
		if t.count >= len(t.calls) {
			t.t.Errorf("received an unexpected client call: %v", spec.Procedure)
			t.t.FailNow()
		}

		call := t.calls[t.count]
		t.count++

		return &testClientStream{t: t.t, call: call}, nil
	}
}

type testClientStream struct {
	t    *testing.T
	call ClientCall
}

func (s *testClientStream) SendHeaders() error { return nil }

func (s *testClientStream) Send(msg any) error {
	req, ok := msg.(proto.Message)
	if !ok {
		s.t.Errorf("request is not a proto.Message: %T", msg)
		s.t.FailNow()
	}

	if diff := cmp.Diff(s.call.WantRequest, req, protocmp.Transform(), IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
		s.t.Errorf("request diff (+got -want):\n %s", diff)
		s.t.FailNow()
	}

	return nil
}

func (s *testClientStream) CloseSend() error { return nil }

func (s *testClientStream) Receive(msg any) error {
	if s.call.WantError != nil {
		return s.call.WantError
	}

	res, ok := msg.(proto.Message)
	if !ok {
		return connect.NewError(connect.CodeInternal, fmt.Sprintf("response is not a proto.Message: %T", msg))
	}

	proto.Merge(res, s.call.WantResponse())

	return nil
}

func (s *testClientStream) Close() error { return nil }

func IgnoreUnexported() cmp.Option {
	// the exporter opt allows all unexported fields: https://github.com/google/go-cmp/pull/176
	return cmp.Exporter(func(reflect.Type) bool { return true })
}
