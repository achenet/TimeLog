package auto_suggest

import "testing"

type calcDiffTestCase struct {
	a        string
	b        string
	expected int
}

type autoSuggestTestCase struct {
	input      string
	dictionary []string
	expected   []string
}

var calcDiffTestCaseList = []calcDiffTestCase{
	{
		a:        "yo",
		b:        "yoyo",
		expected: 2,
	},
	{
		a:        "yo",
		b:        "bobo",
		expected: 3,
	},
	{
		a:        "yo",
		b:        "",
		expected: 2,
	},
	{
		a:        "yo",
		b:        "thug",
		expected: 4,
	},
}

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestCalculateDiff(t *testing.T) {
	for _, tc := range calcDiffTestCaseList {
		got := CalculateStringDiff(tc.a, tc.b)
		if got != tc.expected {
			t.Error("got:", got, "expected", tc.expected)
		}
	}
}
