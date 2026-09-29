package validation

import "errors"

// Rule- checking function

func ValidationAmount(amount float64) error {
	if amount <= 0{
		return errors.New("payment amount must be greater than zero")
	}
	return nil
}

