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

type TransactionRequest struct {
	SenderAddress    *string  `json:"sender_address"`
	RecipientAddress *string  `json:"recipient_address"`
	SenderPublicKey  *string  `json:"sender_public_key"`
	Value            *float32 `json:"value"`
	Signature        *string  `json:"signature"`
}

func (tr *TransactionRequest) Validate() string {
	if tr.Signature == nil {
		return "Signature is missing in the request"
	}

	if tr.SenderAddress == nil {
		return "SenderAddress is missing in the request"
	}

	if tr.RecipientAddress == nil {
		return "RecipientAddress is missing in the request"
	}

	if tr.SenderPublicKey == nil {
		return "SenderPublicKey is missing in the request"
	}

	if tr.Value == nil {
		return "Value is missing in the request"
	}

	return ""
}

type AmountResponse struct {
	Amount float32 `json:"amount"`
}
