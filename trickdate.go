package trickdate

import (
	//"time"
	"regexp"
)

var reDate *regexp.Regexp

/*Func to starts regexp
Date format accepted: dd/mm/yyyy without requiring separator character
* */
func init() {
	reDate = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*(?:19|20)\d{2}?$`)
}

//Verifies if the string is a valid date
func IsDate(date string) bool {
	return reDate.MatchString(date)
}

func FormatDate(date string) (string, error) {
	return date, nil
}
