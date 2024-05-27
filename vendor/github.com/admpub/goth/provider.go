package goth

import (
	"context"
	"net/http"
	"sync"

	"golang.org/x/oauth2"
)

// Provider needs to be implemented for each 3rd party authentication provider
// e.g. Facebook, Twitter, etc...
type Provider interface {
	Name() string
	SetName(name string)
	BeginAuth(state string) (Session, error)
	UnmarshalSession(string) (Session, error)
	FetchUser(Session) (User, error)
	Debug(bool)
	RefreshToken(refreshToken string) (*oauth2.Token, error) // Get new access token based on the refresh token
	RefreshTokenAvailable() bool                             // Refresh token is provided by auth provider or not
}

const NoAuthUrlErrorMessage = "an AuthURL has not been set"

// Providers is list of known/available providers.
type Providers struct {
	m map[string]Provider
	l sync.RWMutex
}

var providers = NewProviders()

// UseProviders adds a list of available providers for use with Goth.
// Can be called multiple times. If you pass the same provider more
// than once, the last will be used.
func UseProviders(viders ...Provider) {
	providers.Add(viders...)
}

// GetProviders returns a list of all the providers currently in use.
func GetProviders() map[string]Provider {
	return providers.All()
}

// GetProvider returns a previously created provider. If Goth has not
// been told to use the named provider it will return an error.
func GetProvider(name string) (Provider, error) {
	return providers.Get(name)
}

// ClearProviders will remove all providers currently in use.
// This is useful, mostly, for testing purposes.
func ClearProviders() {
	providers.Clear()
}

// ContextForClient provides a context for use with oauth2.
func ContextForClient(h *http.Client) context.Context {
	if h == nil {
		return context.Background()
	}
	return context.WithValue(context.Background(), oauth2.HTTPClient, h)
}

// HTTPClientWithFallBack to be used in all fetch operations.
func HTTPClientWithFallBack(h *http.Client) *http.Client {
	if h != nil {
		return h
	}
	return http.DefaultClient
}
