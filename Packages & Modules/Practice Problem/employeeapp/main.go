package main

import "employeeapp/employee"

func main()  {
	emp:=employee.Employee{
		ID: 501,
		Name: "Priya",
		Department: "IT",
		Salary: 60000,

	}
	employee.DisplayEmployee(emp)
}

// Employee ID: 501
// Name: Priya
// Total Salary (with Bonus): 66000



// [ Command: go run main.go ]
//              │
//              ▼
// [ Resolves Module: employeeapp ]
//              │
//              ▼
// [ Imports package "employeeapp/employee" ]
//              │
//              ▼
// [ Enters func main() ]
//              │
//              ├─► Instantiates Employee struct: emp (Salary: 60000)
//              │
//              └─► Calls: employee.DisplayEmployee(emp)
//                             │
//                             ├─► Prints: Employee ID: 501
//                             ├─► Prints: Name: Priya
//                             └─► Calls: emp.CalculateSalary()
//                                            │
//                                            ├─► Reads: e.Salary (60000)
//                                            ├─► Calls: calculateBonus(60000) ──► Returns: 6000
//                                            └─► Adds: 60000 + 6000           ──► Returns: 66000
//                             │
//                             └─► Prints: Total Salary (with Bonus): 66000
//              │
//              ▼
// [ Program Ends Successfully ]


