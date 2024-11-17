package lexer

import (
)

// lexer
type Scanner struct {
	line, col    int
	name         string
	input        []byte
	ch 	         byte
	position     int
	readPosition int
}

func NewScanner(name string, input []byte) *Scanner {
	s := &Scanner{
		line: 1,
		col:  0,
		name: name,
		input: input,
	}
	s.readChar()
	return s
}

func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.ch = 0
	} else {
		s.ch = s.input[s.readPosition]
		s.col++
		if s.ch == '\n' {
			s.line++
			s.col = 0
		}
	}
	s.position = s.readPosition
	s.readPosition += 1
}

func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input) {
		return 0
	} else {
		return s.input[s.readPosition]
	}
}

func (s *Scanner) NextToken() Token {
	s.skipWhitespaceAndComment()
	var tok Token
	switch s.ch {
	case ';':
		tok = newToken(TOKEN_SEP_SEMI, "")
	case '(':
		tok = newToken(TOKEN_SEP_LPAREN, "")
	case ')':
		tok = newToken(TOKEN_SEP_RPAREN, "")
	case ',':
		tok = newToken(TOKEN_SEP_COMMA, "")
	case '+':
		tok = newToken(TOKEN_OP_ADD, "")
	case '{':
		tok = newToken(TOKEN_SEP_LBRACK, "")
	case '}':
		tok = newToken(TOKEN_SEP_RBRACK, "")
	case '[':
		tok = newToken(TOKEN_SEP_LCURLY, "")
	case ']':
		tok = newToken(TOKEN_SEP_RCURLY, "")
	case ':':
		if s.peekChar() == '=' {
			s.readChar()
			tok = newToken(TOKEN_OP_DEFINE, "")
		} else {
			tok = newToken(TOKEN_SEP_COLON, "")
		}
	case '=':
		if s.peekChar() == '=' {
			s.readChar()
			tok = newToken(TOKEN_OP_EQ, "")
		} else {
			tok = newToken(TOKEN_OP_ASSIGN, "")
		}
	case '!':
		if s.peekChar() == '=' {
			s.readChar()
			tok = newToken(TOKEN_OP_NE, "")
		} else {
			tok = newToken(TOKEN_OP_NOT, "")
		}
	case '<':
		if s.peekChar() == '=' {
			s.readChar()
			tok = newToken(TOKEN_OP_LE, "")
		} else if s.peekChar() == '<' {
			s.readChar()
			tok = newToken(TOKEN_OP_SHL, "")
		} else {
			tok = newToken(TOKEN_OP_LT, "")
		}
	case '>':
		if s.peekChar() == '=' {
			s.readChar()
			tok = newToken(TOKEN_OP_GE, "")
		} else if s.peekChar() == '>' {
			s.readChar()
			tok = newToken(TOKEN_OP_SHR, "")
		} else {
			tok = newToken(TOKEN_OP_GT, "")
		}
	case '/':
		if s.peekChar() == '/' {
			s.readChar()
			tok = newToken(TOKEN_OP_IDIV, "")
		} else {
			tok = newToken(TOKEN_OP_DIV, "")
		}
	case '*':
		tok = newToken(TOKEN_OP_MUL, "")
	case '%':
		tok = newToken(TOKEN_OP_MOD, "")
	case '&':
		if s.peekChar() == '&' {
			s.readChar()
			tok = newToken(TOKEN_OP_AND, "")
		} else {
			tok = newToken(TOKEN_OP_BAND, "")
		}
	case '|':
		if s.peekChar() == '|' {
			s.readChar()
			tok = newToken(TOKEN_OP_OR, "")
		} else {
			tok = newToken(TOKEN_OP_BOR, "")
		}
	case '~':
		tok = newToken(TOKEN_OP_BNOT, "")
	case '^':
		tok = newToken(TOKEN_OP_POW, "")
	case 0:
		tok = newToken(TOKEN_EOF, "")
	default:
		if isLetter(s.ch) {
			return s.identifier()
		} else if isDigit(s.ch) {
			return s.number()
		} else {
			tok = newToken(TOKEN_ILLEGAL, string(s.ch))
		}
	}
	s.readChar()
	return tok
}

func (s *Scanner) skipWhitespaceAndComment() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' || s.ch == '#' {
		if s.ch == '#' {
			for s.ch != '\n' && s.ch != 0 {
				s.readChar()
			}
		}
		s.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) identifier() Token {
	position := s.position
	for isLetter(s.ch) || isDigit(s.ch) {
		s.readChar()
	}
	// maybe keyword
	ident := string(s.input[position:s.position])
	if tok, ok := keywords[ident]; ok {
		return newToken(tok, "")
	}
	return newToken(TOKEN_IDENTIFIER, string(ident))
}

// integer or float
func (s *Scanner) number() Token {
	position := s.position
	for isDigit(s.ch) {
		s.readChar()
	}
	
	if s.ch == '.' && isDigit(s.peekChar()) {
		s.readChar()
		for isDigit(s.ch) {
			s.readChar()
		}
		return newToken(TOKEN_NUMBER, string(s.input[position:s.position]))
	}
	return newToken(TOKEN_NUMBER, string(s.input[position:s.position]))
}
