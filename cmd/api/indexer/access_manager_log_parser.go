package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/raylsnetwork/rayls-privacy-ops-api/contracts/RaylsAccessManagerV1"
)

// ContractLog is a decoded on-chain event ready for publishing to NATS.
type ContractLog struct {
	ContractName    string          `json:"contract_name"`
	ContractAddress string          `json:"contract_address"`
	EventName       string          `json:"event_name"`
	RawEventData    json.RawMessage `json:"raw_event_data"`
	BlockNumber     uint64          `json:"block_number"`
	TransactionHash string          `json:"transaction_hash"`
	LogIndex        uint64          `json:"log_index"`
	BlockTimestamp  time.Time       `json:"block_timestamp"`
}

type eventParser func(log ethtypes.Log) (json.RawMessage, error)

// AccessManagerLogParser decodes logs from the RaylsAccessManagerV1 contract.
type AccessManagerLogParser struct {
	client   *ethclient.Client
	contract *RaylsAccessManagerV1.RaylsAccessManagerV1
	addr     common.Address
	parsers  map[common.Hash]namedParser
}

type namedParser struct {
	name   string
	parser eventParser
}

func NewAccessManagerLogParser(
	client *ethclient.Client,
	contract *RaylsAccessManagerV1.RaylsAccessManagerV1,
	addr common.Address,
) *AccessManagerLogParser {
	p := &AccessManagerLogParser{
		client:   client,
		contract: contract,
		addr:     addr,
	}
	p.parsers = p.buildDispatchTable()
	return p
}

func (p *AccessManagerLogParser) buildDispatchTable() map[common.Hash]namedParser {
	c := p.contract
	return map[common.Hash]namedParser{
		common.HexToHash("0x132ed6dc8f3ad0a2622797dd549e9ce93930e97a46ef2d6eb22d869ffdbbb16a"): {
			name: "RoleGranted",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleGrantedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x4ee4aaa38790a48e6b8f5ce86c7caaefda507fe314fbfc735fc80b9199cf78dd"): {
			name: "RoleRevoked",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleRevokedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0xf6c6045d1b47f46b319a8eeab190df3ca82b778042b5821afddeeb79cbb876ae"): {
			name: "RoleRegistered",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleRegisteredEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x0934b892db1e7f222186cc5d99854056a432751edfdc3958fa8ac202189b0190"): {
			name: "RoleLabelSet",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleLabelSetEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x1fd6dd7631312dfac2205b52913f99de03b4d7e381d5d27d3dbfe0713e6e6340"): {
			name: "RoleAdminChanged",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleAdminChangedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x7a8059630b897b5de4c08ade69f8b90c3ead1f8596d62d10b6c4d14a0afb4ae2"): {
			name: "RoleGuardianChanged",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleGuardianChangedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x6f8c9cc9fbd18ffcca61bdcfb39fd6b1221ef57d4bddae3b7b1cd29d0182ec4d"): {
			name: "RoleGrantDelayChanged",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackRoleGrantDelayChangedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x858dcc343ec01ddecf10dcc66e22e9f94aba129db1c5b9f7ef55e9bbeb8ed1ed"): {
			name: "FunctionAllowedRoleAdded",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackFunctionAllowedRoleAddedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0xbcb61a373603c95a0bbade2811911058483e005aeb9018e41be1974cef3052be"): {
			name: "FunctionAllowedRoleRemoved",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackFunctionAllowedRoleRemovedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x65d966c6719248a487c5aeec39d4843bd07fee0bda699752325c1dc8a24e7bcb"): {
			name: "ContractPauseUpdated",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackContractPauseUpdatedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x07c54e4e8f451da0b5da0982df73f3bd4e908907e69fe590a0bf6512cd02b5f7"): {
			name: "ManagedContractRegistered",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackManagedContractRegisteredEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x19ba4f5a77fda5170703b28869d1cbd8dbe901087628a8427b329c44d6698408"): {
			name: "OperationScheduled",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackOperationScheduledEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x001d1d6b8d664ddc069841e54d003c4a632a7392aa33ef7f83806380c2bc6873"): {
			name: "OperationExecuted",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackOperationExecutedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x7a168f30d6a7a3b7aa2e211c067c04bde31169a0f3171a98084508204bc56094"): {
			name: "OperationCanceled",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackOperationCanceledEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x8edb51102c468e6b1e099e8298c7bd09b83314acb26db8cf8457d6bb57804b96"): {
			name: "ContractScopedRoleGranted",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackContractScopedRoleGrantedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
		common.HexToHash("0x69a55e61ec515c001bf41d27de9fa01997a08631aee812c6cf4b0ede4ff33a92"): {
			name: "ContractScopedRoleRevoked",
			parser: func(l ethtypes.Log) (json.RawMessage, error) {
				e, err := c.UnpackContractScopedRoleRevokedEvent(&l)
				if err != nil {
					return nil, err
				}
				return marshalEvent(e)
			},
		},
	}
}

// ParseRange fetches and decodes all AccessManager logs in [fromBlock, toBlock].
func (p *AccessManagerLogParser) ParseRange(ctx context.Context, fromBlock, toBlock uint64) ([]ContractLog, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{p.addr},
	}

	rawLogs, err := p.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("filter logs [%d, %d]: %w", fromBlock, toBlock, err)
	}

	// Cache block timestamps to avoid repeated RPC calls for same block.
	tsCache := make(map[uint64]time.Time)

	var out []ContractLog
	for _, l := range rawLogs {
		if len(l.Topics) == 0 {
			continue
		}
		np, ok := p.parsers[l.Topics[0]]
		if !ok {
			continue
		}

		data, err := np.parser(l)
		if err != nil {
			// Log decode errors are non-fatal — skip the event.
			continue
		}

		ts, err := p.blockTimestamp(ctx, l.BlockNumber, tsCache)
		if err != nil {
			ts = time.Time{}
		}

		out = append(out, ContractLog{
			ContractName:    "RaylsAccessManagerV1",
			ContractAddress: l.Address.Hex(),
			EventName:       np.name,
			RawEventData:    data,
			BlockNumber:     l.BlockNumber,
			TransactionHash: l.TxHash.Hex(),
			LogIndex:        uint64(l.Index),
			BlockTimestamp:  ts,
		})
	}
	return out, nil
}

func (p *AccessManagerLogParser) blockTimestamp(
	ctx context.Context,
	blockNum uint64,
	cache map[uint64]time.Time,
) (time.Time, error) {
	if ts, ok := cache[blockNum]; ok {
		return ts, nil
	}
	block, err := p.client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return time.Time{}, err
	}
	ts := time.Unix(int64(block.Time()), 0).UTC()
	cache[blockNum] = ts
	return ts, nil
}

// marshalEvent marshals the decoded event struct, stripping the Raw field which
// contains non-serialisable types (e.g. Topics []common.Hash).
func marshalEvent(v any) (json.RawMessage, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	// Remove the Raw field from the JSON output — it holds the original log.
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return b, nil
	}
	delete(m, "Raw")
	return json.Marshal(m)
}
