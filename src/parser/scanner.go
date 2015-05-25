// CAUTION: Generated file - DO NOT EDIT.

package parser

import (
	"fmt"
	"log"
)

var (
	src          string
	buf          []byte
	current      byte
	currentIndex int
)

func getc() byte {
	if current != 0 {
		buf = append(buf, current)
	}
	current = 0
	if currentIndex < len(src) {
		current = src[currentIndex]
		currentIndex++
	}

	return current
}

//    %yyc is a "macro" to access the "current" character.
//
//    %yyn is a "macro" to move to the "next" character.
//
//    %yyb is a "macro" to return the begining-of-line status (a bool typed value).
//        It is used for patterns like `^re`.
//        Example: %yyb prev == 0 || prev == '\n'
//
//    %yyt is a "macro" to return the top/current start condition (an int typed value).
//        It is used when there are patterns with conditions like `<cond>re`.
//        Example: %yyt startCond

func Tokenize(in string) (tokens []string) { // This left brace is closed by *1
	src = in
	current = 0
	currentIndex = 0
	buf = make([]byte, 100)
	c := getc() // init

yystate0:

	buf = buf[:0] // Code before the first rule is executed before every scan cycle (state 0 action)

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = getc()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate4
	case c == '+' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084' || c == '\u0096' || c == '\u009c' || c == '¤' || c == '¶' || c == '¼' || c == 'Ã':
		goto yystate9
	case c == '<':
		goto yystate10
	case c == '>':
		goto yystate12
	case c == '\'':
		goto yystate7
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c == '|':
		goto yystate13
	}

yystate2:
	c = getc()
	goto yyrule6

yystate3:
	c = getc()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	}

yystate4:
	c = getc()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate5
	case c == '\\':
		goto yystate6
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = getc()
	goto yyrule4

yystate6:
	c = getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate4
	}

yystate7:
	c = getc()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate5
	case c == '\\':
		goto yystate8
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = getc()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate7
	}

yystate9:
	c = getc()
	switch {
	default:
		goto yyrule5
	case c == '+' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0084' || c == '\u0096' || c == '\u009c' || c == '¤' || c == '¶' || c == '¼' || c == 'Ã':
		goto yystate9
	}

yystate10:
	c = getc()
	switch {
	default:
		goto yyrule3
	case c == '<':
		goto yystate11
	}

yystate11:
	c = getc()
	goto yyrule3

yystate12:
	c = getc()
	switch {
	default:
		goto yyrule3
	case c == '>':
		goto yystate11
	}

yystate13:
	c = getc()
	goto yyrule2

yyrule1: // [ \t\n\r]+
	{
		//addToken()
		goto yystate0
	}
yyrule2: // {PIPE}
	{
		tokens = append(tokens, fmt.Sprintf("%s", buf))
		goto yystate0
	}
yyrule3: // {REDIRECTION}
	{
		tokens = append(tokens, fmt.Sprintf("%s", buf))

		goto yystate0
	}
yyrule4: // {STRING}
	{
		tokens = append(tokens, fmt.Sprintf("%s", buf))
		goto yystate0
	}
yyrule5: // {SYM}
	{
		tokens = append(tokens, fmt.Sprintf("%s", buf))
		goto yystate0
	}
yyrule6: // \0
	{
		return // Exit on EOF or any other error
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	// The rendered scanner enters top of the user code section when
	// lexem recongition fails. In this example it should never happen.
	log.Fatal("scanner internal error")

	return
} // *1 this right brace

// func main() {
//      tokens := Tokenize("ls |wc")
//      fmt.Println(tokens)
// }
