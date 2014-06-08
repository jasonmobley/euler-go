// A unit fraction contains 1 in the numerator.
// The decimal representation of the unit fractions with denominators 2 to 10 are given:
// 
// 1/2	= 	0.5
// 1/3	= 	0.(3)
// 1/4	= 	0.25
// 1/5	= 	0.2
// 1/6	= 	0.1(6)
// 1/7	= 	0.(142857)
// 1/8	= 	0.125
// 1/9	= 	0.(1)
// 1/10	= 	0.1
//
// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle.
// It can be seen that 1/7 has a 6-digit recurring cycle.
// 
// Find the value of d < 1000 for which 1/d contains the longest
// recurring cycle in its decimal fraction part.
//
package main

import "fmt"
import "math/big"
import "github.com/jasonmobley/euler-go/primes"

func main() {
    var d, currentCycleLength, maxCycleLength, dMax int
    var primeSieve chan int = primes.Sieve()
    defer close(primeSieve)

    for d = <-primeSieve; d < 1000; d = <-primeSieve {
        currentCycleLength = cycleLength(d)

        if currentCycleLength > maxCycleLength {
            maxCycleLength = currentCycleLength
            dMax = d
        }
        //fmt.Printf("%d -> %d\n", d, cycleLength(d))
    }

    fmt.Println(dMax)
}

func cycleLength(n int) int {
    if n == 2 || n == 5 {
        return 0
    }

    var exp int

    test := big.NewInt(0)
    nBig := big.NewInt(int64(n))

    for exp = 1; true; exp++ {
        if test.Mod(nines(exp), nBig).Int64() == 0 {
            return exp
        }
    }

    return 0
}

func nines(n int) *big.Int {
    val := big.NewInt(0)
    one := big.NewInt(1)
    ten := big.NewInt(10)

    val.Exp(ten, big.NewInt(int64(n)), nil)

    return val.Sub(val, one)
}
