// Yet another lexical analyzer, recognizing the following lexemes:
//
// digit   ->  [0-9]
// letter  ->  [a-zA-Z]
// id      ->  letter (letter|digit)*
// relop   ->  < | > | <= | >= | = | <>
package main

import "unicode"

const (
	IF = iota
	ELSE
	ID
	RELOP
)

var (
	keywords = map[string]bool{
		"if":   true,
		"else": true,
	}
)

type Token struct {
	Type  int
	Value string
}

func main() {
	Scan("abc34 321 <= >= >32 dsa< mnn32 <> sdea if else  ")
}

// Scan scans input and returns a list of Token's.
func Scan(input string) []Token {
	var state, pos int
	var t []Token

	for i := 0; i < len(input); i++ {
		c := rune(input[i])

		switch state {
		case 0:
			if c == '<' {
				state = 1
				pos = i
			} else if c == '=' {
				state = 5
				pos = i
			} else if c == '>' {
				state = 6
				pos = i
			} else if unicode.IsLetter(c) {
				state = 9
				pos = i
			}
		case 1:
			if c == '=' {
				state = 2
			} else if c == '>' {
				state = 3
			} else {
				state = 4
			}
		case 2:
			t = append(t, Token{RELOP, "LE"})
			state = 0
		case 3:
			t = append(t, Token{RELOP, "NE"})
			state = 0
		case 4:
			i--
			t = append(t, Token{RELOP, "LT"})
			state = 0
		case 5:
			t = append(t, Token{RELOP, "EQ"})
			state = 0
		case 6:
			if c == '=' {
				state = 7
			} else {
				state = 8
			}
		case 7:
			t = append(t, Token{RELOP, "GE"})
			state = 0
		case 8:
			i--
			t = append(t, Token{RELOP, "GT"})
			state = 0
		case 9:
			if !(unicode.IsLetter(c) || unicode.IsDigit(c)) {
				state = 10
			}
		case 10:
			i--
			if keywords[input[pos:i]] {
				t = append(t, Token{IF, input[pos:i]})
			} else {
				t = append(t, Token{ID, input[pos:i]})
			}
			state = 0
		}
	}

	return t
}
