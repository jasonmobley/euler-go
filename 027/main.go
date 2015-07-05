// Euler discovered the remarkable quadratic formula:
//
// n² + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
// However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly
// when n = 41, 41² + 41 + 41 is clearly divisible by 41.
//
// The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the
// consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.
//
// Considering quadratics of the form:
//
// n² + an + b, where |a| < 1000 and |b| < 1000
//
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |−4| = 4
//
// Find the product of the coefficients, a and b, for the quadratic expression that produces
// the maximum number of primes for consecutive values of n, starting with n = 0.
package main

import "fmt"
import "github.com/jasonmobley/euler-go/primes"

func main() {
    var primeSieve chan int = primes.Sieve()
    defer close(primeSieve)

    primeSlice := make([]int, 1000, 1000)

    for i := 0; i < 1000; i++ {
        primeSlice[i] = <-primeSieve
    }

    var a, b, n, result int
    var maxN, maxA, maxB int
    var stillPrime = true

    for a = -999; a < 1000; a++ {
        for b = -999; b < 1000; b++ {
            stillPrime = true
            for n = 0; stillPrime; n++ {
                result = (n * n) + (a * n) + b;
                if (isPrime(result, &primeSlice)) {
                    if (n > maxN) {
                        maxN = n
                        maxA = a
                        maxB = b
                    }
                } else {
                    stillPrime = false
                }
            }
            fmt.Println()
        }
    }

    fmt.Printf("Got %d consecutive primes with a = %d and b = %d\n", maxN, maxA, maxB)
}

func abs(value int) int {
    if value < 0 {
        return -1 * value;
    }
    return value;
}

func isPrime(value int, listOfPrimes *[]int) bool {
    if value == 0 {
        return true;
    }

    var mid int
    first := 0
    last := len(*listOfPrimes) - 1

    for first <= last {
        mid = (first + last) / 2

        if value < (*listOfPrimes)[mid] {
            last = mid - 1
        } else if value > (*listOfPrimes)[mid] {
            first = mid + 1
        } else {
            return true
        }
    }
    return false
}

