package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.

// if else
// func FizzBuzz(n int) {
// 	var result string;
// 	for i := 1; i <= n; i++ {
// 		if(i % 3 == 0 && i % 5 == 0) {
// 			result += "Fizz Buzz, ";
// 		} else if(i % 3 == 0) {
// 			result += "Fizz, ";
// 		} else if(i % 5 == 0) {
// 			result += "Buzz, ";
// 		} else {
// 			result += fmt.Sprintf("%d, ", i);
// 		}
// 	}
// 	fmt.Println(result[:len(result) - 2]);
// }

// switch case
func FizzBuzz(n int) {
	var result string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result += "Fizz Buzz, "
		case i%3 == 0:
			result += "Fizz, "
		case i%5 == 0:
			result += "Buzz, "
		default:
			result += fmt.Sprintf("%d, ", i)
		}
	}
	fmt.Println(result[:len(result)-2])
}

// switch case in go can have a logic, like if else and doesn't need a break statement
