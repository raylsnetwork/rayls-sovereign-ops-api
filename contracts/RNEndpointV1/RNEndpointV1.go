// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNEndpointV1

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

// RNEndpointV1MetaData contains all meta data concerning the RNEndpointV1 contract.
var RNEndpointV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureContracts\",\"inputs\":[{\"name\":\"_messageExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_userGovernance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_messageDispatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"currentChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageDispatcherAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenRegistryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserGovernanceAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_publicChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messageDispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMessageDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageExecutor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractRNMessageExecutorV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receivePayload\",\"inputs\":[{\"name\":\"_srcChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_srcAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dstAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_raylsMessage\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessage\",\"components\":[{\"name\":\"messageMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeMessageMetadata\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newResourceMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeNewResourceMetadata\",\"components\":[{\"name\":\"resourceDeployType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeResourceDeployType\"},{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"factoryTemplate\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"initializerParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"revertPayloadData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"_dstChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendToAddress\",\"inputs\":[{\"name\":\"_dstChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_privateChainAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_revertDataPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferMetadata\",\"type\":\"tuple\",\"internalType\":\"structRaylsNodeBridgedTransferMetadata\",\"components\":[{\"name\":\"assetType\",\"type\":\"uint8\",\"internalType\":\"enumRaylsNodeBridgeableERC\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userGovernance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUserGovernance\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__CalledByUnauthorizedAddress\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNEndpointV1__InvalidMessageDispatcherAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__InvalidMessageExecutorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__InvalidTokenRegistryAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__InvalidUserGovernanceAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__NoPublicAddressMapping\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNEndpointV1__SourceAndDestinationChainsSame\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNEndpointV1__TokenNotActive\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNEndpointV1__TokenNotFound\",\"inputs\":[{\"name\":\"privateAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNEndpointV1__TokenUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "RNEndpointV1",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051611eb06100fd60003960008181610e5501528181610e7e0152610fb70152611eb06000f3fe6080604052600436106101205760003560e01c80639d23c4c7116100a65780639d23c4c714610283578063a0a8e460146102a3578063a6ab36f2146102b7578063ad3cb1cc146102d7578063ad552e3214610308578063affed0e014610328578063b8b871381461033e578063bf7e214f1461035c578063caed989714610371578063d087d2881461038f578063e0d9d846146103a4578063ea77c671146103c457600080fd5b8063150b375f146101255780633408e4701461015857806337ffde051461016d57806347219ccc1461018f57806348628eed146101a55780634f1ef286146101d257806352d1902d146101e557806354fd4d50146101fa5780635d31f35f1461022f5780636431087a1461024f5780636cbadbfa1461026d575b600080fd5b34801561013157600080fd5b5061014561014036600461140b565b6103e4565b6040519081526020015b60405180910390f35b34801561016457600080fd5b50600054610145565b34801561017957600080fd5b5061018d610188366004611643565b6105bc565b005b34801561019b57600080fd5b5061014560015481565b3480156101b157600080fd5b506006546101c5906001600160a01b031681565b60405161014f91906117f1565b61018d6101e0366004611805565b610647565b3480156101f157600080fd5b50610145610666565b34801561020657600080fd5b5060408051808201909152600381526219171b60e91b60208201525b60405161014f91906118a4565b34801561023b57600080fd5b506005546101c5906001600160a01b031681565b34801561025b57600080fd5b506006546001600160a01b03166101c5565b34801561027957600080fd5b5061014560005481565b34801561028f57600080fd5b506004546101c5906001600160a01b031681565b3480156102af57600080fd5b506001610145565b3480156102c357600080fd5b5061018d6102d23660046118b7565b610683565b3480156102e357600080fd5b50610222604051806040016040528060058152602001640352e302e360dc1b81525081565b34801561031457600080fd5b506003546101c5906001600160a01b031681565b34801561033457600080fd5b5061014560025481565b34801561034a57600080fd5b506005546001600160a01b03166101c5565b34801561036857600080fd5b506101c5610791565b34801561037d57600080fd5b506004546001600160a01b03166101c5565b34801561039b57600080fd5b50600254610145565b3480156103b057600080fd5b506101456103bf3660046118f0565b6107a0565b3480156103d057600080fd5b5061018d6103df366004611986565b610afa565b6004805460405163599fbc6560e11b81526000926001600160a01b039092169163b33f78ca91610416913391016117f1565b602060405180830381865afa158015610433573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045791906119f2565b15806104d0575060048054604051633260339d60e11b81526001600160a01b03909116916364c0673a9161048d913391016117f1565b602060405180830381865afa1580156104aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ce91906119f2565b155b156104f95733604051637da4759f60e11b81526004016104f091906117f1565b60405180910390fd5b610501611385565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091526105b16040518060c00160405280898152602001886001600160a01b0316815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250938552505050602080830187905260408051918201815291815290820152606001839052610bfc565b979650505050505050565b6105d2336000356001600160e01b031916610cff565b60035460208301516040516390fe4cff60e01b81526001600160a01b03909216916390fe4cff9161060e9187919086908b908b90600401611a0d565b600060405180830381600087803b15801561062857600080fd5b505af115801561063c573d6000803e3d6000fd5b505050505050505050565b61064f610e4a565b61065882610eda565b6106628282610ef3565b5050565b6000610670610fac565b50600080516020611e5b83398151915290565b600061068d610ff5565b805490915060ff600160401b82041615906001600160401b03166000811580156106b45750825b90506000826001600160401b031660011480156106d05750303b155b9050811580156106de575080155b156106fc5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561072657845460ff60401b1916600160401b1785555b61072e611020565b6000889055600187905561074186611028565b831561078757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b600061079b611069565b905090565b6004805460405163599fbc6560e11b81526000926001600160a01b039092169163b33f78ca916107d2913391016117f1565b602060405180830381865afa1580156107ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081391906119f2565b158061088c575060048054604051633260339d60e11b81526001600160a01b03909116916364c0673a91610849913391016117f1565b602060405180830381865afa158015610866573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088a91906119f2565b155b156108ac5733604051637da4759f60e11b81526004016104f091906117f1565b6004805460405163599fbc6560e11b81526001600160a01b039091169163b33f78ca916108db918a91016117f1565b602060405180830381865afa1580156108f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091c91906119f2565b61093b57856040516353cd3f9360e11b81526004016104f091906117f1565b60048054604051633260339d60e11b81526001600160a01b03909116916364c0673a9161096a918a91016117f1565b602060405180830381865afa158015610987573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ab91906119f2565b6109ca5785604051631da5ffef60e11b81526004016104f091906117f1565b600480546040516348ef6c7d60e11b81526000926001600160a01b03909216916391ded8fa916109fc918b91016117f1565b600060405180830381865afa158015610a19573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a419190810190611ac4565b60a0015190506001600160a01b038116610a705786604051631f1c5c1d60e31b81526004016104f091906117f1565b610a78611385565b610aed6040518060c001604052808b8152602001846001600160a01b0316815260200189898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020810184905260408101889052606001869052610bfc565b9998505050505050505050565b610b10336000356001600160e01b031916610cff565b6001600160a01b038416610b3757604051635e62a26160e01b815260040160405180910390fd5b6001600160a01b038316610b5e576040516366a858c560e01b815260040160405180910390fd5b6001600160a01b038216610b8557604051633463e97160e11b815260040160405180910390fd5b6001600160a01b038116610bac57604051635e59d63f60e11b815260040160405180910390fd5b600380546001600160a01b039586166001600160a01b0319918216179091556004805494861694821694909417909355600580549285169284169290921790915560068054919093169116179055565b600080600260008154610c0e90611c2c565b91829055506040805160c081018252808201838152606087810151908301526080878101519083015260a0808801519083015281529085015160208201526000548551929350909103610c7457604051631ceda0f560e11b815260040160405180910390fd5b6006546000805486516020880151604051630c5a8fc560e21b815293946001600160a01b03169363316a3f1493610cb393909233928990600401611cc0565b6020604051808303816000875af1158015610cd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf69190611dc4565b95945050505050565b6000610d09611082565b80549091506001600160a01b031680610d38576000604051638944034760e01b81526004016104f091906117f1565b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa158015610d9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc09190611ddd565b92509250925082610e41578015610dea5760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff821615610e265760405163a426878960e01b81526001600160a01b038816600482015263ffffffff831660248201526044016104f0565b86604051632ecd3d0360e21b81526004016104f091906117f1565b50505050505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610eba57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610eae6110e4565b6001600160a01b031614155b15610ed85760405163703e46dd60e11b815260040160405180910390fd5b565b610ef0336000356001600160e01b031916610cff565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f4d575060408051601f3d908101601f19168201909252610f4a91810190611dc4565b60015b610f6c5781604051634c9c8ce360e01b81526004016104f091906117f1565b600080516020611e5b8339815191528114610f9d57604051632a87526960e21b8152600481018290526024016104f0565b610fa783836110fa565b505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ed85760405163703e46dd60e11b815260040160405180910390fd5b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b610ed8611150565b6000611032611082565b80549091506001600160a01b0316156110605781604051638944034760e01b81526004016104f091906117f1565b61066282611175565b6000611073611082565b546001600160a01b0316919050565b60008060ff196110b360017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f35611e2b565b6040516020016110c591815260200190565b60408051601f1981840301815291905280516020909101201692915050565b6000600080516020611e5b833981519152611073565b61110382611205565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561114857610fa78282611261565b6106626112ce565b6111586112ed565b610ed857604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b03811661119e5780604051638944034760e01b81526004016104f091906117f1565b60006111a8611082565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b806001600160a01b03163b6000036112325780604051634c9c8ce360e01b81526004016104f091906117f1565b600080516020611e5b83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161127e9190611e3e565b600060405180830381855af49150503d80600081146112b9576040519150601f19603f3d011682016040523d82523d6000602084013e6112be565b606091505b5091509150610cf6858383611307565b3415610ed85760405163b398979f60e01b815260040160405180910390fd5b60006112f7610ff5565b54600160401b900460ff16919050565b60608261131c576113178261135d565b611356565b815115801561133357506001600160a01b0384163b155b156113535783604051639996b31560e01b81526004016104f091906117f1565b50805b9392505050565b80511561136c57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b604080516080810190915280600081526060602082015260400160008152602001606081525090565b6001600160a01b0381168114610ef057600080fd5b60008083601f8401126113d557600080fd5b5081356001600160401b038111156113ec57600080fd5b60208301915083602082850101111561140457600080fd5b9250929050565b6000806000806060858703121561142157600080fd5b843593506020850135611433816113ae565b925060408501356001600160401b0381111561144e57600080fd5b61145a878288016113c3565b95989497509550505050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561149e5761149e611466565b60405290565b604051608081016001600160401b038111828210171561149e5761149e611466565b6040516101a081016001600160401b038111828210171561149e5761149e611466565b604051601f8201601f191681016001600160401b038111828210171561151157611511611466565b604052919050565b60006001600160401b0382111561153257611532611466565b50601f01601f191660200190565b600082601f83011261155157600080fd5b813561156461155f82611519565b6114e9565b81815284602083860101111561157957600080fd5b816020850160208301376000918101602001919091529392505050565b8035600781106115a557600080fd5b919050565b600060c082840312156115bc57600080fd5b60405160c081018181106001600160401b03821117156115de576115de611466565b6040529050806115ed83611596565b8152602083013560208201526040830135611607816113ae565b6040820152606083013561161a816113ae565b6060820152608083013561162d816113ae565b608082015260a092830135920191909152919050565b600080600080600060a0868803121561165b57600080fd5b85359450602086013561166d816113ae565b9350604086013561167d816113ae565b925060608601356001600160401b038082111561169957600080fd5b908701906040828a0312156116ad57600080fd5b6116b561147c565b81833511156116c357600080fd5b82358301610120818c0312156116d857600080fd5b6116e06114a4565b813581526020820135848111156116f657600080fd5b82016080818e03121561170857600080fd5b6117106114a4565b600282351061171e57600080fd5b8135815260208201358681111561173457600080fd5b6117408f828501611540565b60208301525061175260408301611596565b604082015260608201358681111561176957600080fd5b6117758f828501611540565b60608301525060208301525060408201358481111561179357600080fd5b61179f8d828501611540565b6040830152506117b28c606084016115aa565b60608201528252506020830135828111156117cc57600080fd5b6117d88b828601611540565b6020830152509699959850939660800135949350505050565b6001600160a01b0391909116815260200190565b6000806040838503121561181857600080fd5b8235611823816113ae565b915060208301356001600160401b0381111561183e57600080fd5b61184a85828601611540565b9150509250929050565b60005b8381101561186f578181015183820152602001611857565b50506000910152565b60008151808452611890816020860160208601611854565b601f01601f19169290920160200192915050565b6020815260006113566020830184611878565b6000806000606084860312156118cc57600080fd5b833592506020840135915060408401356118e5816113ae565b809150509250925092565b600080600080600080610140878903121561190a57600080fd5b86359550602087013561191c816113ae565b945060408701356001600160401b038082111561193857600080fd5b6119448a838b016113c3565b9096509450606089013591508082111561195d57600080fd5b5061196a89828a01611540565b92505061197a88608089016115aa565b90509295509295509295565b6000806000806080858703121561199c57600080fd5b84356119a7816113ae565b935060208501356119b7816113ae565b925060408501356119c7816113ae565b915060608501356119d7816113ae565b939692955090935050565b805180151581146115a557600080fd5b600060208284031215611a0457600080fd5b611356826119e2565b600060018060a01b03808816835260a06020840152611a2f60a0840188611878565b6040840196909652606083019490945250911660809091015292915050565b600082601f830112611a5f57600080fd5b8151611a6d61155f82611519565b818152846020838601011115611a8257600080fd5b611a93826020830160208701611854565b949350505050565b80516115a5816113ae565b8051600d81106115a557600080fd5b8051600581106115a557600080fd5b600060208284031215611ad657600080fd5b81516001600160401b0380821115611aed57600080fd5b908301906101a08286031215611b0257600080fd5b611b0a6114c6565b82518152602083015182811115611b2057600080fd5b611b2c87828601611a4e565b602083015250604083015182811115611b4457600080fd5b611b5087828601611a4e565b604083015250606083015182811115611b6857600080fd5b611b7487828601611a4e565b606083015250611b8660808401611a9b565b6080820152611b9760a08401611a9b565b60a082015260c083015160c0820152611bb260e08401611aa6565b60e08201526101009150611bc7828401611ab5565b828201526101209150611bdb828401611ab5565b828201526101409150611bef828401611ab5565b91810191909152610160828101519082015261018091820151918101919091529392505050565b634e487b7160e01b600052601160045260246000fd5b600060018201611c3e57611c3e611c16565b5060010190565b634e487b7160e01b600052602160045260246000fd5b60078110611c6b57611c6b611c45565b9052565b611c7a828251611c5b565b60208101516020830152604081015160018060a01b038082166040850152806060840151166060850152806080840151166080850152505060a081015160a08301525050565b858152600060018060a01b03808716602084015285604084015280851660608401525060a060808301528251604060a0840152805160e0840152602081015161012080610100860152815160028110611d1b57611d1b611c45565b61020086015260208201516080610220870152611d3c610280870182611878565b90506040830151611d51610240880182611c5b565b50606083015192506101ff1986820301610260870152611d718184611878565b925050604083015160df198684030182870152611d8e8382611878565b9250505060608201519150611da7610140850183611c6f565b6020850151848203609f190160c08601529150610aed8183611878565b600060208284031215611dd657600080fd5b5051919050565b600080600060608486031215611df257600080fd5b611dfb846119e2565b9250602084015163ffffffff81168114611e1457600080fd5b9150611e22604085016119e2565b90509250925092565b8181038181111561101a5761101a611c16565b60008251611e50818460208701611854565b919091019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122023f96096abf981cb24c0b9e3b8d6190ebaa107d9f3a0ef9280a701a8b492993e64736f6c63430008180033",
}

// RNEndpointV1 is an auto generated Go binding around an Ethereum contract.
type RNEndpointV1 struct {
	abi abi.ABI
}

// NewRNEndpointV1 creates a new instance of RNEndpointV1.
func NewRNEndpointV1() *RNEndpointV1 {
	parsed, err := RNEndpointV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNEndpointV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNEndpointV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNEndpointV1 *RNEndpointV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := rNEndpointV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNEndpointV1 *RNEndpointV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := rNEndpointV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (rNEndpointV1 *RNEndpointV1) PackAuthority() []byte {
	enc, err := rNEndpointV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("authority", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackConfigureContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea77c671.
//
// Solidity: function configureContracts(address _messageExecutor, address _tokenRegistry, address _userGovernance, address _messageDispatcher) returns()
func (rNEndpointV1 *RNEndpointV1) PackConfigureContracts(messageExecutor common.Address, tokenRegistry common.Address, userGovernance common.Address, messageDispatcher common.Address) []byte {
	enc, err := rNEndpointV1.abi.Pack("configureContracts", messageExecutor, tokenRegistry, userGovernance, messageDispatcher)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackContractVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackContractVersion() []byte {
	enc, err := rNEndpointV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("contractVersion", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCurrentChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackCurrentChainId() []byte {
	enc, err := rNEndpointV1.abi.Pack("currentChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCurrentChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackCurrentChainId(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("currentChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackGetChainId() []byte {
	enc, err := rNEndpointV1.abi.Pack("getChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackGetChainId(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("getChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetMessageDispatcherAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6431087a.
//
// Solidity: function getMessageDispatcherAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackGetMessageDispatcherAddress() []byte {
	enc, err := rNEndpointV1.abi.Pack("getMessageDispatcherAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMessageDispatcherAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6431087a.
//
// Solidity: function getMessageDispatcherAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackGetMessageDispatcherAddress(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("getMessageDispatcherAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackGetNonce() []byte {
	enc, err := rNEndpointV1.abi.Pack("getNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetTokenRegistryAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcaed9897.
//
// Solidity: function getTokenRegistryAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackGetTokenRegistryAddress() []byte {
	enc, err := rNEndpointV1.abi.Pack("getTokenRegistryAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenRegistryAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcaed9897.
//
// Solidity: function getTokenRegistryAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackGetTokenRegistryAddress(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("getTokenRegistryAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetUserGovernanceAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb8b87138.
//
// Solidity: function getUserGovernanceAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackGetUserGovernanceAddress() []byte {
	enc, err := rNEndpointV1.abi.Pack("getUserGovernanceAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUserGovernanceAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb8b87138.
//
// Solidity: function getUserGovernanceAddress() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackGetUserGovernanceAddress(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("getUserGovernanceAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _chainId, uint256 _publicChainId, address authority_) returns()
func (rNEndpointV1 *RNEndpointV1) PackInitialize(chainId *big.Int, publicChainId *big.Int, authority common.Address) []byte {
	enc, err := rNEndpointV1.abi.Pack("initialize", chainId, publicChainId, authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMessageDispatcher is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48628eed.
//
// Solidity: function messageDispatcher() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackMessageDispatcher() []byte {
	enc, err := rNEndpointV1.abi.Pack("messageDispatcher")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMessageDispatcher is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48628eed.
//
// Solidity: function messageDispatcher() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackMessageDispatcher(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("messageDispatcher", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackMessageExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad552e32.
//
// Solidity: function messageExecutor() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackMessageExecutor() []byte {
	enc, err := rNEndpointV1.abi.Pack("messageExecutor")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMessageExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad552e32.
//
// Solidity: function messageExecutor() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackMessageExecutor(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("messageExecutor", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackNonce() []byte {
	enc, err := rNEndpointV1.abi.Pack("nonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackNonce(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("nonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNEndpointV1 *RNEndpointV1) PackProxiableUUID() []byte {
	enc, err := rNEndpointV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNEndpointV1 *RNEndpointV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := rNEndpointV1.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackPublicChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47219ccc.
//
// Solidity: function publicChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) PackPublicChainId() []byte {
	enc, err := rNEndpointV1.abi.Pack("publicChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPublicChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x47219ccc.
//
// Solidity: function publicChainId() view returns(uint256)
func (rNEndpointV1 *RNEndpointV1) UnpackPublicChainId(data []byte) (*big.Int, error) {
	out, err := rNEndpointV1.abi.Unpack("publicChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackReceivePayload is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37ffde05.
//
// Solidity: function receivePayload(uint256 _srcChainId, address _srcAddress, address _dstAddress, ((uint256,(uint8,bytes,uint8,bytes),bytes,(uint8,uint256,address,address,address,uint256)),bytes) _raylsMessage, bytes32 _messageId) returns()
func (rNEndpointV1 *RNEndpointV1) PackReceivePayload(srcChainId *big.Int, srcAddress common.Address, dstAddress common.Address, raylsMessage RaylsNodeMessage, messageId [32]byte) []byte {
	enc, err := rNEndpointV1.abi.Pack("receivePayload", srcChainId, srcAddress, dstAddress, raylsMessage, messageId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSend is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b375f.
//
// Solidity: function send(uint256 _dstChainId, address _destination, bytes _payload) returns(bytes32 messageId)
func (rNEndpointV1 *RNEndpointV1) PackSend(dstChainId *big.Int, destination common.Address, payload []byte) []byte {
	enc, err := rNEndpointV1.abi.Pack("send", dstChainId, destination, payload)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSend is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b375f.
//
// Solidity: function send(uint256 _dstChainId, address _destination, bytes _payload) returns(bytes32 messageId)
func (rNEndpointV1 *RNEndpointV1) UnpackSend(data []byte) ([32]byte, error) {
	out, err := rNEndpointV1.abi.Unpack("send", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackSendToAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0d9d846.
//
// Solidity: function sendToAddress(uint256 _dstChainId, address _privateChainAddress, bytes _payload, bytes _revertDataPayload, (uint8,uint256,address,address,address,uint256) transferMetadata) returns(bytes32 messageId)
func (rNEndpointV1 *RNEndpointV1) PackSendToAddress(dstChainId *big.Int, privateChainAddress common.Address, payload []byte, revertDataPayload []byte, transferMetadata RaylsNodeBridgedTransferMetadata) []byte {
	enc, err := rNEndpointV1.abi.Pack("sendToAddress", dstChainId, privateChainAddress, payload, revertDataPayload, transferMetadata)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSendToAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe0d9d846.
//
// Solidity: function sendToAddress(uint256 _dstChainId, address _privateChainAddress, bytes _payload, bytes _revertDataPayload, (uint8,uint256,address,address,address,uint256) transferMetadata) returns(bytes32 messageId)
func (rNEndpointV1 *RNEndpointV1) UnpackSendToAddress(data []byte) ([32]byte, error) {
	out, err := rNEndpointV1.abi.Unpack("sendToAddress", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackTokenRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackTokenRegistry() []byte {
	enc, err := rNEndpointV1.abi.Pack("tokenRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackTokenRegistry(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("tokenRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (rNEndpointV1 *RNEndpointV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := rNEndpointV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUserGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d31f35f.
//
// Solidity: function userGovernance() view returns(address)
func (rNEndpointV1 *RNEndpointV1) PackUserGovernance() []byte {
	enc, err := rNEndpointV1.abi.Pack("userGovernance")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUserGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d31f35f.
//
// Solidity: function userGovernance() view returns(address)
func (rNEndpointV1 *RNEndpointV1) UnpackUserGovernance(data []byte) (common.Address, error) {
	out, err := rNEndpointV1.abi.Unpack("userGovernance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (rNEndpointV1 *RNEndpointV1) PackVersion() []byte {
	enc, err := rNEndpointV1.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (rNEndpointV1 *RNEndpointV1) UnpackVersion(data []byte) (string, error) {
	out, err := rNEndpointV1.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// RNEndpointV1AuthorityUpdated represents a AuthorityUpdated event raised by the RNEndpointV1 contract.
type RNEndpointV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNEndpointV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RNEndpointV1AuthorityUpdated) ContractEventName() string {
	return RNEndpointV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (rNEndpointV1 *RNEndpointV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RNEndpointV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != rNEndpointV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNEndpointV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := rNEndpointV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNEndpointV1.abi.Events[event].Inputs {
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

// RNEndpointV1Initialized represents a Initialized event raised by the RNEndpointV1 contract.
type RNEndpointV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNEndpointV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNEndpointV1Initialized) ContractEventName() string {
	return RNEndpointV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNEndpointV1 *RNEndpointV1) UnpackInitializedEvent(log *types.Log) (*RNEndpointV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNEndpointV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNEndpointV1Initialized)
	if len(log.Data) > 0 {
		if err := rNEndpointV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNEndpointV1.abi.Events[event].Inputs {
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

// RNEndpointV1Upgraded represents a Upgraded event raised by the RNEndpointV1 contract.
type RNEndpointV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNEndpointV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RNEndpointV1Upgraded) ContractEventName() string {
	return RNEndpointV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (rNEndpointV1 *RNEndpointV1) UnpackUpgradedEvent(log *types.Log) (*RNEndpointV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != rNEndpointV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNEndpointV1Upgraded)
	if len(log.Data) > 0 {
		if err := rNEndpointV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNEndpointV1.abi.Events[event].Inputs {
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
func (rNEndpointV1 *RNEndpointV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1CalledByUnauthorizedAddress"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1CalledByUnauthorizedAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1InvalidMessageDispatcherAddress"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1InvalidMessageDispatcherAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1InvalidMessageExecutorAddress"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1InvalidMessageExecutorAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1InvalidTokenRegistryAddress"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1InvalidTokenRegistryAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1InvalidUserGovernanceAddress"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1InvalidUserGovernanceAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1NoPublicAddressMapping"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1NoPublicAddressMappingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1SourceAndDestinationChainsSame"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1SourceAndDestinationChainsSameError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1TokenNotActive"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1TokenNotActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1TokenNotFound"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1TokenNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RNEndpointV1TokenUnauthorizedAccount"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRNEndpointV1TokenUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNEndpointV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return rNEndpointV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNEndpointV1AddressEmptyCode represents a AddressEmptyCode error raised by the RNEndpointV1 contract.
type RNEndpointV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RNEndpointV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (rNEndpointV1 *RNEndpointV1) UnpackAddressEmptyCodeError(raw []byte) (*RNEndpointV1AddressEmptyCode, error) {
	out := new(RNEndpointV1AddressEmptyCode)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RNEndpointV1 contract.
type RNEndpointV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RNEndpointV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (rNEndpointV1 *RNEndpointV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RNEndpointV1ERC1967InvalidImplementation, error) {
	out := new(RNEndpointV1ERC1967InvalidImplementation)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RNEndpointV1 contract.
type RNEndpointV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RNEndpointV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (rNEndpointV1 *RNEndpointV1) UnpackERC1967NonPayableError(raw []byte) (*RNEndpointV1ERC1967NonPayable, error) {
	out := new(RNEndpointV1ERC1967NonPayable)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1FailedCall represents a FailedCall error raised by the RNEndpointV1 contract.
type RNEndpointV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RNEndpointV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (rNEndpointV1 *RNEndpointV1) UnpackFailedCallError(raw []byte) (*RNEndpointV1FailedCall, error) {
	out := new(RNEndpointV1FailedCall)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1InvalidInitialization represents a InvalidInitialization error raised by the RNEndpointV1 contract.
type RNEndpointV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNEndpointV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNEndpointV1 *RNEndpointV1) UnpackInvalidInitializationError(raw []byte) (*RNEndpointV1InvalidInitialization, error) {
	out := new(RNEndpointV1InvalidInitialization)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1NotInitializing represents a NotInitializing error raised by the RNEndpointV1 contract.
type RNEndpointV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNEndpointV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNEndpointV1 *RNEndpointV1) UnpackNotInitializingError(raw []byte) (*RNEndpointV1NotInitializing, error) {
	out := new(RNEndpointV1NotInitializing)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1CalledByUnauthorizedAddress represents a RNEndpointV1__CalledByUnauthorizedAddress error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1CalledByUnauthorizedAddress struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__CalledByUnauthorizedAddress(address account)
func RNEndpointV1RNEndpointV1CalledByUnauthorizedAddressErrorID() common.Hash {
	return common.HexToHash("0x622ea64df6bfc9d552336a17a9bded06666c1b016cb0ce1dd1d741049fcd8902")
}

// UnpackRNEndpointV1CalledByUnauthorizedAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__CalledByUnauthorizedAddress(address account)
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1CalledByUnauthorizedAddressError(raw []byte) (*RNEndpointV1RNEndpointV1CalledByUnauthorizedAddress, error) {
	out := new(RNEndpointV1RNEndpointV1CalledByUnauthorizedAddress)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1CalledByUnauthorizedAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1InvalidMessageDispatcherAddress represents a RNEndpointV1__InvalidMessageDispatcherAddress error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1InvalidMessageDispatcherAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__InvalidMessageDispatcherAddress()
func RNEndpointV1RNEndpointV1InvalidMessageDispatcherAddressErrorID() common.Hash {
	return common.HexToHash("0xbcb3ac7ed7b91fcfe7f55fcd74bb2d264cf13a036556881f9a7052794e7cdaeb")
}

// UnpackRNEndpointV1InvalidMessageDispatcherAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__InvalidMessageDispatcherAddress()
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1InvalidMessageDispatcherAddressError(raw []byte) (*RNEndpointV1RNEndpointV1InvalidMessageDispatcherAddress, error) {
	out := new(RNEndpointV1RNEndpointV1InvalidMessageDispatcherAddress)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1InvalidMessageDispatcherAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1InvalidMessageExecutorAddress represents a RNEndpointV1__InvalidMessageExecutorAddress error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1InvalidMessageExecutorAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__InvalidMessageExecutorAddress()
func RNEndpointV1RNEndpointV1InvalidMessageExecutorAddressErrorID() common.Hash {
	return common.HexToHash("0x5e62a26119cf127519068d0231ee1c01494813ee84c0bc5b5007ea644d44c514")
}

// UnpackRNEndpointV1InvalidMessageExecutorAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__InvalidMessageExecutorAddress()
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1InvalidMessageExecutorAddressError(raw []byte) (*RNEndpointV1RNEndpointV1InvalidMessageExecutorAddress, error) {
	out := new(RNEndpointV1RNEndpointV1InvalidMessageExecutorAddress)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1InvalidMessageExecutorAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1InvalidTokenRegistryAddress represents a RNEndpointV1__InvalidTokenRegistryAddress error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1InvalidTokenRegistryAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__InvalidTokenRegistryAddress()
func RNEndpointV1RNEndpointV1InvalidTokenRegistryAddressErrorID() common.Hash {
	return common.HexToHash("0x66a858c5feef81d6740209a2594c0ae954aab5bdf5ebe2978aa2dd540e77a4b3")
}

// UnpackRNEndpointV1InvalidTokenRegistryAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__InvalidTokenRegistryAddress()
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1InvalidTokenRegistryAddressError(raw []byte) (*RNEndpointV1RNEndpointV1InvalidTokenRegistryAddress, error) {
	out := new(RNEndpointV1RNEndpointV1InvalidTokenRegistryAddress)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1InvalidTokenRegistryAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1InvalidUserGovernanceAddress represents a RNEndpointV1__InvalidUserGovernanceAddress error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1InvalidUserGovernanceAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__InvalidUserGovernanceAddress()
func RNEndpointV1RNEndpointV1InvalidUserGovernanceAddressErrorID() common.Hash {
	return common.HexToHash("0x68c7d2e2fd70b01806f2e3117feea42ad9e0503e657f6e9fab95a352c3273d2f")
}

// UnpackRNEndpointV1InvalidUserGovernanceAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__InvalidUserGovernanceAddress()
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1InvalidUserGovernanceAddressError(raw []byte) (*RNEndpointV1RNEndpointV1InvalidUserGovernanceAddress, error) {
	out := new(RNEndpointV1RNEndpointV1InvalidUserGovernanceAddress)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1InvalidUserGovernanceAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1NoPublicAddressMapping represents a RNEndpointV1__NoPublicAddressMapping error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1NoPublicAddressMapping struct {
	PrivateAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__NoPublicAddressMapping(address privateAddress)
func RNEndpointV1RNEndpointV1NoPublicAddressMappingErrorID() common.Hash {
	return common.HexToHash("0xf8e2e0e8b84b32879835a60258d9e5ebd2c764f0b4f0b4d43b3c2e8f848fa8b7")
}

// UnpackRNEndpointV1NoPublicAddressMappingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__NoPublicAddressMapping(address privateAddress)
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1NoPublicAddressMappingError(raw []byte) (*RNEndpointV1RNEndpointV1NoPublicAddressMapping, error) {
	out := new(RNEndpointV1RNEndpointV1NoPublicAddressMapping)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1NoPublicAddressMapping", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1SourceAndDestinationChainsSame represents a RNEndpointV1__SourceAndDestinationChainsSame error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1SourceAndDestinationChainsSame struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__SourceAndDestinationChainsSame()
func RNEndpointV1RNEndpointV1SourceAndDestinationChainsSameErrorID() common.Hash {
	return common.HexToHash("0x39db41eacf737f778ad9d69984c5a72e558ad178326d69c4efe94e8cb3794150")
}

// UnpackRNEndpointV1SourceAndDestinationChainsSameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__SourceAndDestinationChainsSame()
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1SourceAndDestinationChainsSameError(raw []byte) (*RNEndpointV1RNEndpointV1SourceAndDestinationChainsSame, error) {
	out := new(RNEndpointV1RNEndpointV1SourceAndDestinationChainsSame)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1SourceAndDestinationChainsSame", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1TokenNotActive represents a RNEndpointV1__TokenNotActive error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1TokenNotActive struct {
	PrivateAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__TokenNotActive(address privateAddress)
func RNEndpointV1RNEndpointV1TokenNotActiveErrorID() common.Hash {
	return common.HexToHash("0x3b4bffdeaa80b87762981559a3cf1cae68fb183461ad2f88fbb8d6afec319f6f")
}

// UnpackRNEndpointV1TokenNotActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__TokenNotActive(address privateAddress)
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1TokenNotActiveError(raw []byte) (*RNEndpointV1RNEndpointV1TokenNotActive, error) {
	out := new(RNEndpointV1RNEndpointV1TokenNotActive)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1TokenNotActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1TokenNotFound represents a RNEndpointV1__TokenNotFound error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1TokenNotFound struct {
	PrivateAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__TokenNotFound(address privateAddress)
func RNEndpointV1RNEndpointV1TokenNotFoundErrorID() common.Hash {
	return common.HexToHash("0xa79a7f268a92081c20f92b251f2c4343c67d43780edb0921b1f581a188b65388")
}

// UnpackRNEndpointV1TokenNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__TokenNotFound(address privateAddress)
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1TokenNotFoundError(raw []byte) (*RNEndpointV1RNEndpointV1TokenNotFound, error) {
	out := new(RNEndpointV1RNEndpointV1TokenNotFound)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1TokenNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RNEndpointV1TokenUnauthorizedAccount represents a RNEndpointV1__TokenUnauthorizedAccount error raised by the RNEndpointV1 contract.
type RNEndpointV1RNEndpointV1TokenUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNEndpointV1__TokenUnauthorizedAccount(address account)
func RNEndpointV1RNEndpointV1TokenUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xfb48eb3e832134099b94f0b0dbc353e26c605926174e48d3226ad546efa0624c")
}

// UnpackRNEndpointV1TokenUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNEndpointV1__TokenUnauthorizedAccount(address account)
func (rNEndpointV1 *RNEndpointV1) UnpackRNEndpointV1TokenUnauthorizedAccountError(raw []byte) (*RNEndpointV1RNEndpointV1TokenUnauthorizedAccount, error) {
	out := new(RNEndpointV1RNEndpointV1TokenUnauthorizedAccount)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RNEndpointV1TokenUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RNEndpointV1 contract.
type RNEndpointV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RNEndpointV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (rNEndpointV1 *RNEndpointV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RNEndpointV1RaylsAccessManagedContractPaused, error) {
	out := new(RNEndpointV1RaylsAccessManagedContractPaused)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RNEndpointV1 contract.
type RNEndpointV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RNEndpointV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (rNEndpointV1 *RNEndpointV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RNEndpointV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RNEndpointV1RaylsAccessManagedInvalidAuthority)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RNEndpointV1 contract.
type RNEndpointV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RNEndpointV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (rNEndpointV1 *RNEndpointV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RNEndpointV1RaylsAccessManagedMustSchedule, error) {
	out := new(RNEndpointV1RaylsAccessManagedMustSchedule)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RNEndpointV1 contract.
type RNEndpointV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RNEndpointV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (rNEndpointV1 *RNEndpointV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RNEndpointV1RaylsAccessManagedUnauthorized, error) {
	out := new(RNEndpointV1RaylsAccessManagedUnauthorized)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RNEndpointV1 contract.
type RNEndpointV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RNEndpointV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (rNEndpointV1 *RNEndpointV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RNEndpointV1UUPSUnauthorizedCallContext, error) {
	out := new(RNEndpointV1UUPSUnauthorizedCallContext)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNEndpointV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RNEndpointV1 contract.
type RNEndpointV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RNEndpointV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (rNEndpointV1 *RNEndpointV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RNEndpointV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RNEndpointV1UUPSUnsupportedProxiableUUID)
	if err := rNEndpointV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
