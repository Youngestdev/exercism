// Package leap indicates whether a year is a leap year or not.
package leap

// IsLeapYear should have a comment documenting it.

func IsLeapYear(year int) bool {
	// Overkill!
// Return true if 'year' passes the conditions below.
	if (year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
		return true
	}

	return false
}
