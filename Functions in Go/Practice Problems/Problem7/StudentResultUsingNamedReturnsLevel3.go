package main

import (
	"fmt"
)

func calculateResult(math,Physics,cs int) (total int,average float64,result string) {
	total = math + Physics + cs 
	average = float64(total)/3.0

	if math >= 40 && Physics >= 40 && cs >= 40{
		result = "Pass"
	}else{
		result = "Fail"
	}
	return
}

func main()  {
	var m,p,c int 
	fmt.Scan(&m,&p,&c)
	total , avg , res := calculateResult(m,p,c)
	fmt.Println("Total =",total)
	fmt.Println("Average = %2f",avg)
	fmt.Println("Result",res)

}


