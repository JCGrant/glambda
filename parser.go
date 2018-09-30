package glambda

import (
	"fmt"
)

type parser struct {
	lexer     *lexer
	input     string
	peekCount int
	tokens    []token
}

func newParser(input string) *parser {
	return &parser{
		lexer:  lex(input),
		input:  input,
		tokens: make([]token, 1, 1),
	}
}

func parse(input string) []node {
	p := newParser(input)
	return p.parse()
}

func (p *parser) peek() token {
	if p.peekCount > 0 {
		return p.tokens[p.peekCount-1]
	}
	p.peekCount = 1
	p.tokens[0] = p.lexer.nextItem()
	return p.tokens[0]
}

func (p *parser) expect(expectedType tokenType) bool {
	peeked := p.peek()
	return peeked.tokenType == expectedType
}

func (p *parser) next() token {
	if p.peekCount > 0 {
		p.peekCount--
	} else {
		p.tokens[0] = p.lexer.nextItem()
	}
	return p.tokens[p.peekCount]
}

func (p *parser) consume(expectedType tokenType) token {
	token := p.next()
	if token.tokenType == expectedType {
		return token
	}
	panic(fmt.Sprintf(
		"expected '%s' token, got '%s' at:\n%s",
		expectedType, token, p.input[p.lexer.position-1:],
	))
}

func (p *parser) parse() []node {
	var nodes []node
	for {
		for p.expect(tokenNewLine) {
			p.consume(tokenNewLine)
		}
		if p.expect(tokenEOF) {
			break
		}
		nodes = append(nodes, p.parseDefinition())
	}
	return nodes
}

func (p *parser) parseDefinition() definitionNode {
	name := p.consume(tokenIdentifier).value
	p.consume(tokenEquals)
	abstraction := p.parseAbstraction()
	return definitionNode{
		name,
		abstraction,
	}
}

func (p *parser) parseLambdaTerm() term {
	return p.parseApplication()
}

func (p *parser) parseAbstraction() abstractionNode {
	p.consume(tokenLambda)
	var variables []variableNode
	for p.expect(tokenIdentifier) {
		variables = append(variables, p.parseVariable())
	}
	p.consume(tokenDot)
	var abstraction term = p.parseLambdaTerm()
	for i := len(variables) - 1; i >= 0; i-- {
		variable := variables[i]
		abstraction = abstractionNode{
			variable,
			abstraction,
		}
	}
	return abstraction.(abstractionNode)
}

func (p *parser) parseApplication() term {
	left := p.parseAtom()
	for {
		if p.expect(tokenNewLine) {
			return left
		}
		right := p.parseAtom()
		if right == nil {
			return left
		}
		left = applicationNode{left, right}
	}
}

func (p *parser) parseAtom() term {
	if p.expect(tokenLeftParen) {
		p.consume(tokenLeftParen)
		term := p.parseLambdaTerm()
		p.consume(tokenRightParen)
		return term
	}
	if p.expect(tokenLambda) {
		return p.parseAbstraction()
	}
	if p.expect(tokenIdentifier) {
		return p.parseVariable()
	}
	return nil
}

func (p *parser) parseVariable() variableNode {
	name := p.consume(tokenIdentifier).value
	return variableNode{
		name,
	}
}
