package trickdate

import (
	"testing"
	"time"
)

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
		t.Logf("#%d Date validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := IsDate(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestIsDateFull(t *testing.T) {
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

		// Valid.
		{"31/12/2017", true},
		{"31 12 2017", true},
		{"31-12-2017", true},
		{"31122017", true},
		{"31      12     2017", true},
		{"31/////12/////2017", true},

		// Accepts years
		{"31/12/1889", true},
		{"31/12/3123", true},
	} {
		t.Logf("#%d Data validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := IsDateFull(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestHasDate(t *testing.T) {
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

		//String contains no date
		{"the big being occurred in", false},

		// Valid.
		{"Christmas is in 25/12/2017", true},
		{"Christmas is in 25 12 2017", true},
		{"Happy new year in 31-12-2017", true},
		{"Happy new year in 31122017", true},
		{"After 31      12     2017 -> New year", true},
		{"31/////12/////2017 -> happy", true},
	} {
		t.Logf("#%d Has date validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := HasDate(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestFormatDate(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		{"31/12/2017", "2017-12-31"},
		{"31 12 2017", "2017-12-31"},
		{"31-12-2017", "2017-12-31"},
		{"31122017", "2017-12-31"},
		{"31      12     2017", "2017-12-31"},
		{"31/////12/////2017", "2017-12-31"},
		{"19-05-1981", "1981-05-19"},
	} {
		t.Logf("#%d Conversion date of %s should return %v: ", i, v.valueTested, v.expected)
		got := FormatDate(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestConvertDate(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    time.Time
	}{
		{"2017-12-31", time.Date(2017, 12, 31, 3, 0, 0, 0, time.UTC)},
		{"1900-12-31", time.Date(1900, 12, 31, 3, 0, 0, 0, time.UTC)},
		{"1968-09-25", time.Date(1968, 9, 25, 3, 0, 0, 0, time.UTC)},
		{"1981-09-08", time.Date(1981, 9, 8, 3, 0, 0, 0, time.UTC)},
		{"1991-09-02", time.Date(1991, 9, 2, 3, 0, 0, 0, time.UTC)},
		{"1982-11-16", time.Date(1982, 11, 16, 3, 0, 0, 0, time.UTC)},
		{"1981-05-19", time.Date(1981, 5, 19, 3, 0, 0, 0, time.UTC)},
	} {
		t.Logf("#%d Conversion date of %s should return %v: ", i, v.valueTested, v.expected)
		got := ConvertDate(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestGetDay(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		// Invalid format.
		{"11111111111", ""},
		{"", ""},
		{"#$%¨&*(ABCDEF", ""},
		{"2017-12-31", ""},

		// Common invalid patterns.
		{"00000000", ""},
		{"11111111", ""},

		// Requires two-digit daytime representation
		{"1/10/1900", ""},

		// Requires two-digit month representation
		{"10/1/1900", ""},

		// Requires four-digit year representation
		{"10/01/17", ""},

		// The day can not be greater than 31
		{"32/12/1900", ""},

		// The month can not be greater than 12
		{"31/13/1900", ""},

		// Accepts years from the twentieth to the twenty-first century
		{"31/12/1889", ""},
		{"31/12/3123", ""},

		// Valid.
		{"31/12/2017", "31"},
		{"31 12 2017", "31"},
		{"31-12-2017", "31"},
		{"31122017", "31"},
		{"31      12     2017", "31"},
		{"31/////12/////2017", "31"},
	} {
		t.Logf("#%d GetDay validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := GetDay(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestGetMonth(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		// Invalid format.
		{"11111111111", ""},
		{"", ""},
		{"#$%¨&*(ABCDEF", ""},
		{"2017-12-31", ""},

		// Common invalid patterns.
		{"00000000", ""},
		{"11111111", ""},

		// Requires two-digit daytime representation
		{"1/10/1900", ""},

		// Requires two-digit month representation
		{"10/1/1900", ""},

		// Requires four-digit year representation
		{"10/01/17", ""},

		// The day can not be greater than 31
		{"32/12/1900", ""},

		// The month can not be greater than 12
		{"31/13/1900", ""},

		// Accepts years from the twentieth to the twenty-first century
		{"31/12/1889", ""},
		{"31/12/3123", ""},

		// Valid.
		{"31/12/2017", "12"},
		{"31 12 2017", "12"},
		{"31-12-2017", "12"},
		{"31122017", "12"},
		{"31      12     2017", "12"},
		{"31/////12/////2017", "12"},
	} {
		t.Logf("#%d GetMonth validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := GetMonth(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestGetYear(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		// Invalid format.
		{"11111111111", ""},
		{"", ""},
		{"#$%¨&*(ABCDEF", ""},
		{"2017-12-31", ""},

		// Common invalid patterns.
		{"00000000", ""},
		{"11111111", ""},

		// Requires two-digit daytime representation
		{"1/10/1900", ""},

		// Requires two-digit month representation
		{"10/1/1900", ""},

		// Requires four-digit year representation
		{"10/01/17", ""},

		// The day can not be greater than 31
		{"32/12/1900", ""},

		// The month can not be greater than 12
		{"31/13/1900", ""},

		// Accepts years from the twentieth to the twenty-first century
		{"31/12/1889", ""},
		{"31/12/3123", ""},

		// Valid.
		{"31/12/2017", "2017"},
		{"31 12 2017", "2017"},
		{"31-12-2017", "2017"},
		{"31122017", "2017"},
		{"31      12     2017", "2017"},
		{"31/////12/////2017", "2017"},
	} {
		t.Logf("#%d GetYear validation of %s should return %v: ", i, v.valueTested, v.expected)
		got := GetYear(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestGetTZ(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		{"1965-12-17T01:00:00.000Z", "T01:00:00.000Z"},
		{"1965-12-17T02:00:00.000Z", "T02:00:00.000Z"},
		{"1964-03-01T03:00:00.000Z", "T03:00:00.000Z"},
		{"1964-03-01T04:00:00.000Z", "T04:00:00.000Z"},
		{"1964-03-01T05:00:00.000Z", "T05:00:00.000Z"},
		{"1964-03-01T06:00:00.000Z", "T06:00:00.000Z"},
		{"1964-03-01T07:00:00.000Z", "T07:00:00.000Z"},
		{"1964-03-01T08:00:00.000Z", "T08:00:00.000Z"},
		{"1964-03-01T09:00:00.000Z", "T09:00:00.000Z"},
		{"1964-03-01T10:00:00.000Z", "T10:00:00.000Z"},
		{"1964-03-01T11:00:00.000Z", "T11:00:00.000Z"},
		{"1964-03-01T12:00:00.000Z", "T12:00:00.000Z"},
	} {
		t.Logf("#%d Get time zone date of %s should return %v: ", i, v.valueTested, v.expected)
		got := GetTZ(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}

func TestChangeTZ(t *testing.T) {
	for i, v := range []struct {
		valueTested string
		expected    string
	}{
		{"1965-12-17T01:00:00.000Z", "1965-12-17T03:00:00.000Z"},
		{"1965-12-17T02:00:00.000Z", "1965-12-17T03:00:00.000Z"},
		{"1964-03-01T03:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T04:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T05:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T06:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T07:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T08:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T09:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T10:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T11:00:00.000Z", "1964-03-01T03:00:00.000Z"},
		{"1964-03-01T12:00:00.000Z", "1964-03-01T03:00:00.000Z"},
	} {
		t.Logf("#%d Change time zone date of %s should return %v: ", i, v.valueTested, v.expected)
		got := ChangeTZ(v.valueTested)
		if got != v.expected {
			t.Fatal(ballotX)
		}
		t.Log(checkMark)
	}
}
