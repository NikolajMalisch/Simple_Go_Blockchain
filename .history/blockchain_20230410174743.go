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
	timestamp  int64
