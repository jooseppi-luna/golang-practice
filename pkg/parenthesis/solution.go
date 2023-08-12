package parenthesis

import "fmt"

const openParen = "("
const closedParen = ")"

type Stack []string

func ParenthesisBalanced(parens []string) (bool, error) {
	var stack Stack

	if len(parens) == 0 {
		return true, nil
	}

	for _, paren := range parens {
		if paren == openParen {
			stack = append(stack, paren)
		} else if paren == closedParen {
			if len(stack) == 0 {
				return false, nil
			}
			stack = stack[:len(stack)-1]
		} else {
			return false, fmt.Errorf("unrecognized character in paren list: %s", paren)
		}
	}

	if len(stack) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

