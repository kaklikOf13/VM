package vm

import (
	"strconv"
	"strings"
)

const (
	TT_VALUE      = "Value"
	TT_PLUS       = "Plus"
	TT_MINUS      = "Minus"
	TT_MUL        = "Mul"
	TT_DIV        = "Div"
	TT_LPAREN     = "Lparen"
	TT_RPAREN     = "Rparen"
	TT_NEWLINE    = "Newline"
	TT_BREAKPOINT = "Breakpoint"
)

type Token struct {
	Type  string
	Value Value
}

func (t Token) String() string {
	if t.Value == nil {
		return "<" + t.Type + ">"
	}
	return "<" + t.Type + ":" + t.Value.String() + ">"
}

type Lexer struct {
	Pos    uint64
	Ch     rune
	Input  string
	Tokens []Token
}

func (l *Lexer) NextToken() {
	if l.Pos >= uint64(len(l.Input)) {
		l.Ch = 0
	} else {
		l.Ch = rune(l.Input[l.Pos])
		l.Pos += 1
	}
}
func GerateTokens(input string) []Token {
	l := &Lexer{0, 0, input, []Token{}}
	l.NextToken()
	for l.Ch != 0 {
		switch l.Ch {
		case '+':
			l.Tokens = append(l.Tokens, Token{TT_PLUS, nil})
		case '-':
			l.Tokens = append(l.Tokens, Token{TT_MINUS, nil})
		case '*':
			l.Tokens = append(l.Tokens, Token{TT_MUL, nil})
		case '/':
			l.Tokens = append(l.Tokens, Token{TT_DIV, nil})
		case '\n':
			l.Tokens = append(l.Tokens, Token{TT_NEWLINE, nil})
		case ';':
			l.Tokens = append(l.Tokens, Token{TT_BREAKPOINT, nil})
		}
		if strings.ContainsRune("1234567890.", l.Ch) {
			var value string
			for strings.ContainsRune("1234567890.", l.Ch) && l.Ch != 0 {
				value += string(l.Ch)
				l.NextToken()
			}
			v, _ := strconv.Atoi(value)
			l.Tokens = append(l.Tokens, Token{TT_VALUE, &Int{value: int64(v)}})
			continue
		}
		l.NextToken()
	}
	return l.Tokens
}
