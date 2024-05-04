package block

import (
	"fmt"
	"strings"
)

type Transaction struct {
	SenderAddress    string  `json:"sender_address"`
	RecipientAddress string  `json:"recipient_address"`
	Value            float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) PrintTransaction() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_address      %s\n", t.SenderAddress)
	fmt.Printf(" recipient_address   %s\n", t.RecipientAddress)
	fmt.Printf(" value               %.1f\n", t.Value)
}
