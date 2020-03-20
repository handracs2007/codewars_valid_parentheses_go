package main

import "fmt"

func ValidParentheses(parens string) bool {
	var openCount = 0

	for _, val := range parens {
		if val == '(' {
			openCount++
		} else {
			if openCount == 0 {
				return false
			} else {
				openCount--
			}
		}
	}

	return openCount == 0
}

func main() {
	fmt.Println(ValidParentheses("()"))
	fmt.Println(ValidParentheses(")(()))"))
	fmt.Println(ValidParentheses("("))
	fmt.Println(ValidParentheses("(())((()())())"))
}
