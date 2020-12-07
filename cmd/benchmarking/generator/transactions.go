package generator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Create generates a transaction offline and stores it in a JSON file.
func Create(filePath string) {

	fromAddress := "sif1rwmv253fqsvcg483vrax9xe3nwcu7nc5fzunur"
	toAddress := "sif1qvt8nm5s40xl9axahe33xzat7jxc8jnwktvgyh"
	amount := "10000rowan"
	chainId := "monkey-bars"

	args := []string{"tx",
					 "send", 			fromAddress, toAddress, amount,
					 "--chain-id", 		chainId,
					 "--generate-only"}

	cmd := exec.Command("sifnodecli", args...)

	outfile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
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

	// yes s6c8kXlIzjxDohE4QdY7HbfSt3VUwnOZ | sifnodecli keys show user1 -a

	fromAddress := "sif1rwmv253fqsvcg483vrax9xe3nwcu7nc5fzunur"

	args := []string{"tx",
					 "sign",       			unsignedTxPath,
					 "--from",     			fromAddress,
					 "--keyring-backend", 	"test"}

	cmd := exec.Command("sifnodecli", args...)

	outfile, err := os.Create(signedTxOutputFile)
	if err != nil {
		log.Fatal(err)
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