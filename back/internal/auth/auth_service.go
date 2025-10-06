package auth

import (
	"back/internal/config"
	"context"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Service struct {
	verifier *oidc.IDTokenVerifier
	config   *oauth2.Config
}

func NewService(cfg *config.Config, ctx context.Context) *Service {
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: cfg.OAuth2.ClientID,
	}
	v := provider.Verifier(oidcConfig)

	c := oauth2.Config{
		ClientID:     cfg.OAuth2.ClientID,
		ClientSecret: cfg.OAuth2.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5556/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Service{
		verifier: v,
		config:   &c,
	}
}
