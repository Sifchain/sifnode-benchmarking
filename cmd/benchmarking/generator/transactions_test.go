package generator

import (
	g "github.com/Sifchain/sifnode-benchmarking/cmd/benchmarking/generator"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TESTS
func TestCreate(t *testing.T) {
	testFileName := "1.tx"
	fileExists := false

	g.Create(testFileName)

	if _, err := os.Stat(testFileName); err == nil || os.IsExist(err) {
		fileExists = true
	} else {
		fileExists = false
	}

	assert.True(t, fileExists)
}
