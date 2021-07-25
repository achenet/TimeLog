// A package to handle suggest strings similar to the string entered.
package auto_suggest

const MaximumDiff = 5

func AutoSuggest(input string, dictionary []string) []string {
	possibleStrings := make([]string, 0)
	for _, word := range dictionary {
		if CalculateStringDiff(input, word) < MaximumDiff {
			possibleStrings = append(possibleStrings, word)
		}
	}
	return possibleStrings
}

// Calculate difference between the two strings.
// Difference is number of position differences plus the number of runes in first string not in second string,
// plus number of runes in second string not in first string.
// For example, "run" and "rnu" have a difference of 2, while "run" and "cat" have a difference of 6.
func CalculateStringDiff(firstStr, secondStr string) int {
	numberOfDifferentPositions := calculateNumberOfDifferentPositions(firstStr, secondStr)
	numberOfDifferentRunes := calculateNumberOfDifferentRunes(firstStr, secondStr)
	return numberOfDifferentPositions + numberOfDifferentRunes
}

func calculateNumberOfDifferentPositions(firstStr, secondStr string) int {
	diff := 0
	for i := 0; i < min(len(firstStr), len(secondStr)); i++ {
		if firstStr[i] != secondStr[i] {
			diff++
		}
	}
	diff += abs(len(firstStr) - len(secondStr))
	return diff
}

func calculateNumberOfDifferentRunes(firstStr, secondStr string) int {
	// Analyse the character frequency of both strings
	m1, m2 := analyseString(firstStr), analyseString(secondStr)
	diff := 0
	for r := range m1 {
		if _, exists := m2[r]; !exists {
			m2[r] = 0
		}
		diff += abs(m1[r] - m2[r])
	}

	// Add letters that are only in the second string
	for r := range m2 {
		if _, exists := m1[r]; !exists {
			diff += m2[r]
		}
	}
	return diff
}

func analyseString(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range []rune(s) {
		if _, exists := m[r]; !exists {
			m[r] = 0
		}
		m[r]++
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return (-1 * a)
	}
	return a
}
