package lexer

// TokenType represents the type of the token.
type TokenType struct {
	Name  string
	Regex string
	Class string
}

// getTokenTypesList returns a list of all token types.
func getTokenTypesList() []TokenType {
	keywords := []TokenType{
		{"let", `\blet\b`, "keyword"},
		{"in", `\bin\b`, "keyword"},
		{"where", `\bwhere\b`, "keyword"},
		{"if", `\bif\b`, "keyword"},
		{"then", `\bthen\b`, "keyword"},
		{"else", `\belse\b`, "keyword"},
		{"case", `\bcase\b`, "keyword"},
		{"of", `\bof\b`, "keyword"},
		{"data", `\bdata\b`, "keyword"},
		{"type", `\btype\b`, "keyword"},
		{"class", `\bclass\b`, "keyword"},
		{"instance", `\binstance\b`, "keyword"},
		{"deriving", `\bderiving\b`, "keyword"},
		{"do", `\bdo\b`, "keyword"},
		{"module", `\bmodule\b`, "keyword"},
		{"import", `\bimport\b`, "keyword"},
		{"as", `\bas\b`, "keyword"},
		{"hiding", `\bhiding\b`, "keyword"},
		{"qualified", `\bqualified\b`, "keyword"},
		{"infix", `\binfix\b`, "keyword"},
		{"infixl", `\binfixl\b`, "keyword"},
		{"infixr", `\binfixr\b`, "keyword"},
		{"foreign", `\bforeign\b`, "keyword"},
		{"export", `\bexport\b`, "keyword"},
		{"safe", `\bsafe\b`, "keyword"},
		{"unsafe", `\bunsafe\b`, "keyword"},
		{"mdo", `\bmdo\b`, "keyword"},
		{"family", `\bfamily\b`, "keyword"},
		{"role", `\brole\b`, "keyword"},
		{"group", `\bgroup\b`, "keyword"},
		{"pattern", `\bpattern\b`, "keyword"},
		{"static", `\bstatic\b`, "keyword"},
		{"stock", `\bstock\b`, "keyword"},
		{"anyclass", `\banyclass\b`, "keyword"},
		{"via", `\bvia\b`, "keyword"},
		{"default", `\bdefault\b`, "keyword"},
		{"forall", `\bforall\b`, "keyword"},
		{"newtype", `\bnewtype\b`, "keyword"},
	}

	operators := []TokenType{
		{"DOUBLE_COLON", `::`, "operator"},
		{"ARROW", `->`, "operator"},
		{"FAT_ARROW", `=>`, "operator"},
		{"EQUALS", `==`, "operator"},
		{"NOT_EQUALS", `/=`, "operator"},
		{"BIND", `>>=`, "operator"},
		{"RIGHT_SHIFT", `>>`, "operator"},
		{"LEFT_SHIFT", `<<`, "operator"},
		{"REVERSE_BIND", `=<<`, "operator"},
		{"FUNCTOR_APPLY", `<\$>`, "operator"},
		{"APPLICATIVE_APPLY", `<*>`, "operator"},
		{"ALTERNATIVE", `<\|>`, "operator"},
		{"CONCAT", `\+\+`, "operator"},
		{"INDEX", `!!`, "operator"},
		{"COLON", `:`, "operator"},
		{"ASSIGN", `=`, "operator"},
		{"PLUS", `\+`, "operator"},
		{"MINUS", `-`, "operator"},
		{"MULTIPLY", `\*`, "operator"},
		{"DIVIDE", `/`, "operator"},
		{"LESS", `<`, "operator"},
		{"GREATER", `>`, "operator"},
		{"AND", `&&`, "operator"},
		{"OR", `\|\|`, "operator"},
		{"NOT", `!`, "operator"},
	}

	variables := []TokenType{
		{"ident", `[a-z_][a-zA-Z0-9_']*`, "variable"},
	}

	constructors := []TokenType{
		{"constructor", `[A-Z][a-zA-Z0-9_']*`, "constructor"},
	}

	constants := []TokenType{
		{"integer", `\d+`, "constant"},
		{"float", `\d+\.\d+`, "constant"},
		{"string", `"[^"]*"`, "constant"},
		{"char", `'[^']'`, "constant"},
		{"multiline_string", `"""[\s\S]*?"""`, "constant"},
	}

	punctuations := []TokenType{
		{"DOT", `\.`, "punctuation"},
		{"COMMA", `,`, "punctuation"},
		{"SEMICOLON", `;`, "punctuation"},
		{"LPAREN", `\(`, "punctuation"},
		{"RPAREN", `\)`, "punctuation"},
		{"LBRACE", `\{`, "punctuation"},
		{"RBRACE", `\}`, "punctuation"},
		{"LBRACKET", `\[`, "punctuation"},
		{"RBRACKET", `\]`, "punctuation"},
		{"BACKTICK", "`", "punctuation"},
		{"BACKSLASH", `\\`, "punctuation"},
		{"AT", `@`, "punctuation"},
		{"HASH", `#`, "punctuation"},
		{"DOLLAR", `\$`, "punctuation"},
		{"PERCENT", `%`, "punctuation"},
		{"CARET", `\^`, "punctuation"},
		{"AMPERSAND", `&`, "punctuation"},
		{"PIPE", `\|`, "punctuation"},
		{"QUESTION", `\?`, "punctuation"},
		{"TILDE", `~`, "punctuation"},
	}

	skip := []TokenType{
		{"SPACE", `\s`, "skip"},
		{"COMMENT", `--.*|\{-[\s\S]*?-\}`, "skip"},
	}

	preprocessor := []TokenType{
		{"PRAGMA", `\{-#[\s\S]*?#-\}`, "preprocessor"},
	}

	tokenTypesList := append(keywords, operators...)
	tokenTypesList = append(tokenTypesList, variables...)
	tokenTypesList = append(tokenTypesList, constructors...)
	tokenTypesList = append(tokenTypesList, constants...)
	tokenTypesList = append(tokenTypesList, punctuations...)
	tokenTypesList = append(tokenTypesList, skip...)
	tokenTypesList = append(tokenTypesList, preprocessor...)

	return tokenTypesList
}
