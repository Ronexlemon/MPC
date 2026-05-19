package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ronexlemon/mpc/enclave/internal/handler"
	"github.com/ronexlemon/mpc/enclave/pkg/config"
)

func main() {

	var curveOrder, success = new(big.Int).SetString(
		"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141",
		16,
	)

	if !success {
		log.Fatal("Failed to initialize curve order")
	}

	nodeName, nodeIndexStr := config.LoadEnv()

	var nodeIndex int64
	_, err := fmt.Sscanf(nodeIndexStr, "%d", &nodeIndex)
	if err != nil {
		log.Fatal(err)
	}

	server := handler.NewServer(nodeName, nodeIndex, curveOrder)

	log.Printf("Headless Enclave Server [%s] on :9090", nodeName)

	log.Fatal(http.ListenAndServe(":9090", server))
}