package address

import "testing"

func TestFindAddress(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "local-test",
			want:    "192.168.1.5",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindAddress()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
