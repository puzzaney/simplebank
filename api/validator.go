package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/puzzaney/simplebank/util"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		return util.IsSupportedCurency(currency)
	}

	return false
}
