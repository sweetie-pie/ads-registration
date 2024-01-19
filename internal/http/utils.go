package http

import "encoding/base64"

func generateJWT(key string, claims *UserClaims) string {
	return ""
}

func parseJWT(key string, token string) *UserClaims {
	return &UserClaims{}
}

func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
