package lexer

import (
	"fmt"
	"haskell-interpreter/src/internal/lexer"
	"os"
)

func ExecLexer() {
	code, err := os.ReadFile("lexer/input.txt")
	if err != nil {
		fmt.Println("Error during reading input file:", err)
	}

	lex := lexer.NewLexer(string(code))
	tokens := lex.LexAnalyze()

	writeToFile("results/result.txt", tokens)
	writeToFile("results/keywords.txt", lex.KeywordsTokens)
	writeToFile("results/operators.txt", lex.OperatorsTokens)
	writeToFile("results/names.txt", lex.NamesTokens)
	writeToFile("results/punctuations.txt", lex.PunctuationsTokens)
}

func writeToFile(filename string, tokens []lexer.Token) {
	// Create the results directory if it doesn't exist
	if err := os.MkdirAll("results", os.ModePerm); err != nil {
		fmt.Println("Error creating results directory:", err)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Lexeme", "Token type", "Row", "Column", "Id"))
	file.WriteString("=========================================================\n")

	for _, token := range tokens {
		file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", token.Value, token.Type, token.Line, token.Column, token.ID))
	}
}
