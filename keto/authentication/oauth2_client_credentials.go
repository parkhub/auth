package authentication

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/clientcredentials"
)

// swagger:model authenticationOAuth2ClientCredentialsSession
type OAuth2ClientCredentialsSession struct {
	// Here, it's subject
	*DefaultSession
}

type OAuth2ClientCredentialsAuthentication struct {
	tokenURL string
}

// swagger:model AuthenticationOAuth2ClientCredentialsRequest
type AuthenticationOAuth2ClientCredentialsRequest struct {
	// Token is the token to introspect.
	ClientID string `json:"client_id"`

	ClientSecret string `json:"client_secret"`

	// Scope is an array of scopes that are required.
	Scopes []string `json:"scope"`
}

func NewOAuth2ClientCredentialsSession() *OAuth2ClientCredentialsSession {
	return &OAuth2ClientCredentialsSession{
		DefaultSession: new(DefaultSession),
	}
}

func NewOAuth2ClientCredentialsAuthentication(tokenURL string) *OAuth2ClientCredentialsAuthentication {
	return &OAuth2ClientCredentialsAuthentication{
		tokenURL: tokenURL,
	}
}

func (a *OAuth2ClientCredentialsAuthentication) Authenticate(r *http.Request) (Session, error) {
	var auth AuthenticationOAuth2ClientCredentialsRequest

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		return nil, errors.WithStack(err)
	}

	c := &clientcredentials.Config{
		TokenURL:     a.tokenURL,
		ClientID:     auth.ClientID,
		ClientSecret: auth.ClientSecret,
		Scopes:       auth.Scopes,
	}

	token, err := c.Token(context.Background())
	if err != nil {
		return nil, errors.WithStack(ErrUnauthorized)
	} else if token.AccessToken == "" {
		return nil, errors.WithStack(ErrUnauthorized)
	}

	return &OAuth2ClientCredentialsSession{
		DefaultSession: &DefaultSession{
			Subject: auth.ClientID,
		},
	}, nil
}
