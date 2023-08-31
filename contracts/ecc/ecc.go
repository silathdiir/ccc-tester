// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EccMetaData contains all meta data concerning the Ecc contract.
var EccMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p0\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"p1\",\"type\":\"uint256[2]\"}],\"name\":\"ecAdd\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecAdds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"ecMul\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"retP\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecMuls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecPairings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"ecRecovers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a867d0ca": "ecAdd(uint256[2],uint256[2])",
		"7de7133d": "ecAdds(uint256)",
		"71de3724": "ecMul(uint256[2],uint256)",
		"f602cf71": "ecMuls(uint256)",
		"459be292": "ecPairings(uint256)",
		"112355a4": "ecRecovers(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610c7f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063112355a414610067578063459be2921461007c57806371de37241461008f5780637de7133d146100b8578063a867d0ca146100cb578063f602cf71146100de575b600080fd5b61007a610075366004610a8f565b6100f1565b005b61007a61008a366004610a8f565b6101d3565b6100a261009d366004610b3c565b6105bd565b6040516100af9190610b67565b60405180910390f35b61007a6100c6366004610a8f565b61062e565b6100a26100d9366004610b98565b61067c565b61007a6100ec366004610a8f565b610719565b7fb94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde97fe742ff452d41413616a5bf43fe15dd88294e983d3d36206c2712f39083d638bd601b7fe0a0fc89be718fbc1033e1d30d78be1c68081562ed2e97af876f286f3453231d60005b858110156101cb5760408051600081526020810180835287905260ff851691810191909152606081018590526080810183905260019060a0016020604051602081039080840390855afa1580156101b4573d6000803e3d6000fd5b508291506101c3905081610be3565b915050610159565b505050505050565b60408051600280825260608201909252600091816020015b60408051808201909152600080825260208201528152602001906001900390816101eb5750506040805160028082526060820190925291925060009190602082015b610235610a2e565b81526020019060019003908161022d57905050905060018260008151811061025f5761025f610bfc565b6020026020010151600001818152505060028260008151811061028457610284610bfc565b602002602001015160200181815250507f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed816000815181106102c8576102c8610bfc565b60209081029190910181015151015280517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c290829060009061030c5761030c610bfc565b6020908102919091010151515280517f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa90829060009061034e5761034e610bfc565b60200260200101516020015160016002811061036c5761036c610bfc565b602002015280517f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b9082906000906103a6576103a6610bfc565b6020026020010151602001516000600281106103c4576103c4610bfc565b602002015281517f1aa125a22bd902874034e67868aed40267e5575d5919677987e3bc6dd42a32fe908390600190811061040057610400610bfc565b602002602001015160000181815250507f1bacc186725464068956d9a191455c2d6f6db282d83645c610510d8d4efbaee08260018151811061044457610444610bfc565b602002602001015160200181815250507f1b7734c80605f71f1e2de61e998ce5854ff2abebb76537c3d67e50d71422a8528160018151811061048857610488610bfc565b60209081029190910181015151015280517f10d5a1e34b2388a5ebe266033a5e0e63c89084203784da0c6bd9b052a78a2cac90829060019081106104ce576104ce610bfc565b6020908102919091010151515280517f275739c5c2cdbc72e37c689e2ab441ea76c1d284b9c46ae8f5c42ead937819e1908290600190811061051257610512610bfc565b60200260200101516020015160016002811061053057610530610bfc565b602002015280517f018de34c5b7c3d3d75428bbe050f1449ea3d9961d563291f307a1874f7332e65908290600190811061056c5761056c610bfc565b60200260200101516020015160006002811061058a5761058a610bfc565b602002015260005b838110156105b7576105a4838361075a565b50806105af81610be3565b915050610592565b50505050565b6105c5610a53565b60006040518060600160405280856000600281106105e5576105e5610bfc565b602002015181526020018560016002811061060257610602610bfc565b602002015181526020018481525090506040826060836007600019fa61062757600080fd5b5092915050565b610636610a53565b6001815260026020820152610649610a53565b600181526002602082015260005b838110156105b757610669838361067c565b508061067481610be3565b915050610657565b610684610a53565b60006040518060800160405280856000600281106106a4576106a4610bfc565b60200201518152602001856001600281106106c1576106c1610bfc565b60200201518152602001846000600281106106de576106de610bfc565b60200201518152602001846001600281106106fb576106fb610bfc565b6020020151905290506040826080836006600019fa61062757600080fd5b610721610a53565b600181526002602082015260005b82811015610755576107428260036105bd565b508061074d81610be3565b91505061072f565b505050565b6000808351600661076b9190610c12565b905060008167ffffffffffffffff81111561078857610788610aa8565b6040519080825280602002602001820160405280156107b1578160200160208202803683370190505b5090506107bc610a71565b600085518751146107cc57600080fd5b60005b8751811015610a05578781815181106107ea576107ea610bfc565b602002602001015160000151848260066108049190610c12565b61080f906000610c31565b8151811061081f5761081f610bfc565b60200260200101818152505087818151811061083d5761083d610bfc565b602002602001015160200151848260066108579190610c12565b610862906001610c31565b8151811061087257610872610bfc565b60200260200101818152505086818151811061089057610890610bfc565b60209081029190910101515151846108a9836006610c12565b6108b4906002610c31565b815181106108c4576108c4610bfc565b6020026020010181815250508681815181106108e2576108e2610bfc565b602090810291909101810151510151846108fd836006610c12565b610908906003610c31565b8151811061091857610918610bfc565b60200260200101818152505086818151811061093657610936610bfc565b60200260200101516020015160006002811061095457610954610bfc565b602002015184610965836006610c12565b610970906004610c31565b8151811061098057610980610bfc565b60200260200101818152505086818151811061099e5761099e610bfc565b6020026020010151602001516001600281106109bc576109bc610bfc565b6020020151846109cd836006610c12565b6109d8906005610c31565b815181106109e8576109e8610bfc565b6020908102919091010152806109fd81610be3565b9150506107cf565b50602082602086026020860160085afa905080610a2157600080fd5b5051151595945050505050565b6040518060400160405280610a41610a53565b8152602001610a4e610a53565b905290565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b600060208284031215610aa157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610acf57600080fd5b6040516040810181811067ffffffffffffffff82111715610b0057634e487b7160e01b600052604160045260246000fd5b8060405250806040840185811115610b1757600080fd5b845b81811015610b31578035835260209283019201610b19565b509195945050505050565b60008060608385031215610b4f57600080fd5b610b598484610abe565b946040939093013593505050565b60408101818360005b6002811015610b8f578151835260209283019290910190600101610b70565b50505092915050565b60008060808385031215610bab57600080fd5b610bb58484610abe565b9150610bc48460408501610abe565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b600060018201610bf557610bf5610bcd565b5060010190565b634e487b7160e01b600052603260045260246000fd5b6000816000190483118215151615610c2c57610c2c610bcd565b500290565b60008219821115610c4457610c44610bcd565b50019056fea2646970667358221220e1d05de6dc17217060252992e1c31cd0cd5cf7b830057f590f61e5d73e12b72464736f6c634300080f0033",
}

// EccABI is the input ABI used to generate the binding from.
// Deprecated: Use EccMetaData.ABI instead.
var EccABI = EccMetaData.ABI

// Deprecated: Use EccMetaData.Sigs instead.
// EccFuncSigs maps the 4-byte function signature to its string representation.
var EccFuncSigs = EccMetaData.Sigs

// EccBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EccMetaData.Bin instead.
var EccBin = EccMetaData.Bin

// DeployEcc deploys a new Ethereum contract, binding an instance of Ecc to it.
func DeployEcc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ecc, error) {
	parsed, err := EccMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EccBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecc{EccCaller: EccCaller{contract: contract}, EccTransactor: EccTransactor{contract: contract}, EccFilterer: EccFilterer{contract: contract}}, nil
}

// Ecc is an auto generated Go binding around an Ethereum contract.
type Ecc struct {
	EccCaller     // Read-only binding to the contract
	EccTransactor // Write-only binding to the contract
	EccFilterer   // Log filterer for contract events
}

// EccCaller is an auto generated read-only Go binding around an Ethereum contract.
type EccCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EccTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EccFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EccSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EccSession struct {
	Contract     *Ecc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EccCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EccCallerSession struct {
	Contract *EccCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EccTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EccTransactorSession struct {
	Contract     *EccTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EccRaw is an auto generated low-level Go binding around an Ethereum contract.
type EccRaw struct {
	Contract *Ecc // Generic contract binding to access the raw methods on
}

// EccCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EccCallerRaw struct {
	Contract *EccCaller // Generic read-only contract binding to access the raw methods on
}

// EccTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EccTransactorRaw struct {
	Contract *EccTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcc creates a new instance of Ecc, bound to a specific deployed contract.
func NewEcc(address common.Address, backend bind.ContractBackend) (*Ecc, error) {
	contract, err := bindEcc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecc{EccCaller: EccCaller{contract: contract}, EccTransactor: EccTransactor{contract: contract}, EccFilterer: EccFilterer{contract: contract}}, nil
}

// NewEccCaller creates a new read-only instance of Ecc, bound to a specific deployed contract.
func NewEccCaller(address common.Address, caller bind.ContractCaller) (*EccCaller, error) {
	contract, err := bindEcc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EccCaller{contract: contract}, nil
}

// NewEccTransactor creates a new write-only instance of Ecc, bound to a specific deployed contract.
func NewEccTransactor(address common.Address, transactor bind.ContractTransactor) (*EccTransactor, error) {
	contract, err := bindEcc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EccTransactor{contract: contract}, nil
}

// NewEccFilterer creates a new log filterer instance of Ecc, bound to a specific deployed contract.
func NewEccFilterer(address common.Address, filterer bind.ContractFilterer) (*EccFilterer, error) {
	contract, err := bindEcc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EccFilterer{contract: contract}, nil
}

// bindEcc binds a generic wrapper to an already deployed contract.
func bindEcc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EccABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecc *EccRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecc.Contract.EccCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecc *EccRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecc.Contract.EccTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecc *EccRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecc.Contract.EccTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecc *EccCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecc *EccTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecc *EccTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecc.Contract.contract.Transact(opts, method, params...)
}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccCaller) EcAdd(opts *bind.CallOpts, p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Ecc.contract.Call(opts, &out, "ecAdd", p0, p1)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccSession) EcAdd(p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcAdd(&_Ecc.CallOpts, p0, p1)
}

// EcAdd is a free data retrieval call binding the contract method 0xa867d0ca.
//
// Solidity: function ecAdd(uint256[2] p0, uint256[2] p1) view returns(uint256[2] retP)
func (_Ecc *EccCallerSession) EcAdd(p0 [2]*big.Int, p1 [2]*big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcAdd(&_Ecc.CallOpts, p0, p1)
}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccCaller) EcMul(opts *bind.CallOpts, p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Ecc.contract.Call(opts, &out, "ecMul", p, s)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccSession) EcMul(p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcMul(&_Ecc.CallOpts, p, s)
}

// EcMul is a free data retrieval call binding the contract method 0x71de3724.
//
// Solidity: function ecMul(uint256[2] p, uint256 s) view returns(uint256[2] retP)
func (_Ecc *EccCallerSession) EcMul(p [2]*big.Int, s *big.Int) ([2]*big.Int, error) {
	return _Ecc.Contract.EcMul(&_Ecc.CallOpts, p, s)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccTransactor) EcAdds(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecAdds", n)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccSession) EcAdds(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcAdds(&_Ecc.TransactOpts, n)
}

// EcAdds is a paid mutator transaction binding the contract method 0x7de7133d.
//
// Solidity: function ecAdds(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcAdds(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcAdds(&_Ecc.TransactOpts, n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccTransactor) EcMuls(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecMuls", n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccSession) EcMuls(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcMuls(&_Ecc.TransactOpts, n)
}

// EcMuls is a paid mutator transaction binding the contract method 0xf602cf71.
//
// Solidity: function ecMuls(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcMuls(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcMuls(&_Ecc.TransactOpts, n)
}

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccTransactor) EcPairings(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecPairings", n)
}

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccSession) EcPairings(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcPairings(&_Ecc.TransactOpts, n)
}

// EcPairings is a paid mutator transaction binding the contract method 0x459be292.
//
// Solidity: function ecPairings(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcPairings(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcPairings(&_Ecc.TransactOpts, n)
}

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccTransactor) EcRecovers(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Ecc.contract.Transact(opts, "ecRecovers", n)
}

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccSession) EcRecovers(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcRecovers(&_Ecc.TransactOpts, n)
}

// EcRecovers is a paid mutator transaction binding the contract method 0x112355a4.
//
// Solidity: function ecRecovers(uint256 n) returns()
func (_Ecc *EccTransactorSession) EcRecovers(n *big.Int) (*types.Transaction, error) {
	return _Ecc.Contract.EcRecovers(&_Ecc.TransactOpts, n)
}
