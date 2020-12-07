package generator

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"log"
	"os/exec"
)

type Configuration struct {
	FromAddress string
	ToAddress   string
	Amount      string
	ChainId     string
}

// Create generates a transaction offline and stores it in a JSON file.
func Create(filePath string) {
	configuration := Configuration{}
	err := gonfig.GetConf("../config.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := configuration.FromAddress
	toAddress := configuration.ToAddress
	amount := configuration.Amount
	chainId := configuration.ChainId

	args := []string{"tx", "send", fromAddress, toAddress, amount,
					 "--generate-only",
					 "--chain-id", chainId,
					 ">", filePath}

	_, err = exec.Command("sifnodecli", args...).Output()
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