package generator

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TESTS
func TestCreate(t *testing.T) {
	testFileName := "TestCreate.tx"
	fileExists := false

	Create(testFileName)

	if _, err := os.Stat(testFileName); err == nil || os.IsExist(err) {
		fileExists = true
	} else {
		fileExists = false
	}

	assert.True(t, fileExists)
}

func TestSign(t *testing.T) {
	testFileName := "TestSign.tx"
	signedFileName := "TestSign.signed.tx"
	fileExists := false

	Create(testFileName)
	Sign(testFileName, signedFileName)

	if _, err := os.Stat(signedFileName); err == nil || os.IsExist(err) {
		fileExists = true
	} else {
		fileExists = false
	}

	assert.True(t, fileExists)
}
