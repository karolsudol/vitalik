// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maker

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// TravelSaverPaymentPlan is an auto generated low-level Go binding around an user-defined struct.
type TravelSaverPaymentPlan struct {
	TravelPlanID       *big.Int
	ID                 *big.Int
	TotalAmount        *big.Int
	AmountSent         *big.Int
	AmountPerInterval  *big.Int
	TotalIntervals     *big.Int
	IntervalsProcessed *big.Int
	NextTransferOn     *big.Int
	Interval           *big.Int
	Sender             common.Address
	Alive              bool
}

// TravelSaverTravelPlan is an auto generated low-level Go binding around an user-defined struct.
type TravelSaverTravelPlan struct {
	Owner             common.Address
	ID                *big.Int
	OperatorPlanID    *big.Int
	OperatorUserID    *big.Int
	ContributedAmount *big.Int
	CreatedAt         *big.Int
	ClaimedAt         *big.Int
	Claimed           bool
}

// TravelSaverMetaData contains all meta data concerning the TravelSaver contract.
var TravelSaverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorWallet_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalsProcessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextTransferOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"alive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTravelSaver.PaymentPlan\",\"name\":\"paymentPlan\",\"type\":\"tuple\"}],\"name\":\"CancelPaymentPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"ClaimTravelPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ContributeToTravelPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalsProcessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextTransferOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"alive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTravelSaver.PaymentPlan\",\"name\":\"paymentPlan\",\"type\":\"tuple\"}],\"name\":\"CreatedPaymentPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorUserID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTravelSaver.TravelPlan\",\"name\":\"travelPlan\",\"type\":\"tuple\"}],\"name\":\"CreatedTravelPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalsProcessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextTransferOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"alive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTravelSaver.PaymentPlan\",\"name\":\"paymentPlan\",\"type\":\"tuple\"}],\"name\":\"EndPaymentPlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"intervalNo\",\"type\":\"uint256\"}],\"name\":\"PaymentPlanIntervalEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"callableOn\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intervalNo\",\"type\":\"uint256\"}],\"name\":\"StartPaymentPlanInterval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"cancelPaymentPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"claimTravelPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"contributeToTravelPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalLength\",\"type\":\"uint256\"}],\"name\":\"createPaymentPlan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorPlanID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorUserID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalLength\",\"type\":\"uint256\"}],\"name\":\"createTravelPaymentPlan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentPlanID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorPlanID_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorUserID_\",\"type\":\"uint256\"}],\"name\":\"createTravelPlan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"getPaymentPlanDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"travelPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalIntervals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intervalsProcessed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextTransferOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"alive\",\"type\":\"bool\"}],\"internalType\":\"structTravelSaver.PaymentPlan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"getTravelPlanDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorUserID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structTravelSaver.TravelPlan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"runInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"IDs\",\"type\":\"uint256[]\"}],\"name\":\"runIntervals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"travelPlans\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorPlanID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorUserID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contributedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TravelSaverABI is the input ABI used to generate the binding from.
// Deprecated: Use TravelSaverMetaData.ABI instead.
var TravelSaverABI = TravelSaverMetaData.ABI

// TravelSaver is an auto generated Go binding around an Ethereum contract.
type TravelSaver struct {
	TravelSaverCaller     // Read-only binding to the contract
	TravelSaverTransactor // Write-only binding to the contract
	TravelSaverFilterer   // Log filterer for contract events
}

// TravelSaverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TravelSaverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TravelSaverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TravelSaverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TravelSaverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TravelSaverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TravelSaverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TravelSaverSession struct {
	Contract     *TravelSaver      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TravelSaverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TravelSaverCallerSession struct {
	Contract *TravelSaverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TravelSaverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TravelSaverTransactorSession struct {
	Contract     *TravelSaverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TravelSaverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TravelSaverRaw struct {
	Contract *TravelSaver // Generic contract binding to access the raw methods on
}

// TravelSaverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TravelSaverCallerRaw struct {
	Contract *TravelSaverCaller // Generic read-only contract binding to access the raw methods on
}

// TravelSaverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TravelSaverTransactorRaw struct {
	Contract *TravelSaverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTravelSaver creates a new instance of TravelSaver, bound to a specific deployed contract.
func NewTravelSaver(address common.Address, backend bind.ContractBackend) (*TravelSaver, error) {
	contract, err := bindTravelSaver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TravelSaver{TravelSaverCaller: TravelSaverCaller{contract: contract}, TravelSaverTransactor: TravelSaverTransactor{contract: contract}, TravelSaverFilterer: TravelSaverFilterer{contract: contract}}, nil
}

// NewTravelSaverCaller creates a new read-only instance of TravelSaver, bound to a specific deployed contract.
func NewTravelSaverCaller(address common.Address, caller bind.ContractCaller) (*TravelSaverCaller, error) {
	contract, err := bindTravelSaver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TravelSaverCaller{contract: contract}, nil
}

// NewTravelSaverTransactor creates a new write-only instance of TravelSaver, bound to a specific deployed contract.
func NewTravelSaverTransactor(address common.Address, transactor bind.ContractTransactor) (*TravelSaverTransactor, error) {
	contract, err := bindTravelSaver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TravelSaverTransactor{contract: contract}, nil
}

// NewTravelSaverFilterer creates a new log filterer instance of TravelSaver, bound to a specific deployed contract.
func NewTravelSaverFilterer(address common.Address, filterer bind.ContractFilterer) (*TravelSaverFilterer, error) {
	contract, err := bindTravelSaver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TravelSaverFilterer{contract: contract}, nil
}

// bindTravelSaver binds a generic wrapper to an already deployed contract.
func bindTravelSaver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TravelSaverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TravelSaver *TravelSaverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TravelSaver.Contract.TravelSaverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TravelSaver *TravelSaverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TravelSaver.Contract.TravelSaverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TravelSaver *TravelSaverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TravelSaver.Contract.TravelSaverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TravelSaver *TravelSaverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TravelSaver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TravelSaver *TravelSaverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TravelSaver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TravelSaver *TravelSaverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TravelSaver.Contract.contract.Transact(opts, method, params...)
}

// ContributedAmount is a free data retrieval call binding the contract method 0x6dcfb52b.
//
// Solidity: function contributedAmount(uint256 , address ) view returns(uint256)
func (_TravelSaver *TravelSaverCaller) ContributedAmount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "contributedAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContributedAmount is a free data retrieval call binding the contract method 0x6dcfb52b.
//
// Solidity: function contributedAmount(uint256 , address ) view returns(uint256)
func (_TravelSaver *TravelSaverSession) ContributedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _TravelSaver.Contract.ContributedAmount(&_TravelSaver.CallOpts, arg0, arg1)
}

// ContributedAmount is a free data retrieval call binding the contract method 0x6dcfb52b.
//
// Solidity: function contributedAmount(uint256 , address ) view returns(uint256)
func (_TravelSaver *TravelSaverCallerSession) ContributedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _TravelSaver.Contract.ContributedAmount(&_TravelSaver.CallOpts, arg0, arg1)
}

// GetPaymentPlanDetails is a free data retrieval call binding the contract method 0xd955bdc9.
//
// Solidity: function getPaymentPlanDetails(uint256 ID) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_TravelSaver *TravelSaverCaller) GetPaymentPlanDetails(opts *bind.CallOpts, ID *big.Int) (TravelSaverPaymentPlan, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "getPaymentPlanDetails", ID)

	if err != nil {
		return *new(TravelSaverPaymentPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(TravelSaverPaymentPlan)).(*TravelSaverPaymentPlan)

	return out0, err

}

// GetPaymentPlanDetails is a free data retrieval call binding the contract method 0xd955bdc9.
//
// Solidity: function getPaymentPlanDetails(uint256 ID) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_TravelSaver *TravelSaverSession) GetPaymentPlanDetails(ID *big.Int) (TravelSaverPaymentPlan, error) {
	return _TravelSaver.Contract.GetPaymentPlanDetails(&_TravelSaver.CallOpts, ID)
}

// GetPaymentPlanDetails is a free data retrieval call binding the contract method 0xd955bdc9.
//
// Solidity: function getPaymentPlanDetails(uint256 ID) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_TravelSaver *TravelSaverCallerSession) GetPaymentPlanDetails(ID *big.Int) (TravelSaverPaymentPlan, error) {
	return _TravelSaver.Contract.GetPaymentPlanDetails(&_TravelSaver.CallOpts, ID)
}

// GetTravelPlanDetails is a free data retrieval call binding the contract method 0xddbf07f9.
//
// Solidity: function getTravelPlanDetails(uint256 ID) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_TravelSaver *TravelSaverCaller) GetTravelPlanDetails(opts *bind.CallOpts, ID *big.Int) (TravelSaverTravelPlan, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "getTravelPlanDetails", ID)

	if err != nil {
		return *new(TravelSaverTravelPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(TravelSaverTravelPlan)).(*TravelSaverTravelPlan)

	return out0, err

}

// GetTravelPlanDetails is a free data retrieval call binding the contract method 0xddbf07f9.
//
// Solidity: function getTravelPlanDetails(uint256 ID) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_TravelSaver *TravelSaverSession) GetTravelPlanDetails(ID *big.Int) (TravelSaverTravelPlan, error) {
	return _TravelSaver.Contract.GetTravelPlanDetails(&_TravelSaver.CallOpts, ID)
}

// GetTravelPlanDetails is a free data retrieval call binding the contract method 0xddbf07f9.
//
// Solidity: function getTravelPlanDetails(uint256 ID) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool))
func (_TravelSaver *TravelSaverCallerSession) GetTravelPlanDetails(ID *big.Int) (TravelSaverTravelPlan, error) {
	return _TravelSaver.Contract.GetTravelPlanDetails(&_TravelSaver.CallOpts, ID)
}

// OperatorWallet is a free data retrieval call binding the contract method 0xf90ef4ac.
//
// Solidity: function operatorWallet() view returns(address)
func (_TravelSaver *TravelSaverCaller) OperatorWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "operatorWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorWallet is a free data retrieval call binding the contract method 0xf90ef4ac.
//
// Solidity: function operatorWallet() view returns(address)
func (_TravelSaver *TravelSaverSession) OperatorWallet() (common.Address, error) {
	return _TravelSaver.Contract.OperatorWallet(&_TravelSaver.CallOpts)
}

// OperatorWallet is a free data retrieval call binding the contract method 0xf90ef4ac.
//
// Solidity: function operatorWallet() view returns(address)
func (_TravelSaver *TravelSaverCallerSession) OperatorWallet() (common.Address, error) {
	return _TravelSaver.Contract.OperatorWallet(&_TravelSaver.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TravelSaver *TravelSaverCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TravelSaver *TravelSaverSession) Token() (common.Address, error) {
	return _TravelSaver.Contract.Token(&_TravelSaver.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TravelSaver *TravelSaverCallerSession) Token() (common.Address, error) {
	return _TravelSaver.Contract.Token(&_TravelSaver.CallOpts)
}

// TravelPlans is a free data retrieval call binding the contract method 0x8818e165.
//
// Solidity: function travelPlans(uint256 ) view returns(address owner, uint256 ID, uint256 operatorPlanID, uint256 operatorUserID, uint256 contributedAmount, uint256 createdAt, uint256 claimedAt, bool claimed)
func (_TravelSaver *TravelSaverCaller) TravelPlans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner             common.Address
	ID                *big.Int
	OperatorPlanID    *big.Int
	OperatorUserID    *big.Int
	ContributedAmount *big.Int
	CreatedAt         *big.Int
	ClaimedAt         *big.Int
	Claimed           bool
}, error) {
	var out []interface{}
	err := _TravelSaver.contract.Call(opts, &out, "travelPlans", arg0)

	outstruct := new(struct {
		Owner             common.Address
		ID                *big.Int
		OperatorPlanID    *big.Int
		OperatorUserID    *big.Int
		ContributedAmount *big.Int
		CreatedAt         *big.Int
		ClaimedAt         *big.Int
		Claimed           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OperatorPlanID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OperatorUserID = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ContributedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ClaimedAt = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// TravelPlans is a free data retrieval call binding the contract method 0x8818e165.
//
// Solidity: function travelPlans(uint256 ) view returns(address owner, uint256 ID, uint256 operatorPlanID, uint256 operatorUserID, uint256 contributedAmount, uint256 createdAt, uint256 claimedAt, bool claimed)
func (_TravelSaver *TravelSaverSession) TravelPlans(arg0 *big.Int) (struct {
	Owner             common.Address
	ID                *big.Int
	OperatorPlanID    *big.Int
	OperatorUserID    *big.Int
	ContributedAmount *big.Int
	CreatedAt         *big.Int
	ClaimedAt         *big.Int
	Claimed           bool
}, error) {
	return _TravelSaver.Contract.TravelPlans(&_TravelSaver.CallOpts, arg0)
}

// TravelPlans is a free data retrieval call binding the contract method 0x8818e165.
//
// Solidity: function travelPlans(uint256 ) view returns(address owner, uint256 ID, uint256 operatorPlanID, uint256 operatorUserID, uint256 contributedAmount, uint256 createdAt, uint256 claimedAt, bool claimed)
func (_TravelSaver *TravelSaverCallerSession) TravelPlans(arg0 *big.Int) (struct {
	Owner             common.Address
	ID                *big.Int
	OperatorPlanID    *big.Int
	OperatorUserID    *big.Int
	ContributedAmount *big.Int
	CreatedAt         *big.Int
	ClaimedAt         *big.Int
	Claimed           bool
}, error) {
	return _TravelSaver.Contract.TravelPlans(&_TravelSaver.CallOpts, arg0)
}

// CancelPaymentPlan is a paid mutator transaction binding the contract method 0x90513e1a.
//
// Solidity: function cancelPaymentPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactor) CancelPaymentPlan(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "cancelPaymentPlan", ID)
}

// CancelPaymentPlan is a paid mutator transaction binding the contract method 0x90513e1a.
//
// Solidity: function cancelPaymentPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverSession) CancelPaymentPlan(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CancelPaymentPlan(&_TravelSaver.TransactOpts, ID)
}

// CancelPaymentPlan is a paid mutator transaction binding the contract method 0x90513e1a.
//
// Solidity: function cancelPaymentPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactorSession) CancelPaymentPlan(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CancelPaymentPlan(&_TravelSaver.TransactOpts, ID)
}

// ClaimTravelPlan is a paid mutator transaction binding the contract method 0xcf67dc41.
//
// Solidity: function claimTravelPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactor) ClaimTravelPlan(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "claimTravelPlan", ID)
}

// ClaimTravelPlan is a paid mutator transaction binding the contract method 0xcf67dc41.
//
// Solidity: function claimTravelPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverSession) ClaimTravelPlan(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.ClaimTravelPlan(&_TravelSaver.TransactOpts, ID)
}

// ClaimTravelPlan is a paid mutator transaction binding the contract method 0xcf67dc41.
//
// Solidity: function claimTravelPlan(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactorSession) ClaimTravelPlan(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.ClaimTravelPlan(&_TravelSaver.TransactOpts, ID)
}

// ContributeToTravelPlan is a paid mutator transaction binding the contract method 0x2b06fa74.
//
// Solidity: function contributeToTravelPlan(uint256 ID, uint256 amount) returns()
func (_TravelSaver *TravelSaverTransactor) ContributeToTravelPlan(opts *bind.TransactOpts, ID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "contributeToTravelPlan", ID, amount)
}

// ContributeToTravelPlan is a paid mutator transaction binding the contract method 0x2b06fa74.
//
// Solidity: function contributeToTravelPlan(uint256 ID, uint256 amount) returns()
func (_TravelSaver *TravelSaverSession) ContributeToTravelPlan(ID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.ContributeToTravelPlan(&_TravelSaver.TransactOpts, ID, amount)
}

// ContributeToTravelPlan is a paid mutator transaction binding the contract method 0x2b06fa74.
//
// Solidity: function contributeToTravelPlan(uint256 ID, uint256 amount) returns()
func (_TravelSaver *TravelSaverTransactorSession) ContributeToTravelPlan(ID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.ContributeToTravelPlan(&_TravelSaver.TransactOpts, ID, amount)
}

// CreatePaymentPlan is a paid mutator transaction binding the contract method 0x223f83e2.
//
// Solidity: function createPaymentPlan(uint256 _travelPlanID, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256)
func (_TravelSaver *TravelSaverTransactor) CreatePaymentPlan(opts *bind.TransactOpts, _travelPlanID *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "createPaymentPlan", _travelPlanID, amountPerInterval, totalIntervals, intervalLength)
}

// CreatePaymentPlan is a paid mutator transaction binding the contract method 0x223f83e2.
//
// Solidity: function createPaymentPlan(uint256 _travelPlanID, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256)
func (_TravelSaver *TravelSaverSession) CreatePaymentPlan(_travelPlanID *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreatePaymentPlan(&_TravelSaver.TransactOpts, _travelPlanID, amountPerInterval, totalIntervals, intervalLength)
}

// CreatePaymentPlan is a paid mutator transaction binding the contract method 0x223f83e2.
//
// Solidity: function createPaymentPlan(uint256 _travelPlanID, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256)
func (_TravelSaver *TravelSaverTransactorSession) CreatePaymentPlan(_travelPlanID *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreatePaymentPlan(&_TravelSaver.TransactOpts, _travelPlanID, amountPerInterval, totalIntervals, intervalLength)
}

// CreateTravelPaymentPlan is a paid mutator transaction binding the contract method 0xaaec77e5.
//
// Solidity: function createTravelPaymentPlan(uint256 operatorPlanID_, uint256 operatorUserID_, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256 travelPlanID, uint256 paymentPlanID)
func (_TravelSaver *TravelSaverTransactor) CreateTravelPaymentPlan(opts *bind.TransactOpts, operatorPlanID_ *big.Int, operatorUserID_ *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "createTravelPaymentPlan", operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
}

// CreateTravelPaymentPlan is a paid mutator transaction binding the contract method 0xaaec77e5.
//
// Solidity: function createTravelPaymentPlan(uint256 operatorPlanID_, uint256 operatorUserID_, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256 travelPlanID, uint256 paymentPlanID)
func (_TravelSaver *TravelSaverSession) CreateTravelPaymentPlan(operatorPlanID_ *big.Int, operatorUserID_ *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreateTravelPaymentPlan(&_TravelSaver.TransactOpts, operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
}

// CreateTravelPaymentPlan is a paid mutator transaction binding the contract method 0xaaec77e5.
//
// Solidity: function createTravelPaymentPlan(uint256 operatorPlanID_, uint256 operatorUserID_, uint256 amountPerInterval, uint256 totalIntervals, uint256 intervalLength) returns(uint256 travelPlanID, uint256 paymentPlanID)
func (_TravelSaver *TravelSaverTransactorSession) CreateTravelPaymentPlan(operatorPlanID_ *big.Int, operatorUserID_ *big.Int, amountPerInterval *big.Int, totalIntervals *big.Int, intervalLength *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreateTravelPaymentPlan(&_TravelSaver.TransactOpts, operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
}

// CreateTravelPlan is a paid mutator transaction binding the contract method 0x85fcc9a2.
//
// Solidity: function createTravelPlan(uint256 operatorPlanID_, uint256 operatorUserID_) returns(uint256)
func (_TravelSaver *TravelSaverTransactor) CreateTravelPlan(opts *bind.TransactOpts, operatorPlanID_ *big.Int, operatorUserID_ *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "createTravelPlan", operatorPlanID_, operatorUserID_)
}

// CreateTravelPlan is a paid mutator transaction binding the contract method 0x85fcc9a2.
//
// Solidity: function createTravelPlan(uint256 operatorPlanID_, uint256 operatorUserID_) returns(uint256)
func (_TravelSaver *TravelSaverSession) CreateTravelPlan(operatorPlanID_ *big.Int, operatorUserID_ *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreateTravelPlan(&_TravelSaver.TransactOpts, operatorPlanID_, operatorUserID_)
}

// CreateTravelPlan is a paid mutator transaction binding the contract method 0x85fcc9a2.
//
// Solidity: function createTravelPlan(uint256 operatorPlanID_, uint256 operatorUserID_) returns(uint256)
func (_TravelSaver *TravelSaverTransactorSession) CreateTravelPlan(operatorPlanID_ *big.Int, operatorUserID_ *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.CreateTravelPlan(&_TravelSaver.TransactOpts, operatorPlanID_, operatorUserID_)
}

// RunInterval is a paid mutator transaction binding the contract method 0x0cf9e555.
//
// Solidity: function runInterval(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactor) RunInterval(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "runInterval", ID)
}

// RunInterval is a paid mutator transaction binding the contract method 0x0cf9e555.
//
// Solidity: function runInterval(uint256 ID) returns()
func (_TravelSaver *TravelSaverSession) RunInterval(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.RunInterval(&_TravelSaver.TransactOpts, ID)
}

// RunInterval is a paid mutator transaction binding the contract method 0x0cf9e555.
//
// Solidity: function runInterval(uint256 ID) returns()
func (_TravelSaver *TravelSaverTransactorSession) RunInterval(ID *big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.RunInterval(&_TravelSaver.TransactOpts, ID)
}

// RunIntervals is a paid mutator transaction binding the contract method 0x4e581547.
//
// Solidity: function runIntervals(uint256[] IDs) returns()
func (_TravelSaver *TravelSaverTransactor) RunIntervals(opts *bind.TransactOpts, IDs []*big.Int) (*types.Transaction, error) {
	return _TravelSaver.contract.Transact(opts, "runIntervals", IDs)
}

// RunIntervals is a paid mutator transaction binding the contract method 0x4e581547.
//
// Solidity: function runIntervals(uint256[] IDs) returns()
func (_TravelSaver *TravelSaverSession) RunIntervals(IDs []*big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.RunIntervals(&_TravelSaver.TransactOpts, IDs)
}

// RunIntervals is a paid mutator transaction binding the contract method 0x4e581547.
//
// Solidity: function runIntervals(uint256[] IDs) returns()
func (_TravelSaver *TravelSaverTransactorSession) RunIntervals(IDs []*big.Int) (*types.Transaction, error) {
	return _TravelSaver.Contract.RunIntervals(&_TravelSaver.TransactOpts, IDs)
}

// TravelSaverCancelPaymentPlanIterator is returned from FilterCancelPaymentPlan and is used to iterate over the raw logs and unpacked data for CancelPaymentPlan events raised by the TravelSaver contract.
type TravelSaverCancelPaymentPlanIterator struct {
	Event *TravelSaverCancelPaymentPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverCancelPaymentPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverCancelPaymentPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverCancelPaymentPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverCancelPaymentPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverCancelPaymentPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverCancelPaymentPlan represents a CancelPaymentPlan event raised by the TravelSaver contract.
type TravelSaverCancelPaymentPlan struct {
	ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelPaymentPlan is a free log retrieval operation binding the contract event 0x1508b4f1d043dd53494e9f162026083edd68e32aa1e6e2d8818021c384fe080a.
//
// Solidity: event CancelPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) FilterCancelPaymentPlan(opts *bind.FilterOpts, ID []*big.Int, owner []common.Address) (*TravelSaverCancelPaymentPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "CancelPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverCancelPaymentPlanIterator{contract: _TravelSaver.contract, event: "CancelPaymentPlan", logs: logs, sub: sub}, nil
}

// WatchCancelPaymentPlan is a free log subscription operation binding the contract event 0x1508b4f1d043dd53494e9f162026083edd68e32aa1e6e2d8818021c384fe080a.
//
// Solidity: event CancelPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) WatchCancelPaymentPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverCancelPaymentPlan, ID []*big.Int, owner []common.Address) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "CancelPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverCancelPaymentPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "CancelPaymentPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelPaymentPlan is a log parse operation binding the contract event 0x1508b4f1d043dd53494e9f162026083edd68e32aa1e6e2d8818021c384fe080a.
//
// Solidity: event CancelPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) ParseCancelPaymentPlan(log types.Log) (*TravelSaverCancelPaymentPlan, error) {
	event := new(TravelSaverCancelPaymentPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "CancelPaymentPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverClaimTravelPlanIterator is returned from FilterClaimTravelPlan and is used to iterate over the raw logs and unpacked data for ClaimTravelPlan events raised by the TravelSaver contract.
type TravelSaverClaimTravelPlanIterator struct {
	Event *TravelSaverClaimTravelPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverClaimTravelPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverClaimTravelPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverClaimTravelPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverClaimTravelPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverClaimTravelPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverClaimTravelPlan represents a ClaimTravelPlan event raised by the TravelSaver contract.
type TravelSaverClaimTravelPlan struct {
	ID  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimTravelPlan is a free log retrieval operation binding the contract event 0xed3f56ffe7a310d05b1c1e6ff348c01c760245c8cbf4fd85dcd0bc1c2e26f6ff.
//
// Solidity: event ClaimTravelPlan(uint256 indexed ID)
func (_TravelSaver *TravelSaverFilterer) FilterClaimTravelPlan(opts *bind.FilterOpts, ID []*big.Int) (*TravelSaverClaimTravelPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "ClaimTravelPlan", IDRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverClaimTravelPlanIterator{contract: _TravelSaver.contract, event: "ClaimTravelPlan", logs: logs, sub: sub}, nil
}

// WatchClaimTravelPlan is a free log subscription operation binding the contract event 0xed3f56ffe7a310d05b1c1e6ff348c01c760245c8cbf4fd85dcd0bc1c2e26f6ff.
//
// Solidity: event ClaimTravelPlan(uint256 indexed ID)
func (_TravelSaver *TravelSaverFilterer) WatchClaimTravelPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverClaimTravelPlan, ID []*big.Int) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "ClaimTravelPlan", IDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverClaimTravelPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "ClaimTravelPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimTravelPlan is a log parse operation binding the contract event 0xed3f56ffe7a310d05b1c1e6ff348c01c760245c8cbf4fd85dcd0bc1c2e26f6ff.
//
// Solidity: event ClaimTravelPlan(uint256 indexed ID)
func (_TravelSaver *TravelSaverFilterer) ParseClaimTravelPlan(log types.Log) (*TravelSaverClaimTravelPlan, error) {
	event := new(TravelSaverClaimTravelPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "ClaimTravelPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverContributeToTravelPlanIterator is returned from FilterContributeToTravelPlan and is used to iterate over the raw logs and unpacked data for ContributeToTravelPlan events raised by the TravelSaver contract.
type TravelSaverContributeToTravelPlanIterator struct {
	Event *TravelSaverContributeToTravelPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverContributeToTravelPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverContributeToTravelPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverContributeToTravelPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverContributeToTravelPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverContributeToTravelPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverContributeToTravelPlan represents a ContributeToTravelPlan event raised by the TravelSaver contract.
type TravelSaverContributeToTravelPlan struct {
	ID          *big.Int
	Contributor common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContributeToTravelPlan is a free log retrieval operation binding the contract event 0x20c5b66cd5e4a4148323852541307bd30d5674f62e2247aefbd3dbae8be5e918.
//
// Solidity: event ContributeToTravelPlan(uint256 indexed ID, address indexed contributor, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) FilterContributeToTravelPlan(opts *bind.FilterOpts, ID []*big.Int, contributor []common.Address) (*TravelSaverContributeToTravelPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "ContributeToTravelPlan", IDRule, contributorRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverContributeToTravelPlanIterator{contract: _TravelSaver.contract, event: "ContributeToTravelPlan", logs: logs, sub: sub}, nil
}

// WatchContributeToTravelPlan is a free log subscription operation binding the contract event 0x20c5b66cd5e4a4148323852541307bd30d5674f62e2247aefbd3dbae8be5e918.
//
// Solidity: event ContributeToTravelPlan(uint256 indexed ID, address indexed contributor, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) WatchContributeToTravelPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverContributeToTravelPlan, ID []*big.Int, contributor []common.Address) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "ContributeToTravelPlan", IDRule, contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverContributeToTravelPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "ContributeToTravelPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContributeToTravelPlan is a log parse operation binding the contract event 0x20c5b66cd5e4a4148323852541307bd30d5674f62e2247aefbd3dbae8be5e918.
//
// Solidity: event ContributeToTravelPlan(uint256 indexed ID, address indexed contributor, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) ParseContributeToTravelPlan(log types.Log) (*TravelSaverContributeToTravelPlan, error) {
	event := new(TravelSaverContributeToTravelPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "ContributeToTravelPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverCreatedPaymentPlanIterator is returned from FilterCreatedPaymentPlan and is used to iterate over the raw logs and unpacked data for CreatedPaymentPlan events raised by the TravelSaver contract.
type TravelSaverCreatedPaymentPlanIterator struct {
	Event *TravelSaverCreatedPaymentPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverCreatedPaymentPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverCreatedPaymentPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverCreatedPaymentPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverCreatedPaymentPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverCreatedPaymentPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverCreatedPaymentPlan represents a CreatedPaymentPlan event raised by the TravelSaver contract.
type TravelSaverCreatedPaymentPlan struct {
	ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreatedPaymentPlan is a free log retrieval operation binding the contract event 0xfd671f7182c83d1e07d1bd1748354402aa8a9193629cad7a0fbaac5c8ffe2818.
//
// Solidity: event CreatedPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) FilterCreatedPaymentPlan(opts *bind.FilterOpts, ID []*big.Int, owner []common.Address) (*TravelSaverCreatedPaymentPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "CreatedPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverCreatedPaymentPlanIterator{contract: _TravelSaver.contract, event: "CreatedPaymentPlan", logs: logs, sub: sub}, nil
}

// WatchCreatedPaymentPlan is a free log subscription operation binding the contract event 0xfd671f7182c83d1e07d1bd1748354402aa8a9193629cad7a0fbaac5c8ffe2818.
//
// Solidity: event CreatedPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) WatchCreatedPaymentPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverCreatedPaymentPlan, ID []*big.Int, owner []common.Address) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "CreatedPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverCreatedPaymentPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "CreatedPaymentPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedPaymentPlan is a log parse operation binding the contract event 0xfd671f7182c83d1e07d1bd1748354402aa8a9193629cad7a0fbaac5c8ffe2818.
//
// Solidity: event CreatedPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) ParseCreatedPaymentPlan(log types.Log) (*TravelSaverCreatedPaymentPlan, error) {
	event := new(TravelSaverCreatedPaymentPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "CreatedPaymentPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverCreatedTravelPlanIterator is returned from FilterCreatedTravelPlan and is used to iterate over the raw logs and unpacked data for CreatedTravelPlan events raised by the TravelSaver contract.
type TravelSaverCreatedTravelPlanIterator struct {
	Event *TravelSaverCreatedTravelPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverCreatedTravelPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverCreatedTravelPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverCreatedTravelPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverCreatedTravelPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverCreatedTravelPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverCreatedTravelPlan represents a CreatedTravelPlan event raised by the TravelSaver contract.
type TravelSaverCreatedTravelPlan struct {
	ID         *big.Int
	Owner      common.Address
	TravelPlan TravelSaverTravelPlan
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedTravelPlan is a free log retrieval operation binding the contract event 0x71acf40162e8307c5fd5ec4e08cd622f5e5d29c2b1982af5421d95c0b1d117cf.
//
// Solidity: event CreatedTravelPlan(uint256 indexed ID, address indexed owner, (address,uint256,uint256,uint256,uint256,uint256,uint256,bool) travelPlan)
func (_TravelSaver *TravelSaverFilterer) FilterCreatedTravelPlan(opts *bind.FilterOpts, ID []*big.Int, owner []common.Address) (*TravelSaverCreatedTravelPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "CreatedTravelPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverCreatedTravelPlanIterator{contract: _TravelSaver.contract, event: "CreatedTravelPlan", logs: logs, sub: sub}, nil
}

// WatchCreatedTravelPlan is a free log subscription operation binding the contract event 0x71acf40162e8307c5fd5ec4e08cd622f5e5d29c2b1982af5421d95c0b1d117cf.
//
// Solidity: event CreatedTravelPlan(uint256 indexed ID, address indexed owner, (address,uint256,uint256,uint256,uint256,uint256,uint256,bool) travelPlan)
func (_TravelSaver *TravelSaverFilterer) WatchCreatedTravelPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverCreatedTravelPlan, ID []*big.Int, owner []common.Address) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "CreatedTravelPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverCreatedTravelPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "CreatedTravelPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedTravelPlan is a log parse operation binding the contract event 0x71acf40162e8307c5fd5ec4e08cd622f5e5d29c2b1982af5421d95c0b1d117cf.
//
// Solidity: event CreatedTravelPlan(uint256 indexed ID, address indexed owner, (address,uint256,uint256,uint256,uint256,uint256,uint256,bool) travelPlan)
func (_TravelSaver *TravelSaverFilterer) ParseCreatedTravelPlan(log types.Log) (*TravelSaverCreatedTravelPlan, error) {
	event := new(TravelSaverCreatedTravelPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "CreatedTravelPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverEndPaymentPlanIterator is returned from FilterEndPaymentPlan and is used to iterate over the raw logs and unpacked data for EndPaymentPlan events raised by the TravelSaver contract.
type TravelSaverEndPaymentPlanIterator struct {
	Event *TravelSaverEndPaymentPlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverEndPaymentPlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverEndPaymentPlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverEndPaymentPlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverEndPaymentPlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverEndPaymentPlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverEndPaymentPlan represents a EndPaymentPlan event raised by the TravelSaver contract.
type TravelSaverEndPaymentPlan struct {
	ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEndPaymentPlan is a free log retrieval operation binding the contract event 0x5e2829107023b345d1e5d6a90c0485ea550de003591f622a48dd48034b80486b.
//
// Solidity: event EndPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) FilterEndPaymentPlan(opts *bind.FilterOpts, ID []*big.Int, owner []common.Address) (*TravelSaverEndPaymentPlanIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "EndPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverEndPaymentPlanIterator{contract: _TravelSaver.contract, event: "EndPaymentPlan", logs: logs, sub: sub}, nil
}

// WatchEndPaymentPlan is a free log subscription operation binding the contract event 0x5e2829107023b345d1e5d6a90c0485ea550de003591f622a48dd48034b80486b.
//
// Solidity: event EndPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) WatchEndPaymentPlan(opts *bind.WatchOpts, sink chan<- *TravelSaverEndPaymentPlan, ID []*big.Int, owner []common.Address) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "EndPaymentPlan", IDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverEndPaymentPlan)
				if err := _TravelSaver.contract.UnpackLog(event, "EndPaymentPlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEndPaymentPlan is a log parse operation binding the contract event 0x5e2829107023b345d1e5d6a90c0485ea550de003591f622a48dd48034b80486b.
//
// Solidity: event EndPaymentPlan(uint256 indexed ID, address indexed owner, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool) paymentPlan)
func (_TravelSaver *TravelSaverFilterer) ParseEndPaymentPlan(log types.Log) (*TravelSaverEndPaymentPlan, error) {
	event := new(TravelSaverEndPaymentPlan)
	if err := _TravelSaver.contract.UnpackLog(event, "EndPaymentPlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverPaymentPlanIntervalEndedIterator is returned from FilterPaymentPlanIntervalEnded and is used to iterate over the raw logs and unpacked data for PaymentPlanIntervalEnded events raised by the TravelSaver contract.
type TravelSaverPaymentPlanIntervalEndedIterator struct {
	Event *TravelSaverPaymentPlanIntervalEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverPaymentPlanIntervalEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverPaymentPlanIntervalEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverPaymentPlanIntervalEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverPaymentPlanIntervalEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverPaymentPlanIntervalEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverPaymentPlanIntervalEnded represents a PaymentPlanIntervalEnded event raised by the TravelSaver contract.
type TravelSaverPaymentPlanIntervalEnded struct {
	ID         *big.Int
	IntervalNo *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentPlanIntervalEnded is a free log retrieval operation binding the contract event 0xbec3f63cb2d17ab9edd65a1a7ab167efd9802d37681c490ed1b215c6995bb3cb.
//
// Solidity: event PaymentPlanIntervalEnded(uint256 indexed ID, uint256 indexed intervalNo)
func (_TravelSaver *TravelSaverFilterer) FilterPaymentPlanIntervalEnded(opts *bind.FilterOpts, ID []*big.Int, intervalNo []*big.Int) (*TravelSaverPaymentPlanIntervalEndedIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var intervalNoRule []interface{}
	for _, intervalNoItem := range intervalNo {
		intervalNoRule = append(intervalNoRule, intervalNoItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "PaymentPlanIntervalEnded", IDRule, intervalNoRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverPaymentPlanIntervalEndedIterator{contract: _TravelSaver.contract, event: "PaymentPlanIntervalEnded", logs: logs, sub: sub}, nil
}

// WatchPaymentPlanIntervalEnded is a free log subscription operation binding the contract event 0xbec3f63cb2d17ab9edd65a1a7ab167efd9802d37681c490ed1b215c6995bb3cb.
//
// Solidity: event PaymentPlanIntervalEnded(uint256 indexed ID, uint256 indexed intervalNo)
func (_TravelSaver *TravelSaverFilterer) WatchPaymentPlanIntervalEnded(opts *bind.WatchOpts, sink chan<- *TravelSaverPaymentPlanIntervalEnded, ID []*big.Int, intervalNo []*big.Int) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var intervalNoRule []interface{}
	for _, intervalNoItem := range intervalNo {
		intervalNoRule = append(intervalNoRule, intervalNoItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "PaymentPlanIntervalEnded", IDRule, intervalNoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverPaymentPlanIntervalEnded)
				if err := _TravelSaver.contract.UnpackLog(event, "PaymentPlanIntervalEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentPlanIntervalEnded is a log parse operation binding the contract event 0xbec3f63cb2d17ab9edd65a1a7ab167efd9802d37681c490ed1b215c6995bb3cb.
//
// Solidity: event PaymentPlanIntervalEnded(uint256 indexed ID, uint256 indexed intervalNo)
func (_TravelSaver *TravelSaverFilterer) ParsePaymentPlanIntervalEnded(log types.Log) (*TravelSaverPaymentPlanIntervalEnded, error) {
	event := new(TravelSaverPaymentPlanIntervalEnded)
	if err := _TravelSaver.contract.UnpackLog(event, "PaymentPlanIntervalEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverStartPaymentPlanIntervalIterator is returned from FilterStartPaymentPlanInterval and is used to iterate over the raw logs and unpacked data for StartPaymentPlanInterval events raised by the TravelSaver contract.
type TravelSaverStartPaymentPlanIntervalIterator struct {
	Event *TravelSaverStartPaymentPlanInterval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverStartPaymentPlanIntervalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverStartPaymentPlanInterval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverStartPaymentPlanInterval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverStartPaymentPlanIntervalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverStartPaymentPlanIntervalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverStartPaymentPlanInterval represents a StartPaymentPlanInterval event raised by the TravelSaver contract.
type TravelSaverStartPaymentPlanInterval struct {
	ID         *big.Int
	CallableOn *big.Int
	Amount     *big.Int
	IntervalNo *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStartPaymentPlanInterval is a free log retrieval operation binding the contract event 0x7191b4f9ba2ec111b9ae2a916ba78702dfbfe79ca993c2ce4321ff79a52a4c9c.
//
// Solidity: event StartPaymentPlanInterval(uint256 indexed ID, uint256 indexed callableOn, uint256 indexed amount, uint256 intervalNo)
func (_TravelSaver *TravelSaverFilterer) FilterStartPaymentPlanInterval(opts *bind.FilterOpts, ID []*big.Int, callableOn []*big.Int, amount []*big.Int) (*TravelSaverStartPaymentPlanIntervalIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var callableOnRule []interface{}
	for _, callableOnItem := range callableOn {
		callableOnRule = append(callableOnRule, callableOnItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "StartPaymentPlanInterval", IDRule, callableOnRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverStartPaymentPlanIntervalIterator{contract: _TravelSaver.contract, event: "StartPaymentPlanInterval", logs: logs, sub: sub}, nil
}

// WatchStartPaymentPlanInterval is a free log subscription operation binding the contract event 0x7191b4f9ba2ec111b9ae2a916ba78702dfbfe79ca993c2ce4321ff79a52a4c9c.
//
// Solidity: event StartPaymentPlanInterval(uint256 indexed ID, uint256 indexed callableOn, uint256 indexed amount, uint256 intervalNo)
func (_TravelSaver *TravelSaverFilterer) WatchStartPaymentPlanInterval(opts *bind.WatchOpts, sink chan<- *TravelSaverStartPaymentPlanInterval, ID []*big.Int, callableOn []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}
	var callableOnRule []interface{}
	for _, callableOnItem := range callableOn {
		callableOnRule = append(callableOnRule, callableOnItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "StartPaymentPlanInterval", IDRule, callableOnRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverStartPaymentPlanInterval)
				if err := _TravelSaver.contract.UnpackLog(event, "StartPaymentPlanInterval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStartPaymentPlanInterval is a log parse operation binding the contract event 0x7191b4f9ba2ec111b9ae2a916ba78702dfbfe79ca993c2ce4321ff79a52a4c9c.
//
// Solidity: event StartPaymentPlanInterval(uint256 indexed ID, uint256 indexed callableOn, uint256 indexed amount, uint256 intervalNo)
func (_TravelSaver *TravelSaverFilterer) ParseStartPaymentPlanInterval(log types.Log) (*TravelSaverStartPaymentPlanInterval, error) {
	event := new(TravelSaverStartPaymentPlanInterval)
	if err := _TravelSaver.contract.UnpackLog(event, "StartPaymentPlanInterval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TravelSaverTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TravelSaver contract.
type TravelSaverTransferIterator struct {
	Event *TravelSaverTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TravelSaverTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TravelSaverTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TravelSaverTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TravelSaverTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TravelSaverTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TravelSaverTransfer represents a Transfer event raised by the TravelSaver contract.
type TravelSaverTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TravelSaverTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TravelSaver.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TravelSaverTransferIterator{contract: _TravelSaver.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TravelSaverTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TravelSaver.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TravelSaverTransfer)
				if err := _TravelSaver.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_TravelSaver *TravelSaverFilterer) ParseTransfer(log types.Log) (*TravelSaverTransfer, error) {
	event := new(TravelSaverTransfer)
	if err := _TravelSaver.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
