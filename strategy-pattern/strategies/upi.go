package strategies

import "fmt"

type UPIPayment struct {
	UpiId string
}

func (p *UPIPayment) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using UPI account %s\n", amount, p.UpiId)
	return nil
}
