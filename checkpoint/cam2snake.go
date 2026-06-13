// Instructions

// Write a function that converts a string from camelCase to snake_case.

//     If the string is empty, return an empty string.
//     If the string is not camelCase, return the string unchanged.
//     If the string is camelCase, return the snake_case version of the string.

// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

//     lowerCamelCase
//     UpperCamelCase

// Rules for writing in camelCase:

//     The word does not end on a capitalized letter (CamelCasE).
//     No two capitalized letters shall follow directly each other (CamelCAse).
//     Numbers or punctuation are not allowed in the word anywhere (camelCase1).

// Expected function

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }

// And its output:

// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2

package main

import (
	"fmt"
	"unicode"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	var result []rune
	runes := []rune(s)
	for i, r := range runes {
		if i == len(runes)-1 && unicode.IsUpper(runes[i]) {
			return s
		}
		if !unicode.IsLetter(r) {
			return s
		}
		// if i < len(runes)-1 && (unicode.IsUpper(runes[i])) && (unicode.IsUpper(runes[i+1])) { // manual thinking to handle consecutive uppercae letters
		// 	runes[i] = runes[i] + ('a' - 'A')
		// }
		if i > 0 && (unicode.IsLower(runes[i-1])) && (unicode.IsUpper(r)) { // places underscore just before the index under consideration
			result = append(result, '_')
		}
		result = append(result, (r))

	}
	return string(result)
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASe"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
