package ast

import "monkey/token"

type Identifier struct {
	Token token.Token
	Value string
}

var _ Expression = (*Identifier)(nil)

func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
