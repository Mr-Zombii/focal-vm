package main

import (
	"fmt"
	"focal-lang/internal/lexer"
)

func main() {
	ll := lexer.NewLexer("+=-=")
	for {
		var t = ll.NextToken()
		fmt.Println(t.GetContents())
		if t == lexer.TOKEN_EOF {
			break
		}
	}
}
