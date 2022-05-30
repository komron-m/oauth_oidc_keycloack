package internal

import (
	"context"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"net/http"
	"net/url"
	"strings"
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

	keyProvider := jwks.NewCachingProvider(issuerURL, time.Hour*24)

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

type Scope struct {
	Scopes string `json:"scope"`
}

func (s *Scope) Validate(ctx context.Context) error {
	return nil
}

func (s *Scope) HasScope(scope string) bool {
	givenScopes := strings.Split(s.Scopes, " ")
	for _, gs := range givenScopes {
		if gs == scope {
			return true
		}
	}
	return false
}

func ScopeControlMid(allowedScope string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

			claims, ok := token.CustomClaims.(*Scope)
			if !ok || !claims.HasScope(allowedScope) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message":"Insufficient scope."}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
