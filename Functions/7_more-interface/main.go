package main

import "fmt"

type paymentMethod interface {
	processPayment(amount float64) string
}

type creditCard struct {
	cardNumber string
}

type payPal struct {
	email string
}

func (c creditCard) processPayment(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using CreditCard.", amount)
}

func (p payPal) processPayment(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

func processPayment(p paymentMethod, amount float64) {
	fmt.Println(p.processPayment(amount))
}

func main() {
//creating different payment methods
  cc := creditCard{cardNumber: "1234-5678-9876-5432"}
  pp := payPal{email: "user@example.com"}

 // process payments
 processPayment(cc, 100.50)
 processPayment(pp, 75.25)

}
