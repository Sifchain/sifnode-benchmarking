package generator

import (
	"fmt"
	"log"
	"os/exec"
)

// Create generates a transaction offline and stores it in a JSON file.
func Create(filePath string) {
	fromAddress := "sif1gr9tkt5k3y9expemamt9dqc26agstw4s8j7l0y"
	toAddress := "sif18eawx2rdkddewhr5umy86g3hvsx0k9t8nq9kt7"
	amount := "10000rowan"
	chainId := "monkey-bars"

	args := []string{"tx", "send", fromAddress, toAddress, amount,
					 "--generate-only",
					 "--chain-id", chainId,
					 ">", filePath}

	_, err := exec.Command("sifnodecli", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
}

// Sign signs a transaction that was generated offline.
func Sign(filePath string) {
	// Sign Transaction
}

// Send broadcasts the transaction to a target node.
func Send() {
	// Broadcast Transaction
	// Iterate over folder of transactions
}

// GenerateTransactions creates, signs, and sends transactions in bulk.
func GenerateTransactions(n int) {
	// Create test transactions
	for i := 0; i < n; i++ {
		filePath := fmt.Sprintf("data/%d.tx", i)
		Create(filePath)
		Sign(filePath)
	}

	// Submit test transactions
	for i := 0; i < n; i++ {
		Send()
	}
}