package domain

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ImmutableModel is the base for entities that are never updated after creation.
// Use this instead of Model when the corresponding table has no updated_at column.
type ImmutableModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

// TokenStatus represents the on-chain lifecycle state of a Rayls token.
// Values sourced from ITokenGovernance.TokenStatus in
// rayls-node/rayls-privacy-ledger/interfaces/ITokenGovernanceV1.sol.
type TokenStatus uint8

const (
	TokenStatusInactive   TokenStatus = 0
	TokenStatusActive     TokenStatus = 1
	TokenStatusPaused     TokenStatus = 2
	TokenStatusFrozen     TokenStatus = 3
	TokenStatusDeprecated TokenStatus = 4

	// TokenStatusInternal is an off-chain lifecycle state (not part of the on-chain
	// ITokenGovernance enum). A token deployed through POST /api/tokens is stored with this
	// status and stays internal — it is NOT promoted to Active when the indexer discovers it.
	TokenStatusInternal TokenStatus = 10
)

// Label returns a lowercase string label for the status (used in API/SSE payloads).
func (s TokenStatus) Label() string {
	switch s {
	case TokenStatusInactive:
		return "inactive"
	case TokenStatusActive:
		return "active"
	case TokenStatusPaused:
		return "paused"
	case TokenStatusFrozen:
		return "frozen"
	case TokenStatusDeprecated:
		return "deprecated"
	case TokenStatusInternal:
		return "internal"
	default:
		return "unknown"
	}
}

// PrivacyNodeStatus is the PN-controlled lifecycle status of a token in the on-chain TokenRegistry
// (PNTokenCoreV1). It is one of the three independent status fields the registry tracks; the ops-api,
// running on the Privacy Node operator side, drives and reads this one. Values mirror
// TokenStructs.PrivacyNodeStatus in the TokenRegistry contract.
//
// This is distinct from TokenStatus, which is the off-chain, indexer-assigned label stored in the
// tokens table and exposed by GET /api/tokens.
type PrivacyNodeStatus uint8

const (
	PrivacyNodeStatusUndefined       PrivacyNodeStatus = 0 // default, never stored
	PrivacyNodeStatusWaitingApproval PrivacyNodeStatus = 1 // registerToken() called; awaiting operator review
	PrivacyNodeStatusAuthorized      PrivacyNodeStatus = 2 // operator approved; operational on the PN
	PrivacyNodeStatusUnauthorized    PrivacyNodeStatus = 3 // operator rejected
	PrivacyNodeStatusFrozen          PrivacyNodeStatus = 4 // frozen by local admin; blocks all operations
)

// IsKnown reports whether the status is one this build understands. The registry contract is
// upgradeable, so it can start returning a value added after this API was built; callers use
// this to surface that drift (Label would just say "unknown") instead of silently passing it on.
func (s PrivacyNodeStatus) IsKnown() bool {
	switch s {
	case PrivacyNodeStatusUndefined,
		PrivacyNodeStatusWaitingApproval,
		PrivacyNodeStatusAuthorized,
		PrivacyNodeStatusUnauthorized,
		PrivacyNodeStatusFrozen:
		return true
	default:
		return false
	}
}

// Label returns a lowercase string label for the status (used in API/SSE payloads).
func (s PrivacyNodeStatus) Label() string {
	switch s {
	case PrivacyNodeStatusUndefined:
		return "undefined"
	case PrivacyNodeStatusWaitingApproval:
		return "waiting_approval"
	case PrivacyNodeStatusAuthorized:
		return "authorized"
	case PrivacyNodeStatusUnauthorized:
		return "unauthorized"
	case PrivacyNodeStatusFrozen:
		return "frozen"
	default:
		return "unknown"
	}
}

// ParseSettableStatus maps a request "status" string to the PrivacyNodeStatus it names, accepting
// only the two an operator may set directly. It mirrors ParseFreezeLayer/ParseSubmitTarget so all
// three registry endpoints take a self-documenting string enum rather than a raw on-chain number.
//
// "waiting_approval" and "undefined" are registration-assigned, and "frozen" has dedicated
// freeze/unfreeze endpoints (routing it through updatePrivacyNodeStatus would bypass them), so all
// three return ok=false.
func ParseSettableStatus(s string) (PrivacyNodeStatus, bool) {
	switch s {
	case PrivacyNodeStatusAuthorized.Label():
		return PrivacyNodeStatusAuthorized, true
	case PrivacyNodeStatusUnauthorized.Label():
		return PrivacyNodeStatusUnauthorized, true
	default:
		return PrivacyNodeStatusUndefined, false
	}
}

// FreezeLayer identifies which of the TokenRegistry's independent freeze layers an admin
// freeze/unfreeze targets. Each layer has its own on-chain FROZEN state and owner (see
// TokenRegistrySDD.doc). The ops-api operator drives the two address-based layers below; the Hub
// (PNH) layer is intentionally not yet supported — its contract methods take a
// FrozenToken{resourceId, chainIds} and are cross-chain PNH/relayer callbacks, not an operator action.
type FreezeLayer string

const (
	// FreezeLayerPrivacyNode freezes at the PN level (freezeOnPrivacyNode). Blocks ALL operations.
	FreezeLayerPrivacyNode FreezeLayer = "privacy_node"
	// FreezeLayerPublicChain freezes at the public-chain level (freezeOnPublicChain). Blocks public
	// chain operations only.
	FreezeLayerPublicChain FreezeLayer = "public_chain"
)

// ParseFreezeLayer maps a request "layer" string to a supported FreezeLayer. It returns ok=false for
// unknown values and for the Hub (PNH) layer, which is not yet supported.
func ParseFreezeLayer(s string) (FreezeLayer, bool) {
	switch FreezeLayer(s) {
	case FreezeLayerPrivacyNode:
		return FreezeLayerPrivacyNode, true
	case FreezeLayerPublicChain:
		return FreezeLayerPublicChain, true
	default:
		return "", false
	}
}

// SubmitTarget identifies which layer an authorized token is being submitted to. After a token is
// AUTHORIZED on the Privacy Node, the operator can submit it to the Hub and/or the Public Chain (see
// TokenRegistrySDD.doc). Submitting only initiates the flow — activation on the Hub and deployment on
// the Public Chain complete later via cross-chain PNH / relayer callbacks, not through the ops-api.
type SubmitTarget string

const (
	// SubmitTargetHub submits the token to the Private Hub (submitToHub): sends addToken() to the PNH
	// and moves hubStatus to WAITING_APPROVAL.
	SubmitTargetHub SubmitTarget = "hub"
	// SubmitTargetPublicChain submits the token to the Public Chain (submitToPublicChain): moves
	// publicChainStatus to PENDING_DEPLOYMENT.
	SubmitTargetPublicChain SubmitTarget = "public_chain"
)

// ParseSubmitTarget maps a request "target" string to a supported SubmitTarget. It returns ok=false
// for unknown values.
func ParseSubmitTarget(s string) (SubmitTarget, bool) {
	switch SubmitTarget(s) {
	case SubmitTargetHub:
		return SubmitTargetHub, true
	case SubmitTargetPublicChain:
		return SubmitTargetPublicChain, true
	default:
		return "", false
	}
}

// ErcStandard represents the token standard reported by the TokenRegistry contract.
// Values sourced from RaylsNodeBridgeableERC in
// rayls-node/rayls-privacy-ledger/RNMessageLib.sol.
type ErcStandard uint8

const (
	ErcStandardCustom       ErcStandard = 0
	ErcStandardERC20        ErcStandard = 1
	ErcStandardERC721       ErcStandard = 2
	ErcStandardERC1155      ErcStandard = 3
	ErcStandardEnygma       ErcStandard = 4
	ErcStandardZkDvpERC721  ErcStandard = 5
	ErcStandardZkDvpERC1155 ErcStandard = 6

	// ErcStandardStableCoin is a Rayls-native, Circle-style stablecoin (pause + blacklist +
	// master-minter/controllers) deployed through the RNContractFactory under the
	// "RAYLS_STABLECOIN" key. NOT stock Circle FiatTokenV2 bytecode — see
	// openspec/changes/add-stablecoin-token-type/proposal.md.
	ErcStandardStableCoin ErcStandard = 7
)

// Label returns the protocol's canonical name for the standard (the RNContractFactory key),
// e.g. "RAYLS_ERC20". Unrecognized/Custom standards return "CUSTOM".
func (e ErcStandard) Label() string {
	switch e {
	case ErcStandardERC20:
		return "RAYLS_ERC20"
	case ErcStandardERC721:
		return "RAYLS_ERC721"
	case ErcStandardERC1155:
		return "RAYLS_ERC1155"
	case ErcStandardEnygma:
		return "RAYLS_ENYGMA"
	case ErcStandardZkDvpERC721:
		return "RAYLS_ERC721_DVP"
	case ErcStandardZkDvpERC1155:
		return "RAYLS_ERC1155_DVP"
	case ErcStandardStableCoin:
		return "RAYLS_STABLECOIN"
	default:
		return "CUSTOM"
	}
}

// ParseErcStandard maps a token "type" string to the corresponding ErcStandard. It accepts
// every form the string can arrive in:
//
//   - Blockscout's RaylsTokenDiscovery fetcher output ("Rayls-ERC-20", "Rayls-Enygma", …),
//     which is what actually lands in the Blockscout `tokens.type` column and flows through the
//     indexer. These are the authoritative classification of a Rayls-native token.
//   - Blockscout's plain ERC fetcher output ("ERC-20"/"ERC-721"/"ERC-1155") for non-Rayls tokens
//     the Rayls fetcher never claimed.
//   - The protocol canonical labels ("RAYLS_ERC20", …) as returned by ErcStandard.Label(), so an
//     API `?type=` filter can be expressed in the same vocabulary the responses use.
//
// Unknown values map to ErcStandardCustom.
func ParseErcStandard(t string) ErcStandard {
	switch t {
	// Rayls-native fetcher strings (authoritative — these are what reach the indexer).
	case "Rayls-ERC-20":
		return ErcStandardERC20
	case "Rayls-ERC-721":
		return ErcStandardERC721
	case "Rayls-ERC-1155":
		return ErcStandardERC1155
	case "Rayls-Enygma":
		return ErcStandardEnygma
	case "Rayls-ERC-721-DVP":
		return ErcStandardZkDvpERC721
	case "Rayls-ERC-1155-DVP":
		return ErcStandardZkDvpERC1155
	case "Rayls-StableCoin":
		return ErcStandardStableCoin

	// Plain ERC strings for non-Rayls tokens the Rayls fetcher never overwrote.
	case "ERC-20":
		return ErcStandardERC20
	case "ERC-721":
		return ErcStandardERC721
	case "ERC-1155":
		return ErcStandardERC1155

	// Canonical labels, so an API filter can use the same vocabulary as responses.
	case "RAYLS_ERC20":
		return ErcStandardERC20
	case "RAYLS_ERC721":
		return ErcStandardERC721
	case "RAYLS_ERC1155":
		return ErcStandardERC1155
	case "RAYLS_ENYGMA":
		return ErcStandardEnygma
	case "RAYLS_ERC721_DVP":
		return ErcStandardZkDvpERC721
	case "RAYLS_ERC1155_DVP":
		return ErcStandardZkDvpERC1155
	case "RAYLS_STABLECOIN":
		return ErcStandardStableCoin

	default:
		return ErcStandardCustom
	}
}
