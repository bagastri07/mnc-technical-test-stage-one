package task

func isValid(s string) bool {
	stack := []rune{}

	// Mapping of opening to closing brackets
	bracketMap := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if char == '<' || char == '{' || char == '[' {
			// Push opening brackets onto the stack
			stack = append(stack, char)
		} else if char == '>' || char == '}' || char == ']' {
			// If closing bracket appears without a corresponding opening bracket
			if len(stack) == 0 {
				return false
			}
			// Check if the closing bracket matches the last opened bracket
			last := stack[len(stack)-1]
			expected, exists := bracketMap[last]
			if !exists || expected != char {
				return false
			}
			// Pop the last opened bracket
			stack = stack[:len(stack)-1]
		}
	}

	// Valid if all opened brackets are closed
	return len(stack) == 0
}
