package ast

import (
	"bytes"
	"monkey/token"
)

type IfExpression struct {
	Alternative *BlockStatement
	Condition   Expression
	Consequence *BlockStatement
	Token       token.Token
}

func (ie *IfExpression) expressionNode() {
}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
