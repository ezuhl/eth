package data

import (
	"github.com/davidebianchi/go-jsonclient"
	"github.com/eth/internal/data/model"
	"testing"
)

func Test_prysmClient_GetChainedHead(t *testing.T) {
	type fields struct {
		client *jsonclient.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.ChainHeadResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful request",
			fields: fields{
				client: func() *jsonclient.Client {
					opts := jsonclient.Options{
						BaseURL: "http://3.70.154.20:3500/eth/v1alpha1/",
					}
					client, err := jsonclient.New(opts)
					if err != nil {
						t.Fatal(err)
					}
					return client
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &prysmClient{
				client: tt.fields.client,
			}
			got, err := p.GetChainedHead()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChainedHead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("GetChainedHead() got = %v, want %v", got, tt.want)
			}
		})
	}
}
