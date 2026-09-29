package main

import "fmt"

func main()  {

								// Simple Switch Statement 
	// day := 2 
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")

	// case 4:
	// 	fmt.Println("Thursday")

	// default:
	// 	fmt.Println("Invalid day")
	// }
	

						// Multiple Values in a Case 
		//👉 Ek single case ke andar multiple values specify ki ja sakti hain.
	day := 6
	switch day{
	case 1,2,3,4,5:
		fmt.Println("Weekday")

	case 6,7:
		fmt.Println("Weekend")

	default:
		fmt.Println("Invalid day")
	}


						//switch Without an Expression
	
	marks := 85 
	switch{
	case marks >= 90:
		fmt.Println("Grade A")

	case marks >= 80:
		fmt.Println("Grade B")

	case marks >= 70:
		fmt.Println("Grade C")

	default:
		fmt.Println("Fail")
	}



						// fallthrough:- keyword next wale case ko bhi jabardasti execute karwa deta hai.


	num := 1 
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
	
}
