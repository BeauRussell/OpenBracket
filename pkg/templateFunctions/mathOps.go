package templateFunctions

import (
	"fmt"
	"reflect"
	"text/template"
)

var MathOps template.FuncMap = template.FuncMap{
	"len": GenericLen,
	"sub": func(a, b int) int {
		return a - b
	},
	"add": func(a, b int) int {
		return a + b
	},
	"mod": func(a, b int) int {
		return a % b
	},
	"div": func(a, b int) int {
		return a / b
	},
	"mul": func(a, b int) int {
		return a * b
	},
}

func GenericLen(slice interface{}) int {
	v := reflect.ValueOf(slice)

	// Check if the input is a slice
	if v.Kind() == reflect.Slice {
		return v.Len()
	}
	// If it's not a slice, return 0 or handle error
	fmt.Printf("Invalid type: expected slice, got %s\n", v.Kind())
	return 0
}
