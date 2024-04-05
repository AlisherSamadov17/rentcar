package check

import (
	"errors"
	"regexp"
	"time"
	"unicode"
)

func ValidateCarYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}
	return nil
}

func ValidateGmailCustomer(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,3}$`)
    return emailRegex.MatchString(e)
}


func ValidatePhoneNumberOfCustomer(phone string) bool {
	if 12 < len(phone) && len(phone) <= 13{
		phoneregex:= regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
		return phoneregex.MatchString(phone)
	}
	return false
}

func ValidatePassword(password string) error {
    if len(password) < 8 {
      return errors.New("password length must be at least 8 characters")
    }
  
    var hasUppercase, hasLowercase, hasDigit, hasSymbol bool
  
    for _, char := range password {
      switch {
      case unicode.IsUpper(char):
        hasUppercase = true
      case unicode.IsLower(char):
        hasLowercase = true
      case unicode.IsNumber(char):
        hasDigit = true
      case unicode.IsPunct(char) || unicode.IsSymbol(char):
        hasSymbol = true
      }
    }
  
    if !hasUppercase {
      return errors.New("password must contain at least one uppercase letter")
    }
    if !hasLowercase {
      return errors.New("password must contain at least one lowercase letter")
    }
    if !hasDigit {
      return errors.New("password must contain at least one digit")
    }
    if !hasSymbol {
      return errors.New("password must contain at least one symbol")
    }
  
    return nil
  }
