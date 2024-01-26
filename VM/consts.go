package vm

func splitWithSeparators(input []Token, separators []string) [][]Token {
	result := make([][]Token, 0)
	currentSegment := make([]Token, 0)

	for _, item := range input {
		containsSeparator := false

		for _, separator := range separators {
			if item.Type == separator {
				containsSeparator = true
				break
			}
		}

		if containsSeparator {
			if len(currentSegment) > 0 {
				result = append(result, currentSegment)
				currentSegment = make([]Token, 0)
			}
		}

		currentSegment = append(currentSegment, item)
	}

	if len(currentSegment) > 0 {
		result = append(result, currentSegment)
	}

	return result
}
func containsString(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}
