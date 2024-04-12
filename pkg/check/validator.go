package check

import (
	"errors"
	"time"
	"regexp"
	// "clone/rent_car_us/config"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}


	
	func ValidateEmail(email string) error {
		emailRegex := `^[a-zA-Z0-9._%+-]+@(?:gmail|email)+(?:com|ru)$`
		regex := regexp.MustCompile(emailRegex)
		if regex.MatchString(email) {
			return nil
		} else {
			return errors.New("email is not valid")
		}
	}

	
	
	func ValidatePassword(password string) error {
		lowercaseRegex := `[a-z]`
		hasLowercase, _ := regexp.MatchString(lowercaseRegex, password)
		
		uppercaseRegex := `[A-Z]`
		hasUppercase, _ := regexp.MatchString(uppercaseRegex, password)

		digitRegex := `[0-9]`
		hasDigit, _ := regexp.MatchString(digitRegex, password)
		
		symbolRegex := `[!@#$%^&*()-_+=~\[\]{}|\\:;"'<>,.?\/]`
		hasSymbol, _ := regexp.MatchString(symbolRegex, password)
	
		if hasLowercase && hasUppercase && hasDigit && hasSymbol && len(password) >= 8 {
			return nil
		}
	
		return errors.New("password does not meet the criteria")
	}
	
	
	
	
	func ValidatePhone(phone string) error {
		
		phoneRegex := `^\+998\d{9}$`
	
		regex := regexp.MustCompile(phoneRegex)
	
		if regex.MatchString(phone) {
			return nil
		} else {
			return errors.New("phone number is not valid")
		}
	}
	
	var ORDER_STATUS = []string{
		"new", "in-process", "finished", "canceled",
	}


	func CheckOrderStatus(status string) error {
		for i:=0;i<4;i++ {
			if ORDER_STATUS[i] == status {
				return nil
			}
		}
		return errors.New("error: Invalid order status")
	}
