package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var validNationalID validator.Func = func(fl validator.FieldLevel) bool {
	reg, err := regexp.Compile("/[^0-9]/")
	if err != nil {
		return false
	}
	code := fl.Field().String()
	code = reg.ReplaceAllString(code, "")
	if len(code) != 10 {
		return false
	}
	codes := strings.Split(code, "")
	last, err := strconv.Atoi(codes[9])
	i := 10
	sum := 0
	for in, el := range codes {
		temp, err := strconv.Atoi(el)
		if err != nil {
			return false
		}
		if in == 9 {
			break
		}
		sum += temp * i
		i -= 1
	}
	mod := sum % 11
	if mod >= 2 {
		mod = 11 - mod
	}
	return mod == last
}

var validBirthDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if date.Before(today.AddDate(-7, 0, 0)) && date.After(today.AddDate(-100, 0, 0)) {
			return true
		}
	}
	return false
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("birthdate", validBirthDate)
		v.RegisterValidation("nationalID", validNationalID)
	}
}
