package glambda

import "fmt"

type node interface {
	node()
}

type term interface {
	node
	term()
}

// Variable
type variableNode struct {
	name string
}

func (n variableNode) node() {}
func (n variableNode) term() {}

func (n variableNode) String() string {
	return n.name
}

// Abstraction
type abstractionNode struct {
	variable variableNode
	body     term
}

func (n abstractionNode) node() {}
func (n abstractionNode) term() {}

func (n abstractionNode) String() string {
	return fmt.Sprintf("(λ %s . %s)", n.variable, n.body)
}

// Application
type applicationNode struct {
	left  term
	right term
}

func (n applicationNode) node() {}
func (n applicationNode) term() {}

func (n applicationNode) String() string {
	return fmt.Sprintf("(%s %s)", n.left, n.right)
}

// Definition
type definitionNode struct {
	name        string
	abstraction abstractionNode
}

func (n definitionNode) node() {}

func (n definitionNode) String() string {
	return fmt.Sprintf("%s = %s", n.name, n.abstraction)
}
