package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/raylsnetwork/rayls-privacy-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-privacy-ops-api/domain"
	"github.com/raylsnetwork/rayls-privacy-ops-api/logger"
)

var _ core.TeleportService = (*teleportService)(nil)

// teleportRegistry is the minimal registry read the teleport preflight needs: confirm the token is
// registered and active. Satisfied by core.TokenRegistryService.
type teleportRegistry interface {
	Exists(ctx context.Context, tokenAddress string) (bool, error)
}

// teleportService owns the teleport business rules: only ERC20/ERC721/ERC1155 are eligible, and a
// mandatory preflight (token registered & active, caller balance/ownership) runs before the
// teleportToPublicChain transaction is signed. It owns the destination (public) chain, passed to the
// TokenChainClient as destinationChainId. On-chain interaction is delegated to a TokenChainClient.
type teleportService struct {
	tokenClient   core.TokenChainClient
	registry      teleportRegistry
	publicChainID *big.Int
	log           logger.Logger
}

func NewTeleportService(
	tokenClient core.TokenChainClient,
	registry teleportRegistry,
	publicChainID *big.Int,
	log logger.Logger,
) core.TeleportService {
	return &teleportService{
		tokenClient:   tokenClient,
		registry:      registry,
		publicChainID: publicChainID,
		log:           log,
	}
}

// Teleport validates the standard, runs the registry + balance/ownership preflight, then delegates
// the signed teleportToPublicChain transaction to the chain client.
func (s *teleportService) Teleport(
	ctx context.Context,
	tokenAddress string,
	standard domain.ErcStandard,
	in core.TeleportInput,
) (string, error) {
	exists, err := s.registry.Exists(ctx, tokenAddress)
	if err != nil {
		return "", fmt.Errorf("check token registry: %w", err)
	}
	if !exists {
		return "", core.NewValidationError("", "token does not exist or inactive")
	}

	var txHash string
	switch standard {
	case domain.ErcStandardERC20:
		if err = s.checkERC20Balance(ctx, tokenAddress, in.From, in.Amount); err != nil {
			return "", err
		}
		txHash, err = s.tokenClient.TeleportERC20(ctx, in.From, tokenAddress, in.To, in.Amount, s.publicChainID)
	case domain.ErcStandardERC721:
		if err = s.checkERC721Ownership(ctx, tokenAddress, in.From, in.TokenID); err != nil {
			return "", err
		}
		txHash, err = s.tokenClient.TeleportERC721(ctx, in.From, tokenAddress, in.To, in.TokenID, s.publicChainID)
	case domain.ErcStandardERC1155:
		if err = s.checkERC1155Balance(ctx, tokenAddress, in.From, in.TokenID, in.Amount); err != nil {
			return "", err
		}
		txHash, err = s.tokenClient.TeleportERC1155(
			ctx,
			in.From,
			tokenAddress,
			in.To,
			in.TokenID,
			in.Amount,
			s.publicChainID,
			in.Data,
		)
	default:
		return "", core.NewValidationError("standard", "teleport supports only ERC20, ERC721 and ERC1155")
	}
	if err != nil {
		return "", err
	}

	s.log.Info("token teleported", "address", tokenAddress, "standard", standard.Label(), "txHash", txHash)
	return txHash, nil
}

func (s *teleportService) checkERC20Balance(ctx context.Context, tokenAddress, account string, amount *big.Int) error {
	if amount == nil {
		return core.NewValidationError("amount", "amount is required")
	}
	balance, err := s.tokenClient.ERC20Balance(ctx, tokenAddress, account)
	if err != nil {
		return fmt.Errorf("query ERC20 balance: %w", err)
	}
	if balance.Cmp(amount) < 0 {
		return core.NewValidationError(
			"amount",
			fmt.Sprintf("insufficient balance: have %s, need %s", balance, amount),
		)
	}
	return nil
}

func (s *teleportService) checkERC721Ownership(
	ctx context.Context,
	tokenAddress, account string,
	tokenID *big.Int,
) error {
	if tokenID == nil {
		return core.NewValidationError("tokenId", "tokenId is required")
	}
	owner, err := s.tokenClient.ERC721Owner(ctx, tokenAddress, tokenID)
	if err != nil {
		return fmt.Errorf("query ERC721 owner: %w", err)
	}
	if domain.NormalizeAddress(owner) != domain.NormalizeAddress(account) {
		return core.NewValidationError("tokenId", fmt.Sprintf("caller does not own token %s", tokenID))
	}
	return nil
}

func (s *teleportService) checkERC1155Balance(
	ctx context.Context,
	tokenAddress, account string,
	tokenID, amount *big.Int,
) error {
	if tokenID == nil || amount == nil {
		return core.NewValidationError("amount", "id and amount are required")
	}
	balance, err := s.tokenClient.ERC1155Balance(ctx, tokenAddress, account, tokenID)
	if err != nil {
		return fmt.Errorf("query ERC1155 balance: %w", err)
	}
	if balance.Cmp(amount) < 0 {
		return core.NewValidationError(
			"amount",
			fmt.Sprintf("insufficient balance for token %s: have %s, need %s", tokenID, balance, amount),
		)
	}
	return nil
}
