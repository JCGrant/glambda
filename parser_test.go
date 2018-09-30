package glambda

import (
	"testing"
)

func TestParseLambdaTerm(t *testing.T) {
	tests := []struct {
		input    string
		expected node
	}{
		// Variables
		{
			`x`,
			variableNode{"x"},
		},
		{
			`(x)`,
			variableNode{"x"},
		},

		// Abstractions
		{
			`\x . x`,
			abstractionNode{
				variableNode{"x"},
				variableNode{"x"},
			},
		},
		{
			`(\x . x)`,
			abstractionNode{
				variableNode{"x"},
				variableNode{"x"},
			},
		},
		{
			`\x . (\y . x)`,
			abstractionNode{
				variableNode{"x"},
				abstractionNode{
					variableNode{"y"},
					variableNode{"x"},
				},
			},
		},
		{
			`\x . N M`,
			abstractionNode{
				variableNode{"x"},
				applicationNode{
					variableNode{"N"},
					variableNode{"M"},
				},
			},
		},
		{
			`\x . (\y . x) y`,
			abstractionNode{
				variableNode{"x"},
				applicationNode{
					abstractionNode{
						variableNode{"y"},
						variableNode{"x"},
					},
					variableNode{"y"},
				},
			},
		},
		{
			`\f x . x`,
			abstractionNode{
				variableNode{"f"},
				abstractionNode{
					variableNode{"x"},
					variableNode{"x"},
				},
			},
		},
		{
			`\f x . f (f (f x))`,
			abstractionNode{
				variableNode{"f"},
				abstractionNode{
					variableNode{"x"},
					applicationNode{
						variableNode{"f"},
						applicationNode{
							variableNode{"f"},
							applicationNode{
								variableNode{"f"},
								variableNode{"x"},
							},
						},
					},
				},
			},
		},
		{
			`\a b c . a b c`,
			abstractionNode{
				variableNode{"a"},
				abstractionNode{
					variableNode{"b"},
					abstractionNode{
						variableNode{"c"},
						applicationNode{
							applicationNode{
								variableNode{"a"},
								variableNode{"b"},
							},
							variableNode{"c"},
						},
					},
				},
			},
		},

		// Applications
		{
			`N M`,
			applicationNode{
				variableNode{"N"},
				variableNode{"M"},
			},
		},
		{
			`(N M)`,
			applicationNode{
				variableNode{"N"},
				variableNode{"M"},
			},
		},
		{
			`N M O`,
			applicationNode{
				applicationNode{
					variableNode{"N"},
					variableNode{"M"},
				},
				variableNode{"O"},
			},
		},
		{
			`N (M O)`,
			applicationNode{
				variableNode{"N"},
				applicationNode{
					variableNode{"M"},
					variableNode{"O"},
				},
			},
		},
		{
			`(\x . N) M`,
			applicationNode{
				abstractionNode{
					variableNode{"x"},
					variableNode{"N"},
				},
				variableNode{"M"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			p := newParser(test.input)
			actual := p.parseLambdaTerm()
			assertEqual(t, test.expected, actual)
		})
	}
}
