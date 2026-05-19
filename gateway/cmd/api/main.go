package main

import (
	"log"
	"math/big"
	"net/http"
	"github.com/ronexlemon/mpc/gateway/internal/handler"
)



var curveOrder *big.Int

func init() {
	var success bool
	// Constant order for EVM cryptography matching secp256k1
	curveOrder, success = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	if !success {
		log.Fatal("Failed to set cryptographic parameters")
	}
}





func main() {
	server := handler.NewServer(curveOrder)
	log.Println("WaaS Production Gateway active on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", server))
}