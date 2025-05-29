package main

import (
	"log"

	"github.com/unicornultrafoundation/subnet-ray-manager/cluster"
	"github.com/unicornultrafoundation/subnet-ray-manager/config"
	"github.com/unicornultrafoundation/subnet-ray-manager/market"
)

func main() {
	log.Println("Starting Subnet Ray Manager...")

	// Load config
	cfg := config.LoadConfig()

	// Create onchain event handler (implement onchain.Service as needed)
	handler := cluster.NewService()

	// Create onchain service
	onchainService, err := market.NewClusterMarketService(cfg, handler)
	if err != nil {
		log.Fatalf("Failed to initialize onchain service: %v", err)
	}

	// Start listening for OrderCreated events
	onchainService.Start()
	onchainService.Stop()
}
