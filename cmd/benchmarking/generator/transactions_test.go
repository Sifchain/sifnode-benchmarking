package generator

import (
	"github.com/sifnode-benchmarking/cmd/benchmarking/generator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

// SETUP TESTIFY
type TransactionsTestSuite struct {
	suite.Suite
}

// MANAGE TEMPORARY FILES
func (suite *TransactionsTestSuite) SetupTestSuite() {
	os.MkdirAll("test_files/", os.ModePerm)
}
func (suite *TransactionsTestSuite) TearDownTestSuite() {
	os.RemoveAll("test_files/")
}

// TESTS
func (suite *TransactionsTestSuite) TestCreate() {
	testFileName := "tmp/1.tx"
	fileExists := false

	generator.Create(testFileName)

	if _, err := os.Stat("tmp/1.tx"); err == nil || os.IsExist(err) {
		fileExists = true
	} else {
		fileExists = false
	}

	assert.True(suite.T(), fileExists)
}

// LINK TESTIFY AND TESTING PACKAGES
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsTestSuite))
}