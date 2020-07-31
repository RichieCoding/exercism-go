//Package leap reports if a year is a leap or not
package leap

// IsLeapYear returns true if year is a leap, false if not
func IsLeapYear(y int) bool {
	if y%4 == 0 && y%400 == 0 || y%4 == 0 && y%100 != 0 {
		return true
	}
	return false
}
