package main

import (
	"fmt"
	"mgol-go/src/parser"
)

var separator = "=================="

func main() {
	// filePath := os.Args[1]

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// symbolTable := lexer.GetSymbolTableInstance()

	// lexer.FillSymbolTable(symbolTable)
	// defer symbolTable.Cleanup()

	// scanner := lexer.NewScanner(file, symbolTable)
	// for {
	// 	token := scanner.Scan()
	// 	fmt.Println(token)
	// 	if token == lexer.EOF_TOKEN {
	// 		break
	// 	}
	// }
	// fmt.Printf(separator + "\nTabela de símbolos\n" + separator + "\n")
	// symbolTable.Print()

	fmt.Println(parser.GetRulesMap("./parser/grammar.json"))

}
