// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import "fmt"
import "math"

func main() {
    primesChannel := make(chan int)
    sum := 0

    go eratosthenes(2000000, primesChannel)
    for p := range primesChannel {
        sum += p
    }

    fmt.Println(sum)
}

func eratosthenes(limit int, primes chan int) {
    limitRoot := int(math.Sqrt(float64(limit)))
    composite := make([]bool, limit + 1)
    for i := 2; i <= limitRoot; i++ {
        if !composite[i] {
            for j := i*i; j <= limit; j += i {
                composite[j] = true
            }
        }
    }
    for k := 2; k < limit; k++ {
        if !composite[k] {
            primes <- k
        }
    }
    close(primes)
}
