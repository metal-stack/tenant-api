package client_test

import (
	"log/slog"
	"testing"

	"connectrpc.com/connect/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	client "github.com/metal-stack/tenant-api/go/client"
	apiv1 "github.com/metal-stack/tenant-api/go/tenant/api/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestInterceptor(t *testing.T) {
	cl, err := client.New(&client.DialConfig{
		BaseURL: "http://this-is-just-for-testing",
		Interceptors: []connect.ClientInterceptor{
			client.NewTestInterceptor(t, []client.ClientCall{
				{
					WantRequest: &apiv1.TenantServiceGetRequest{
						Id: "t1",
					},
					WantResponse: func() proto.Message {
						return &apiv1.TenantServiceGetResponse{
							Tenant: &apiv1.Tenant{
								Meta: &apiv1.Meta{Id: "t1"},
								Name: "T1",
							},
						}
					},
				},
			}),
		},
		UserAgent: "cli-test",
		Log:       slog.Default(),
	})
	require.NoError(t, err)

	resp, err := cl.Apiv1().Tenant().Get(t.Context(), &apiv1.TenantServiceGetRequest{
		Id: "t1",
	})
	require.NoError(t, err)

	if diff := cmp.Diff(&apiv1.TenantServiceGetResponse{
		Tenant: &apiv1.Tenant{
			Meta: &apiv1.Meta{Id: "t1"},
			Name: "T1",
		},
	}, resp, protocmp.Transform(), client.IgnoreUnexported(), cmpopts.IgnoreTypes(protoimpl.MessageState{})); diff != "" {
		t.Errorf("diff = %s", diff)
	}
}
