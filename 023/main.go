// A perfect number is a number for which the sum of its proper divisors is
// exactly equal to the number. For example, the sum of the proper divisors
// of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect
// number.
// 
// A number n is called deficient if the sum of its proper divisors is less
// than n and it is called abundant if this sum exceeds n.
// 
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
// the smallest number that can be written as the sum of two abundant
// numbers is 24. By mathematical analysis, it can be shown that all
// integers greater than 28123 can be written as the sum of two abundant
// numbers. However, this upper limit cannot be reduced any further by
// analysis even though it is known that the greatest number that cannot be
// expressed as the sum of two abundant numbers is less than this limit.
// 
// Find the sum of all the positive integers which cannot be written as the
// sum of two abundant numbers.
package main

import "fmt"

func main() {
    limit := 28123
    abundantNumbers := make([]int, 0, 100)

    for num := 1; num <= limit; num++ {
        if isAbundant(num) {
            abundantNumbers = append(abundantNumbers, num)
        }
    }

    abundantSums := map[int]bool{}

    for _, val1 := range abundantNumbers {
        for _, val2 := range abundantNumbers {
            if (val1 + val2) <= limit {
                abundantSums[val1 + val2] = true
            }
        }
    }

    result := 0
    for i := 0; i <= limit; i++ {
        if !abundantSums[i] {
            result += i
            //fmt.Println(i)
        }
    }

    fmt.Println(result)
}

func isAbundant(n int) bool {
    return sumOfSlice(properDivisors(n)) > n
}

func properDivisors(n int) []int {
    divisors := make([]int, 0, 10)

    for i := 1; i <= (n/2); i++ {
        if n % i == 0 {
            divisors = append(divisors, i)
        }
    }

    return divisors
}

func sumOfSlice(numbers []int) int {
    sum := 0
    for _, value := range numbers {
        sum += value
    }

    return sum
}
