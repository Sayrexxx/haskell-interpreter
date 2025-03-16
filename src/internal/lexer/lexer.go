package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

// Token represents a token with type and value.
type Token struct {
	Type   string
	Value  string
	Line   int
	Column int
	ID     string
}

// Lexer presents a lexical analyzer.
type Lexer struct {
	code               string
	pos                int
	tokens             []Token
	NamesTokens        []Token
	KeywordsTokens     []Token
	OperatorsTokens    []Token
	PunctuationsTokens []Token

	namesIndex        int
	tableNames        map[string]int
	operatorsIndex    int
	tableOperators    map[string]int
	punctuationsIndex int
	tablePunctuations map[string]int
	keywordsIndex     int
	tableKeywords     map[string]int
}

// NewLexer creates a new lexical analyzer.
func NewLexer(code string) *Lexer {
	return &Lexer{
		code:              code,
		pos:               0,
		tableNames:        make(map[string]int),
		tableOperators:    make(map[string]int),
		tablePunctuations: make(map[string]int),
		tableKeywords:     make(map[string]int),
	}
}

// LexAnalyze performs lexical analysis.
func (l *Lexer) LexAnalyze() []Token {
	var errors []string
	for {
		success, err := l.nextToken()
		if !success {
			break
		}
		if err != "" {
			errors = append(errors, err)
		}
	}

	for _, err := range errors {
		fmt.Println(err)
	}

	return l.tokens
}

// nextToken finds the next token.
func (l *Lexer) nextToken() (bool, string) {
	if l.pos >= len(l.code) {
		return false, ""
	}

	text := l.code[l.pos:]

	// Skip multi-line comments
	if strings.HasPrefix(text, "{-") {
		endCommentPos := strings.Index(text, "-}")
		if endCommentPos == -1 {
			line, col := l.charToLineCol(l.pos)
			l.pos = len(l.code)
			return true, fmt.Sprintf("Error: unclosed multiline comment at position %d:%d", line, col)
		}
		l.pos += endCommentPos + 2
		return true, ""
	}

	// Skip single line comments
	if strings.HasPrefix(text, "--") {
		nextLinePos := strings.Index(text, "\n")
		if nextLinePos == -1 {
			l.pos = len(l.code)
		} else {
			l.pos += nextLinePos + 1
		}
		return true, ""
	}

	// Search for tokens
	for _, tokenType := range getTokenTypesList() {
		regex := regexp.MustCompile("^" + tokenType.Regex)
		match := regex.FindString(text)

		if match != "" {
			line, col := l.charToLineCol(l.pos)
			token := Token{
				Type:   tokenType.Name,
				Value:  match,
				Line:   line,
				Column: col,
			}

			switch tokenType.Class {
			case "keyword":
				if _, exists := l.tableKeywords[match]; !exists {
					l.tableKeywords[match] = l.keywordsIndex
					l.keywordsIndex++
				}
				token.ID = fmt.Sprintf("K:%d", l.tableKeywords[match])
				l.KeywordsTokens = append(l.KeywordsTokens, token)
			case "operator":
				if _, exists := l.tableOperators[match]; !exists {
					l.tableOperators[match] = l.operatorsIndex
					l.operatorsIndex++
				}
				token.ID = fmt.Sprintf("O:%d", l.tableOperators[match])
				l.OperatorsTokens = append(l.OperatorsTokens, token)
			case "variable", "constant":
				if _, exists := l.tableNames[match]; !exists {
					l.tableNames[match] = l.namesIndex
					l.namesIndex++
				}
				token.ID = fmt.Sprintf("N:%d", l.tableNames[match])
				l.NamesTokens = append(l.NamesTokens, token)
			case "punctuation":
				if _, exists := l.tablePunctuations[match]; !exists {
					l.tablePunctuations[match] = l.punctuationsIndex
					l.punctuationsIndex++
				}
				token.ID = fmt.Sprintf("P:%d", l.tablePunctuations[match])
				l.PunctuationsTokens = append(l.PunctuationsTokens, token)
			case "preprocessor":
				token.ID = "PREPROCESSOR"
			}

			l.tokens = append(l.tokens, token)
			l.pos += len(match)
			return true, ""
		}
	}

	// Error handling
	line, col := l.charToLineCol(l.pos)
	err := fmt.Sprintf("Error: unknown character at position %d:%d", line, col)
	l.pos++
	return true, err
}

// charToLineCol converts the position in the row and column.
func (l *Lexer) charToLineCol(pos int) (int, int) {
	if pos < 0 || pos >= len(l.code) {
		return -1, -1
	}

	line := 1
	col := 1

	for i := 0; i < pos; i++ {
		if l.code[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}

	return line, col
}
