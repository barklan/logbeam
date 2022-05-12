package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer_input(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		bodyFile    string
		statusCode  int
	}{
		{
			"simple ndjson",
			applicationNDJSON,
			"simplendjson",
			http.StatusAccepted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTestServer(t)
			rr := httptest.NewRecorder()
			body, err := os.Open(filepath.Join("testdata", tt.bodyFile))
			if err != nil {
				t.Fatalf("failed to open test body\n")
			}
			r := httptest.NewRequest(http.MethodPost, "/", body)
			r.SetBasicAuth("logbeam", "logbeam")
			s.input(rr, r)
			require.Equal(t, tt.statusCode, rr.Code)
		})
	}
}
