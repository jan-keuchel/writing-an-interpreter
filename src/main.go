package main

import (
	"fmt"
	"os"

	"github.com/jan-keuchel/writing-an-interpreter/src/lexer"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong format. use './binary <file>'")
		return
	}
	file := os.Args[1]

	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Erro reading file: %s", err)
	}
	code := string(bytes)

	fmt.Printf("Code: \n%s\n", code)

	lexer := lexer.NewLexer(code)
	lexer.LexCode()

}

