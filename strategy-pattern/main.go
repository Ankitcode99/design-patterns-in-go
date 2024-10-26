package main

import (
	payment "github.com/Ankitcode99/design-patterns-in-go/strategy-pattern/strategies"
)

func main() {
	paymentCtx := NewPaymentContext()

	// Use credit card payment
	creditCard := &payment.CreditCardPayment{
		CardNumber: "1234-5678-9012-3456",
		Cvv:        "123",
	}
	paymentCtx.SetStrategy(creditCard)
	paymentCtx.ExecutePayment(100.50)

	// Switch to PayPal payment
	upi := &payment.UPIPayment{
		UpiId: "someone@okhdfcbank",
	}
	paymentCtx.SetStrategy(upi)
	paymentCtx.ExecutePayment(50.75)

	// Switch to bank transfer payment
	bankTransfer := &payment.BankTransferPayment{
		AccountNumber: "987654321",
	}
	paymentCtx.SetStrategy(bankTransfer)
	paymentCtx.ExecutePayment(75.25)
}
