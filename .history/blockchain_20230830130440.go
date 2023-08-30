package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Definition einer Struktur für einen Block im Blockchain
type Block struct {
	data      map[string]interface{} // Daten für den Block (Transaktion)
	hash      string                 // Hash-Wert des Blocks
	prevHash  string                 // Hash-Wert des vorherigen Blocks
	timestamp time.Time              // Zeitstempel für die Blockerstellung
	pow       int                    // Proof-of-Work (Beweis der Arbeit)
}

// Definition einer Struktur für die gesamte Blockchain
type Blockchain struct {
	genesisBlock Block   // Genesis-Block (erster Block)
	chain        []Block // Kette von Blöcken
	difficulty   int     // Schwierigkeitsgrad für das Mining
}

// Methode zur Berechnung des Hash-Werts eines Blocks
func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.prevHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// Methode zum "Mining" eines Blocks
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

// Funktion zum Erstellen einer Blockchain
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	genesisBlock.hash = genesisBlock.CalculateHash()
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

// Methode zum Hinzufügen eines Blocks zur Blockchain
func (bc *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		data:      blockData,
		hash:      "",
		prevHash:  lastBlock.hash,
		timestamp: time.Now(),
		pow:       0,
	}
	newBlock.Mine(bc.difficulty)
	newBlock.hash = newBlock.CalculateHash()
	bc.chain = append(bc.chain, newBlock)
}

// Methode zur Überprüfung der Integrität der Blockchain
func (bc Blockchain) IsValid() bool {
	for i := 1; i < len(bc.chain); i++ {
		previousBlock := bc.chain[i-1]
		currentBlock := bc.chain[i]
		if currentBlock.hash != currentBlock.CalculateHash() || currentBlock.prevHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	difficulty := 2
	blockchain := CreateBlockchain(difficulty)

	blockchain.AddBlock("Alice", "Bob", 1.5)
	blockchain.AddBlock("Bob", "Charlie", 0.7)

	fmt.Println("Blockchain ist gültig:", blockchain.IsValid())
}
