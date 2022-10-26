package mailtrap

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	token := "ddsdsdsd"

	type args struct {
		host  *string
		token *string
	}

	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "Create client",
			args: args{
				host:  nil,
				token: &token,
			},
			want:    &Client{HTTPClient: &http.Client{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(tt.args.host, tt.args.token)
			got := c.HTTPClient
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
