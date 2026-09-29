package payment

import "fmt"

func ProcessFee(studentName string,amount float64) bool {
	fmt.Printf("[PAYMENT] Processing ₹%.2f fee for student: %s\n", amount,studentName)
	return true
}
