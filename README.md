# Subnet Ray Manager

A Go service that listens to onchain events (such as `OrderCreated`) from a smart contract and automatically provisions Ray clusters for tenants.

## Features

- Listen to onchain events from a specified smart contract.
- Provision, renew, and upgrade Ray clusters based on blockchain events.
- Modular service structure for easy extension and integration.

## Installation

```bash
git clone git@github.com:unicornultrafoundation/subnet-ray-manager.git
cd subnet-ray-manager
go mod tidy
```

## Configuration

Set the following environment variables:

- `ETH_RPC_URL`: Ethereum node URL (e.g., Infura, Alchemy, or your own node)
- `CONTRACT_ADDRESS`: Smart contract address to listen for events

## Usage

### Run Locally

```bash
export ETH_RPC_URL=your_rpc_url
export CONTRACT_ADDRESS=your_contract_address
go run ./cmd/subnet-ray-manager
```

### Run with Docker

```bash
docker-compose build
docker-compose up
```

## Project Structure

- `cmd/subnet-ray-manager/`: Application entry point
- `market/`: Onchain (market) event listener and service logic
- `cluster/`: Ray cluster management logic
- `config/`: Centralized configuration loading

## Contribution

Pull requests and issues are welcome!

## License

MIT
