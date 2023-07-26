package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//Recursive	solution
// func Fibonacci(n int) int {
// 	if n == 0 {
// 		return 0;
// 	}
// 	if n == 1 {
// 		return 1;
// 	}
// 	return Fibonacci(n - 1) + Fibonacci(n - 2);
// }
//Iterative solution more confusing but faster
func Fibonacci(n int) int {
	if n == 0 {
		return 0;
	}
	if n == 1 {
		return 1;
	}
	var prev1 int = 1;
	var prev2 int = 0;
	var result int = 0;
	for i := 2; i <= n; i++ {
		result = prev1 + prev2;
		prev2 = prev1;
		prev1 = result;
	}
	return result;
}
