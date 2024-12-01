package validator

import (
	"regexp"

	"maulanadityaa/laundry-app-rest-api/config"
	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/entity"

	"github.com/go-playground/validator/v10"
)

func UniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	var existingEmail entity.Account

	result := config.DB.Where("email = ?", email).First(&existingEmail)

	return result.Error != nil
}

func UniquePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	var existingCustomer entity.Customer
	result := config.DB.Where("phone_number = ?", phoneNumber).First(&existingCustomer)
	if result.Error == nil {
		return false
	}

	var existingEmployee entity.Employee
	result = config.DB.Where("phone_number = ?", phoneNumber).First(&existingEmployee)

	return result.Error != nil
}

func IndonesianPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	phoneNumber = helper.FormatPhoneNumber(phoneNumber)
	internationalRegex := regexp.MustCompile(`^\+628[1-9][0-9]{7,11}$`)

	return internationalRegex.MatchString(phoneNumber)
}
