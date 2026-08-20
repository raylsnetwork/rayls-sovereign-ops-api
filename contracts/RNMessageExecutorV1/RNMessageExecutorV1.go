// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RNMessageExecutorV1

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

// RNMessageLibMessage is an auto generated low-level Go binding around an user-defined struct.
type RNMessageLibMessage struct {
	To   common.Address
	Data []byte
}

// RNMessageExecutorV1MetaData contains all meta data concerning the RNMessageExecutorV1 contract.
var RNMessageExecutorV1MetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authority\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizedEndpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"currentChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeMessage\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeMessageBatch\",\"inputs\":[{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structRNMessageLib.Message[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executed\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAuthorizedEndpoint\",\"inputs\":[{\"name\":\"_authorizedEndpoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuthorityUpdated\",\"inputs\":[{\"name\":\"oldAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAuthority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageIdExecuted\",\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__InvalidEndpointAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__MessageBatchFailure\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"errorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__MessageFailure\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"errorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__MessageIdAlreadyExecuted\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__NoContractAtAddress\",\"inputs\":[{\"name\":\"contract_address\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNMessageExecutorV1__UnauthorizedEndpoint\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RNReentrancyGuardV1__ReceiveReentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RNReentrancyGuardV1__SendReentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__InvalidAuthority\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__MustSchedule\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"RaylsAccessManaged__Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	ID:  "RNMessageExecutorV1",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516114656100fd600039600081816108aa015281816108d30152610a0c01526114656000f3fe6080604052600436106100a25760003560e01c80632f8d242f146100a75780634f1ef286146100dd57806352d1902d146100f25780636cbadbfa146101155780637a6865211461012b5780638129fc1c1461014b57806390fe4cff14610160578063a0a8e46014610180578063a9fcfb3314610194578063ad3cb1cc146101d4578063bf7e214f14610212578063c4d66de814610227578063c5c37a0014610247575b600080fd5b3480156100b357600080fd5b506002546100c7906001600160a01b031681565b6040516100d49190610f42565b60405180910390f35b6100f06100eb366004611014565b610267565b005b3480156100fe57600080fd5b50610107610286565b6040519081526020016100d4565b34801561012157600080fd5b5061010760015481565b34801561013757600080fd5b506100f0610146366004611061565b6102a3565b34801561015757600080fd5b506100f0610302565b34801561016c57600080fd5b506100f061017b36600461107c565b610401565b34801561018c57600080fd5b506001610107565b3480156101a057600080fd5b506101c46101af366004611121565b60036020526000908152604090205460ff1681565b60405190151581526020016100d4565b3480156101e057600080fd5b50610205604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516100d4919061118a565b34801561021e57600080fd5b506100c7610586565b34801561023357600080fd5b506100f0610242366004611061565b61059f565b34801561025357600080fd5b506100f061026236600461119d565b6106ad565b61026f61089f565b6102788261092f565b6102828282610948565b5050565b6000610290610a01565b5060008051602061141083398151915290565b6102b9336000356001600160e01b031916610a4a565b6001600160a01b0381166102e057604051630a577ba960e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061030c610b95565b805490915060ff600160401b82041615906001600160401b03166000811580156103335750825b90506000826001600160401b0316600114801561034f5750303b155b90508115801561035d575080155b1561037b5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103a557845460ff60401b1916600160401b1785555b6000805461ffff191661010117905583156103fa57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050565b6002546001600160a01b03163314610437573360405163347e2eb560e21b815260040161042e9190610f42565b60405180910390fd5b60005460011961010090910460ff1601610464576040516317326c6360e01b815260040160405180910390fd5b6000805461ff001916610200178155838152600360205260409020805460ff19811660011790915560ff1680156104b15760405163bb29ff9f60e01b81526004810185905260240161042e565b6104ba87610bc0565b600080886001600160a01b031688886040516104d7929190611232565b6000604051808303816000865af19150503d8060008114610514576040519150601f19603f3d011682016040523d82523d6000602084013e610519565b606091505b5091509150816105405785816040516303dd878360e41b815260040161042e929190611242565b604051869086907e769f3f82cb2a521c5b72f211aff687dae3cebd0b4631790417d1b17e15689a90600090a350506000805461ff00191661010017905550505050505050565b6000610590610bed565b546001600160a01b0316919050565b60006105a9610b95565b805490915060ff600160401b82041615906001600160401b03166000811580156105d05750825b90506000826001600160401b031660011480156105ec5750303b155b9050811580156105fa575080155b156106185760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561064257845460ff60401b1916600160401b1785555b61064a610c4f565b610652610302565b4660015561065f86610c57565b83156106a557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6002546001600160a01b031633146106da573360405163347e2eb560e21b815260040161042e9190610f42565b60005460011961010090910460ff1601610707576040516317326c6360e01b815260040160405180910390fd5b6000805461ff001916610200178155838152600360205260409020805460ff19811660011790915560ff1680156107545760405163bb29ff9f60e01b81526004810185905260240161042e565b8460005b8181101561085a57600088888381811061077457610774611263565b90506020028101906107869190611279565b61078f90611299565b905061079e8160000151610bc0565b60008082600001516001600160a01b031683602001518a8a8a6040516020016107ca949392919061130a565b60408051601f19818403018152908290526107e491611346565b6000604051808303816000865af19150503d8060008114610821576040519150601f19603f3d011682016040523d82523d6000602084013e610826565b606091505b50915091508161084f57888482604051634755189760e01b815260040161042e93929190611358565b505050600101610758565b50604051859085907e769f3f82cb2a521c5b72f211aff687dae3cebd0b4631790417d1b17e15689a90600090a350506000805461ff0019166101001790555050505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061090f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610903610c98565b6001600160a01b031614155b1561092d5760405163703e46dd60e11b815260040160405180910390fd5b565b610945336000356001600160e01b031916610a4a565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156109a2575060408051601f3d908101601f1916820190925261099f91810190611377565b60015b6109c15781604051634c9c8ce360e01b815260040161042e9190610f42565b60008051602061141083398151915281146109f257604051632a87526960e21b81526004810182905260240161042e565b6109fc8383610cae565b505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461092d5760405163703e46dd60e11b815260040160405180910390fd5b6000610a54610bed565b80549091506001600160a01b031680610a83576000604051638944034760e01b815260040161042e9190610f42565b60405163b700961360e01b81526001600160a01b0385811660048301523060248301526001600160e01b031985166044830152600091829182919085169063b700961390606401606060405180830381865afa158015610ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b91906113a0565b92509250925082610b8c578015610b355760405163cc9855ad60e01b815260040160405180910390fd5b63ffffffff821615610b715760405163a426878960e01b81526001600160a01b038816600482015263ffffffff8316602482015260440161042e565b86604051632ecd3d0360e21b815260040161042e9190610f42565b50505050505050565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b806001600160a01b03163b60000361094557806040516336289bd560e11b815260040161042e9190610f42565b60008060ff19610c1e60017f2d9a26166e6ed6927c3e14d89c4d572a357f984344fcaed3c38c3ff5cdb83f356113ee565b604051602001610c3091815260200190565b60408051601f1981840301815291905280516020909101201692915050565b61092d610d04565b6000610c61610bed565b80549091506001600160a01b031615610c8f5781604051638944034760e01b815260040161042e9190610f42565b61028282610d29565b6000600080516020611410833981519152610590565b610cb782610db9565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115610cfc576109fc8282610e15565b610282610e8b565b610d0c610eaa565b61092d57604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b038116610d525780604051638944034760e01b815260040161042e9190610f42565b6000610d5c610bed565b80546040519192506001600160a01b03808516929116907fa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b7638998019890600090a380546001600160a01b0319166001600160a01b0392909216919091179055565b806001600160a01b03163b600003610de65780604051634c9c8ce360e01b815260040161042e9190610f42565b60008051602061141083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051610e329190611346565b600060405180830381855af49150503d8060008114610e6d576040519150601f19603f3d011682016040523d82523d6000602084013e610e72565b606091505b5091509150610e82858383610ec4565b95945050505050565b341561092d5760405163b398979f60e01b815260040160405180910390fd5b6000610eb4610b95565b54600160401b900460ff16919050565b606082610ed957610ed482610f1a565b610f13565b8151158015610ef057506001600160a01b0384163b155b15610f105783604051639996b31560e01b815260040161042e9190610f42565b50805b9392505050565b805115610f2957805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b6001600160a01b0391909116815260200190565b80356001600160a01b0381168114610f6d57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610f9957600080fd5b81356001600160401b0380821115610fb357610fb3610f72565b604051601f8301601f19908116603f01168101908282118183101715610fdb57610fdb610f72565b81604052838152866020858801011115610ff457600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561102757600080fd5b61103083610f56565b915060208301356001600160401b0381111561104b57600080fd5b61105785828601610f88565b9150509250929050565b60006020828403121561107357600080fd5b610f1382610f56565b60008060008060008060a0878903121561109557600080fd5b61109e87610f56565b955060208701356001600160401b03808211156110ba57600080fd5b818901915089601f8301126110ce57600080fd5b8135818111156110dd57600080fd5b8a60208285010111156110ef57600080fd5b602083019750809650505050604087013592506060870135915061111560808801610f56565b90509295509295509295565b60006020828403121561113357600080fd5b5035919050565b60005b8381101561115557818101518382015260200161113d565b50506000910152565b6000815180845261117681602086016020860161113a565b601f01601f19169290920160200192915050565b602081526000610f13602083018461115e565b6000806000806000608086880312156111b557600080fd5b85356001600160401b03808211156111cc57600080fd5b818801915088601f8301126111e057600080fd5b8135818111156111ef57600080fd5b8960208260051b850101111561120457600080fd5b6020928301975095505086013592506040860135915061122660608701610f56565b90509295509295909350565b8183823760009101908152919050565b82815260406020820152600061125b604083018461115e565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235603e1983360301811261128f57600080fd5b9190910192915050565b6000604082360312156112ab57600080fd5b604051604081016001600160401b0382821081831117156112ce576112ce610f72565b816040526112db85610f56565b835260208501359150808211156112f157600080fd5b506112fe36828601610f88565b60208301525092915050565b6000855161131c818460208a0161113a565b9190910193845250602083019190915260601b6001600160601b0319166040820152605401919050565b6000825161128f81846020870161113a565b838152826020820152606060408201526000610e82606083018461115e565b60006020828403121561138957600080fd5b5051919050565b80518015158114610f6d57600080fd5b6000806000606084860312156113b557600080fd5b6113be84611390565b9250602084015163ffffffff811681146113d757600080fd5b91506113e560408501611390565b90509250925092565b81810381811115610bba57634e487b7160e01b600052601160045260246000fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220e57a09b1986d3e237b136d88b933ad03783207d7c47b178725e64f1c27ab4b4564736f6c63430008180033",
}

// RNMessageExecutorV1 is an auto generated Go binding around an Ethereum contract.
type RNMessageExecutorV1 struct {
	abi abi.ABI
}

// NewRNMessageExecutorV1 creates a new instance of RNMessageExecutorV1.
func NewRNMessageExecutorV1() *RNMessageExecutorV1 {
	parsed, err := RNMessageExecutorV1MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &RNMessageExecutorV1{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *RNMessageExecutorV1) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackAuthority() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("authority")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthority is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackAuthority(data []byte) (common.Address, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("authority", data)
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackAuthorizedEndpoint() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("authorizedEndpoint")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAuthorizedEndpoint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f8d242f.
//
// Solidity: function authorizedEndpoint() view returns(address)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackAuthorizedEndpoint(data []byte) (common.Address, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("authorizedEndpoint", data)
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackContractVersion() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("contractVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackContractVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint256)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackContractVersion(data []byte) (*big.Int, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("contractVersion", data)
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackCurrentChainId() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("currentChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCurrentChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cbadbfa.
//
// Solidity: function currentChainId() view returns(uint256)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackCurrentChainId(data []byte) (*big.Int, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("currentChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackExecuteMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90fe4cff.
//
// Solidity: function executeMessage(address to, bytes data, bytes32 messageId, uint256 fromChainId, address from) returns()
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackExecuteMessage(to common.Address, data []byte, messageId [32]byte, fromChainId *big.Int, from common.Address) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("executeMessage", to, data, messageId, fromChainId, from)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecuteMessageBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5c37a00.
//
// Solidity: function executeMessageBatch((address,bytes)[] messages, bytes32 messageId, uint256 fromChainId, address from) returns()
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackExecuteMessageBatch(messages []RNMessageLibMessage, messageId [32]byte, fromChainId *big.Int, from common.Address) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("executeMessageBatch", messages, messageId, fromChainId, from)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecuted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackExecuted(arg0 [32]byte) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("executed", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecuted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackExecuted(data []byte) (bool, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("executed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackInitialize() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("initialize")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address authority_) returns()
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackInitialize0(authority common.Address) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("initialize0", authority)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackProxiableUUID() []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := rNMessageExecutorV1.abi.Unpack("proxiableUUID", data)
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackSetAuthorizedEndpoint(authorizedEndpoint common.Address) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("setAuthorizedEndpoint", authorizedEndpoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (rNMessageExecutorV1 *RNMessageExecutorV1) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := rNMessageExecutorV1.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// RNMessageExecutorV1AuthorityUpdated represents a AuthorityUpdated event raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1AuthorityUpdated struct {
	OldAuthority common.Address
	NewAuthority common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const RNMessageExecutorV1AuthorityUpdatedEventName = "AuthorityUpdated"

// ContractEventName returns the user-defined event name.
func (RNMessageExecutorV1AuthorityUpdated) ContractEventName() string {
	return RNMessageExecutorV1AuthorityUpdatedEventName
}

// UnpackAuthorityUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorityUpdated(address indexed oldAuthority, address indexed newAuthority)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackAuthorityUpdatedEvent(log *types.Log) (*RNMessageExecutorV1AuthorityUpdated, error) {
	event := "AuthorityUpdated"
	if log.Topics[0] != rNMessageExecutorV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageExecutorV1AuthorityUpdated)
	if len(log.Data) > 0 {
		if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageExecutorV1.abi.Events[event].Inputs {
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

// RNMessageExecutorV1Initialized represents a Initialized event raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1Initialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const RNMessageExecutorV1InitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (RNMessageExecutorV1Initialized) ContractEventName() string {
	return RNMessageExecutorV1InitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackInitializedEvent(log *types.Log) (*RNMessageExecutorV1Initialized, error) {
	event := "Initialized"
	if log.Topics[0] != rNMessageExecutorV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageExecutorV1Initialized)
	if len(log.Data) > 0 {
		if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageExecutorV1.abi.Events[event].Inputs {
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

// RNMessageExecutorV1MessageIdExecuted represents a MessageIdExecuted event raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1MessageIdExecuted struct {
	FromChainId *big.Int
	MessageId   [32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const RNMessageExecutorV1MessageIdExecutedEventName = "MessageIdExecuted"

// ContractEventName returns the user-defined event name.
func (RNMessageExecutorV1MessageIdExecuted) ContractEventName() string {
	return RNMessageExecutorV1MessageIdExecutedEventName
}

// UnpackMessageIdExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessageIdExecuted(uint256 indexed fromChainId, bytes32 indexed messageId)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackMessageIdExecutedEvent(log *types.Log) (*RNMessageExecutorV1MessageIdExecuted, error) {
	event := "MessageIdExecuted"
	if log.Topics[0] != rNMessageExecutorV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageExecutorV1MessageIdExecuted)
	if len(log.Data) > 0 {
		if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageExecutorV1.abi.Events[event].Inputs {
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

// RNMessageExecutorV1Upgraded represents a Upgraded event raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1Upgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const RNMessageExecutorV1UpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (RNMessageExecutorV1Upgraded) ContractEventName() string {
	return RNMessageExecutorV1UpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackUpgradedEvent(log *types.Log) (*RNMessageExecutorV1Upgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != rNMessageExecutorV1.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RNMessageExecutorV1Upgraded)
	if len(log.Data) > 0 {
		if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range rNMessageExecutorV1.abi.Events[event].Inputs {
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
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1InvalidEndpointAddress"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1InvalidEndpointAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1MessageBatchFailure"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1MessageBatchFailureError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1MessageFailure"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1MessageFailureError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1MessageIdAlreadyExecuted"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1MessageIdAlreadyExecutedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1NoContractAtAddress"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1NoContractAtAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNMessageExecutorV1UnauthorizedEndpoint"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNMessageExecutorV1UnauthorizedEndpointError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNReentrancyGuardV1ReceiveReentrancy"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNReentrancyGuardV1ReceiveReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RNReentrancyGuardV1SendReentrancy"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRNReentrancyGuardV1SendReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RaylsAccessManagedContractPaused"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRaylsAccessManagedContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RaylsAccessManagedInvalidAuthority"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRaylsAccessManagedInvalidAuthorityError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RaylsAccessManagedMustSchedule"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRaylsAccessManagedMustScheduleError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["RaylsAccessManagedUnauthorized"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackRaylsAccessManagedUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], rNMessageExecutorV1.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return rNMessageExecutorV1.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// RNMessageExecutorV1AddressEmptyCode represents a AddressEmptyCode error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1AddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func RNMessageExecutorV1AddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackAddressEmptyCodeError(raw []byte) (*RNMessageExecutorV1AddressEmptyCode, error) {
	out := new(RNMessageExecutorV1AddressEmptyCode)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1ERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1ERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func RNMessageExecutorV1ERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackERC1967InvalidImplementationError(raw []byte) (*RNMessageExecutorV1ERC1967InvalidImplementation, error) {
	out := new(RNMessageExecutorV1ERC1967InvalidImplementation)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1ERC1967NonPayable represents a ERC1967NonPayable error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1ERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func RNMessageExecutorV1ERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackERC1967NonPayableError(raw []byte) (*RNMessageExecutorV1ERC1967NonPayable, error) {
	out := new(RNMessageExecutorV1ERC1967NonPayable)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1FailedCall represents a FailedCall error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1FailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func RNMessageExecutorV1FailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackFailedCallError(raw []byte) (*RNMessageExecutorV1FailedCall, error) {
	out := new(RNMessageExecutorV1FailedCall)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1InvalidInitialization represents a InvalidInitialization error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1InvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func RNMessageExecutorV1InvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackInvalidInitializationError(raw []byte) (*RNMessageExecutorV1InvalidInitialization, error) {
	out := new(RNMessageExecutorV1InvalidInitialization)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1NotInitializing represents a NotInitializing error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1NotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func RNMessageExecutorV1NotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackNotInitializingError(raw []byte) (*RNMessageExecutorV1NotInitializing, error) {
	out := new(RNMessageExecutorV1NotInitializing)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1InvalidEndpointAddress represents a RNMessageExecutorV1__InvalidEndpointAddress error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1InvalidEndpointAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__InvalidEndpointAddress()
func RNMessageExecutorV1RNMessageExecutorV1InvalidEndpointAddressErrorID() common.Hash {
	return common.HexToHash("0x0a577ba955d5f7b5c82dfc80fcd1b20d3490efd6e3dd15a3cf2ab61582913ac5")
}

// UnpackRNMessageExecutorV1InvalidEndpointAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__InvalidEndpointAddress()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1InvalidEndpointAddressError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1InvalidEndpointAddress, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1InvalidEndpointAddress)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1InvalidEndpointAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1MessageBatchFailure represents a RNMessageExecutorV1__MessageBatchFailure error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1MessageBatchFailure struct {
	MessageId    [32]byte
	MessageIndex *big.Int
	ErrorData    []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__MessageBatchFailure(bytes32 messageId, uint256 messageIndex, bytes errorData)
func RNMessageExecutorV1RNMessageExecutorV1MessageBatchFailureErrorID() common.Hash {
	return common.HexToHash("0x47551897d68186e9787ae04cca74a64539ed64aa6c5ac1b0673be2a354971dac")
}

// UnpackRNMessageExecutorV1MessageBatchFailureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__MessageBatchFailure(bytes32 messageId, uint256 messageIndex, bytes errorData)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1MessageBatchFailureError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1MessageBatchFailure, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1MessageBatchFailure)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1MessageBatchFailure", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1MessageFailure represents a RNMessageExecutorV1__MessageFailure error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1MessageFailure struct {
	MessageId [32]byte
	ErrorData []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__MessageFailure(bytes32 messageId, bytes errorData)
func RNMessageExecutorV1RNMessageExecutorV1MessageFailureErrorID() common.Hash {
	return common.HexToHash("0x3dd87830f9e4081d282caff0bd4944a66288cf40428e0472aaf527f462e6c910")
}

// UnpackRNMessageExecutorV1MessageFailureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__MessageFailure(bytes32 messageId, bytes errorData)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1MessageFailureError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1MessageFailure, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1MessageFailure)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1MessageFailure", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1MessageIdAlreadyExecuted represents a RNMessageExecutorV1__MessageIdAlreadyExecuted error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1MessageIdAlreadyExecuted struct {
	MessageId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__MessageIdAlreadyExecuted(bytes32 messageId)
func RNMessageExecutorV1RNMessageExecutorV1MessageIdAlreadyExecutedErrorID() common.Hash {
	return common.HexToHash("0xbb29ff9f5a9983ec8b5fc5ff095a92bcad20fae0fdd73d48ad52a00125cab04b")
}

// UnpackRNMessageExecutorV1MessageIdAlreadyExecutedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__MessageIdAlreadyExecuted(bytes32 messageId)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1MessageIdAlreadyExecutedError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1MessageIdAlreadyExecuted, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1MessageIdAlreadyExecuted)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1MessageIdAlreadyExecuted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1NoContractAtAddress represents a RNMessageExecutorV1__NoContractAtAddress error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1NoContractAtAddress struct {
	ContractAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__NoContractAtAddress(address contract_address)
func RNMessageExecutorV1RNMessageExecutorV1NoContractAtAddressErrorID() common.Hash {
	return common.HexToHash("0x6c5137aa4f8d941c5780cea92039d402609b6edef8c6a718a81fe9fe1df2be21")
}

// UnpackRNMessageExecutorV1NoContractAtAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__NoContractAtAddress(address contract_address)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1NoContractAtAddressError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1NoContractAtAddress, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1NoContractAtAddress)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1NoContractAtAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNMessageExecutorV1UnauthorizedEndpoint represents a RNMessageExecutorV1__UnauthorizedEndpoint error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNMessageExecutorV1UnauthorizedEndpoint struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNMessageExecutorV1__UnauthorizedEndpoint(address caller)
func RNMessageExecutorV1RNMessageExecutorV1UnauthorizedEndpointErrorID() common.Hash {
	return common.HexToHash("0xd1f8bad4f965b134722e9de91c6b900308e6611a188a5f64668857ca6881c607")
}

// UnpackRNMessageExecutorV1UnauthorizedEndpointError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNMessageExecutorV1__UnauthorizedEndpoint(address caller)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNMessageExecutorV1UnauthorizedEndpointError(raw []byte) (*RNMessageExecutorV1RNMessageExecutorV1UnauthorizedEndpoint, error) {
	out := new(RNMessageExecutorV1RNMessageExecutorV1UnauthorizedEndpoint)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNMessageExecutorV1UnauthorizedEndpoint", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNReentrancyGuardV1ReceiveReentrancy represents a RNReentrancyGuardV1__ReceiveReentrancy error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNReentrancyGuardV1ReceiveReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNReentrancyGuardV1__ReceiveReentrancy()
func RNMessageExecutorV1RNReentrancyGuardV1ReceiveReentrancyErrorID() common.Hash {
	return common.HexToHash("0x17326c63e33303d77a4e1a51b7a04424a96a05b70c445492bc1f97ec9e33b2e8")
}

// UnpackRNReentrancyGuardV1ReceiveReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNReentrancyGuardV1__ReceiveReentrancy()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNReentrancyGuardV1ReceiveReentrancyError(raw []byte) (*RNMessageExecutorV1RNReentrancyGuardV1ReceiveReentrancy, error) {
	out := new(RNMessageExecutorV1RNReentrancyGuardV1ReceiveReentrancy)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNReentrancyGuardV1ReceiveReentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RNReentrancyGuardV1SendReentrancy represents a RNReentrancyGuardV1__SendReentrancy error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RNReentrancyGuardV1SendReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RNReentrancyGuardV1__SendReentrancy()
func RNMessageExecutorV1RNReentrancyGuardV1SendReentrancyErrorID() common.Hash {
	return common.HexToHash("0x5c3d889ea3f0e73f68a749d7795ddad4411fe30d6bf3a5394d8355213c911258")
}

// UnpackRNReentrancyGuardV1SendReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RNReentrancyGuardV1__SendReentrancy()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRNReentrancyGuardV1SendReentrancyError(raw []byte) (*RNMessageExecutorV1RNReentrancyGuardV1SendReentrancy, error) {
	out := new(RNMessageExecutorV1RNReentrancyGuardV1SendReentrancy)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RNReentrancyGuardV1SendReentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RaylsAccessManagedContractPaused represents a RaylsAccessManaged__ContractPaused error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RaylsAccessManagedContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func RNMessageExecutorV1RaylsAccessManagedContractPausedErrorID() common.Hash {
	return common.HexToHash("0xcc9855adb54c3eae80b5cb29441d3eb81c3fcdc928abeaba0ac7979b3072c3ab")
}

// UnpackRaylsAccessManagedContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__ContractPaused()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRaylsAccessManagedContractPausedError(raw []byte) (*RNMessageExecutorV1RaylsAccessManagedContractPaused, error) {
	out := new(RNMessageExecutorV1RaylsAccessManagedContractPaused)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RaylsAccessManagedInvalidAuthority represents a RaylsAccessManaged__InvalidAuthority error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RaylsAccessManagedInvalidAuthority struct {
	Authority common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func RNMessageExecutorV1RaylsAccessManagedInvalidAuthorityErrorID() common.Hash {
	return common.HexToHash("0x894403477abbba3df1c9b757907cda2d4b4d9ba52242d5e6bdae2dc6a343662b")
}

// UnpackRaylsAccessManagedInvalidAuthorityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__InvalidAuthority(address authority)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRaylsAccessManagedInvalidAuthorityError(raw []byte) (*RNMessageExecutorV1RaylsAccessManagedInvalidAuthority, error) {
	out := new(RNMessageExecutorV1RaylsAccessManagedInvalidAuthority)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedInvalidAuthority", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RaylsAccessManagedMustSchedule represents a RaylsAccessManaged__MustSchedule error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RaylsAccessManagedMustSchedule struct {
	Caller common.Address
	Delay  uint32
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func RNMessageExecutorV1RaylsAccessManagedMustScheduleErrorID() common.Hash {
	return common.HexToHash("0xa42687899abecbc3886d9231154d03927d3e61df458dda1e50129311f272684f")
}

// UnpackRaylsAccessManagedMustScheduleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__MustSchedule(address caller, uint32 delay)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRaylsAccessManagedMustScheduleError(raw []byte) (*RNMessageExecutorV1RaylsAccessManagedMustSchedule, error) {
	out := new(RNMessageExecutorV1RaylsAccessManagedMustSchedule)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedMustSchedule", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1RaylsAccessManagedUnauthorized represents a RaylsAccessManaged__Unauthorized error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1RaylsAccessManagedUnauthorized struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func RNMessageExecutorV1RaylsAccessManagedUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0xbb34f40c15462e9dd86b312158515c7ee80e2406cfa90a92a6fb6fad1bc48a48")
}

// UnpackRaylsAccessManagedUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RaylsAccessManaged__Unauthorized(address caller)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackRaylsAccessManagedUnauthorizedError(raw []byte) (*RNMessageExecutorV1RaylsAccessManagedUnauthorized, error) {
	out := new(RNMessageExecutorV1RaylsAccessManagedUnauthorized)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "RaylsAccessManagedUnauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1UUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1UUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func RNMessageExecutorV1UUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*RNMessageExecutorV1UUPSUnauthorizedCallContext, error) {
	out := new(RNMessageExecutorV1UUPSUnauthorizedCallContext)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// RNMessageExecutorV1UUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the RNMessageExecutorV1 contract.
type RNMessageExecutorV1UUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func RNMessageExecutorV1UUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (rNMessageExecutorV1 *RNMessageExecutorV1) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*RNMessageExecutorV1UUPSUnsupportedProxiableUUID, error) {
	out := new(RNMessageExecutorV1UUPSUnsupportedProxiableUUID)
	if err := rNMessageExecutorV1.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
