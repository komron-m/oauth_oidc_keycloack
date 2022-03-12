package internal

import (
	"context"
	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/coreos/go-oidc/v3/oidc"
	"net/http"
)

type OpenIDCMid func(next http.Handler) http.Handler

func NewOpenIDCMid(issuer, clientID string) (OpenIDCMid, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawIDToken, err := jwtmiddleware.AuthHeaderTokenExtractor(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			_, err = verifier.Verify(r.Context(), rawIDToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}, nil
}
