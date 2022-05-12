package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/barklan/logbeam/pkg/security"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestServer_getAuthToken(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		username   string
		password   string
		statusCode int
	}{
		{
			"good creds",
			"logbeam",
			"logbeam",
			http.StatusOK,
		},
		{
			"wrong username",
			"bobik",
			"logbeam",
			http.StatusNotFound,
		},
		{
			"wrong password",
			"logbeam",
			"logbeam2",
			http.StatusForbidden,
		},
	}
	for _, tt := range tests { // nolint:paralleltest
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := NewTestServer(t)
			rr := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("username", tt.username)
			rctx.URLParams.Add("password", tt.password)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			s.getAuthToken(rr, r)
			require.Equal(t, tt.statusCode, rr.Code)

			if rr.Code == http.StatusOK {
				var auth AuthToken
				err := json.Unmarshal(rr.Body.Bytes(), &auth)
				require.NoError(t, err)
				t.Logf("access_token: %q", auth.AccessToken)
				ok, err := security.ValidateJWT(auth.AccessToken, s.conf.Secret)
				require.NoError(t, err)
				require.True(t, ok)
			} else {
				var errResp ErrorResp
				err := json.Unmarshal(rr.Body.Bytes(), &errResp)
				require.NoError(t, err)
				require.GreaterOrEqual(t, len(errResp.Error.Message), 10)
			}
		})
	}
}
