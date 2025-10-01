package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validNationalID validator.Func = func(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(string)
	if ok {
		idList := strings.Split(id, "")
		if len(idList) != 10 {
			return false
		}
		idListInt := make([]int, 10)
		for i := 0; i < 10; i++ {
			idListInt[i], _ = strconv.Atoi(idList[i])
		}
		controlNumber := idListInt[9] 

		sum := 0
		for i := 8; i >= 0; i-- { 
			position := 10 - i 
			sum += idListInt[i] * position
		}
		
		remainder := sum % 11
		if remainder > 2 {
			return remainder + controlNumber == 11
		} else {
			return remainder == controlNumber
		}
	}
	return false
}

var validBirthDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		now := time.Now()
		minDate := now.Add(-100 * 365 * 24 * time.Hour)
		maxDate := now.Add(-7 * 365 * 24 * time.Hour)
		return date.After(minDate) && date.Before(maxDate)
	}
	return false
}


func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_birth_date", validBirthDate)
		v.RegisterValidation("valid_national_id", validNationalID)
	}
}
