package validator

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidateProduceCode(produceCode string) error {
	if len(produceCode) != 19 {
		fmt.Printf("len %d", len(produceCode))
		return fmt.Errorf("invalid produce code")
	}
	r, _ := regexp.Compile("[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}")
	if valid := r.MatchString(produceCode); !valid {
		return fmt.Errorf("invalid produce code")
	}

	return nil
}

func ValidatePrice(price float64) error {
	precisionPrice := fmt.Sprintf("%.2f", price)
	f, _ := strconv.ParseFloat(precisionPrice, 64)
	if f != price {
		return fmt.Errorf("invalid price")
	}
	return nil
}

func ValidateName(name string) error {
	r, _ := regexp.Compile("^[a-zA-Z0-9_]+( [a-zA-Z0-9_]+)*$")
	if valid := r.MatchString(name); !valid {
		return fmt.Errorf("name must be alphanumeric")
	}
	return nil

}
