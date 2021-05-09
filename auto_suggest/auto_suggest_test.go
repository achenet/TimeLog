package auto_suggest

import "testing"

type calcDiffTestCase struct {
	firstStr  string
	secondStr string
	expected  int
}

type autoSuggestTestCase struct {
	input      string
	dictionary []string
	expected   []string
}

var calcDiffMultipleTestCases = []calcDiffTestCase{
	{
		firstStr:  "yo",
		secondStr: "yoyo",
		expected:  2,
	},
	{
		firstStr:  "yo",
		secondStr: "bobo",
		expected:  3,
	},
	{
		firstStr:  "yo",
		secondStr: "",
		expected:  2,
	},
	{
		firstStr:  "yo",
		secondStr: "thug",
		expected:  4,
	},
}

var autoSuggestMultipleTestCases = []autoSuggestTestCase{
	{
		input:      "run",
		dictionary: []string{"run", "jump", "walk"},
		expected:   []string{},
	},
	{
		input:      "runm",
		dictionary: []string{"run", "jump", "walk"},
		expected:   []string{"run"},
	},
	{
		input:      "rnu",
		dictionary: []string{"run", "jump", "walk"},
		expected:   []string{"run"},
	},
}

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestCalculateDiff(t *testing.T) {
	for _, tc := range calcDiffMultipleTestCases {
		got := CalculateStringDiff(tc.firstStr, tc.secondStr)
		if got != tc.expected {
			t.Error("got:", got, "expected", tc.expected)
		}
	}
}

func TestAutoSuggest(t *testing.T) {
	for _, tc := range autoSuggestMultipleTestCases {
		got := AutoSuggest(tc.input, tc.dictionary)
		if !AreEqual(got, tc.expected) {
			t.Error("got:", got, "expected", tc.expected)
		}
	}
}

func AreEqual(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
