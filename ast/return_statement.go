package ast

import (
	"bytes"
	"monkey/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

var _ Statement = (*ReturnStatement)(nil)

func (r *ReturnStatement) statementNode() {
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	out := bytes.Buffer{}
	out.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
