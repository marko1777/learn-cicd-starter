package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "No Header",
			headers: map[string][]string{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No ApiKey",
			headers: map[string][]string{"Authorization": {"foo bar"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No length",
			headers: map[string][]string{"Authorization": {"ApiKey"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Good",
			headers: map[string][]string{"Authorization": {"ApiKey bar"}},
			want:    "bar",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
