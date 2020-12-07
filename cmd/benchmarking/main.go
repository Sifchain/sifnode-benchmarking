package main

import (
	"fmt"
	g "github.com/Sifchain/sifnode-benchmarking/cmd/benchmarking/generator"
)

func main() {
	fmt.Printf("Hello, world!")
	g.Create("test.tx")
}