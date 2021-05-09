// A package to handle suggest strings similar to the string entered.
package auto_suggest

const MaximumDiff = 3

func AutoSuggest(input string, dictionary []string) []string {
	possibleStrings := make([]string, 0)
	for _, word := range dictionary {
		if CalculateStringDiff(input, word) < MaximumDiff && word != input {
			possibleStrings = append(possibleStrings, word)
		}
	}
	return possibleStrings
}

// Returns the number of bytes in "firstStr" that are not in "secondStr".
func CalculateStringDiff(firstStr, secondStr string) (diff int) {
	// Find the longer of the two strings
	var longerStr, shorterStr string
	if len(firstStr) > len(secondStr) {
		longerStr, shorterStr = firstStr, secondStr
	} else {
		longerStr, shorterStr = secondStr, firstStr
	}

	// Add the difference
	diff += len(longerStr) - len(shorterStr)

	// Examine the characters of short string to see where they differ
	// from those of longer string.
	for i := range shorterStr {
		if shorterStr[i] != longerStr[i] {
			diff++
		}
	}

	return diff
}
