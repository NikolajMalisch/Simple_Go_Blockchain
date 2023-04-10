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


