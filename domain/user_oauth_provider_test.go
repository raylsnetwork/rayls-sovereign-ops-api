package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOAuthProvider_ReturnsGoogleProvider(t *testing.T) {
	// "google" maps to OAuthProviderGoogle
	provider, err := ParseOAuthProvider("google")

	require.NoError(t, err)
	assert.Equal(t, OAuthProviderGoogle, provider)
}

func TestParseOAuthProvider_ReturnsMicrosoftProvider(t *testing.T) {
	// "microsoft" maps to OAuthProviderMicrosoft
	provider, err := ParseOAuthProvider("microsoft")

	require.NoError(t, err)
	assert.Equal(t, OAuthProviderMicrosoft, provider)
}

func TestParseOAuthProvider_ReturnsErrorForUnknownProvider(t *testing.T) {
	// An unrecognised provider name returns an error
	_, err := ParseOAuthProvider("facebook")

	require.Error(t, err)
}
