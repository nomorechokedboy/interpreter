package ast

import "monkey/token"

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

var _ Expression = (*IntegerLiteral)(nil)

func (i *IntegerLiteral) expressionNode() {
}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.TokenLiteral()
}
