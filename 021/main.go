// Let d(n) be defined as the sum of proper divisors of n
// (numbers less than n which divide evenly into n).
// 
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are
// an amicable pair and each of a and b are called amicable numbers.
// 
// For example, the proper divisors of 220 are
// 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110;
// therefore d(220) = 284.
// 
// The proper divisors of 284 are 1, 2, 4, 71 and 142;
// so d(284) = 220.
// 
// Evaluate the sum of all the amicable numbers under 10000.
package main

import "fmt"

func main() {
    amicableNumbers := make(map[int]bool)
    sumOfAmicableNumbers := 0

    var possibleAmicable int

    for i := 1; i < 10000; i++ {
        possibleAmicable = sumOfSlice(properDivisors(i))

        if i != possibleAmicable && possibleAmicable < 10000 && sumOfSlice(properDivisors(possibleAmicable)) == i {
            amicableNumbers[i] = true
            amicableNumbers[possibleAmicable] = true
        }
    }

    for num := range amicableNumbers {
        sumOfAmicableNumbers += num
    }

    fmt.Println(sumOfAmicableNumbers)
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

