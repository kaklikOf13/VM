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
				continue
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

func combineMaps(map1, map2 map[string]string) map[string]string {
	map12 := make(map[string]string)
	for key, value := range map1 {
		map12[key] = value
	}
	for key, value := range map2 {
		map12[key] = value
	}

	return map12
}

const VARS_NAME = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

var mainTypes = map[string]string{
	"int64": "long int",
	"int32": "int",
	"int16":"short int",
	"int8":"signed char",
	"int":   "long int",

	"uint64": "unsigned long int",
	"uint32": "unsigned int",
	"uint16":"unsigned short int",
	"uint8":"unsigned char",
	"uint":   "unsigned long int",
}
