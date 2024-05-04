package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        int64          `json:"nonce"`
	PreviousHash [32]byte       `json:"previous_hash"`
	Transactions []*Transaction `json:"transactions"`
	Timestamp    int64          `json:"timestamp"`
}

func NewBlock(nonce int64, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Transactions: transactions,
		Timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) PrintBlock() {
	fmt.Printf("timestamp       %d\n", b.Timestamp)
	fmt.Printf("nonce           %d\n", b.Nonce)
	fmt.Printf("previous_hash   %x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.PrintTransaction()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}
