package main

import (
	"fmt"
	"goprojects2/database"
	"goprojects2/payment"
	"goprojects2/students"
	"goprojects2/validation"
)

func main()  {
	fmt.Println("-------Starting Student Addmission Process--------")

	// 1. Create Student object
	s := students.Student{
		ID: 101,
		Name: "Rahul Sharma",
		Age: 20,
	}

	// Validation Data 
	err := validation.ValidationStudent(s.Name,s.Age)
	if err != nil{
		fmt.Println("[VALIDATION ERROR]:",err)
	}

	// Process Payment 
	paymentSuccess := payment.ProcessFee(s.Name,25000.00)

	// 4. Save to Database if Payment Succeeded
	if paymentSuccess{
		database.SaveToDB(s.Name)
		fmt.Println("[STATUS] Admission Completed Successfully!")
	}

}

// -------Starting Student Addmission Process--------
// [PAYMENT] Processing ₹25000.00 fee for student: Rahul Sharma
// [DB SUCCESS] Data succesfully saved Database: Rahul Sharma
// [STATUS] Admission Completed Successfully!


