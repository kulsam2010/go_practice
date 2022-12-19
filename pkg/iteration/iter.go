package iteration

/*
Given a string and count, return the string reepeated count times
*/
func Repeat(str string, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += str
	}
	return result
}
