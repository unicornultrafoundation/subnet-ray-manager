// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// SubnetClusterMarketCluster is an auto generated low-level Go binding around an user-defined struct.
type SubnetClusterMarketCluster struct {
	Owner       common.Address
	OrderId     *big.Int
	NodeIps     []*big.Int
	Active      bool
	Expiration  *big.Int
	Renter      common.Address
	RenterIp    *big.Int
	Ip          *big.Int
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}

// ClusterMarketMetaData contains all meta data concerning the ClusterMarket contract.
var ClusterMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"}],\"name\":\"ClusterInactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeIp\",\"type\":\"uint256\"}],\"name\":\"ClusterNodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newNodeIps\",\"type\":\"uint256[]\"}],\"name\":\"ClusterNodesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"DiscountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"OrderConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discountPrice\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"additionalDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"additionalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discountPrice\",\"type\":\"uint256\"}],\"name\":\"OrderExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"discountPrice\",\"type\":\"uint256\"}],\"name\":\"OrderScaled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"ResourcePriceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"addDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"newNodeIps\",\"type\":\"uint256[]\"}],\"name\":\"addNodesToCluster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeIp1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeIp2\",\"type\":\"uint256\"}],\"name\":\"areNodesInAnySameCluster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"sameCluster\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clusters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"renterIp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nodeIps\",\"type\":\"uint256[]\"}],\"name\":\"confirmOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"confirmScaleOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentalDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"discounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"}],\"name\":\"getCluster\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nodeIps\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"renterIp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"internalType\":\"structSubnetClusterMarket.Cluster\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeIp\",\"type\":\"uint256\"}],\"name\":\"getClustersOfNode\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"clusterIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rentalDuration\",\"type\":\"uint256\"}],\"name\":\"getDiscountPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"clusterIds\",\"type\":\"uint256[]\"}],\"name\":\"inactiveExpiredClusters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gpuPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cpuPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_memoryBytesPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_diskPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_networkPrice\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextClusterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeIpToClusterIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumSubnetClusterMarket.OrderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentalDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumSubnetClusterMarket.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nodeIps\",\"type\":\"uint256[]\"}],\"name\":\"recreateCluster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeIp\",\"type\":\"uint256\"}],\"name\":\"removeNodeFromCluster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nodeIps\",\"type\":\"uint256[]\"}],\"name\":\"removeNodesFromCluster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resourcePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"scale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryBytes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"}],\"name\":\"setResourcePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newIp\",\"type\":\"uint256\"}],\"name\":\"updateClusterIp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"updateDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClusterMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use ClusterMarketMetaData.ABI instead.
var ClusterMarketABI = ClusterMarketMetaData.ABI

// ClusterMarket is an auto generated Go binding around an Ethereum contract.
type ClusterMarket struct {
	ClusterMarketCaller     // Read-only binding to the contract
	ClusterMarketTransactor // Write-only binding to the contract
	ClusterMarketFilterer   // Log filterer for contract events
}

// ClusterMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClusterMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClusterMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClusterMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClusterMarketSession struct {
	Contract     *ClusterMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClusterMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClusterMarketCallerSession struct {
	Contract *ClusterMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClusterMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClusterMarketTransactorSession struct {
	Contract     *ClusterMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClusterMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClusterMarketRaw struct {
	Contract *ClusterMarket // Generic contract binding to access the raw methods on
}

// ClusterMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClusterMarketCallerRaw struct {
	Contract *ClusterMarketCaller // Generic read-only contract binding to access the raw methods on
}

// ClusterMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClusterMarketTransactorRaw struct {
	Contract *ClusterMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClusterMarket creates a new instance of ClusterMarket, bound to a specific deployed contract.
func NewClusterMarket(address common.Address, backend bind.ContractBackend) (*ClusterMarket, error) {
	contract, err := bindClusterMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClusterMarket{ClusterMarketCaller: ClusterMarketCaller{contract: contract}, ClusterMarketTransactor: ClusterMarketTransactor{contract: contract}, ClusterMarketFilterer: ClusterMarketFilterer{contract: contract}}, nil
}

// NewClusterMarketCaller creates a new read-only instance of ClusterMarket, bound to a specific deployed contract.
func NewClusterMarketCaller(address common.Address, caller bind.ContractCaller) (*ClusterMarketCaller, error) {
	contract, err := bindClusterMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketCaller{contract: contract}, nil
}

// NewClusterMarketTransactor creates a new write-only instance of ClusterMarket, bound to a specific deployed contract.
func NewClusterMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*ClusterMarketTransactor, error) {
	contract, err := bindClusterMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketTransactor{contract: contract}, nil
}

// NewClusterMarketFilterer creates a new log filterer instance of ClusterMarket, bound to a specific deployed contract.
func NewClusterMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*ClusterMarketFilterer, error) {
	contract, err := bindClusterMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketFilterer{contract: contract}, nil
}

// bindClusterMarket binds a generic wrapper to an already deployed contract.
func bindClusterMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClusterMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClusterMarket *ClusterMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClusterMarket.Contract.ClusterMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClusterMarket *ClusterMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ClusterMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClusterMarket *ClusterMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ClusterMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClusterMarket *ClusterMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClusterMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClusterMarket *ClusterMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClusterMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClusterMarket *ClusterMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClusterMarket.Contract.contract.Transact(opts, method, params...)
}

// AreNodesInAnySameCluster is a free data retrieval call binding the contract method 0x1e672888.
//
// Solidity: function areNodesInAnySameCluster(uint256 nodeIp1, uint256 nodeIp2) view returns(bool sameCluster)
func (_ClusterMarket *ClusterMarketCaller) AreNodesInAnySameCluster(opts *bind.CallOpts, nodeIp1 *big.Int, nodeIp2 *big.Int) (bool, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "areNodesInAnySameCluster", nodeIp1, nodeIp2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreNodesInAnySameCluster is a free data retrieval call binding the contract method 0x1e672888.
//
// Solidity: function areNodesInAnySameCluster(uint256 nodeIp1, uint256 nodeIp2) view returns(bool sameCluster)
func (_ClusterMarket *ClusterMarketSession) AreNodesInAnySameCluster(nodeIp1 *big.Int, nodeIp2 *big.Int) (bool, error) {
	return _ClusterMarket.Contract.AreNodesInAnySameCluster(&_ClusterMarket.CallOpts, nodeIp1, nodeIp2)
}

// AreNodesInAnySameCluster is a free data retrieval call binding the contract method 0x1e672888.
//
// Solidity: function areNodesInAnySameCluster(uint256 nodeIp1, uint256 nodeIp2) view returns(bool sameCluster)
func (_ClusterMarket *ClusterMarketCallerSession) AreNodesInAnySameCluster(nodeIp1 *big.Int, nodeIp2 *big.Int) (bool, error) {
	return _ClusterMarket.Contract.AreNodesInAnySameCluster(&_ClusterMarket.CallOpts, nodeIp1, nodeIp2)
}

// Clusters is a free data retrieval call binding the contract method 0x919500c0.
//
// Solidity: function clusters(uint256 ) view returns(address owner, uint256 orderId, bool active, uint256 expiration, address renter, uint256 renterIp, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketCaller) Clusters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner       common.Address
	OrderId     *big.Int
	Active      bool
	Expiration  *big.Int
	Renter      common.Address
	RenterIp    *big.Int
	Ip          *big.Int
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "clusters", arg0)

	outstruct := new(struct {
		Owner       common.Address
		OrderId     *big.Int
		Active      bool
		Expiration  *big.Int
		Renter      common.Address
		RenterIp    *big.Int
		Ip          *big.Int
		Gpu         *big.Int
		Cpu         *big.Int
		MemoryBytes *big.Int
		Disk        *big.Int
		Network     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OrderId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Expiration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Renter = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.RenterIp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Ip = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Gpu = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Cpu = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.MemoryBytes = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Disk = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Network = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Clusters is a free data retrieval call binding the contract method 0x919500c0.
//
// Solidity: function clusters(uint256 ) view returns(address owner, uint256 orderId, bool active, uint256 expiration, address renter, uint256 renterIp, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketSession) Clusters(arg0 *big.Int) (struct {
	Owner       common.Address
	OrderId     *big.Int
	Active      bool
	Expiration  *big.Int
	Renter      common.Address
	RenterIp    *big.Int
	Ip          *big.Int
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	return _ClusterMarket.Contract.Clusters(&_ClusterMarket.CallOpts, arg0)
}

// Clusters is a free data retrieval call binding the contract method 0x919500c0.
//
// Solidity: function clusters(uint256 ) view returns(address owner, uint256 orderId, bool active, uint256 expiration, address renter, uint256 renterIp, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketCallerSession) Clusters(arg0 *big.Int) (struct {
	Owner       common.Address
	OrderId     *big.Int
	Active      bool
	Expiration  *big.Int
	Renter      common.Address
	RenterIp    *big.Int
	Ip          *big.Int
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	return _ClusterMarket.Contract.Clusters(&_ClusterMarket.CallOpts, arg0)
}

// Discounts is a free data retrieval call binding the contract method 0xbae6a690.
//
// Solidity: function discounts(uint256 ) view returns(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketCaller) Discounts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MinDuration *big.Int
	Percent     *big.Int
}, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "discounts", arg0)

	outstruct := new(struct {
		MinDuration *big.Int
		Percent     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinDuration = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Percent = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Discounts is a free data retrieval call binding the contract method 0xbae6a690.
//
// Solidity: function discounts(uint256 ) view returns(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketSession) Discounts(arg0 *big.Int) (struct {
	MinDuration *big.Int
	Percent     *big.Int
}, error) {
	return _ClusterMarket.Contract.Discounts(&_ClusterMarket.CallOpts, arg0)
}

// Discounts is a free data retrieval call binding the contract method 0xbae6a690.
//
// Solidity: function discounts(uint256 ) view returns(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketCallerSession) Discounts(arg0 *big.Int) (struct {
	MinDuration *big.Int
	Percent     *big.Int
}, error) {
	return _ClusterMarket.Contract.Discounts(&_ClusterMarket.CallOpts, arg0)
}

// GetCluster is a free data retrieval call binding the contract method 0x0179240c.
//
// Solidity: function getCluster(uint256 clusterId) view returns((address,uint256,uint256[],bool,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ClusterMarket *ClusterMarketCaller) GetCluster(opts *bind.CallOpts, clusterId *big.Int) (SubnetClusterMarketCluster, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "getCluster", clusterId)

	if err != nil {
		return *new(SubnetClusterMarketCluster), err
	}

	out0 := *abi.ConvertType(out[0], new(SubnetClusterMarketCluster)).(*SubnetClusterMarketCluster)

	return out0, err

}

// GetCluster is a free data retrieval call binding the contract method 0x0179240c.
//
// Solidity: function getCluster(uint256 clusterId) view returns((address,uint256,uint256[],bool,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ClusterMarket *ClusterMarketSession) GetCluster(clusterId *big.Int) (SubnetClusterMarketCluster, error) {
	return _ClusterMarket.Contract.GetCluster(&_ClusterMarket.CallOpts, clusterId)
}

// GetCluster is a free data retrieval call binding the contract method 0x0179240c.
//
// Solidity: function getCluster(uint256 clusterId) view returns((address,uint256,uint256[],bool,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ClusterMarket *ClusterMarketCallerSession) GetCluster(clusterId *big.Int) (SubnetClusterMarketCluster, error) {
	return _ClusterMarket.Contract.GetCluster(&_ClusterMarket.CallOpts, clusterId)
}

// GetClustersOfNode is a free data retrieval call binding the contract method 0x8f4e4b1e.
//
// Solidity: function getClustersOfNode(uint256 nodeIp) view returns(uint256[] clusterIds)
func (_ClusterMarket *ClusterMarketCaller) GetClustersOfNode(opts *bind.CallOpts, nodeIp *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "getClustersOfNode", nodeIp)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetClustersOfNode is a free data retrieval call binding the contract method 0x8f4e4b1e.
//
// Solidity: function getClustersOfNode(uint256 nodeIp) view returns(uint256[] clusterIds)
func (_ClusterMarket *ClusterMarketSession) GetClustersOfNode(nodeIp *big.Int) ([]*big.Int, error) {
	return _ClusterMarket.Contract.GetClustersOfNode(&_ClusterMarket.CallOpts, nodeIp)
}

// GetClustersOfNode is a free data retrieval call binding the contract method 0x8f4e4b1e.
//
// Solidity: function getClustersOfNode(uint256 nodeIp) view returns(uint256[] clusterIds)
func (_ClusterMarket *ClusterMarketCallerSession) GetClustersOfNode(nodeIp *big.Int) ([]*big.Int, error) {
	return _ClusterMarket.Contract.GetClustersOfNode(&_ClusterMarket.CallOpts, nodeIp)
}

// GetDiscountPercent is a free data retrieval call binding the contract method 0x1cf93b49.
//
// Solidity: function getDiscountPercent(uint256 rentalDuration) view returns(uint256)
func (_ClusterMarket *ClusterMarketCaller) GetDiscountPercent(opts *bind.CallOpts, rentalDuration *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "getDiscountPercent", rentalDuration)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDiscountPercent is a free data retrieval call binding the contract method 0x1cf93b49.
//
// Solidity: function getDiscountPercent(uint256 rentalDuration) view returns(uint256)
func (_ClusterMarket *ClusterMarketSession) GetDiscountPercent(rentalDuration *big.Int) (*big.Int, error) {
	return _ClusterMarket.Contract.GetDiscountPercent(&_ClusterMarket.CallOpts, rentalDuration)
}

// GetDiscountPercent is a free data retrieval call binding the contract method 0x1cf93b49.
//
// Solidity: function getDiscountPercent(uint256 rentalDuration) view returns(uint256)
func (_ClusterMarket *ClusterMarketCallerSession) GetDiscountPercent(rentalDuration *big.Int) (*big.Int, error) {
	return _ClusterMarket.Contract.GetDiscountPercent(&_ClusterMarket.CallOpts, rentalDuration)
}

// NextClusterId is a free data retrieval call binding the contract method 0x4983dc44.
//
// Solidity: function nextClusterId() view returns(uint256)
func (_ClusterMarket *ClusterMarketCaller) NextClusterId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "nextClusterId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextClusterId is a free data retrieval call binding the contract method 0x4983dc44.
//
// Solidity: function nextClusterId() view returns(uint256)
func (_ClusterMarket *ClusterMarketSession) NextClusterId() (*big.Int, error) {
	return _ClusterMarket.Contract.NextClusterId(&_ClusterMarket.CallOpts)
}

// NextClusterId is a free data retrieval call binding the contract method 0x4983dc44.
//
// Solidity: function nextClusterId() view returns(uint256)
func (_ClusterMarket *ClusterMarketCallerSession) NextClusterId() (*big.Int, error) {
	return _ClusterMarket.Contract.NextClusterId(&_ClusterMarket.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_ClusterMarket *ClusterMarketCaller) NextOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "nextOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_ClusterMarket *ClusterMarketSession) NextOrderId() (*big.Int, error) {
	return _ClusterMarket.Contract.NextOrderId(&_ClusterMarket.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_ClusterMarket *ClusterMarketCallerSession) NextOrderId() (*big.Int, error) {
	return _ClusterMarket.Contract.NextOrderId(&_ClusterMarket.CallOpts)
}

// NodeIpToClusterIds is a free data retrieval call binding the contract method 0x949a549d.
//
// Solidity: function nodeIpToClusterIds(uint256 , uint256 ) view returns(uint256)
func (_ClusterMarket *ClusterMarketCaller) NodeIpToClusterIds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "nodeIpToClusterIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeIpToClusterIds is a free data retrieval call binding the contract method 0x949a549d.
//
// Solidity: function nodeIpToClusterIds(uint256 , uint256 ) view returns(uint256)
func (_ClusterMarket *ClusterMarketSession) NodeIpToClusterIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _ClusterMarket.Contract.NodeIpToClusterIds(&_ClusterMarket.CallOpts, arg0, arg1)
}

// NodeIpToClusterIds is a free data retrieval call binding the contract method 0x949a549d.
//
// Solidity: function nodeIpToClusterIds(uint256 , uint256 ) view returns(uint256)
func (_ClusterMarket *ClusterMarketCallerSession) NodeIpToClusterIds(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _ClusterMarket.Contract.NodeIpToClusterIds(&_ClusterMarket.CallOpts, arg0, arg1)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ClusterMarket *ClusterMarketCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ClusterMarket *ClusterMarketSession) Operator() (common.Address, error) {
	return _ClusterMarket.Contract.Operator(&_ClusterMarket.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ClusterMarket *ClusterMarketCallerSession) Operator() (common.Address, error) {
	return _ClusterMarket.Contract.Operator(&_ClusterMarket.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address user, uint8 status, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, address paymentToken, uint256 clusterId, uint256 paidAmount, uint256 discountAmount, uint8 orderType)
func (_ClusterMarket *ClusterMarketCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User           common.Address
	Status         uint8
	Ip             *big.Int
	Gpu            *big.Int
	Cpu            *big.Int
	MemoryBytes    *big.Int
	Disk           *big.Int
	Network        *big.Int
	RentalDuration *big.Int
	PaymentToken   common.Address
	ClusterId      *big.Int
	PaidAmount     *big.Int
	DiscountAmount *big.Int
	OrderType      uint8
}, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		User           common.Address
		Status         uint8
		Ip             *big.Int
		Gpu            *big.Int
		Cpu            *big.Int
		MemoryBytes    *big.Int
		Disk           *big.Int
		Network        *big.Int
		RentalDuration *big.Int
		PaymentToken   common.Address
		ClusterId      *big.Int
		PaidAmount     *big.Int
		DiscountAmount *big.Int
		OrderType      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Ip = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Gpu = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Cpu = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MemoryBytes = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Disk = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Network = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.RentalDuration = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.PaymentToken = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.ClusterId = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.PaidAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.DiscountAmount = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.OrderType = *abi.ConvertType(out[13], new(uint8)).(*uint8)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address user, uint8 status, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, address paymentToken, uint256 clusterId, uint256 paidAmount, uint256 discountAmount, uint8 orderType)
func (_ClusterMarket *ClusterMarketSession) Orders(arg0 *big.Int) (struct {
	User           common.Address
	Status         uint8
	Ip             *big.Int
	Gpu            *big.Int
	Cpu            *big.Int
	MemoryBytes    *big.Int
	Disk           *big.Int
	Network        *big.Int
	RentalDuration *big.Int
	PaymentToken   common.Address
	ClusterId      *big.Int
	PaidAmount     *big.Int
	DiscountAmount *big.Int
	OrderType      uint8
}, error) {
	return _ClusterMarket.Contract.Orders(&_ClusterMarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address user, uint8 status, uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, address paymentToken, uint256 clusterId, uint256 paidAmount, uint256 discountAmount, uint8 orderType)
func (_ClusterMarket *ClusterMarketCallerSession) Orders(arg0 *big.Int) (struct {
	User           common.Address
	Status         uint8
	Ip             *big.Int
	Gpu            *big.Int
	Cpu            *big.Int
	MemoryBytes    *big.Int
	Disk           *big.Int
	Network        *big.Int
	RentalDuration *big.Int
	PaymentToken   common.Address
	ClusterId      *big.Int
	PaidAmount     *big.Int
	DiscountAmount *big.Int
	OrderType      uint8
}, error) {
	return _ClusterMarket.Contract.Orders(&_ClusterMarket.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClusterMarket *ClusterMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClusterMarket *ClusterMarketSession) Owner() (common.Address, error) {
	return _ClusterMarket.Contract.Owner(&_ClusterMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClusterMarket *ClusterMarketCallerSession) Owner() (common.Address, error) {
	return _ClusterMarket.Contract.Owner(&_ClusterMarket.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_ClusterMarket *ClusterMarketCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_ClusterMarket *ClusterMarketSession) PaymentToken() (common.Address, error) {
	return _ClusterMarket.Contract.PaymentToken(&_ClusterMarket.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_ClusterMarket *ClusterMarketCallerSession) PaymentToken() (common.Address, error) {
	return _ClusterMarket.Contract.PaymentToken(&_ClusterMarket.CallOpts)
}

// ResourcePrice is a free data retrieval call binding the contract method 0xe4ccfc15.
//
// Solidity: function resourcePrice() view returns(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketCaller) ResourcePrice(opts *bind.CallOpts) (struct {
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	var out []interface{}
	err := _ClusterMarket.contract.Call(opts, &out, "resourcePrice")

	outstruct := new(struct {
		Gpu         *big.Int
		Cpu         *big.Int
		MemoryBytes *big.Int
		Disk        *big.Int
		Network     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gpu = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Cpu = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MemoryBytes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Disk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Network = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ResourcePrice is a free data retrieval call binding the contract method 0xe4ccfc15.
//
// Solidity: function resourcePrice() view returns(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketSession) ResourcePrice() (struct {
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	return _ClusterMarket.Contract.ResourcePrice(&_ClusterMarket.CallOpts)
}

// ResourcePrice is a free data retrieval call binding the contract method 0xe4ccfc15.
//
// Solidity: function resourcePrice() view returns(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketCallerSession) ResourcePrice() (struct {
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
}, error) {
	return _ClusterMarket.Contract.ResourcePrice(&_ClusterMarket.CallOpts)
}

// AddDiscount is a paid mutator transaction binding the contract method 0x0ab68776.
//
// Solidity: function addDiscount(uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketTransactor) AddDiscount(opts *bind.TransactOpts, minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "addDiscount", minDuration, percent)
}

// AddDiscount is a paid mutator transaction binding the contract method 0x0ab68776.
//
// Solidity: function addDiscount(uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketSession) AddDiscount(minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.AddDiscount(&_ClusterMarket.TransactOpts, minDuration, percent)
}

// AddDiscount is a paid mutator transaction binding the contract method 0x0ab68776.
//
// Solidity: function addDiscount(uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) AddDiscount(minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.AddDiscount(&_ClusterMarket.TransactOpts, minDuration, percent)
}

// AddNodesToCluster is a paid mutator transaction binding the contract method 0x4ffadbce.
//
// Solidity: function addNodesToCluster(uint256 clusterId, uint256[] newNodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactor) AddNodesToCluster(opts *bind.TransactOpts, clusterId *big.Int, newNodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "addNodesToCluster", clusterId, newNodeIps)
}

// AddNodesToCluster is a paid mutator transaction binding the contract method 0x4ffadbce.
//
// Solidity: function addNodesToCluster(uint256 clusterId, uint256[] newNodeIps) returns()
func (_ClusterMarket *ClusterMarketSession) AddNodesToCluster(clusterId *big.Int, newNodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.AddNodesToCluster(&_ClusterMarket.TransactOpts, clusterId, newNodeIps)
}

// AddNodesToCluster is a paid mutator transaction binding the contract method 0x4ffadbce.
//
// Solidity: function addNodesToCluster(uint256 clusterId, uint256[] newNodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) AddNodesToCluster(clusterId *big.Int, newNodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.AddNodesToCluster(&_ClusterMarket.TransactOpts, clusterId, newNodeIps)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.CancelOrder(&_ClusterMarket.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.CancelOrder(&_ClusterMarket.TransactOpts, orderId)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xc0c10db2.
//
// Solidity: function confirmOrder(uint256 orderId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactor) ConfirmOrder(opts *bind.TransactOpts, orderId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "confirmOrder", orderId, nodeIps)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xc0c10db2.
//
// Solidity: function confirmOrder(uint256 orderId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketSession) ConfirmOrder(orderId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ConfirmOrder(&_ClusterMarket.TransactOpts, orderId, nodeIps)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0xc0c10db2.
//
// Solidity: function confirmOrder(uint256 orderId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) ConfirmOrder(orderId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ConfirmOrder(&_ClusterMarket.TransactOpts, orderId, nodeIps)
}

// ConfirmScaleOrder is a paid mutator transaction binding the contract method 0xb609caeb.
//
// Solidity: function confirmScaleOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketTransactor) ConfirmScaleOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "confirmScaleOrder", orderId)
}

// ConfirmScaleOrder is a paid mutator transaction binding the contract method 0xb609caeb.
//
// Solidity: function confirmScaleOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketSession) ConfirmScaleOrder(orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ConfirmScaleOrder(&_ClusterMarket.TransactOpts, orderId)
}

// ConfirmScaleOrder is a paid mutator transaction binding the contract method 0xb609caeb.
//
// Solidity: function confirmScaleOrder(uint256 orderId) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) ConfirmScaleOrder(orderId *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.ConfirmScaleOrder(&_ClusterMarket.TransactOpts, orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb58a917b.
//
// Solidity: function createOrder(uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactor) CreateOrder(opts *bind.TransactOpts, ip *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, rentalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "createOrder", ip, gpu, cpu, memoryBytes, disk, network, rentalDuration, maxPrice)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb58a917b.
//
// Solidity: function createOrder(uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketSession) CreateOrder(ip *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, rentalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.CreateOrder(&_ClusterMarket.TransactOpts, ip, gpu, cpu, memoryBytes, disk, network, rentalDuration, maxPrice)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xb58a917b.
//
// Solidity: function createOrder(uint256 ip, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 rentalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) CreateOrder(ip *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, rentalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.CreateOrder(&_ClusterMarket.TransactOpts, ip, gpu, cpu, memoryBytes, disk, network, rentalDuration, maxPrice)
}

// Extend is a paid mutator transaction binding the contract method 0xddddca11.
//
// Solidity: function extend(uint256 clusterId, uint256 additionalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactor) Extend(opts *bind.TransactOpts, clusterId *big.Int, additionalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "extend", clusterId, additionalDuration, maxPrice)
}

// Extend is a paid mutator transaction binding the contract method 0xddddca11.
//
// Solidity: function extend(uint256 clusterId, uint256 additionalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketSession) Extend(clusterId *big.Int, additionalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Extend(&_ClusterMarket.TransactOpts, clusterId, additionalDuration, maxPrice)
}

// Extend is a paid mutator transaction binding the contract method 0xddddca11.
//
// Solidity: function extend(uint256 clusterId, uint256 additionalDuration, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) Extend(clusterId *big.Int, additionalDuration *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Extend(&_ClusterMarket.TransactOpts, clusterId, additionalDuration, maxPrice)
}

// InactiveExpiredClusters is a paid mutator transaction binding the contract method 0x27072835.
//
// Solidity: function inactiveExpiredClusters(uint256[] clusterIds) returns()
func (_ClusterMarket *ClusterMarketTransactor) InactiveExpiredClusters(opts *bind.TransactOpts, clusterIds []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "inactiveExpiredClusters", clusterIds)
}

// InactiveExpiredClusters is a paid mutator transaction binding the contract method 0x27072835.
//
// Solidity: function inactiveExpiredClusters(uint256[] clusterIds) returns()
func (_ClusterMarket *ClusterMarketSession) InactiveExpiredClusters(clusterIds []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.InactiveExpiredClusters(&_ClusterMarket.TransactOpts, clusterIds)
}

// InactiveExpiredClusters is a paid mutator transaction binding the contract method 0x27072835.
//
// Solidity: function inactiveExpiredClusters(uint256[] clusterIds) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) InactiveExpiredClusters(clusterIds []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.InactiveExpiredClusters(&_ClusterMarket.TransactOpts, clusterIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address _initialOwner, address _operator, uint256 _gpuPrice, uint256 _cpuPrice, uint256 _memoryBytesPrice, uint256 _diskPrice, uint256 _networkPrice) returns()
func (_ClusterMarket *ClusterMarketTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _operator common.Address, _gpuPrice *big.Int, _cpuPrice *big.Int, _memoryBytesPrice *big.Int, _diskPrice *big.Int, _networkPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "initialize", _initialOwner, _operator, _gpuPrice, _cpuPrice, _memoryBytesPrice, _diskPrice, _networkPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address _initialOwner, address _operator, uint256 _gpuPrice, uint256 _cpuPrice, uint256 _memoryBytesPrice, uint256 _diskPrice, uint256 _networkPrice) returns()
func (_ClusterMarket *ClusterMarketSession) Initialize(_initialOwner common.Address, _operator common.Address, _gpuPrice *big.Int, _cpuPrice *big.Int, _memoryBytesPrice *big.Int, _diskPrice *big.Int, _networkPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Initialize(&_ClusterMarket.TransactOpts, _initialOwner, _operator, _gpuPrice, _cpuPrice, _memoryBytesPrice, _diskPrice, _networkPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xedbf4ac2.
//
// Solidity: function initialize(address _initialOwner, address _operator, uint256 _gpuPrice, uint256 _cpuPrice, uint256 _memoryBytesPrice, uint256 _diskPrice, uint256 _networkPrice) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) Initialize(_initialOwner common.Address, _operator common.Address, _gpuPrice *big.Int, _cpuPrice *big.Int, _memoryBytesPrice *big.Int, _diskPrice *big.Int, _networkPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Initialize(&_ClusterMarket.TransactOpts, _initialOwner, _operator, _gpuPrice, _cpuPrice, _memoryBytesPrice, _diskPrice, _networkPrice)
}

// RecreateCluster is a paid mutator transaction binding the contract method 0x3b17f0a7.
//
// Solidity: function recreateCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactor) RecreateCluster(opts *bind.TransactOpts, clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "recreateCluster", clusterId, nodeIps)
}

// RecreateCluster is a paid mutator transaction binding the contract method 0x3b17f0a7.
//
// Solidity: function recreateCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketSession) RecreateCluster(clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RecreateCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIps)
}

// RecreateCluster is a paid mutator transaction binding the contract method 0x3b17f0a7.
//
// Solidity: function recreateCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) RecreateCluster(clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RecreateCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIps)
}

// RemoveNodeFromCluster is a paid mutator transaction binding the contract method 0xd39e77b0.
//
// Solidity: function removeNodeFromCluster(uint256 clusterId, uint256 nodeIp) returns()
func (_ClusterMarket *ClusterMarketTransactor) RemoveNodeFromCluster(opts *bind.TransactOpts, clusterId *big.Int, nodeIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "removeNodeFromCluster", clusterId, nodeIp)
}

// RemoveNodeFromCluster is a paid mutator transaction binding the contract method 0xd39e77b0.
//
// Solidity: function removeNodeFromCluster(uint256 clusterId, uint256 nodeIp) returns()
func (_ClusterMarket *ClusterMarketSession) RemoveNodeFromCluster(clusterId *big.Int, nodeIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RemoveNodeFromCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIp)
}

// RemoveNodeFromCluster is a paid mutator transaction binding the contract method 0xd39e77b0.
//
// Solidity: function removeNodeFromCluster(uint256 clusterId, uint256 nodeIp) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) RemoveNodeFromCluster(clusterId *big.Int, nodeIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RemoveNodeFromCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIp)
}

// RemoveNodesFromCluster is a paid mutator transaction binding the contract method 0xbf0f08bb.
//
// Solidity: function removeNodesFromCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactor) RemoveNodesFromCluster(opts *bind.TransactOpts, clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "removeNodesFromCluster", clusterId, nodeIps)
}

// RemoveNodesFromCluster is a paid mutator transaction binding the contract method 0xbf0f08bb.
//
// Solidity: function removeNodesFromCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketSession) RemoveNodesFromCluster(clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RemoveNodesFromCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIps)
}

// RemoveNodesFromCluster is a paid mutator transaction binding the contract method 0xbf0f08bb.
//
// Solidity: function removeNodesFromCluster(uint256 clusterId, uint256[] nodeIps) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) RemoveNodesFromCluster(clusterId *big.Int, nodeIps []*big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.RemoveNodesFromCluster(&_ClusterMarket.TransactOpts, clusterId, nodeIps)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClusterMarket *ClusterMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClusterMarket *ClusterMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClusterMarket.Contract.RenounceOwnership(&_ClusterMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ClusterMarket *ClusterMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ClusterMarket.Contract.RenounceOwnership(&_ClusterMarket.TransactOpts)
}

// Scale is a paid mutator transaction binding the contract method 0xb906aa96.
//
// Solidity: function scale(uint256 clusterId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactor) Scale(opts *bind.TransactOpts, clusterId *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "scale", clusterId, gpu, cpu, memoryBytes, disk, network, maxPrice)
}

// Scale is a paid mutator transaction binding the contract method 0xb906aa96.
//
// Solidity: function scale(uint256 clusterId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketSession) Scale(clusterId *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Scale(&_ClusterMarket.TransactOpts, clusterId, gpu, cpu, memoryBytes, disk, network, maxPrice)
}

// Scale is a paid mutator transaction binding the contract method 0xb906aa96.
//
// Solidity: function scale(uint256 clusterId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 maxPrice) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) Scale(clusterId *big.Int, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.Scale(&_ClusterMarket.TransactOpts, clusterId, gpu, cpu, memoryBytes, disk, network, maxPrice)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_ClusterMarket *ClusterMarketTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_ClusterMarket *ClusterMarketSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetOperator(&_ClusterMarket.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetOperator(&_ClusterMarket.TransactOpts, _operator)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address token) returns()
func (_ClusterMarket *ClusterMarketTransactor) SetPaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "setPaymentToken", token)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address token) returns()
func (_ClusterMarket *ClusterMarketSession) SetPaymentToken(token common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetPaymentToken(&_ClusterMarket.TransactOpts, token)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address token) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) SetPaymentToken(token common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetPaymentToken(&_ClusterMarket.TransactOpts, token)
}

// SetResourcePrice is a paid mutator transaction binding the contract method 0x7332c548.
//
// Solidity: function setResourcePrice(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network) returns()
func (_ClusterMarket *ClusterMarketTransactor) SetResourcePrice(opts *bind.TransactOpts, gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "setResourcePrice", gpu, cpu, memoryBytes, disk, network)
}

// SetResourcePrice is a paid mutator transaction binding the contract method 0x7332c548.
//
// Solidity: function setResourcePrice(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network) returns()
func (_ClusterMarket *ClusterMarketSession) SetResourcePrice(gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetResourcePrice(&_ClusterMarket.TransactOpts, gpu, cpu, memoryBytes, disk, network)
}

// SetResourcePrice is a paid mutator transaction binding the contract method 0x7332c548.
//
// Solidity: function setResourcePrice(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) SetResourcePrice(gpu *big.Int, cpu *big.Int, memoryBytes *big.Int, disk *big.Int, network *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.SetResourcePrice(&_ClusterMarket.TransactOpts, gpu, cpu, memoryBytes, disk, network)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClusterMarket *ClusterMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClusterMarket *ClusterMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.TransferOwnership(&_ClusterMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClusterMarket.Contract.TransferOwnership(&_ClusterMarket.TransactOpts, newOwner)
}

// UpdateClusterIp is a paid mutator transaction binding the contract method 0x5414323f.
//
// Solidity: function updateClusterIp(uint256 clusterId, uint256 newIp) returns()
func (_ClusterMarket *ClusterMarketTransactor) UpdateClusterIp(opts *bind.TransactOpts, clusterId *big.Int, newIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "updateClusterIp", clusterId, newIp)
}

// UpdateClusterIp is a paid mutator transaction binding the contract method 0x5414323f.
//
// Solidity: function updateClusterIp(uint256 clusterId, uint256 newIp) returns()
func (_ClusterMarket *ClusterMarketSession) UpdateClusterIp(clusterId *big.Int, newIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.UpdateClusterIp(&_ClusterMarket.TransactOpts, clusterId, newIp)
}

// UpdateClusterIp is a paid mutator transaction binding the contract method 0x5414323f.
//
// Solidity: function updateClusterIp(uint256 clusterId, uint256 newIp) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) UpdateClusterIp(clusterId *big.Int, newIp *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.UpdateClusterIp(&_ClusterMarket.TransactOpts, clusterId, newIp)
}

// UpdateDiscount is a paid mutator transaction binding the contract method 0x166cfddd.
//
// Solidity: function updateDiscount(uint256 index, uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketTransactor) UpdateDiscount(opts *bind.TransactOpts, index *big.Int, minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "updateDiscount", index, minDuration, percent)
}

// UpdateDiscount is a paid mutator transaction binding the contract method 0x166cfddd.
//
// Solidity: function updateDiscount(uint256 index, uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketSession) UpdateDiscount(index *big.Int, minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.UpdateDiscount(&_ClusterMarket.TransactOpts, index, minDuration, percent)
}

// UpdateDiscount is a paid mutator transaction binding the contract method 0x166cfddd.
//
// Solidity: function updateDiscount(uint256 index, uint256 minDuration, uint256 percent) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) UpdateDiscount(index *big.Int, minDuration *big.Int, percent *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.UpdateDiscount(&_ClusterMarket.TransactOpts, index, minDuration, percent)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x3ff681e4.
//
// Solidity: function withdrawFund(address token, address to, uint256 amount) returns()
func (_ClusterMarket *ClusterMarketTransactor) WithdrawFund(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.contract.Transact(opts, "withdrawFund", token, to, amount)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x3ff681e4.
//
// Solidity: function withdrawFund(address token, address to, uint256 amount) returns()
func (_ClusterMarket *ClusterMarketSession) WithdrawFund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.WithdrawFund(&_ClusterMarket.TransactOpts, token, to, amount)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x3ff681e4.
//
// Solidity: function withdrawFund(address token, address to, uint256 amount) returns()
func (_ClusterMarket *ClusterMarketTransactorSession) WithdrawFund(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ClusterMarket.Contract.WithdrawFund(&_ClusterMarket.TransactOpts, token, to, amount)
}

// ClusterMarketClusterInactiveIterator is returned from FilterClusterInactive and is used to iterate over the raw logs and unpacked data for ClusterInactive events raised by the ClusterMarket contract.
type ClusterMarketClusterInactiveIterator struct {
	Event *ClusterMarketClusterInactive // Event containing the contract specifics and raw log

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
func (it *ClusterMarketClusterInactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketClusterInactive)
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
		it.Event = new(ClusterMarketClusterInactive)
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
func (it *ClusterMarketClusterInactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketClusterInactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketClusterInactive represents a ClusterInactive event raised by the ClusterMarket contract.
type ClusterMarketClusterInactive struct {
	ClusterId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClusterInactive is a free log retrieval operation binding the contract event 0x8c459d0832b883bf56848f326554cd247915959623cb54c5b5bfc6cecdb39a65.
//
// Solidity: event ClusterInactive(uint256 indexed clusterId)
func (_ClusterMarket *ClusterMarketFilterer) FilterClusterInactive(opts *bind.FilterOpts, clusterId []*big.Int) (*ClusterMarketClusterInactiveIterator, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "ClusterInactive", clusterIdRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketClusterInactiveIterator{contract: _ClusterMarket.contract, event: "ClusterInactive", logs: logs, sub: sub}, nil
}

// WatchClusterInactive is a free log subscription operation binding the contract event 0x8c459d0832b883bf56848f326554cd247915959623cb54c5b5bfc6cecdb39a65.
//
// Solidity: event ClusterInactive(uint256 indexed clusterId)
func (_ClusterMarket *ClusterMarketFilterer) WatchClusterInactive(opts *bind.WatchOpts, sink chan<- *ClusterMarketClusterInactive, clusterId []*big.Int) (event.Subscription, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "ClusterInactive", clusterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketClusterInactive)
				if err := _ClusterMarket.contract.UnpackLog(event, "ClusterInactive", log); err != nil {
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

// ParseClusterInactive is a log parse operation binding the contract event 0x8c459d0832b883bf56848f326554cd247915959623cb54c5b5bfc6cecdb39a65.
//
// Solidity: event ClusterInactive(uint256 indexed clusterId)
func (_ClusterMarket *ClusterMarketFilterer) ParseClusterInactive(log types.Log) (*ClusterMarketClusterInactive, error) {
	event := new(ClusterMarketClusterInactive)
	if err := _ClusterMarket.contract.UnpackLog(event, "ClusterInactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketClusterNodeRemovedIterator is returned from FilterClusterNodeRemoved and is used to iterate over the raw logs and unpacked data for ClusterNodeRemoved events raised by the ClusterMarket contract.
type ClusterMarketClusterNodeRemovedIterator struct {
	Event *ClusterMarketClusterNodeRemoved // Event containing the contract specifics and raw log

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
func (it *ClusterMarketClusterNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketClusterNodeRemoved)
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
		it.Event = new(ClusterMarketClusterNodeRemoved)
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
func (it *ClusterMarketClusterNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketClusterNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketClusterNodeRemoved represents a ClusterNodeRemoved event raised by the ClusterMarket contract.
type ClusterMarketClusterNodeRemoved struct {
	ClusterId *big.Int
	NodeIp    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClusterNodeRemoved is a free log retrieval operation binding the contract event 0x321612790897bd90d2d7b5818c7a2b9a734f3a4b5f70799c4991ba510dfd6074.
//
// Solidity: event ClusterNodeRemoved(uint256 indexed clusterId, uint256 indexed nodeIp)
func (_ClusterMarket *ClusterMarketFilterer) FilterClusterNodeRemoved(opts *bind.FilterOpts, clusterId []*big.Int, nodeIp []*big.Int) (*ClusterMarketClusterNodeRemovedIterator, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}
	var nodeIpRule []interface{}
	for _, nodeIpItem := range nodeIp {
		nodeIpRule = append(nodeIpRule, nodeIpItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "ClusterNodeRemoved", clusterIdRule, nodeIpRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketClusterNodeRemovedIterator{contract: _ClusterMarket.contract, event: "ClusterNodeRemoved", logs: logs, sub: sub}, nil
}

// WatchClusterNodeRemoved is a free log subscription operation binding the contract event 0x321612790897bd90d2d7b5818c7a2b9a734f3a4b5f70799c4991ba510dfd6074.
//
// Solidity: event ClusterNodeRemoved(uint256 indexed clusterId, uint256 indexed nodeIp)
func (_ClusterMarket *ClusterMarketFilterer) WatchClusterNodeRemoved(opts *bind.WatchOpts, sink chan<- *ClusterMarketClusterNodeRemoved, clusterId []*big.Int, nodeIp []*big.Int) (event.Subscription, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}
	var nodeIpRule []interface{}
	for _, nodeIpItem := range nodeIp {
		nodeIpRule = append(nodeIpRule, nodeIpItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "ClusterNodeRemoved", clusterIdRule, nodeIpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketClusterNodeRemoved)
				if err := _ClusterMarket.contract.UnpackLog(event, "ClusterNodeRemoved", log); err != nil {
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

// ParseClusterNodeRemoved is a log parse operation binding the contract event 0x321612790897bd90d2d7b5818c7a2b9a734f3a4b5f70799c4991ba510dfd6074.
//
// Solidity: event ClusterNodeRemoved(uint256 indexed clusterId, uint256 indexed nodeIp)
func (_ClusterMarket *ClusterMarketFilterer) ParseClusterNodeRemoved(log types.Log) (*ClusterMarketClusterNodeRemoved, error) {
	event := new(ClusterMarketClusterNodeRemoved)
	if err := _ClusterMarket.contract.UnpackLog(event, "ClusterNodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketClusterNodesAddedIterator is returned from FilterClusterNodesAdded and is used to iterate over the raw logs and unpacked data for ClusterNodesAdded events raised by the ClusterMarket contract.
type ClusterMarketClusterNodesAddedIterator struct {
	Event *ClusterMarketClusterNodesAdded // Event containing the contract specifics and raw log

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
func (it *ClusterMarketClusterNodesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketClusterNodesAdded)
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
		it.Event = new(ClusterMarketClusterNodesAdded)
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
func (it *ClusterMarketClusterNodesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketClusterNodesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketClusterNodesAdded represents a ClusterNodesAdded event raised by the ClusterMarket contract.
type ClusterMarketClusterNodesAdded struct {
	ClusterId  *big.Int
	NewNodeIps []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClusterNodesAdded is a free log retrieval operation binding the contract event 0xa80b5484eb9aac37d936efc3d8e24049008b11e89daa6495d318baf3105ff563.
//
// Solidity: event ClusterNodesAdded(uint256 indexed clusterId, uint256[] newNodeIps)
func (_ClusterMarket *ClusterMarketFilterer) FilterClusterNodesAdded(opts *bind.FilterOpts, clusterId []*big.Int) (*ClusterMarketClusterNodesAddedIterator, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "ClusterNodesAdded", clusterIdRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketClusterNodesAddedIterator{contract: _ClusterMarket.contract, event: "ClusterNodesAdded", logs: logs, sub: sub}, nil
}

// WatchClusterNodesAdded is a free log subscription operation binding the contract event 0xa80b5484eb9aac37d936efc3d8e24049008b11e89daa6495d318baf3105ff563.
//
// Solidity: event ClusterNodesAdded(uint256 indexed clusterId, uint256[] newNodeIps)
func (_ClusterMarket *ClusterMarketFilterer) WatchClusterNodesAdded(opts *bind.WatchOpts, sink chan<- *ClusterMarketClusterNodesAdded, clusterId []*big.Int) (event.Subscription, error) {

	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "ClusterNodesAdded", clusterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketClusterNodesAdded)
				if err := _ClusterMarket.contract.UnpackLog(event, "ClusterNodesAdded", log); err != nil {
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

// ParseClusterNodesAdded is a log parse operation binding the contract event 0xa80b5484eb9aac37d936efc3d8e24049008b11e89daa6495d318baf3105ff563.
//
// Solidity: event ClusterNodesAdded(uint256 indexed clusterId, uint256[] newNodeIps)
func (_ClusterMarket *ClusterMarketFilterer) ParseClusterNodesAdded(log types.Log) (*ClusterMarketClusterNodesAdded, error) {
	event := new(ClusterMarketClusterNodesAdded)
	if err := _ClusterMarket.contract.UnpackLog(event, "ClusterNodesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketDiscountAddedIterator is returned from FilterDiscountAdded and is used to iterate over the raw logs and unpacked data for DiscountAdded events raised by the ClusterMarket contract.
type ClusterMarketDiscountAddedIterator struct {
	Event *ClusterMarketDiscountAdded // Event containing the contract specifics and raw log

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
func (it *ClusterMarketDiscountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketDiscountAdded)
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
		it.Event = new(ClusterMarketDiscountAdded)
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
func (it *ClusterMarketDiscountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketDiscountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketDiscountAdded represents a DiscountAdded event raised by the ClusterMarket contract.
type ClusterMarketDiscountAdded struct {
	MinDuration *big.Int
	Percent     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDiscountAdded is a free log retrieval operation binding the contract event 0xbb2de15b1fdae29c27610cb7a0e84f181f8df2e1c704e762073d7292f65e74ef.
//
// Solidity: event DiscountAdded(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketFilterer) FilterDiscountAdded(opts *bind.FilterOpts) (*ClusterMarketDiscountAddedIterator, error) {

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "DiscountAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterMarketDiscountAddedIterator{contract: _ClusterMarket.contract, event: "DiscountAdded", logs: logs, sub: sub}, nil
}

// WatchDiscountAdded is a free log subscription operation binding the contract event 0xbb2de15b1fdae29c27610cb7a0e84f181f8df2e1c704e762073d7292f65e74ef.
//
// Solidity: event DiscountAdded(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketFilterer) WatchDiscountAdded(opts *bind.WatchOpts, sink chan<- *ClusterMarketDiscountAdded) (event.Subscription, error) {

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "DiscountAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketDiscountAdded)
				if err := _ClusterMarket.contract.UnpackLog(event, "DiscountAdded", log); err != nil {
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

// ParseDiscountAdded is a log parse operation binding the contract event 0xbb2de15b1fdae29c27610cb7a0e84f181f8df2e1c704e762073d7292f65e74ef.
//
// Solidity: event DiscountAdded(uint256 minDuration, uint256 percent)
func (_ClusterMarket *ClusterMarketFilterer) ParseDiscountAdded(log types.Log) (*ClusterMarketDiscountAdded, error) {
	event := new(ClusterMarketDiscountAdded)
	if err := _ClusterMarket.contract.UnpackLog(event, "DiscountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ClusterMarket contract.
type ClusterMarketInitializedIterator struct {
	Event *ClusterMarketInitialized // Event containing the contract specifics and raw log

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
func (it *ClusterMarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketInitialized)
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
		it.Event = new(ClusterMarketInitialized)
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
func (it *ClusterMarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketInitialized represents a Initialized event raised by the ClusterMarket contract.
type ClusterMarketInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ClusterMarket *ClusterMarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClusterMarketInitializedIterator, error) {

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClusterMarketInitializedIterator{contract: _ClusterMarket.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ClusterMarket *ClusterMarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ClusterMarketInitialized) (event.Subscription, error) {

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketInitialized)
				if err := _ClusterMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ClusterMarket *ClusterMarketFilterer) ParseInitialized(log types.Log) (*ClusterMarketInitialized, error) {
	event := new(ClusterMarketInitialized)
	if err := _ClusterMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the ClusterMarket contract.
type ClusterMarketOperatorUpdatedIterator struct {
	Event *ClusterMarketOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOperatorUpdated)
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
		it.Event = new(ClusterMarketOperatorUpdated)
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
func (it *ClusterMarketOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOperatorUpdated represents a OperatorUpdated event raised by the ClusterMarket contract.
type ClusterMarketOperatorUpdated struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed newOperator)
func (_ClusterMarket *ClusterMarketFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, newOperator []common.Address) (*ClusterMarketOperatorUpdatedIterator, error) {

	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OperatorUpdated", newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOperatorUpdatedIterator{contract: _ClusterMarket.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed newOperator)
func (_ClusterMarket *ClusterMarketFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *ClusterMarketOperatorUpdated, newOperator []common.Address) (event.Subscription, error) {

	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OperatorUpdated", newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOperatorUpdated)
				if err := _ClusterMarket.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
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

// ParseOperatorUpdated is a log parse operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed newOperator)
func (_ClusterMarket *ClusterMarketFilterer) ParseOperatorUpdated(log types.Log) (*ClusterMarketOperatorUpdated, error) {
	event := new(ClusterMarketOperatorUpdated)
	if err := _ClusterMarket.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the ClusterMarket contract.
type ClusterMarketOrderCanceledIterator struct {
	Event *ClusterMarketOrderCanceled // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOrderCanceled)
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
		it.Event = new(ClusterMarketOrderCanceled)
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
func (it *ClusterMarketOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOrderCanceled represents a OrderCanceled event raised by the ClusterMarket contract.
type ClusterMarketOrderCanceled struct {
	OrderId *big.Int
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0x9384174c8517f5537b08e79211fc039e8a098571a3a2b4cb21dfa6f3237e8de1.
//
// Solidity: event OrderCanceled(uint256 indexed orderId, address indexed user)
func (_ClusterMarket *ClusterMarketFilterer) FilterOrderCanceled(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*ClusterMarketOrderCanceledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OrderCanceled", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOrderCanceledIterator{contract: _ClusterMarket.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0x9384174c8517f5537b08e79211fc039e8a098571a3a2b4cb21dfa6f3237e8de1.
//
// Solidity: event OrderCanceled(uint256 indexed orderId, address indexed user)
func (_ClusterMarket *ClusterMarketFilterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *ClusterMarketOrderCanceled, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OrderCanceled", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOrderCanceled)
				if err := _ClusterMarket.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
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

// ParseOrderCanceled is a log parse operation binding the contract event 0x9384174c8517f5537b08e79211fc039e8a098571a3a2b4cb21dfa6f3237e8de1.
//
// Solidity: event OrderCanceled(uint256 indexed orderId, address indexed user)
func (_ClusterMarket *ClusterMarketFilterer) ParseOrderCanceled(log types.Log) (*ClusterMarketOrderCanceled, error) {
	event := new(ClusterMarketOrderCanceled)
	if err := _ClusterMarket.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOrderConfirmedIterator is returned from FilterOrderConfirmed and is used to iterate over the raw logs and unpacked data for OrderConfirmed events raised by the ClusterMarket contract.
type ClusterMarketOrderConfirmedIterator struct {
	Event *ClusterMarketOrderConfirmed // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOrderConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOrderConfirmed)
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
		it.Event = new(ClusterMarketOrderConfirmed)
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
func (it *ClusterMarketOrderConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOrderConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOrderConfirmed represents a OrderConfirmed event raised by the ClusterMarket contract.
type ClusterMarketOrderConfirmed struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderConfirmed is a free log retrieval operation binding the contract event 0xaf50e0923cf2eedda3dc56c3d4d73b4d920cee133cd920ffb873db233abecb0c.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId)
func (_ClusterMarket *ClusterMarketFilterer) FilterOrderConfirmed(opts *bind.FilterOpts, orderId []*big.Int) (*ClusterMarketOrderConfirmedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OrderConfirmed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOrderConfirmedIterator{contract: _ClusterMarket.contract, event: "OrderConfirmed", logs: logs, sub: sub}, nil
}

// WatchOrderConfirmed is a free log subscription operation binding the contract event 0xaf50e0923cf2eedda3dc56c3d4d73b4d920cee133cd920ffb873db233abecb0c.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId)
func (_ClusterMarket *ClusterMarketFilterer) WatchOrderConfirmed(opts *bind.WatchOpts, sink chan<- *ClusterMarketOrderConfirmed, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OrderConfirmed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOrderConfirmed)
				if err := _ClusterMarket.contract.UnpackLog(event, "OrderConfirmed", log); err != nil {
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

// ParseOrderConfirmed is a log parse operation binding the contract event 0xaf50e0923cf2eedda3dc56c3d4d73b4d920cee133cd920ffb873db233abecb0c.
//
// Solidity: event OrderConfirmed(uint256 indexed orderId)
func (_ClusterMarket *ClusterMarketFilterer) ParseOrderConfirmed(log types.Log) (*ClusterMarketOrderConfirmed, error) {
	event := new(ClusterMarketOrderConfirmed)
	if err := _ClusterMarket.contract.UnpackLog(event, "OrderConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the ClusterMarket contract.
type ClusterMarketOrderCreatedIterator struct {
	Event *ClusterMarketOrderCreated // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOrderCreated)
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
		it.Event = new(ClusterMarketOrderCreated)
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
func (it *ClusterMarketOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOrderCreated represents a OrderCreated event raised by the ClusterMarket contract.
type ClusterMarketOrderCreated struct {
	OrderId       *big.Int
	User          common.Address
	Price         *big.Int
	DiscountPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0xf7c110a6973307f2bc91245c2c06344ada13add2c1741e83ac5c0bb332bc85d5.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 price, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*ClusterMarketOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OrderCreated", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOrderCreatedIterator{contract: _ClusterMarket.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0xf7c110a6973307f2bc91245c2c06344ada13add2c1741e83ac5c0bb332bc85d5.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 price, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *ClusterMarketOrderCreated, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OrderCreated", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOrderCreated)
				if err := _ClusterMarket.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0xf7c110a6973307f2bc91245c2c06344ada13add2c1741e83ac5c0bb332bc85d5.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 price, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) ParseOrderCreated(log types.Log) (*ClusterMarketOrderCreated, error) {
	event := new(ClusterMarketOrderCreated)
	if err := _ClusterMarket.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOrderExtendedIterator is returned from FilterOrderExtended and is used to iterate over the raw logs and unpacked data for OrderExtended events raised by the ClusterMarket contract.
type ClusterMarketOrderExtendedIterator struct {
	Event *ClusterMarketOrderExtended // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOrderExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOrderExtended)
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
		it.Event = new(ClusterMarketOrderExtended)
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
func (it *ClusterMarketOrderExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOrderExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOrderExtended represents a OrderExtended event raised by the ClusterMarket contract.
type ClusterMarketOrderExtended struct {
	OrderId            *big.Int
	AdditionalDuration *big.Int
	AdditionalPrice    *big.Int
	DiscountPrice      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOrderExtended is a free log retrieval operation binding the contract event 0xfeba1c697cc59f0331d372a0e8afe2cd20e6f0f80e68e1003decec9fa640f0cc.
//
// Solidity: event OrderExtended(uint256 indexed orderId, uint256 additionalDuration, uint256 additionalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) FilterOrderExtended(opts *bind.FilterOpts, orderId []*big.Int) (*ClusterMarketOrderExtendedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OrderExtended", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOrderExtendedIterator{contract: _ClusterMarket.contract, event: "OrderExtended", logs: logs, sub: sub}, nil
}

// WatchOrderExtended is a free log subscription operation binding the contract event 0xfeba1c697cc59f0331d372a0e8afe2cd20e6f0f80e68e1003decec9fa640f0cc.
//
// Solidity: event OrderExtended(uint256 indexed orderId, uint256 additionalDuration, uint256 additionalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) WatchOrderExtended(opts *bind.WatchOpts, sink chan<- *ClusterMarketOrderExtended, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OrderExtended", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOrderExtended)
				if err := _ClusterMarket.contract.UnpackLog(event, "OrderExtended", log); err != nil {
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

// ParseOrderExtended is a log parse operation binding the contract event 0xfeba1c697cc59f0331d372a0e8afe2cd20e6f0f80e68e1003decec9fa640f0cc.
//
// Solidity: event OrderExtended(uint256 indexed orderId, uint256 additionalDuration, uint256 additionalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) ParseOrderExtended(log types.Log) (*ClusterMarketOrderExtended, error) {
	event := new(ClusterMarketOrderExtended)
	if err := _ClusterMarket.contract.UnpackLog(event, "OrderExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOrderScaledIterator is returned from FilterOrderScaled and is used to iterate over the raw logs and unpacked data for OrderScaled events raised by the ClusterMarket contract.
type ClusterMarketOrderScaledIterator struct {
	Event *ClusterMarketOrderScaled // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOrderScaledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOrderScaled)
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
		it.Event = new(ClusterMarketOrderScaled)
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
func (it *ClusterMarketOrderScaledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOrderScaledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOrderScaled represents a OrderScaled event raised by the ClusterMarket contract.
type ClusterMarketOrderScaled struct {
	OrderId       *big.Int
	Gpu           *big.Int
	Cpu           *big.Int
	MemoryBytes   *big.Int
	Disk          *big.Int
	Network       *big.Int
	TotalPrice    *big.Int
	DiscountPrice *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderScaled is a free log retrieval operation binding the contract event 0xa45388c65f07b9dbc6c6e252989a993c41018892e9e152bb9b8da618c77c09c3.
//
// Solidity: event OrderScaled(uint256 indexed orderId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 totalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) FilterOrderScaled(opts *bind.FilterOpts, orderId []*big.Int) (*ClusterMarketOrderScaledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OrderScaled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOrderScaledIterator{contract: _ClusterMarket.contract, event: "OrderScaled", logs: logs, sub: sub}, nil
}

// WatchOrderScaled is a free log subscription operation binding the contract event 0xa45388c65f07b9dbc6c6e252989a993c41018892e9e152bb9b8da618c77c09c3.
//
// Solidity: event OrderScaled(uint256 indexed orderId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 totalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) WatchOrderScaled(opts *bind.WatchOpts, sink chan<- *ClusterMarketOrderScaled, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OrderScaled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOrderScaled)
				if err := _ClusterMarket.contract.UnpackLog(event, "OrderScaled", log); err != nil {
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

// ParseOrderScaled is a log parse operation binding the contract event 0xa45388c65f07b9dbc6c6e252989a993c41018892e9e152bb9b8da618c77c09c3.
//
// Solidity: event OrderScaled(uint256 indexed orderId, uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network, uint256 totalPrice, uint256 discountPrice)
func (_ClusterMarket *ClusterMarketFilterer) ParseOrderScaled(log types.Log) (*ClusterMarketOrderScaled, error) {
	event := new(ClusterMarketOrderScaled)
	if err := _ClusterMarket.contract.UnpackLog(event, "OrderScaled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ClusterMarket contract.
type ClusterMarketOwnershipTransferredIterator struct {
	Event *ClusterMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClusterMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketOwnershipTransferred)
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
		it.Event = new(ClusterMarketOwnershipTransferred)
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
func (it *ClusterMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketOwnershipTransferred represents a OwnershipTransferred event raised by the ClusterMarket contract.
type ClusterMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClusterMarket *ClusterMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClusterMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketOwnershipTransferredIterator{contract: _ClusterMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClusterMarket *ClusterMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClusterMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketOwnershipTransferred)
				if err := _ClusterMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ClusterMarket *ClusterMarketFilterer) ParseOwnershipTransferred(log types.Log) (*ClusterMarketOwnershipTransferred, error) {
	event := new(ClusterMarketOwnershipTransferred)
	if err := _ClusterMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClusterMarketResourcePriceUpdatedIterator is returned from FilterResourcePriceUpdated and is used to iterate over the raw logs and unpacked data for ResourcePriceUpdated events raised by the ClusterMarket contract.
type ClusterMarketResourcePriceUpdatedIterator struct {
	Event *ClusterMarketResourcePriceUpdated // Event containing the contract specifics and raw log

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
func (it *ClusterMarketResourcePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMarketResourcePriceUpdated)
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
		it.Event = new(ClusterMarketResourcePriceUpdated)
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
func (it *ClusterMarketResourcePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMarketResourcePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMarketResourcePriceUpdated represents a ResourcePriceUpdated event raised by the ClusterMarket contract.
type ClusterMarketResourcePriceUpdated struct {
	Gpu         *big.Int
	Cpu         *big.Int
	MemoryBytes *big.Int
	Disk        *big.Int
	Network     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterResourcePriceUpdated is a free log retrieval operation binding the contract event 0xc05d384d96751ccdcd49a1246e85cf2131ef23bbbe94dffe94b22a352ded9685.
//
// Solidity: event ResourcePriceUpdated(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketFilterer) FilterResourcePriceUpdated(opts *bind.FilterOpts) (*ClusterMarketResourcePriceUpdatedIterator, error) {

	logs, sub, err := _ClusterMarket.contract.FilterLogs(opts, "ResourcePriceUpdated")
	if err != nil {
		return nil, err
	}
	return &ClusterMarketResourcePriceUpdatedIterator{contract: _ClusterMarket.contract, event: "ResourcePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchResourcePriceUpdated is a free log subscription operation binding the contract event 0xc05d384d96751ccdcd49a1246e85cf2131ef23bbbe94dffe94b22a352ded9685.
//
// Solidity: event ResourcePriceUpdated(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketFilterer) WatchResourcePriceUpdated(opts *bind.WatchOpts, sink chan<- *ClusterMarketResourcePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _ClusterMarket.contract.WatchLogs(opts, "ResourcePriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMarketResourcePriceUpdated)
				if err := _ClusterMarket.contract.UnpackLog(event, "ResourcePriceUpdated", log); err != nil {
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

// ParseResourcePriceUpdated is a log parse operation binding the contract event 0xc05d384d96751ccdcd49a1246e85cf2131ef23bbbe94dffe94b22a352ded9685.
//
// Solidity: event ResourcePriceUpdated(uint256 gpu, uint256 cpu, uint256 memoryBytes, uint256 disk, uint256 network)
func (_ClusterMarket *ClusterMarketFilterer) ParseResourcePriceUpdated(log types.Log) (*ClusterMarketResourcePriceUpdated, error) {
	event := new(ClusterMarketResourcePriceUpdated)
	if err := _ClusterMarket.contract.UnpackLog(event, "ResourcePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
