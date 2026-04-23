package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []Block
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s",
		block.Index,
		block.Timestamp,
		block.Data,
		block.PreviousHash,
	)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func newBlock(index int, data string, previousHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().Format(time.RFC3339Nano),
		Data:         data,
		PreviousHash: previousHash,
	}

	block.Hash = calculateHash(block)
	return block
}

func createBlockchain() Blockchain {
	genesisBlock := newBlock(0, "Genesis Block", "")

	return Blockchain{
		Blocks: []Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := newBlock(lastBlock.Index+1, data, lastBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		if calculateHash(currentBlock) != currentBlock.Hash {
			return false
		}
	}

	return true
}

func printBlockchain(blockchain Blockchain) {
	for i, block := range blockchain.Blocks {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("  Index:        %d\n", block.Index)
		fmt.Printf("  Timestamp:    %s\n", block.Timestamp)
		fmt.Printf("  Data:         %s\n", block.Data)
		fmt.Printf("  PreviousHash: %s\n", block.PreviousHash)
		fmt.Printf("  Hash:         %s\n", block.Hash)

		if i > 0 {
			fmt.Printf("  Linked to previous? %t\n",
				block.PreviousHash == blockchain.Blocks[i-1].Hash)
		}

		fmt.Println("--------------------------------------------------")
	}
}

func main() {
	blockchain := createBlockchain()

	blockchain.AddBlock("First block after genesis")
	blockchain.AddBlock("Second block after genesis")
	blockchain.AddBlock("Third block after genesis")

	fmt.Println()
	fmt.Println("=== BLOCKCHAIN ===")
	printBlockchain(blockchain)
	fmt.Println("Blockchain valid?", blockchain.IsValid())

	// Tampering test
	blockchain.Blocks[1].Data = "HACKED BLOCK DATA"
	fmt.Println("After tampering...")
	fmt.Println("=== BLOCKCHAIN ===")
	printBlockchain(blockchain)
	fmt.Println("Blockchain valid?", blockchain.IsValid())
}
