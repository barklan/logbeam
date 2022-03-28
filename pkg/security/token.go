package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJWT(username, signKey string, expire time.Duration) (string, error) {
	signingKey := []byte(signKey)

	expTime := time.Now().Add(expire)
	claims := TokenClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			Issuer:    "logbeam",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("failed to create jwt token: %w", err)
	}

	return ss, nil
}
