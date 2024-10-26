package task

import "strings"

func findFirstMatchingSet(N int, stringsList []string) interface{} {
	var firstKeyFound string
	// make it lowercase, handle in case-sensitive
	for i := 0; i < N; i++ {
		stringsList[i] = strings.ToLower(stringsList[i])
	}

	// map string to cout pair
	stringIndices := make(map[string][]int)

	for i := 0; i < N; i++ {
		str := stringsList[i]
		stringIndices[str] = append(stringIndices[str], i+1)

		if len(stringIndices[str]) > 1 && firstKeyFound == "" {
			firstKeyFound = str
		}
	}

	if value, ok := stringIndices[firstKeyFound]; ok {
		return value
	}

	// not pair match
	return false
}
