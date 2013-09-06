// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10,001st prime number?
package main

import "fmt"
import "math"

func main() {
   var current_value uint64 = 3
   var primes_found uint64 = 2

   for primes_found < 10001 {
      current_value += 2
      if isPrime(current_value) {
         primes_found++
      }
   }

   fmt.Printf("The 10001st prime is %d.\n", current_value)
}

func isPrime(value uint64) bool {
    var i uint64 = 2
    var rootPlusOne uint64 = uint64(math.Sqrt(float64(value))) + 1

    for i < rootPlusOne {
        if value % i == 0 {
            return false
        }
        i += 1
    }
    return true
}
