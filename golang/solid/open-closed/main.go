package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying $%.2f with credit card...\n", amount)
	return nil
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) error {
	fmt.Printf("Paying $%.2f with PayPal...\n", amount)
	return nil
}

type BankTransfer struct{}

func (b BankTransfer) Pay(amount float64) error {
	fmt.Printf("Paying $%.2f with bank transfer...\n", amount)
	return nil
}

type PaymentProcessor struct{}

func (p PaymentProcessor) ProcessPayment(amount float64, method PaymentMethod) error {
	return method.Pay(amount)
}

func main() {
	var processor PaymentProcessor

	err := processor.ProcessPayment(42.99, CreditCard{})
	if err != nil {
		fmt.Println("Payment failed:", err)
	}

	err = processor.ProcessPayment(100.0, PayPal{})
	if err != nil {
		fmt.Println("Payment failed:", err)
	}

	err = processor.ProcessPayment(500.0, BankTransfer{})
	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}
