package generator

import (
	"fmt"
	"os"
	"os/exec"
)

// Create generates a transaction offline and stores it in a JSON file.
func Create(filePath string) {

	fromAddress := "sif1t60c9055s8y2u4m6dyt4fnpa3j2s59al3m283r"
	toAddress := "sif18eawx2rdkddewhr5umy86g3hvsx0k9t8nq9kt7"
	amount := "10000rowan"
	chainId := "monkey-bars"

	args := []string{"tx", "send", fromAddress, toAddress, amount,
					 "--generate-only",
					 "--chain-id", chainId}

	cmd := exec.Command("sifnodecli", args...)

	outfile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start(); if err != nil {
		panic(err)
	}
	cmd.Wait()
}

// Sign signs a transaction that was generated offline.
func Sign(unsignedTxPath string, signedTxOutputFile string) {

	fromAddress := "sif1t60c9055s8y2u4m6dyt4fnpa3j2s59al3m283r"
	chainId := "monkey-bars"

	args := []string{"sign",       unsignedTxPath,
					 "--from",     fromAddress,
					 "--chain-id", chainId}

	cmd := exec.Command("sifnodecli", args...)

	outfile, err := os.Create(signedTxOutputFile)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start(); if err != nil {
		panic(err)
	}
	cmd.Wait()
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
		unsignedTx := fmt.Sprintf("%d.tx", i)
		Create(unsignedTx)

		signedTx := fmt.Sprintf("%d.signed.tx", i)
		Sign(unsignedTx, signedTx)
	}

	// Submit test transactions
	for i := 0; i < n; i++ {
		Send()
	}
}