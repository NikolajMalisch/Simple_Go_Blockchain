package main 

import (
	"fmt"
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data       map[string]interface{}
	hash 	   string
	prevHash   string
	timestamp  time.Time
	pow 	   int
}

// Blockchain represents a custom type that contains our blocks
type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

// CalculateHash derives the block hash for our blockchain by hashing the previous block hash,
// current block data, timestamp, and PoW with the SHA256 algorithm

func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// Mine is a method we created for our type that increments the PoW value and
// calculates the block hash until we get a valid hash
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

// CreateBlockchain is a function that creates a genesis block for our blockchain and returns a new instance of the type
func CreateBlockchain(difficulty int) Blockchain {
	// Create the genesis block
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	// Set the hash value of our genesis block. Since this is the first block in the blockchain,
	// there is no value for the previous hash and the data property is empty.
	genesisBlock.hash = genesisBlock.CalculateHash()

	// Return a new instance of the blockchain type with the genesis block
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// AddBlock is a method we created for a type that does the following:
// captures transaction details (sender, receiver, and transfer amount)
// creates a new block with transaction details
// creates a new block using the hash value of the previous block,
// current block data, and generated PoW; adds the newly created block to the blockchain
func (b *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}

	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}

	// Calculate the PoW for the new block
	newBlock.Mine(b.difficulty)

	// Add the newly created block to the blockchain
	b.chain = append(b.chain, newBlock)
}

// IsValid recalculates the hash value of each block, compares it with the stored hash values of the other blocks,
// and checks if the previous hash value of another block is equal to the hash value of the previous block.
// If either of the two checks fails, the blockchain has been tampered with.
func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]

		// Check if the current block hash matches the calculated hash and if the previous hash of the

		// current block matches the hash value of the previous block
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.previousHash != previousBlock.hash {