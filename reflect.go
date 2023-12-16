package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var strutchField reflect.StructField = valueType.Field(i)
		fmt.Println(strutchField.Name, "With type", strutchField.Type)
		fmt.Println(strutchField.Tag.Get("required"))
		fmt.Println(strutchField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}

		}
	}
	return result
}

func main() {
	readField(Sample{"Eko"})
	readField(Person{"Eko", "", ""})

	person := Person{
		Name:    "Eko",
		Address: "Jakarta",
		Email:   "",
	}

	fmt.Println(IsValid(person))
}
