package market

import (
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unicornultrafoundation/subnet-ray-manager/config"
	"github.com/unicornultrafoundation/subnet-ray-manager/market/contracts"
	"github.com/unicornultrafoundation/subnet-ray-manager/market/types"
)

// Service defines the interface for handling onchain events.
type Service interface {
	OnOrderCreated(order *types.Order)
	// Add more event handler methods as needed (e.g. OnOrderExtended, OnOrderScaled, ...)
}

// ClusterMarketService manages the onchain listener, ethclient, and contract instance for the ClusterMarket.
type ClusterMarketService struct {
	Client       *ethclient.Client
	Contract     *contracts.ClusterMarket
	ContractAddr common.Address
	stopCh       chan struct{}
	wg           sync.WaitGroup
	service      Service
}

// NewClusterMarketService creates and initializes ClusterMarketService from config and service.
// It connects to the Ethereum client, binds the ClusterMarket contract, and prepares to listen for events.
// The service parameter is an implementation of the Service interface that handles onchain events.
// It returns an error if the connection or contract binding fails.
// This service is responsible for listening to onchain events related to the ClusterMarket contract,
// such as OrderCreated, and delegating the handling of these events to the provided service implementation.
// It also provides methods to start and stop the listener, and to access the contract address and client.
// The service should be started by calling Start(), which will begin listening for events in a separate goroutine.
// The Stop() method should be called to gracefully stop the listener and clean up resources.
func NewClusterMarketService(cfg *config.Config, handler Service) (*ClusterMarketService, error) {
	client, err := ethclient.Dial(cfg.EthRpcUrl)
	if err != nil {
		return nil, err
	}
	addr := common.HexToAddress(cfg.ContractAddress)
	contract, err := contracts.NewClusterMarket(addr, client)
	if err != nil {
		return nil, err
	}
	return &ClusterMarketService{
		Client:       client,
		Contract:     contract,
		ContractAddr: addr,
		stopCh:       make(chan struct{}),
		service:      handler,
	}, nil
}

func (s *ClusterMarketService) watchOrderCreated() {
	log.Println("Listening for OrderCreated events...")
	sink := make(chan *contracts.ClusterMarketOrderCreated)
	sub, err := s.Contract.WatchOrderCreated(
		nil, // use default watch opts (no context timeout)
		sink,
		nil, // no filter for orderId
		nil, // no filter for user
	)
	if err != nil {
		log.Fatalf("Failed to subscribe to OrderCreated events: %v", err)
	}
	defer sub.Unsubscribe()

	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-s.stopCh:
			log.Println("Listener stopped.")
			return
		case err := <-sub.Err():
			log.Println("Subscription error:", err)
		case event := <-sink:
			log.Printf("OrderCreated: OrderId=%s, User=%s, Price=%s, DiscountPrice=%s\n",
				event.OrderId.String(),
				event.User.Hex(),
				event.Price.String(),
				event.DiscountPrice.String(),
			)

			order, err := s.GetOrder(event.OrderId)
			if err != nil {
				log.Printf("Failed to get order details for OrderId %s: %v", event.OrderId.String(), err)
				continue
			}
			s.service.OnOrderCreated(order)
		}
	}
}

// Start begins listening for onchain events.
func (s *ClusterMarketService) Start() {
	s.watchOrderCreated()
}

// Stop stops the onchain listener and cleans up resources.
func (s *ClusterMarketService) Stop() {
	close(s.stopCh)
	s.wg.Wait()
	log.Println("Onchain service stopped.")
}

// GetContractAddress returns the address of the onchain contract.
func (s *ClusterMarketService) GetContractAddress() common.Address {
	return s.ContractAddr
}

// GetEthClient returns the Ethereum client used by the onchain service.
func (s *ClusterMarketService) GetEthClient() *ethclient.Client {
	return s.Client
}

// GetContract returns the contract instance used by the onchain service.
func (s *ClusterMarketService) GetContract() *contracts.ClusterMarket {
	return s.Contract
}

// GetOrder retrieves order information from the contract by orderId.
func (s *ClusterMarketService) GetOrder(orderId *big.Int) (*types.Order, error) {
	callOpts := &bind.CallOpts{}
	orderStruct, err := s.Contract.Orders(callOpts, orderId)
	if err != nil {
		return nil, err
	}
	order := &types.Order{
		Id:             orderId,
		User:           orderStruct.User,
		Status:         orderStruct.Status,
		Ip:             orderStruct.Ip,
		Gpu:            orderStruct.Gpu,
		Cpu:            orderStruct.Cpu,
		MemoryBytes:    orderStruct.MemoryBytes,
		Disk:           orderStruct.Disk,
		Network:        orderStruct.Network,
		RentalDuration: orderStruct.RentalDuration,
		PaymentToken:   orderStruct.PaymentToken,
		ClusterId:      orderStruct.ClusterId,
		PaidAmount:     orderStruct.PaidAmount,
		DiscountAmount: orderStruct.DiscountAmount,
		OrderType:      orderStruct.OrderType,
	}
	return order, nil
}
