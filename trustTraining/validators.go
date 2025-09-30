package main

import "github.com/go-playground/validator/v10"


var validNationalID validator.Func = func(fl validator.FieldLevel) bool {
	// TODO
	// کد ملی را اعتبارسنجی کنید
    return true
}

var validBirthDate validator.Func = func(fl validator.FieldLevel) bool {
	// TODO
	// تاریخ تولد را اعتبارسنجی کنید
    return true
}

func registerValidators() {
	// TODO
	// تابع‌های اعتبارسنجی را ثبت کنید
}