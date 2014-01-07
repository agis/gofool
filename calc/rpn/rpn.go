package rpn

import (
	"github.com/Agis-/gofool/structures/stack"
	"strings"
	"strconv"
	"regexp"
	"log"
)

func Eval(exp string) float64 {
	re1, _ := regexp.Compile(`^\d+(\.\d+)?$`)
	re2, _ := regexp.Compile(`\+|-|\*|/`)

	stack := new(stack.Stack)

	for _, tok := range strings.Fields(exp) {
		if re1.MatchString(tok) {
			tok, _ := strconv.ParseFloat(tok, 64)
			stack.Push(tok)
		} else if re2.MatchString(tok) {
			if stack.Length < 2 {
				log.Fatal("Error: Insufficient values for expression")
			}

			a := stack.Pop()
			b := stack.Pop()

			if tok == "+" {
				stack.Push(a + b)
			} else if tok == "-" {
				stack.Push(a - b)
			} else if tok == "*" {
				stack.Push(a * b)
			} else if tok == "/" {
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
