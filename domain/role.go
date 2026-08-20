package domain

import "strings"

// RolePrivacyNodeOperator is the canonical Access Manager role label whose members hold
// operator authority for governance writes. It also gates the admin endpoints via RequireRole.
const RolePrivacyNodeOperator = "PRIVACY_NODE_OPERATOR"

// RoleBankEmployee is the canonical Access Manager role label that gates the tokenization-flow
// governance writes on-chain (addToken, addAddressPair). The operator signer must hold this role in
// addition to RolePrivacyNodeOperator to sign every governance write.
const RoleBankEmployee = "BANK_EMPLOYEE"

// RoleFactoryDeployer is the canonical Access Manager role name/label that RNContractFactory.deploy
// requires the caller to hold. Auto-provisioning grants it to each new custody wallet so the user can
// deploy tokens; it must match the role registered on-chain (registerRole("FACTORY_DEPLOYER")).
const RoleFactoryDeployer = "FACTORY_DEPLOYER"

// NormalizeRoleLabel converts a human-readable role label to a SCREAMING_SNAKE_CASE identifier.
// e.g. "Privacy Node Operator" → "PRIVACY_NODE_OPERATOR". On-chain labels and the values stored by
// the Access Manager indexer are not canonicalized, so callers normalize before comparing.
func NormalizeRoleLabel(label string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(label), " ", "_"))
}
