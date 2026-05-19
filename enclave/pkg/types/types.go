package types

import "math/big"

// Global in-memory storage simulating a secure hardware enclave partition
var SecureEnclaveMemory = make(map[string][]byte)
var curveOrder = big.NewInt(0)

type ProvisionRequest struct {
	WorkspaceID string `json:"workspace_id"`
	ShardHex    string `json:"shard_hex"`
}

type SignRequest struct {
	WorkspaceID string `json:"workspace_id"`
	MessageHash string `json:"message_hash"`
}

type SignResponse struct {
	PartialSignature string `json:"partial_signature"`
	NodeIndex        int64  `json:"node_index"`
}