package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_ReceiveBatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		username   string
		password   string
		statusCode int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := NewMockController(t)
			rr := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.SetBasicAuth(tt.username, tt.password)
			c.ReceiveBatch(rr, r)
		})
	}
}
