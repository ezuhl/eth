package accounttoken

import (
	"gopkg.in/dgrijalva/jwt-go.v3"
	"reflect"
	"testing"
	"time"
)

func Test_tokenAuth_MakeToken(t1 *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "generate",
			args:    args{333999},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tokenAuth{}
			got, err := t.MakeToken(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t1.Errorf("MakeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("MakeToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenAuth_ValidateToken(t1 *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    *SystemClaims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "parsed token",
			args: args{
				tokenString: func() string {
					t := &tokenAuth{}
					tk, err := t.MakeToken(123)
					if err != nil {
						return ""
					}
					return tk
				}(),
			},

			want: &SystemClaims{
				"123",
				jwt.StandardClaims{
					IssuedAt:  time.Now().Unix(),
					ExpiresAt: time.Now().Add(time.Second * 15000).Unix(),
					Issuer:    "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tokenAuth{}
			got, err := t.ValidateToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ValidateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
