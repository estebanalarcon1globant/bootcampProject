package utils

import (
	"fmt"
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

type Transport struct {
	Car
	Bus
}

type Car struct {
	velocity int
}

type Bus struct {
}

func (c *Bus) Run() {
	fmt.Println("Bus is running")
}

func (c *Car) Run() {
	fmt.Println("Car is running")
}
