package cluster

import (
	"github.com/unicornultrafoundation/subnet-ray-manager/market/contracts"
)

// Service handles Ray cluster operations based on onchain events.
type Service struct {
	// Add dependencies/config here if needed (e.g., DB, Ray API client, etc.)
}

// NewService creates a new cluster service instance.
func NewService() *Service {
	return &Service{}
}

// HandleCreateCluster handles the creation of a new Ray cluster for an order.
func (s *Service) OnOrderCreated(event *contracts.ClusterMarketOrderCreated) {
	// Implement logic to create a Ray cluster based on the order details.
	// This could involve calling Ray APIs, setting up resources, etc.
	// Use event.OrderId, event.User, event.ClusterSize, etc. as needed.
}
