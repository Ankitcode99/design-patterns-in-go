package strategies

import "fmt"

type BankTransferPayment struct {
	AccountNumber string
}

func (b *BankTransferPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Bank Transfer from account %s\n", amount, b.AccountNumber)
	return nil
}
