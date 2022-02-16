package utils

import (
	"reflect"
)

func CheckEmptyField(req interface{}) bool {
	v := reflect.ValueOf(req)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return true
		}
	}
	return false
}
