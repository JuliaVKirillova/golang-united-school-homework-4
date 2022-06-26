package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	output = strings.ReplaceAll(input, " ", "")
	if len(output) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	split := strings.LastIndexAny(output, "+-")
	if split < 1 || split > len(output)-2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	oper1 := output[:split]
	oper2 := output[split:]

	if strings.LastIndexAny(oper1, "+-") > 0 || strings.LastIndexAny(oper2, "+-") != 0 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	num1, err := strconv.ParseInt(oper1, 10, 64)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	num2, err := strconv.ParseInt(oper2, 10, 64)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return fmt.Sprintf("%d", num1+num2), nil
}
