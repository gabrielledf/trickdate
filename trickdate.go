package trickdate

import (
	"regexp"
	"time"
)

var recleanDate, reDate, reDateFull, reHasDate, reHasDateFull, reTZ *regexp.Regexp

/*Func to starts regexp
Date format accepted: dd/mm/yyyy without requiring separator character
* */
func init() {
	reDate = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*(19|20)(\d{2})$`)
	reDateFull = regexp.MustCompile(`^(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*\d{4}?$`)
	recleanDate = regexp.MustCompile(`\D`)

	reHasDate = regexp.MustCompile(`(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*(19|20)(\d{2})`)
	reHasDateFull = regexp.MustCompile(`(0[1-9]|[12][0-9]|3[01])[-|\\|/|\s]*(0[1-9]|1[012])[-|\\|/|\s]*\d{4}?`)

	reTZ = regexp.MustCompile(`^(\d{4}\-\d{2}\-\d{2})(T(\d{2})(\:\d{2}\:\d{2}\.\d{3}Z))$`)
}

//Verifies if the string is a valid date - year between 1900-2099
func IsDate(date string) bool {
	return reDate.MatchString(date)
}

//Verifies if the string is a valid date
func IsDateFull(date string) bool {
	return reDateFull.MatchString(date)
}

//Verifies if the string has a valid date in format ddmmaaaa - year between 1900-2099
func HasDate(date string) bool {
	return reHasDate.MatchString(date)
}

//Verifies if the string has a valid date in format ddmmaaaa
func HasDateFull(date string) bool {
	return reHasDateFull.MatchString(date)
}

//Removes non-numeric characters
func cleanDate(date string) string {
	return recleanDate.ReplaceAllString(date, "")
}

/*
 * Checks if string data is time zone T03
 * date RFC3339 time
 */
func GetTZ(date string) string {
	var datePieces []string

	datePieces = reTZ.FindStringSubmatch(date)
	if datePieces != nil {
		return datePieces[2]
	} else {
		return ""
	}
}

/*
 * Change TZ to T03
 * date RFC3339 time
 * */
func ChangeTZ(date string) string {
	var datePieces []string

	datePieces = reTZ.FindStringSubmatch(date)
	if datePieces != nil {
		if datePieces[3] == "03" {
			return date
		} else {
			return datePieces[1] + "T03" + datePieces[4]
		}
	} else {
		return ""
	}
}

//Reverse the order ddmmaaaa to aaaa-mm-dd
func FormatDate(date string) string {
	var day, month, year string
	var datePieces []string

	datePieces = reDate.FindStringSubmatch(cleanDate(date))
	if datePieces != nil {
		day = datePieces[1]
		month = datePieces[2] + "-"
		year = datePieces[3] + datePieces[4] + "-"
	} else {
		return ""
	}

	return year + month + day
}

//Get day from date - ddmmaaaa
func GetDay(date string) string {
	var datePieces []string

	datePieces = reDate.FindStringSubmatch(cleanDate(date))
	if datePieces != nil {
		return datePieces[1]
	} else {
		return ""
	}
}

//Get month from date - ddmmaaaa
func GetMonth(date string) string {
	var datePieces []string
	datePieces = reDate.FindStringSubmatch(cleanDate(date))
	if datePieces != nil {
		return datePieces[2]
	} else {
		return ""
	}
}

//Get year from date - ddmmaaaa
func GetYear(date string) string {
	var datePieces []string

	datePieces = reDate.FindStringSubmatch(cleanDate(date))
	if datePieces != nil {
		return datePieces[3] + datePieces[4]
	} else {
		return ""
	}
}

//Converts string data to RFC3339 time
func ConvertDate(date string) time.Time {
	var dt time.Time
	var err error

	dt, err = time.Parse(time.RFC3339, date+"T03:00:00.000Z")
	if err != nil {
		return dt
	}
	return dt
}

//Converts time.Now() to RFC3339 time
func ConvertDateNow() time.Time {
	var dt time.Time
	var err error
	var date string

	dt = time.Now()
	date = dt.String()
	//Get aaaa-mm-dd
	date = date[0:10] + "T03:00:00.000Z"

	dt, err = time.Parse(time.RFC3339, date)
	if err != nil {
		return dt
	}
	return dt
}
