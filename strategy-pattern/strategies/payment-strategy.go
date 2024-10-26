package strategies

type PaymentStrategy interface {
	Pay(amount float64) error
}
