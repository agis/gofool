package rpn

import (
	"github.com/Agis-/gofool/structures/stack"
	"strings"
	"strconv"
	"regexp"
	"log"
)

func Eval(exp string) float64 {
	reOperand, _ := regexp.Compile(`^\d+(\.\d+)?$`)
	reOperator, _ := regexp.Compile(`\+|-|\*|/`)

	stack := new(stack.Stack)

	for _, tok := range strings.Fields(exp) {
		if reOperand.MatchString(tok) {
			tok, _ := strconv.ParseFloat(tok, 64)
			stack.Push(tok)
		} else if reOperator.MatchString(tok) {
			if stack.Length < 2 {
				log.Fatal("Error: Insufficient values for expression")
			}

			a := stack.Pop()
			b := stack.Pop()

			switch tok {
			case "+":
				stack.Push(a + b)
			case "-":
				stack.Push(a - b)
			case "*":
				stack.Push(a * b)
			case "/":
				stack.Push(a / b)
			}
		} else {
			log.Fatalf("Error: unrecognized token %s", tok)
		}
	}

	if stack.Length != 1 {
		log.Fatal("Error: too many values provided")
	}

	return stack.Pop()
}
