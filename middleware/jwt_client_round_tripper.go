package middleware

import (
	"go.uber.org/zap"
	"net/http"
)

type JwtRoundTripper struct {
	Proxied http.RoundTripper
	Logger  *zap.Logger
	authKey string
}

func NewJwtRoundTripper(proxied http.RoundTripper, logger *zap.Logger, authKey string) *JwtRoundTripper {
	return &JwtRoundTripper{
		Proxied: proxied,
		Logger:  logger,
		authKey: authKey,
	}
}

func (r *JwtRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+r.authKey)

	return r.Proxied.RoundTrip(req)
}
