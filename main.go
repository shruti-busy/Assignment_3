package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr1 interface{}
	arr1 = true

	var arr2 interface{} = []interface{}{4, 5, 6, "hello"}

	merger_val, err := merger(arr1, arr2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(merger_val)
	}

}

func merger(arr1 interface{}, arr2 interface{}) (interface{}, error) {

	if arr1 == nil || arr2 == nil {
		if arr1 == nil {
			return arr2, nil
		} else {
			return arr1, nil
		}
	} else {
		var res interface{}
		var final []interface{}

		first := reflect.ValueOf(arr1)

		switch first.Kind() {

		//For Slices

		case reflect.Slice:
			switch elemType := first.Type().Elem().Kind(); elemType {
			case reflect.Interface:
				// To Handle slice of interface{}
				a := first.Interface().([]interface{})
				for _, val := range a {
					final = append(final, val)
				}
			default:
				// Handle other types
				for i := 0; i < first.Len(); i++ {
					final = append(final, first.Index(i).Interface())
				}
			}
			break

		case reflect.Int:
			a := first.Interface().(int)
			final = append(final, a)
			break

		case reflect.String:
			a := first.Interface().(string)
			final = append(final, a)
			break

		case reflect.Float64:
			a := first.Interface().(float64)
			final = append(final, a)

		case reflect.Bool:
			a := first.Interface().(bool)
			final = append(final, a)

		}

		second := reflect.ValueOf(arr2)

		switch second.Kind() {
		case reflect.Slice:
			switch elemType := second.Type().Elem().Kind(); elemType {
			case reflect.Interface:
				// To Handle slice of interface{}
				a := second.Interface().([]interface{})
				for _, val := range a {
					final = append(final, val)
				}
			default:
				// Handle other types
				for i := 0; i < second.Len(); i++ {
					final = append(final, second.Index(i).Interface())
				}
			}
			break

		case reflect.Float64:
			a := second.Interface().(float64)
			final = append(final, a)

		case reflect.Int:
			a := second.Interface().(int)
			final = append(final, a)

		case reflect.String:
			a := second.Interface().(string)
			final = append(final, a)

		case reflect.Bool:
			a := second.Interface().(bool)
			final = append(final, a)

		}
		res = final
		return res, nil
	}
}
