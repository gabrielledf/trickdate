package trickdate

import "testing"

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestIsDate(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    bool
	}{
		// Invalid format.
		{"11111111111", false},
		{"", false},
		{"#$%¨&*(ABCDEF", false},
		{"2017-12-31", false},

		// Common invalid patterns.
		{"00000000", false},
		{"11111111", false},

		// Requires two-digit daytime representation
		{"1/10/1900", false},

		// Requires two-digit month representation
		{"10/1/1900", false},

		// Requires four-digit year representation
		{"10/01/17", false},

		// The day can not be greater than 31
		{"32/12/1900", false},

		// The month can not be greater than 12
		{"31/13/1900", false},

		// Accepts years from the twentieth to the twenty-first century
		{"31/12/1889", false},
		{"31/12/3123", false},

		// Valid.
		{"31/12/2017", true},
		{"31 12 2017", true},
		{"31-12-2017", true},
		{"31122017", true},
		{"31      12     2017", true},
		{"31/////12/////2017", true},
	} {
		t.Logf("#%d CPF validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := IsDate(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
