package validation

import "errors"


func ValidationStudent(name string,age int) error {
	if name == "" {
		return  errors.New("student name cannot be empty")
	}
	if age < 18 {
		return errors.New("student must be at least 18 years old")
	}
	return nil
}


