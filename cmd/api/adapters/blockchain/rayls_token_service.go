package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/cmd/api/core"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/domain"
)

var (
	_ core.TokenActionService = (*RaylsTokenService)(nil)
	_ core.TokenChainClient   = (*RaylsTokenService)(nil)
)

// abiWord is the size of one ABI-encoded return slot. A single-value return (address, bool)
// occupies exactly one, so a shorter payload means the call did not return what we expect.
const abiWord = 32

// ABI types reused when packing token call arguments.
var (
	addressTy, _ = abi.NewType("address", "", nil)
	uint256Ty, _ = abi.NewType("uint256", "", nil)
	bytesTy, _   = abi.NewType("bytes", "", nil)
	// dvpExtraTy is the SharedObjects.Dvp{721,1155}ExtraData[] tuple array (key,value,isPublic).
	// It is always passed empty, but the tuple component types are required so the function
	// selector (which includes them) is computed correctly.
	dvpExtraTy, _ = abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "key", Type: "string"},
		{Name: "value", Type: "string"},
		{Name: "isPublic", Type: "bool"},
	})
)

// RaylsTokenService signs and broadcasts mint/burn/teleport transactions to a token contract via the
// caller's custody wallet, selecting the function signature per token standard. It is pure on-chain
// interaction: position reads (balanceOf/ownerOf) and transaction signing, with no business rules.
type RaylsTokenService struct {
	client  *ethclient.Client
	custody core.CustodyService
	chainID *big.Int
}

// NewRaylsTokenService reuses an existing RPC client (e.g. from RaylsAccessManagerService).
func NewRaylsTokenService(client *ethclient.Client, custody core.CustodyService) (*RaylsTokenService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chain ID: %w", err)
	}
	return &RaylsTokenService{client: client, custody: custody, chainID: chainID}, nil
}

// Mint builds the standard-specific mint calldata, signs it, and waits for the receipt.
func (s *RaylsTokenService) Mint(
	ctx context.Context,
	signerAddress, tokenAddress string,
	standard domain.ErcStandard,
	in core.MintInput,
) (string, error) {
	calldata, err := mintCalldata(standard, in)
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

// Burn builds the standard-specific burn calldata, signs it, and waits for the receipt.
func (s *RaylsTokenService) Burn(
	ctx context.Context,
	signerAddress, tokenAddress string,
	standard domain.ErcStandard,
	in core.BurnInput,
) (string, error) {
	calldata, err := burnCalldata(standard, in)
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

// SetPaused calls pause() or unpause() on a stablecoin contract, signed with the caller's wallet.
//
// Both are `onlyPauser` — gated on the contract's own `pauser` ADDRESS, not on an AccessManager
// role like mint/burn. Authorization is therefore a msg.sender equality check, which the caller
// must make with Pauser() below; there is no AM function permission to consult.
func (s *RaylsTokenService) SetPaused(
	ctx context.Context,
	signerAddress, tokenAddress string,
	paused bool,
) (string, error) {
	signature := "unpause()"
	if paused {
		signature = "pause()"
	}
	calldata, err := packCall(signature, abi.Arguments{})
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

// Pauser reads the stablecoin's `pauser` address — the only account pause()/unpause() accept.
func (s *RaylsTokenService) Pauser(ctx context.Context, tokenAddress string) (string, error) {
	out, err := s.callContract(ctx, tokenAddress, "pauser()", abi.Arguments{})
	if err != nil {
		return "", fmt.Errorf("query pauser: %w", err)
	}
	if len(out) < abiWord {
		return "", fmt.Errorf("query pauser: short return (%d bytes)", len(out))
	}
	return common.BytesToAddress(out[abiWord-common.AddressLength : abiWord]).Hex(), nil
}

// IsPaused reads the stablecoin's `paused` flag straight from the contract.
//
// Deliberately a live read rather than the am_managed_contracts mirror: that column tracks the
// AccessManager's ContractPauseUpdated event, which is a DIFFERENT pause (AM-level gating of a
// managed contract) from the stablecoin's own `paused` state variable.
func (s *RaylsTokenService) IsPaused(ctx context.Context, tokenAddress string) (bool, error) {
	out, err := s.callContract(ctx, tokenAddress, "paused()", abi.Arguments{})
	if err != nil {
		return false, fmt.Errorf("query paused: %w", err)
	}
	if len(out) < abiWord {
		return false, fmt.Errorf("query paused: short return (%d bytes)", len(out))
	}
	return out[abiWord-1] != 0, nil
}

// TeleportERC20 builds the ERC20 teleportToPublicChain calldata, signs it with the caller's wallet,
// and waits for the receipt. Pure on-chain — eligibility/registration/balance preflight is enforced
// by the TeleportService before this runs.
func (s *RaylsTokenService) TeleportERC20(
	ctx context.Context,
	signerAddress, tokenAddress, to string,
	amount, destinationChainID *big.Int,
) (string, error) {
	calldata, err := teleportERC20Calldata(to, amount, destinationChainID)
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

func (s *RaylsTokenService) TeleportERC721(
	ctx context.Context,
	signerAddress, tokenAddress, to string,
	tokenID, destinationChainID *big.Int,
) (string, error) {
	calldata, err := teleportERC721Calldata(to, tokenID, destinationChainID)
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

func (s *RaylsTokenService) TeleportERC1155(
	ctx context.Context,
	signerAddress, tokenAddress, to string,
	tokenID, amount, destinationChainID *big.Int,
	data []byte,
) (string, error) {
	calldata, err := teleportERC1155Calldata(to, tokenID, amount, destinationChainID, data)
	if err != nil {
		return "", err
	}
	return s.signAndSend(ctx, signerAddress, tokenAddress, calldata)
}

func mintCalldata(standard domain.ErcStandard, in core.MintInput) ([]byte, error) {
	if !common.IsHexAddress(in.To) {
		return nil, fmt.Errorf("invalid recipient address: %q", in.To)
	}
	to := common.HexToAddress(in.To)
	switch standard {
	case domain.ErcStandardERC20, domain.ErcStandardEnygma, domain.ErcStandardStableCoin:
		// The stablecoin inherits RaylsErc20Handler.mint — identical mint(address,uint256) signature.
		if in.Amount == nil {
			return nil, fmt.Errorf("amount is required")
		}
		return packCall("mint(address,uint256)", abi.Arguments{{Type: addressTy}, {Type: uint256Ty}}, to, in.Amount)
	case domain.ErcStandardERC721:
		if in.TokenID == nil {
			return nil, fmt.Errorf("tokenId is required")
		}
		return packCall("mint(address,uint256)", abi.Arguments{{Type: addressTy}, {Type: uint256Ty}}, to, in.TokenID)
	case domain.ErcStandardERC1155:
		if in.TokenID == nil || in.Amount == nil {
			return nil, fmt.Errorf("id and amount are required")
		}
		return packCall("mint(address,uint256,uint256,bytes)",
			abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: bytesTy}},
			to, in.TokenID, in.Amount, in.Data)
	case domain.ErcStandardZkDvpERC721:
		if in.TokenID == nil {
			return nil, fmt.Errorf("tokenId is required")
		}
		return packCall("mint(address,uint256,(string,string,bool)[])",
			abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: dvpExtraTy}},
			to, in.TokenID, emptyDvpExtra())
	case domain.ErcStandardZkDvpERC1155:
		if in.TokenID == nil || in.Amount == nil {
			return nil, fmt.Errorf("id and amount are required")
		}
		return packCall("mint(address,uint256,uint256,bytes,(string,string,bool)[])",
			abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: bytesTy}, {Type: dvpExtraTy}},
			to, in.TokenID, in.Amount, in.Data, emptyDvpExtra())
	default:
		return nil, fmt.Errorf("unsupported token standard for mint: %d", standard)
	}
}

func burnCalldata(standard domain.ErcStandard, in core.BurnInput) ([]byte, error) {
	switch standard {
	case domain.ErcStandardERC20, domain.ErcStandardEnygma, domain.ErcStandardStableCoin:
		// The stablecoin inherits RaylsErc20Handler.burn — identical burn(address,uint256) signature.
		if !common.IsHexAddress(in.From) {
			return nil, fmt.Errorf("invalid from address: %q", in.From)
		}
		if in.Amount == nil {
			return nil, fmt.Errorf("amount is required")
		}
		return packCall(
			"burn(address,uint256)",
			abi.Arguments{{Type: addressTy}, {Type: uint256Ty}},
			common.HexToAddress(in.From),
			in.Amount,
		)
	case domain.ErcStandardERC721, domain.ErcStandardZkDvpERC721:
		if in.TokenID == nil {
			return nil, fmt.Errorf("tokenId is required")
		}
		return packCall("burn(uint256)", abi.Arguments{{Type: uint256Ty}}, in.TokenID)
	case domain.ErcStandardERC1155, domain.ErcStandardZkDvpERC1155:
		if !common.IsHexAddress(in.From) {
			return nil, fmt.Errorf("invalid from address: %q", in.From)
		}
		if in.TokenID == nil || in.Amount == nil {
			return nil, fmt.Errorf("id and amount are required")
		}
		return packCall(
			"burn(address,uint256,uint256)",
			abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}},
			common.HexToAddress(in.From),
			in.TokenID,
			in.Amount,
		)
	default:
		return nil, fmt.Errorf("unsupported token standard for burn: %d", standard)
	}
}

func teleportERC20Calldata(to string, amount, destChainID *big.Int) ([]byte, error) {
	if destChainID == nil {
		return nil, fmt.Errorf("public chain ID is not configured")
	}
	if !common.IsHexAddress(to) {
		return nil, fmt.Errorf("invalid recipient address: %q", to)
	}
	if amount == nil {
		return nil, fmt.Errorf("amount is required")
	}
	return packCall("teleportToPublicChain(address,uint256,uint256)",
		abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}},
		common.HexToAddress(to), amount, destChainID)
}

func teleportERC721Calldata(to string, tokenID, destChainID *big.Int) ([]byte, error) {
	if destChainID == nil {
		return nil, fmt.Errorf("public chain ID is not configured")
	}
	if !common.IsHexAddress(to) {
		return nil, fmt.Errorf("invalid recipient address: %q", to)
	}
	if tokenID == nil {
		return nil, fmt.Errorf("tokenId is required")
	}
	return packCall("teleportToPublicChain(address,uint256,uint256)",
		abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}},
		common.HexToAddress(to), tokenID, destChainID)
}

func teleportERC1155Calldata(to string, tokenID, amount, destChainID *big.Int, data []byte) ([]byte, error) {
	if destChainID == nil {
		return nil, fmt.Errorf("public chain ID is not configured")
	}
	if !common.IsHexAddress(to) {
		return nil, fmt.Errorf("invalid recipient address: %q", to)
	}
	if tokenID == nil || amount == nil {
		return nil, fmt.Errorf("id and amount are required")
	}
	return packCall("teleportToPublicChain(address,uint256,uint256,uint256,bytes)",
		abi.Arguments{{Type: addressTy}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: uint256Ty}, {Type: bytesTy}},
		common.HexToAddress(to), tokenID, amount, destChainID, data)
}

// packCall prepends the 4-byte selector (keccak256(signature)[:4]) to the ABI-encoded arguments.
func packCall(signature string, args abi.Arguments, values ...interface{}) ([]byte, error) {
	enc, err := args.Pack(values...)
	if err != nil {
		return nil, fmt.Errorf("pack %s: %w", signature, err)
	}
	return append(crypto.Keccak256([]byte(signature))[:4], enc...), nil
}

// emptyDvpExtra returns a correctly-typed empty slice for the Dvp*ExtraData[] argument.
func emptyDvpExtra() interface{} {
	return reflect.MakeSlice(reflect.SliceOf(dvpExtraTy.Elem.TupleType), 0, 0).Interface()
}

// signAndSend builds an unsigned tx to `to` with the calldata, delegates signing to custody,
// and polls for the receipt.
func (s *RaylsTokenService) signAndSend(
	ctx context.Context,
	signerAddress, to string,
	calldata []byte,
) (string, error) {
	if !common.IsHexAddress(signerAddress) {
		return "", fmt.Errorf("invalid signer address: %q", signerAddress)
	}
	if !common.IsHexAddress(to) {
		return "", fmt.Errorf("invalid token address: %q", to)
	}
	from := common.HexToAddress(signerAddress)
	toAddr := common.HexToAddress(to)

	nonce, err := s.client.PendingNonceAt(ctx, from)
	if err != nil {
		return "", fmt.Errorf("get pending nonce: %w", err)
	}
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("suggest gas price: %w", err)
	}
	gasLimit, err := s.client.EstimateGas(ctx, ethereum.CallMsg{From: from, To: &toAddr, Data: calldata})
	if err != nil {
		return "", fmt.Errorf("estimate gas: %w", err)
	}

	tx := types.NewTransaction(nonce, toAddr, big.NewInt(0), gasLimit, gasPrice, calldata)
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("encode transaction: %w", err)
	}

	txHash, err := s.custody.SignAndTransact(ctx, txBytes, signerAddress, s.chainID.String())
	if err != nil {
		return "", fmt.Errorf("sign and transact: %w", err)
	}

	// Bound the receipt wait so a dropped or replaced transaction can't hang the request forever.
	waitCtx, cancel := context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	hash := common.HexToHash(txHash)
	for {
		receipt, err := s.client.TransactionReceipt(waitCtx, hash)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-waitCtx.Done():
				return "", fmt.Errorf("timed out waiting for receipt (tx %s): %w", txHash, waitCtx.Err())
			case <-time.After(2 * time.Second):
				continue
			}
		}
		if err != nil {
			return "", fmt.Errorf("get receipt: %w", err)
		}
		if receipt.Status == 0 {
			return "", fmt.Errorf("%w (tx %s)", core.ErrTxReverted, txHash)
		}
		return txHash, nil
	}
}

// ERC20Balance reads the ERC20 balanceOf(address) for account. Raw on-chain read — the caller
// decides whether the balance is sufficient.
func (s *RaylsTokenService) ERC20Balance(ctx context.Context, tokenAddress, account string) (*big.Int, error) {
	if !common.IsHexAddress(account) {
		return nil, fmt.Errorf("invalid account address: %q", account)
	}
	balance, err := s.readUint(
		ctx,
		tokenAddress,
		"balanceOf(address)",
		abi.Arguments{{Type: addressTy}},
		common.HexToAddress(account),
	)
	if err != nil {
		return nil, fmt.Errorf("query ERC20 balance: %w", err)
	}
	return balance, nil
}

// ERC721Owner reads ownerOf(uint256) and returns the owner address as a hex string. Raw on-chain
// read — the caller decides whether the owner matches.
func (s *RaylsTokenService) ERC721Owner(ctx context.Context, tokenAddress string, tokenID *big.Int) (string, error) {
	if tokenID == nil {
		return "", fmt.Errorf("tokenId is required")
	}
	out, err := s.callContract(ctx, tokenAddress, "ownerOf(uint256)", abi.Arguments{{Type: uint256Ty}}, tokenID)
	if err != nil {
		return "", fmt.Errorf("query ERC721 owner: %w", err)
	}
	return common.BytesToAddress(out).Hex(), nil
}

// ERC1155Balance reads balanceOf(address,uint256) for account and tokenID. Raw on-chain read — the
// caller decides whether the balance is sufficient.
func (s *RaylsTokenService) ERC1155Balance(
	ctx context.Context,
	tokenAddress, account string,
	tokenID *big.Int,
) (*big.Int, error) {
	if !common.IsHexAddress(account) {
		return nil, fmt.Errorf("invalid account address: %q", account)
	}
	if tokenID == nil {
		return nil, fmt.Errorf("tokenId is required")
	}
	balance, err := s.readUint(
		ctx,
		tokenAddress,
		"balanceOf(address,uint256)",
		abi.Arguments{{Type: addressTy}, {Type: uint256Ty}},
		common.HexToAddress(account),
		tokenID,
	)
	if err != nil {
		return nil, fmt.Errorf("query ERC1155 balance: %w", err)
	}
	return balance, nil
}

func (s *RaylsTokenService) readUint(
	ctx context.Context,
	tokenAddress, signature string,
	args abi.Arguments,
	values ...interface{},
) (*big.Int, error) {
	out, err := s.callContract(ctx, tokenAddress, signature, args, values...)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(out), nil
}

func (s *RaylsTokenService) callContract(
	ctx context.Context,
	tokenAddress, signature string,
	args abi.Arguments,
	values ...interface{},
) ([]byte, error) {
	if !common.IsHexAddress(tokenAddress) {
		return nil, fmt.Errorf("invalid token address: %q", tokenAddress)
	}
	calldata, err := packCall(signature, args, values...)
	if err != nil {
		return nil, err
	}
	to := common.HexToAddress(tokenAddress)
	out, err := s.client.CallContract(ctx, ethereum.CallMsg{To: &to, Data: calldata}, nil)
	if err != nil {
		return nil, fmt.Errorf("call %s: %w", signature, err)
	}
	return out, nil
}
