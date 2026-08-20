// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNMessageDispatcherV1

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// BatchMessage is an auto generated low-level Go binding around an user-defined struct.
type BatchMessage struct {
	ToChainId *big.Int
	To        common.Address
	Data      RaylsNodeMessage
	MessageId [32]byte
}

// RaylsNodeBridgedTransferMetadata is an auto generated low-level Go binding around an user-defined struct.
type RaylsNodeBridgedTransferMetadata struct {
	AssetType    uint8
	Id           *big.Int
	From         common.Address
	To           common.Address
	TokenAddress common.Address
	Amount       *big.Int
}

// RaylsNodeMessage is an auto generated low-level Go binding around an user-defined struct.
type RaylsNodeMessage struct {
	MessageMetadata RaylsNodeMessageMetadata
	Payload         []byte
}

// RaylsNodeMessageMetadata is an auto generated low-level Go binding around an user-defined struct.
type RaylsNodeMessageMetadata struct {
	Nonce               *big.Int
	NewResourceMetadata RaylsNodeNewResourceMetadata
	RevertPayloadData   []byte
	TransferMetadata    RaylsNodeBridgedTransferMetadata
}

// RaylsNodeNewResourceMetadata is an auto generated low-level Go binding around an user-defined struct.
type RaylsNodeNewResourceMetadata struct {
	ResourceDeployType uint8
	Bytecode           []byte
	FactoryTemplate    uint8
	InitializerParams  []byte
}

// RNMessageDispatcherV1MetaData contains all meta data concerning the RNMessageDispatcherV1 contract.
var RNMessageDispatcherV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizedEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"dispatchMessage\",\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessage\",\"components\":[{\"name\":\"messageMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessageMetadata\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newResourceMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeNewResourceMetadata\",\"components\":[{\"name\":\"resourceDeployType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeResourceDeployType\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"factoryTemplate\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"initializerParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"revertPayloadData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dispatchMessageBatch\",\"inputs\":[{\"name\":\"batchId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structBatchMessage[]\",\"components\":[{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessage\",\"components\":[{\"name\":\"messageMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessageMetadata\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newResourceMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeNewResourceMetadata\",\"components\":[{\"name\":\"resourceDeployType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeResourceDeployType\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"factoryTemplate\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"initializerParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"revertPayloadData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAuthorizedEndpoint\",\"inputs\":[{\"name\":\"_authorizedEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageBatchDispatched\",\"inputs\":[{\"name\":\"batchId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structBatchMessage[]\",\"components\":[{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessage\",\"components\":[{\"name\":\"messageMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessageMetadata\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newResourceMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeNewResourceMetadata\",\"components\":[{\"name\":\"resourceDeployType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeResourceDeployType\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"factoryTemplate\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"initializerParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"revertPayloadData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageDispatched\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRaylsNodeMessage\",\"components\":[{\"name\":\"messageMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessageMetadata\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newResourceMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeNewResourceMetadata\",\"components\":[{\"name\":\"resourceDeployType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeResourceDeployType\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"factoryTemplate\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"initializerParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"revertPayloadData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNMessageDispatcherV1__InvalidEndpointAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNMessageDispatcherV1__UnauthorizedEndpoint\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "RNMessageDispatcherV1",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516114c56100fd6000396000818161053c01528181610565015261069e01526114c56000f3fe60806040526004361061008c5760003560e01c806318499c12146100915780632f8d242f146100c4578063316a3f14146100f15780634f1ef2861461011157806352d1902d1461012657806354fd4d501461013b5780637a68652114610170578063a0a8e46014610190578063ad3cb1cc146101a4578063bf7e214f146101d5578063c4d66de8146101ea575b600080fd5b34801561009d57600080fd5b506100b16100ac366004610eb5565b61020a565b6040519081526020015b60405180910390f35b3480156100d057600080fd5b506000546100e4906001600160a01b031681565b6040516100bb9190610feb565b3480156100fd57600080fd5b506100b161010c366004610fff565b610286565b61012461011f366004611070565b61033c565b005b34801561013257600080fd5b506100b161035b565b34801561014757600080fd5b506040805180820190915260038152620312e360ec1b60208201525b6040516100bb919061110d565b34801561017c57600080fd5b5061012461018b366004611120565b610378565b34801561019c57600080fd5b5060016100b1565b3480156101b057600080fd5b50610163604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156101e157600080fd5b506100e46103d7565b3480156101f657600080fd5b50610124610205366004611120565b6103f0565b600080546001600160a01b031633146102415733604051638af6b95360e01b81526004016102389190610feb565b60405180910390fd5b7f4d296b8802a280920b68424698550abb61dfa35190fcd5fd796e759fc4e2622484848460405161027493929190611288565b60405180910390a150825b9392505050565b600080546001600160a01b031633146102b45733604051638af6b95360e01b81526004016102389190610feb565b60006102eb87878787876000015160000151886040516020016102d79190611330565b6040516020818303038152906040526104f2565b905084866001600160a01b0316827f80d12e982e0c08085735a0ae121a347f45fb2ba93a30e1c3db6a3a39df296b0c878760405161032a929190611343565b60405180910390a49695505050505050565b610344610531565b61034d826105c1565b61035782826105da565b5050565b6000610365610693565b5060008051602061147083398151915290565b61038e336000356001600160e01b0319166106dc565b6001600160a01b0381166103b557604051634d2166a360e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006103e1610827565b546001600160a01b0316919050565b60006103fa610889565b805490915060ff600160401b82041615906001600160401b03166000811580156104215750825b90506000826001600160401b0316600114801561043d5750303b155b90508115801561044b575080155b156104695760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561049357845460ff60401b1916600160401b1785555b61049b6108b4565b6104a4866108bc565b83156104ea57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b600086868686868660405160200161050f9695949392919061136f565b6040516020818303038152906040528051906020012090509695505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806105a157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105956108fd565b6001600160a01b031614155b156105bf5760405163703e46dd60e11b815260040160405180910390fd5b565b6105d7336000356001600160e01b0319166106dc565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610634575060408051601f3d908101601f19168201909252610631918101906113bb565b60015b6106535781604051634c9c8ce360e01b81526004016102389190610feb565b600080516020611470833981519152811461068457604051632a87526960e21b815260048101829052602401610238565b61068e8383610913565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105bf5760405163703e46dd60e11b815260040160405180910390fd5b60006106e6610827565b80549091506001600160a01b031680610715576000604051638944034760e01b81526004016102389190610feb565b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa158015610779573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079d91906113e4565b9250925092508261081e5780156107c75760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff8216156108035760405163a426878960e01b81526001600160a01b038816600482015263ffffffff83166024820152604401610238565b86604051632ecd3d0360e21b81526004016102389190610feb565b50505050505050565b60008060ff1961085860017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f35611432565b60405160200161086a91815260200190565b60408051601f1981840301815291905280516020909101201692915050565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b6105bf610969565b60006108c6610827565b80549091506001600160a01b0316156108f45781604051638944034760e01b81526004016102389190610feb565b6103578261098e565b60006000805160206114708339815191526103e1565b61091c82610a1e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156109615761068e8282610a7a565b610357610af0565b610971610b0f565b6105bf57604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b0381166109b75780604051638944034760e01b81526004016102389190610feb565b60006109c1610827565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b806001600160a01b03163b600003610a4b5780604051634c9c8ce360e01b81526004016102389190610feb565b60008051602061147083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051610a979190611453565b600060405180830381855af49150503d8060008114610ad2576040519150601f19603f3d011682016040523d82523d6000602084013e610ad7565b606091505b5091509150610ae7858383610b29565b95945050505050565b34156105bf5760405163b398979f60e01b815260040160405180910390fd5b6000610b19610889565b54600160401b900460ff16919050565b606082610b3e57610b3982610b7c565b61027f565b8151158015610b5557506001600160a01b0384163b155b15610b755783604051639996b31560e01b81526004016102389190610feb565b508061027f565b805115610b8b57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b80356001600160a01b0381168114610bbb57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610bf857610bf8610bc0565b60405290565b604051608081016001600160401b0381118282101715610bf857610bf8610bc0565b604051601f8201601f191681016001600160401b0381118282101715610c4857610c48610bc0565b604052919050565b600082601f830112610c6157600080fd5b81356001600160401b03811115610c7a57610c7a610bc0565b610c8d601f8201601f1916602001610c20565b818152846020838601011115610ca257600080fd5b816020850160208301376000918101602001919091529392505050565b803560078110610bbb57600080fd5b600060c08284031215610ce057600080fd5b60405160c081018181106001600160401b0382111715610d0257610d02610bc0565b604052905080610d1183610cbf565b815260208301356020820152610d2960408401610ba4565b6040820152610d3a60608401610ba4565b6060820152610d4b60808401610ba4565b608082015260a083013560a08201525092915050565b600060408284031215610d7357600080fd5b610d7b610bd6565b905081356001600160401b0380821115610d9457600080fd5b908301906101208286031215610da957600080fd5b610db1610bfe565b8235815260208084013583811115610dc857600080fd5b840160808189031215610dda57600080fd5b610de2610bfe565b813560028110610df157600080fd5b81528183013585811115610e0457600080fd5b610e108a828501610c50565b8483015250610e2160408301610cbf565b6040820152606082013585811115610e3857600080fd5b610e448a828501610c50565b6060830152508383015250604084013583811115610e6157600080fd5b610e6d88828701610c50565b604084015250610e808760608601610cce565b606083015281855280860135935082841115610e9b57600080fd5b610ea787858801610c50565b818601525050505092915050565b600080600060608486031215610eca57600080fd5b833592506020610edb818601610ba4565b925060408501356001600160401b0380821115610ef757600080fd5b818701915087601f830112610f0b57600080fd5b813581811115610f1d57610f1d610bc0565b8060051b610f2c858201610c20565b918252838101850191858101908b841115610f4657600080fd5b86860192505b83831015610fda57823585811115610f6357600080fd5b86016080818e03601f19011215610f7a5760008081fd5b610f82610bfe565b888201358152610f9460408301610ba4565b89820152606082013587811115610fab5760008081fd5b610fb98f8b83860101610d61565b60408301525060809190910135606082015282529186019190860190610f4c565b809750505050505050509250925092565b6001600160a01b0391909116815260200190565b600080600080600060a0868803121561101757600080fd5b8535945061102760208701610ba4565b93506040860135925061103c60608701610ba4565b915060808601356001600160401b0381111561105757600080fd5b61106388828901610d61565b9150509295509295909350565b6000806040838503121561108357600080fd5b61108c83610ba4565b915060208301356001600160401b038111156110a757600080fd5b6110b385828601610c50565b9150509250929050565b60005b838110156110d85781810151838201526020016110c0565b50506000910152565b600081518084526110f98160208601602086016110bd565b601f01601f19169290920160200192915050565b60208152600061027f60208301846110e1565b60006020828403121561113257600080fd5b61027f82610ba4565b634e487b7160e01b600052602160045260246000fd5b600781106111615761116161113b565b9052565b611170828251611151565b60208101516020830152604081015160018060a01b038082166040850152806060840151166060850152806080840151166080850152505060a081015160a08301525050565b600081516040845280516040850152602081015161012060608601528051600281106111e4576111e461113b565b610160860152602081015160806101808701526112056101e08701826110e1565b9050604082015161121a6101a0880182611151565b506060820151915061015f19868203016101c087015261123a81836110e1565b9150506040820151603f1986830301608087015261125882826110e1565b9150506060820151915061126f60a0860183611165565b602084015191508481036020860152610ae781836110e1565b60006060808301868452602060018060a01b03808816828701526040606060408801528388518086526080955060808901915060808160051b8a0101858b0160005b8381101561131d57607f198c840301855281518051845287898201511689850152868101518a888601526113008b8601826111b6565b918c0151948c0194909452948801949250908701906001016112ca565b50909d9c50505050505050505050505050565b60208152600061027f60208301846111b6565b6001600160a01b0383168152604060208201819052600090611367908301846111b6565b949350505050565b8681526001600160a01b03868116602083015260408201869052841660608201526080810183905260c060a082018190526000906113af908301846110e1565b98975050505050505050565b6000602082840312156113cd57600080fd5b5051919050565b80518015158114610bbb57600080fd5b6000806000606084860312156113f957600080fd5b611402846113d4565b9250602084015163ffffffff8116811461141b57600080fd5b9150611429604085016113d4565b90509250925092565b818103818111156108ae57634e487b7160e01b600052601160045260246000fd5b600082516114658184602087016110bd565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212205426a49914cb062c2d037e1d891b1f950aaaae2f26b0ac37db7eb8fa1a54c2ea64736f6c63430008180033",
}

// RNMessageDispatcherV1 is an auto generated Go binding around an Ethereum contract.
type RNMessageDispatcherV1 struct {
	abi abi.ABI
}

// NewRNMessageDispatcherV1 creates a new instance of RNMessageDispatcherV1.
func NewRNMessageDispatcherV1() *RNMessageDispatcherV1 {
	parsed, err := RNMessageDispatcherV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNMessageDispatcherV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNMessageDispatcherV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackAuthority is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackAuthority() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("authority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackAuthorizedEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f8d242f.
//
// Solidity: function authorizedEndpoint() view returns(address)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackAuthorizedEndpoint() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("authorizedEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthorizedEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f8d242f.
//
// Solidity: function authorizedEndpoint() view returns(address)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackAuthorizedEndpoint(data []byte) (common.Address, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("authorizedEndpoint", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackContractVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackContractVersion() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("contractVersion", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDispatchMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x316a3f14.
//
// Solidity: function dispatchMessage(uint256 fromChainId, address from, uint256 toChainId, address to, ((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes) data) returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackDispatchMessage(fromChainId *big.Int, from common.Address, toChainId *big.Int, to common.Address, data RaylsNodeMessage) []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("dispatchMessage", fromChainId, from, toChainId, to, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDispatchMessage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x316a3f14.
//
// Solidity: function dispatchMessage(uint256 fromChainId, address from, uint256 toChainId, address to, ((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes) data) returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackDispatchMessage(data []byte) ([32]byte, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("dispatchMessage", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDispatchMessageBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18499c12.
//
// Solidity: function dispatchMessageBatch(bytes32 batchId, address from, (uint256,address,((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes),bytes32)[] messages) returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackDispatchMessageBatch(batchId [32]byte, from common.Address, messages []BatchMessage) []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("dispatchMessageBatch", batchId, from, messages)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDispatchMessageBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18499c12.
//
// Solidity: function dispatchMessageBatch(bytes32 batchId, address from, (uint256,address,((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes),bytes32)[] messages) returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackDispatchMessageBatch(data []byte) ([32]byte, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("dispatchMessageBatch", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address authority_) returns()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackInitialize(authority common.Address) []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("initialize", authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackProxiableUUID() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSetAuthorizedEndpoint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a686521.
//
// Solidity: function setAuthorizedEndpoint(address _authorizedEndpoint) returns()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackSetAuthorizedEndpoint(authorizedEndpoint common.Address) []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("setAuthorizedEndpoint", authorizedEndpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) PackVersion() []byte {
	enc, err := rNMessageDispatcherV1.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackVersion(data []byte) (string, error) {
	out, err := rNMessageDispatcherV1.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// RNMessageDispatcherV1AuthorityUpdated represents a AuthorityUpdated event raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNMessageDispatcherV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RNMessageDispatcherV1AuthorityUpdated) ContractEventName() string {
	return RNMessageDispatcherV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RNMessageDispatcherV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != rNMessageDispatcherV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageDispatcherV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageDispatcherV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RNMessageDispatcherV1Initialized represents a Initialized event raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNMessageDispatcherV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNMessageDispatcherV1Initialized) ContractEventName() string {
	return RNMessageDispatcherV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackInitializedEvent(log *types.Log) (*RNMessageDispatcherV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNMessageDispatcherV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageDispatcherV1Initialized)
	if len(log.Data) > 0 {
		if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageDispatcherV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RNMessageDispatcherV1MessageBatchDispatched represents a MessageBatchDispatched event raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1MessageBatchDispatched struct {
	BatchId  [32]byte
	From     common.Address
	Messages []BatchMessage
	Raw      *types.Log // Blockchain specific contextual infos
}

const RNMessageDispatcherV1MessageBatchDispatchedEventName = "MessageBatchDispatched"

// ContractEventName returns the user-defined event name.
func (RNMessageDispatcherV1MessageBatchDispatched) ContractEventName() string {
	return RNMessageDispatcherV1MessageBatchDispatchedEventName
}

// UnpackMessageBatchDispatchedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessageBatchDispatched(bytes32 batchId, address from, (uint256,address,((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes),bytes32)[] messages)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackMessageBatchDispatchedEvent(log *types.Log) (*RNMessageDispatcherV1MessageBatchDispatched, error) {
	event := "MessageBatchDispatched"
	if log.Topics[0] != rNMessageDispatcherV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageDispatcherV1MessageBatchDispatched)
	if len(log.Data) > 0 {
		if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageDispatcherV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RNMessageDispatcherV1MessageDispatched represents a MessageDispatched event raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1MessageDispatched struct {
	MessageId [32]byte
	From      common.Address
	ToChainId *big.Int
	To        common.Address
	Data      RaylsNodeMessage
	Raw       *types.Log // Blockchain specific contextual infos
}

const RNMessageDispatcherV1MessageDispatchedEventName = "MessageDispatched"

// ContractEventName returns the user-defined event name.
func (RNMessageDispatcherV1MessageDispatched) ContractEventName() string {
	return RNMessageDispatcherV1MessageDispatchedEventName
}

// UnpackMessageDispatchedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessageDispatched(bytes32 indexed messageId, address indexed from, uint256 indexed toChainId, address to, ((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes) data)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackMessageDispatchedEvent(log *types.Log) (*RNMessageDispatcherV1MessageDispatched, error) {
	event := "MessageDispatched"
	if log.Topics[0] != rNMessageDispatcherV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageDispatcherV1MessageDispatched)
	if len(log.Data) > 0 {
		if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageDispatcherV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// RNMessageDispatcherV1Upgraded represents a Upgraded event raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNMessageDispatcherV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RNMessageDispatcherV1Upgraded) ContractEventName() string {
	return RNMessageDispatcherV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackUpgradedEvent(log *types.Log) (*RNMessageDispatcherV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != rNMessageDispatcherV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageDispatcherV1Upgraded)
	if len(log.Data) > 0 {
		if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageDispatcherV1.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RNMessageDispatcherV1InvalidEndpointAddress"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRNMessageDispatcherV1InvalidEndpointAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RNMessageDispatcherV1UnauthorizedEndpoint"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRNMessageDispatcherV1UnauthorizedEndpointError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageDispatcherV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return rNMessageDispatcherV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNMessageDispatcherV1AddressEmptyCode represents a AddressEmptyCode error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RNMessageDispatcherV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackAddressEmptyCodeError(raw []byte) (*RNMessageDispatcherV1AddressEmptyCode, error) {
	out := new(RNMessageDispatcherV1AddressEmptyCode)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RNMessageDispatcherV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RNMessageDispatcherV1ERC1967InvalidImplementation, error) {
	out := new(RNMessageDispatcherV1ERC1967InvalidImplementation)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RNMessageDispatcherV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackERC1967NonPayableError(raw []byte) (*RNMessageDispatcherV1ERC1967NonPayable, error) {
	out := new(RNMessageDispatcherV1ERC1967NonPayable)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1FailedCall represents a FailedCall error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RNMessageDispatcherV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackFailedCallError(raw []byte) (*RNMessageDispatcherV1FailedCall, error) {
	out := new(RNMessageDispatcherV1FailedCall)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1InvalidInitialization represents a InvalidInitialization error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNMessageDispatcherV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackInvalidInitializationError(raw []byte) (*RNMessageDispatcherV1InvalidInitialization, error) {
	out := new(RNMessageDispatcherV1InvalidInitialization)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1NotInitializing represents a NotInitializing error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNMessageDispatcherV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackNotInitializingError(raw []byte) (*RNMessageDispatcherV1NotInitializing, error) {
	out := new(RNMessageDispatcherV1NotInitializing)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RNMessageDispatcherV1InvalidEndpointAddress represents a RNMessageDispatcherV1__InvalidEndpointAddress error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RNMessageDispatcherV1InvalidEndpointAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageDispatcherV1__InvalidEndpointAddress()
func RNMessageDispatcherV1RNMessageDispatcherV1InvalidEndpointAddressErrorID() common.Hash {
	return common.HexToHash("0x9a42cd4608441cdc7e0e97bd176bfa586eeb78bc53737c2e5509c7a5b32cf2a3")
}

// UnpackRNMessageDispatcherV1InvalidEndpointAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageDispatcherV1__InvalidEndpointAddress()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRNMessageDispatcherV1InvalidEndpointAddressError(raw []byte) (*RNMessageDispatcherV1RNMessageDispatcherV1InvalidEndpointAddress, error) {
	out := new(RNMessageDispatcherV1RNMessageDispatcherV1InvalidEndpointAddress)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RNMessageDispatcherV1InvalidEndpointAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RNMessageDispatcherV1UnauthorizedEndpoint represents a RNMessageDispatcherV1__UnauthorizedEndpoint error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RNMessageDispatcherV1UnauthorizedEndpoint struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageDispatcherV1__UnauthorizedEndpoint(address caller)
func RNMessageDispatcherV1RNMessageDispatcherV1UnauthorizedEndpointErrorID() common.Hash {
	return common.HexToHash("0x8af6b9537865d1422545bb14d2a37f68d8f287dab30a7df1f1cb17cd0891eb5c")
}

// UnpackRNMessageDispatcherV1UnauthorizedEndpointError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageDispatcherV1__UnauthorizedEndpoint(address caller)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRNMessageDispatcherV1UnauthorizedEndpointError(raw []byte) (*RNMessageDispatcherV1RNMessageDispatcherV1UnauthorizedEndpoint, error) {
	out := new(RNMessageDispatcherV1RNMessageDispatcherV1UnauthorizedEndpoint)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RNMessageDispatcherV1UnauthorizedEndpoint", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RNMessageDispatcherV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RNMessageDispatcherV1RaylsAccessManagedContractPaused, error) {
	out := new(RNMessageDispatcherV1RaylsAccessManagedContractPaused)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RNMessageDispatcherV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RNMessageDispatcherV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RNMessageDispatcherV1RaylsAccessManagedInvalidAuthority)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RNMessageDispatcherV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RNMessageDispatcherV1RaylsAccessManagedMustSchedule, error) {
	out := new(RNMessageDispatcherV1RaylsAccessManagedMustSchedule)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RNMessageDispatcherV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RNMessageDispatcherV1RaylsAccessManagedUnauthorized, error) {
	out := new(RNMessageDispatcherV1RaylsAccessManagedUnauthorized)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RNMessageDispatcherV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RNMessageDispatcherV1UUPSUnauthorizedCallContext, error) {
	out := new(RNMessageDispatcherV1UUPSUnauthorizedCallContext)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageDispatcherV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RNMessageDispatcherV1 contract.
type RNMessageDispatcherV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RNMessageDispatcherV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (rNMessageDispatcherV1 *RNMessageDispatcherV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RNMessageDispatcherV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RNMessageDispatcherV1UUPSUnsupportedProxiableUUID)
	if err := rNMessageDispatcherV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
