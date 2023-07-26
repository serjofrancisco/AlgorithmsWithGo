package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
    charset := "0123456789ABCDEF"
    result := 0
	
	for _, v := range value {
		result = result * base + charsetIndex(charset, v)
	}

	return result;
};


func charsetIndex(charset string, r rune) int {
	for i, v := range charset {
		if v == r {
			return i
		}
	}
	return -1;
};