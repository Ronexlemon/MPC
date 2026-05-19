package handler

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ronexlemon/mpc/enclave/pkg/types"
)

type Server struct {
	NodeName   string
	NodeIndex  int64
	CurveOrder *big.Int
}

func NewServer(nodeName string, nodeIndex int64, curveOrder *big.Int) http.Handler {

	s := &Server{
		NodeName:   nodeName,
		NodeIndex:  nodeIndex,
		CurveOrder: curveOrder,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/internal/provision", s.provision)
	mux.HandleFunc("/internal/sign", s.sign)

	return mux
}


func (s *Server) provision(w http.ResponseWriter, r *http.Request) {

	var req types.ProvisionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	types.SecureEnclaveMemory[req.WorkspaceID] = []byte(req.ShardHex)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,
		`{"status":"provisioned","node":"%s"}`,
		s.NodeName,
	)
}

func (s *Server) sign(w http.ResponseWriter, r *http.Request) {

	var req types.SignRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	shardMaterial, exists := types.SecureEnclaveMemory[req.WorkspaceID]
	if !exists {
		http.Error(w, "Key shard not found", http.StatusNotFound)
		return
	}

	shardInt := new(big.Int).SetBytes(shardMaterial)

	msgBytes, _ := json.Marshal(req.MessageHash)
	msgHash := sha256.Sum256(msgBytes)
	msgInt := new(big.Int).SetBytes(msgHash[:])

	// partial computation
	partial := new(big.Int).Mul(shardInt, msgInt)
	partial.Mod(partial, s.CurveOrder)

	resp := types.SignResponse{
		PartialSignature: fmt.Sprintf("%x", partial),
		NodeIndex:        s.NodeIndex,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}