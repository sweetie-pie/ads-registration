package http

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errSigningMethod = errors.New("error in signing method")
	errInvalidToken  = errors.New("token is invalid")
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

func parseJWT(key string, token string) (*UserClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errSigningMethod
		}

		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	// taking out claims
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		c := &UserClaims{
			Username:    claims["username"].(string),
			IsAdmin:     claims["is_admin"].(bool),
			Banned:      claims["banned"].(bool),
			Active:      claims["active"].(bool),
			AccessLevel: claims["access_level"].(int),
		}

		return c, nil
	}

	return nil, errInvalidToken
}

func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
