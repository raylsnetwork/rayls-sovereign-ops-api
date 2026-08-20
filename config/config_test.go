package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuth_validateDeployed_AllowsUnverifiedSignupInLocalDev(t *testing.T) {
	// Without SECURE_COOKIES the service is plain HTTP (local dev), so the dev-only
	// auth shortcuts stay allowed
	auth := Auth{SecureCookies: false, EmailSignupEnabled: true, BootstrapToken: ""}

	err := auth.validateDeployed()

	require.NoError(t, err)
}

func TestAuth_validateDeployed_RejectsEmailSignupWhenDeployed(t *testing.T) {
	// Unverified email signup is an authentication bypass once the service is reachable
	// over HTTPS, so it must refuse to boot
	auth := Auth{SecureCookies: true, EmailSignupEnabled: true, BootstrapToken: "tok"}

	err := auth.validateDeployed()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "EMAIL_SIGNUP_ENABLED")
}

func TestAuth_validateDeployed_RejectsMissingBootstrapTokenWhenDeployed(t *testing.T) {
	// An unauthenticated /admin/bootstrap lets the first caller claim admin, so a
	// deployed service must set a token
	auth := Auth{SecureCookies: true, EmailSignupEnabled: false, BootstrapToken: ""}

	err := auth.validateDeployed()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "BOOTSTRAP_TOKEN")
}

func TestAuth_validateDeployed_AcceptsHardenedDeployedConfig(t *testing.T) {
	// Signup off and a bootstrap token set is the intended production posture
	auth := Auth{SecureCookies: true, EmailSignupEnabled: false, BootstrapToken: "tok"}

	err := auth.validateDeployed()

	require.NoError(t, err)
}
