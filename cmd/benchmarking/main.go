package main

import (
	"fmt"
	"github.com/Sifchain/sifnode-benchmarking/cmd/benchmarking/generator"
)

func main() {
	fmt.Printf("Hello, world!")
	generator.Create("1.tx")
}