package http

import (
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateJWT(key string, claims *UserClaims) (string, time.Time, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(120 * time.Minute)

	// create claims
	c := token.Claims.(jwt.MapClaims)
	c["exp"] = expireTime.Unix()
	c["username"] = claims.Username
	c["is_admin"] = claims.IsAdmin
	c["banned"] = claims.Banned
	c["active"] = claims.Active
	c["access_level"] = claims.AccessLevel

	// generate token string
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", expireTime, err
	}

	return tokenString, expireTime, nil
}

func parseJWT(key string, token string) *UserClaims {
	return &UserClaims{}
}

func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
