package trickdate

import (
	"regexp"
)

var recleanDate, reDate, reDateFull *regexp.Regexp

/*Func to starts regexp
Date format accepted: dd/mm/yyyy without requiring separator character
* */
func init() {
	//reDate = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*(?:19|20)\d{2}?$`)
	reDate = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*(19|20)(\d{2})$`)
	reDateFull = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*\d{4}?$`)
	recleanDate = regexp.MustCompile(`\D`)
}

//Verifies if the string is a valid date - year between 1900-2099
func IsDate(date string) bool {
	return reDate.MatchString(date)
}

//Verifies if the string is a valid date
func IsDateFull(date string) bool {
	return reDateFull.MatchString(date)
}

//Removes non-numeric characters
func cleanDate(date string) string {
	return recleanDate.ReplaceAllString(date, "")
}

//Reverse the order ddmmaaaa to aaaammdd
func FormatDate(date string) string {
	var day, month, year string
	var datePieces []string

	datePieces = reDate.FindStringSubmatch(cleanDate(date))
	if datePieces != nil {
		day = datePieces[1]
		month = datePieces[2]
		year = datePieces[3] + datePieces[4]
	} else {
		return ""
	}

	return year + month + day
}
