package ast

import "interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	value Expression
}

func (_ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

type Identifier struct {
	Token token.Token // the token.IDENT token 
	Value string 
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}

func (p *Program) TokenLiteral() string  {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}