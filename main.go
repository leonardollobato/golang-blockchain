package main

import (
	"fmt"
	"strconv"

	"github.com/leonardollobato/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First  Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third  Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("\tPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
