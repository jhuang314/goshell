%{
package parser

import (
    "fmt"
    "log"
)

var (
    src      string
    buf      []byte
    current  byte
    currentIndex        int
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

func Tokenize(in string) (tokens []string){ // This left brace is closed by *1
     src = in
     current = 0
     currentIndex = 0
     buf = make([]byte, 100)
     c := getc() // init
%}

%yyc c
%yyn c = getc()

PIPE \|
REDIRECTION     >|>>|<|<<
SYM  [a-zA-Z0-9äöüÄÖÜ+_-]+
STRING  \"(\\.|[^\\"])*\"|\'(\\.|[^\\'])*\'

%%
    buf = buf[:0]   // Code before the first rule is executed before every scan cycle (state 0 action)

[ \t\n\r]+          //addToken()

{PIPE}              tokens = append(tokens, fmt.Sprintf("%s", buf))

{REDIRECTION}       tokens = append(tokens, fmt.Sprintf("%s", buf))
                    
{STRING}            tokens = append(tokens, fmt.Sprintf("%s", buf))

{SYM}               tokens = append(tokens, fmt.Sprintf("%s", buf))

\0                  return // Exit on EOF or any other error



%%
    // The rendered scanner enters top of the user code section when
    // lexem recongition fails. In this example it should never happen.
    log.Fatal("scanner internal error")



    return
} // *1 this right brace



// func main() {
//      tokens := Tokenize("ls |wc")
//      fmt.Println(tokens)
// }
