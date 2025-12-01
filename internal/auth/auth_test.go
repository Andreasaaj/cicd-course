package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "missing Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed header - no ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed header - only prefix",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid ApiKey header",
			headers: http.Header{
				"Authorization": []string{"ApiKey secret123"},
			},
			want:    "secret123",
			wantErr: false,
		},
		// 		{
		// 			name: "extra spacing but still valid",
		// 			headers: http.Header{
		// 				"Authorization": []string{"ApiKey   abc456"},
		// 			},
		// 			want:    "abc456",
		// 			wantErr: false,
		// 		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)

			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() returned error: %v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
