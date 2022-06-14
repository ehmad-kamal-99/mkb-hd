package main

import "fmt"

// Problem Statement
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.

/*
Notes

I think a queue type of solution is required. You basically push each encountered open parenthesis into a queue.
As soon as you encounter a closing parenthesis, you checkNext it to the last one pushed
if matches, pop, else error.

What I missed
- Single Parenthesis Input
- Only open parenthesis
- First parenthesis is closing
- If single parenthesis is passed to validate, runtime error
*/

func main() {
	// call function from here
	fmt.Println(validParenthesis("){"))
}

func validParenthesis(s string) bool {
	var stack []string
	var validation = true

	if len(s) < 2 {
		return false
	}

	for _, each := range s {
		if string(each) == "(" || string(each) == "{" || string(each) == "[" {
			stack = append(stack, string(each))
		} else if string(each) == ")" || string(each) == "}" || string(each) == "]" {
			stack = append(stack, string(each))
			stack, validation = validate(stack, string(each))
			if !validation {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func validate(stack []string, s string) ([]string, bool) {
	if len(stack) < 2 {
		return stack, false
	}

	switch s {
	case ")":
		if stack[len(stack)-2] == "(" {
			stack = stack[:len(stack)-2]
			return stack, true
		}

		return stack, false
	case "}":
		if stack[len(stack)-2] == "{" {
			stack = stack[:len(stack)-2]
			return stack, true
		}

		return stack, false
	case "]":
		if stack[len(stack)-2] == "[" {
			stack = stack[:len(stack)-2]
			return stack, true
		}

		return stack, false
	default:
		return stack, false
	}
}

// Challenge:
