package internal

import (
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"net/http"
	"net/url"
	"time"
)

type AccessTokenMid func(next http.Handler) http.Handler

func NewAccessTokenMid(
	issuer string,
	audience []string,
	signatureAlgorithm validator.SignatureAlgorithm,
	customClaims validator.CustomClaims,
) (AccessTokenMid, error) {
	issuerURL, err := url.Parse(issuer)
	if err != nil {
		return nil, err
	}

	keyProvider := jwks.NewCachingProvider(issuerURL, time.Minute*5)

	jwtValidator, err := validator.New(
		keyProvider.KeyFunc,
		signatureAlgorithm,
		issuerURL.String(),
		audience,
		validator.WithCustomClaims(func() validator.CustomClaims {
			return customClaims
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(jwtmiddleware.DefaultErrorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}, nil
}
