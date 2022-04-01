package api

import (
	"crypto/subtle"
	"net/http"
	"time"

	"github.com/barklan/logbeam/pkg/security"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type AuthToken struct {
	AccessToken string `json:"access_token"` //nolint: tagliatelle
}

func (a *AuthToken) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Controller) getAuthToken(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	password := chi.URLParam(r, "password")
	c.log.Info("usernames", zap.String("conf", c.conf.Username), zap.String("query", username))
	if c.conf.Username == username {
		if !(subtle.ConstantTimeCompare([]byte(password), []byte(c.conf.Password)) == 1) {
			c.errorResp(w, r, "Password is wrong.", http.StatusForbidden)

			return
		}
	} else {
		c.errorResp(w, r, "No user found with this username.", http.StatusNotFound)

		return
	}
	token, err := security.CreateJWT(username, c.conf.Secret, 48*time.Hour)
	if err != nil {
		c.internalError(w, r, "Failed to create token. Please open an issue.", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if render.Render(w, r, &AuthToken{AccessToken: token}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
