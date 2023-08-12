package parenthesis

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParenthesisBalanced(t *testing.T) {
	tests := []struct{
		name string
		arr []string
		matching bool
	}{
		{
			name: "matching base case",
			arr: []string{"(", "(", ")", ")"},
			matching: true,
		},
		{
			name: "nonmatching base case",
			arr: []string{"(", "(", ")", "("},
			matching: false,
		},
		{
			name: "two matching pairs",
			arr: []string{"(", ")", "(", ")"},
			matching: true,
		},
		{
			name: "large set matching",
			arr: []string{"(", "(", ")", "(", "(", ")", "(", ")", ")", ")"},
			matching: true,
		},
		{
			name: "large set nonmatching",
			arr: []string{"(", "(", ")", ")", ")", "(", "(", ")", "(", ")", ")", ")"},
			matching: false,
		},
		{
			name: "large set nonmatching at end",
			arr: []string{"(", "(", ")", "(", "(", ")", "(", ")", ")", ")", ")"},
			matching: false,
		},
		{
			name: "empty",
			arr: []string{},
			matching: true,
		},
		{
			name: "large set matching sets",
			arr: []string{"(", ")", "(", ")", "(", ")", "(", ")", "(", ")"},
			matching: true,
		},
		{
			name: "invalid input",
			arr: []string{"(", "(", "bad", ")", ")"},
			matching: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matching, _ := ParenthesisBalanced(tt.arr)
			assert.Equal(t, tt.matching, matching)
		})
	}
}

