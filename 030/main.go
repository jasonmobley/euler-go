// Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
//
// 1634 = 1^4 + 6^4 + 3^4 + 4^4
// 8208 = 8^4 + 2^4 + 0^4 + 8^4
// 9474 = 9^4 + 4^4 + 7^4 + 4^4
// As 1 = 1^4 is not a sum it is not included.
//
// The sum of these numbers is 1634 + 8208 + 9474 = 19316.
//
// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
package main

import "fmt"
import "math"
import "strconv"
import "strings"

func main() {
    digitFifthPowers := make([]float64, 0, 10)
    var sum, parsedDigit float64

    for i := 2; i < 1000000; i++ {
        sum = 0
        for _, digit := range strings.Split(strconv.Itoa(i), "") {
            parsedDigit, _ = strconv.ParseFloat(digit, 64)
            sum += math.Pow(parsedDigit, 5)
        }
        if (float64(i) == sum) {
            digitFifthPowers = append(digitFifthPowers, sum)
        }
    }

    var sumOfDigitFifthPowers float64 = 0
    for _, digitFifthPower := range digitFifthPowers {
        sumOfDigitFifthPowers += digitFifthPower
    }

    fmt.Println(sumOfDigitFifthPowers)
}
