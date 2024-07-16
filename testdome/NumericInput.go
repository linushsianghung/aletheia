package testdome

import (
	"fmt"
	"unicode"
)

/**
User interface contains NumericInput control, which accepts only digits.

Extend NumericInput structure so that:

It implements UserInput interface.
Add(rune) should add only decimal digits to the input. Other runes should be ignored.
GetValue() should return the current input.
 */
type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(r rune) {
	if unicode.IsDigit(r) {
		n.input = n.input + string(r)

		// i, _ := strconv.Atoi(n.input)
		// n.input = strconv.Itoa(i + int(r - '0'))
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}

func NumericInputs() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}

