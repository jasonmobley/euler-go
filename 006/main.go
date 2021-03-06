// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the
// first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the
// first one hundred natural numbers and the square of the sum.
package main

import "fmt"

func main() {
    var sum, sum_of_squares, square_of_sum, difference, i uint64 = 0, 0, 0, 0, 1

    for i <= 100 {
        sum += i
        sum_of_squares += (i*i)
        i++
    }

    square_of_sum = sum * sum
    difference = square_of_sum - sum_of_squares

    fmt.Printf("The sum of the squares is %d.\n", sum_of_squares)
    fmt.Printf("The square of the sum is %d.\n", square_of_sum)
    fmt.Printf("The difference is %d.\n", difference)
}

