package cluster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RayClient is a simple HTTP client for interacting with a Ray node.
type RayClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewRayClient creates a new RayClient with the given base URL.
func NewRayClient(baseURL string) *RayClient {
	return &RayClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// Request and response structs
type StartHeadRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type StartWorkerRequest struct {
	HeadIP   string `json:"headIP" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// StartNodeRequest represents the payload to start a Ray node.
type StartNodeRequest struct {
	NodeType  string           `json:"node_type"` // e.g. "worker" or "head"
	Resources map[string]int64 `json:"resources"` // e.g. {"CPU": 4, "GPU": 1}
	// Add more fields as needed
}

// StartNodeResponse represents the response from the Ray node API.
type StartNodeResponse struct {
	NodeID string `json:"node_id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// StatusResponse represents the response from the Ray node status API.
type StatusResponse struct {
	Status  string `json:"status"`
	HeadIP  string `json:"head_ip,omitempty"`
	Workers int    `json:"workers,omitempty"`
	Error   string `json:"error,omitempty"`
}

// StopNodeRequest represents the payload to stop a Ray node.
type StopNodeRequest struct {
	NodeID string `json:"node_id" binding:"required"`
}

// StopNodeResponse represents the response from the Ray node stop API.
type StopNodeResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// StartHead sends a request to the Ray node to start a head node.
func (c *RayClient) StartHead(req *StartHeadRequest) (*StartNodeResponse, error) {
	url := fmt.Sprintf("%s/start/head", c.BaseURL)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StartNodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "success" {
		return &result, fmt.Errorf("failed to start head node: %s", result.Error)
	}
	return &result, nil
}

// StartWorker sends a request to the Ray node to start a worker node.
func (c *RayClient) StartWorker(req *StartWorkerRequest) (*StartNodeResponse, error) {
	url := fmt.Sprintf("%s/start/worker", c.BaseURL)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StartNodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "success" {
		return &result, fmt.Errorf("failed to start worker node: %s", result.Error)
	}
	return &result, nil
}

// GetStatus queries the Ray node for its current status.
func (c *RayClient) GetStatus() (*StatusResponse, error) {
	url := fmt.Sprintf("%s/status", c.BaseURL)
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "success" {
		return &result, fmt.Errorf("failed to get status: %s", result.Error)
	}
	return &result, nil
}

// StopNode sends a request to the Ray node to stop a node by node ID.
func (c *RayClient) StopNode(req *StopNodeRequest) (*StopNodeResponse, error) {
	url := fmt.Sprintf("%s/stop", c.BaseURL)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StopNodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if result.Status != "success" {
		return &result, fmt.Errorf("failed to stop node: %s", result.Error)
	}
	return &result, nil
}
