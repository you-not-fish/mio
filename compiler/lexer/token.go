package lexer

const (
	TOKEN_ILLEGAL        = iota           // end-of-file
	
	// separators
	TOKEN_SEP_SEMI                     // ;
	TOKEN_SEP_COMMA                    // ,
	TOKEN_SEP_DOT                      // .
	TOKEN_SEP_COLON                    // :
	TOKEN_SEP_LPAREN                   // (
	TOKEN_SEP_RPAREN                   // )
	TOKEN_SEP_LBRACK                   // [
	TOKEN_SEP_RBRACK                   // ]
	TOKEN_SEP_LCURLY                   // {
	TOKEN_SEP_RCURLY                   // }
	
	// operators and operations
	TOKEN_OP_ASSIGN                    // =
	TOKEN_OP_MINUS                     // - (sub or unm)
	TOKEN_OP_ADD                       // +
	TOKEN_OP_MUL                       // *
	TOKEN_OP_DIV                       // /
	TOKEN_OP_IDIV                      // //
	TOKEN_OP_POW                       // ^
	TOKEN_OP_MOD                       // %
	TOKEN_OP_BAND                      // &
	TOKEN_OP_BOR                       // |
	TOKEN_OP_BNOT                      // ~
	TOKEN_OP_SHR                       // >>
	TOKEN_OP_SHL                       // <<
	TOKEN_OP_LT                        // <
	TOKEN_OP_LE                        // <=
	TOKEN_OP_GT                        // >
	TOKEN_OP_GE                        // >=
	TOKEN_OP_EQ                        // ==
	TOKEN_OP_NE                        // !=
	TOKEN_OP_AND                       // &&
	TOKEN_OP_OR                        // ||
	TOKEN_OP_NOT                       // !
	TOKEN_OP_DEFINE                    // :=

	// keywords
	TOKEN_KW_BREAK                     // break
	TOKEN_KW_ELSE                      // else
	TOKEN_KW_FALSE                     // false
	TOKEN_KW_FOR                       // for
	TOKEN_KW_FUNCTION                  // func
	TOKEN_KW_IF                        // if
	TOKEN_KW_RANGE                     // range
	TOKEN_KW_VAR                       // var
	TOKEN_KW_NIL                       // nil
	TOKEN_KW_RETURN                    // return
	TOKEN_KW_TRUE                      // true
	TOKEN_KW_WHILE                     // while

	// others
	TOKEN_IDENTIFIER                   // identifier
	TOKEN_NUMBER                       // number literal
	TOKEN_STRING                       // string literal
	TOKEN_COMMENT                      // # comment
	TOKEN_EOF						  // end-of-file

	TOKEN_OP_UNM      = TOKEN_OP_MINUS // unary minus
	TOKEN_OP_SUB      = TOKEN_OP_MINUS
)

var keywords = map[string]int{
	"break":    TOKEN_KW_BREAK,
	"else":     TOKEN_KW_ELSE,
	"false":    TOKEN_KW_FALSE,
	"for":      TOKEN_KW_FOR,
	"func":     TOKEN_KW_FUNCTION,
	"if":       TOKEN_KW_IF,
	"range":    TOKEN_KW_RANGE,
	"var":      TOKEN_KW_VAR,
	"nil":      TOKEN_KW_NIL,
	"return":   TOKEN_KW_RETURN,
	"true":     TOKEN_KW_TRUE,
	"while":	TOKEN_KW_WHILE,
}

type Token struct {
	typ     int
	literal string
}

func newToken(typ int, lit string) Token {
	return Token{typ, lit}
}
