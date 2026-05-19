package handler

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ronexlemon/mpc/gateway/internal/helper"
	"github.com/ronexlemon/mpc/gateway/pkg/types"
)



type Server struct{
	CurveOrder *big.Int
}


func NewServer(curveOrder *big.Int)http.Handler{

	mux:=http.NewServeMux()
	s:=&Server{
		CurveOrder: curveOrder,
	}

	mux.HandleFunc("/v1/vaults", s.CreateWalletHandler)
	mux.HandleFunc("/v1/transactions", s.SignTransactionHandler)

	return mux
}


func (s *Server) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") != "b2b_secret_auth_token" {
		http.Error(w, "Unauthorized access token", http.StatusUnauthorized)
		return
	}

	var req types.CreateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate key on the fly to yield clean public tracking attributes
	privKey, _ := crypto.GenerateKey()
	privKeyHex := hex.EncodeToString(crypto.FromECDSA(privKey))
	
	// FIX 1: Explicitly type assert and dereference the pointer using *
	publicKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		http.Error(w, "Failed to cast public key", http.StatusInternalServerError)
		return
	}
	pubAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	shard1, shard2, shard3 := helper.SplitKey(privKeyHex,s.CurveOrder)

	// Programmatically provision nodes over isolated internal docker paths
	nodes := []string{"http://enclave-aws:9090", "http://enclave-gcp:9090", "http://enclave-azure:9090"}
	shards := [][]byte{shard1, shard2, shard3}
	fmt.Println("Shards:",shards)

	for i, url := range nodes {
		payload, _ := json.Marshal(map[string]string{
			"workspace_id": req.WorkspaceID,
			"shard_hex":    string(shards[i]),
		})
		_, err := http.Post(url+"/internal/provision", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			http.Error(w, "Failed to isolate key material to container cluster", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"public_address":"%s","status":"SUCCESS_VAULT_ISOLATED"}`, pubAddress)
}

func(s *Server) SignTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") != "b2b_secret_auth_token" {
		http.Error(w, "Unauthorized access token", http.StatusUnauthorized)
		return
	}

	var req types.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	txPayload := fmt.Sprintf(`{"to":"%s","amount":"%s","nonce":%d}`, req.ToAddress, req.Amount, req.Nonce)

	// Collect partial mathematical components from threshold nodes (e.g., Node 1 and Node 2)
	partialNodes := []string{"http://enclave-aws:9090", "http://enclave-gcp:9090"}
	var signatures []*big.Int

	for _, url := range partialNodes {
		payload, _ := json.Marshal(map[string]string{
			"workspace_id": req.WorkspaceID,
			"message_hash": txPayload,
		})
		resp, err := http.Post(url+"/internal/sign", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			http.Error(w, "Failed network round during threshold math execution", http.StatusInternalServerError)
			return
		}
		
		var signResult map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&signResult); err != nil {
			http.Error(w, "Failed to decode enclave response", http.StatusInternalServerError)
			return
		}

		// FIX 2: Use a type assertion .(string) instead of calling .String() on an interface{}
		sigStr, ok := signResult["partial_signature"].(string)
		if !ok {
			http.Error(w, "Invalid signature format returned from enclave", http.StatusInternalServerError)
			return
		}

		sigBytes, _ := hex.DecodeString(sigStr)
		signatures = append(signatures, new(big.Int).SetBytes(sigBytes))
	}

	// Dynamic Lagrange Interpolation Phase to compute final unified signature safely
	c1 := big.NewInt(2)
	c2 := new(big.Int).Sub(s.CurveOrder, big.NewInt(1)) // -1 mod curveOrder

	term1 := new(big.Int).Mul(signatures[0], c1)
	term2 := new(big.Int).Mul(signatures[1], c2)

	finalSig := new(big.Int).Add(term1, term2)
	finalSig.Mod(finalSig, s.CurveOrder)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"signature_hex":"0x%x","status":"TRANSACTION_SIGNED_BY_CLUSTER"}`, finalSig)
}