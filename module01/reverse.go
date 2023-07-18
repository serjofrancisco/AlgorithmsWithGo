package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// iterative solution
	reversed := "";
	for _, c := range word {
		reversed = string(c) + reversed;
	}
	return reversed;
	
	// recursive solution, doesn't work for non-ascii characters
	// if(len(word) == 0) {
	// 	return "";
	// }
	// return string(word[len(word) - 1]) + Reverse(word[:len(word) - 1]);
}


// for in range is diferent from for using index. 