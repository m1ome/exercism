package leap

const testVersion = 2

/*
Function to implement leap year checking.

Year is leap if at least one of two conditions is true:
- Year that is evenly divisible by 4 except every year that is evenly divisible by 100
- Year is evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	return ((year%4 == 0 && year%100 != 0) || year%400 == 0)
}
