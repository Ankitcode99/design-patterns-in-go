package main

import (
	payment "github.com/Ankitcode99/design-patterns-in-go/strategy-pattern/strategies"
)

type PaymentContext struct {
	strategy payment.PaymentStrategy
}

func NewPaymentContext() *PaymentContext {
	return &PaymentContext{}
}

// SetStrategy changes the payment method
func (pc *PaymentContext) SetStrategy(strategy payment.PaymentStrategy) {
	pc.strategy = strategy
}

// ExecutePayment processes the payment using the selected strategy
func (pc *PaymentContext) ExecutePayment(amount float64) error {
	return pc.strategy.Pay(amount)
}
