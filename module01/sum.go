package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// iterative solution
	// var sum int;
	// for _, n := range numbers {
	// 	sum += n;
	// };
	// return sum;

	// recursive solution
	if(len(numbers) == 0) {
		return 0;
	}
	return numbers[0] + Sum(numbers[1:]);
}
