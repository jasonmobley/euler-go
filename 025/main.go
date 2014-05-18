// The Fibonacci sequence is defined by the recurrence relation:
// 
// F(n) = F(n−1) + F(n−2), where F(1) = 1 and F(2) = 1.
// Hence the first 12 terms will be:
// 
// F(1) = 1
// F(2) = 1
// F(3) = 2
// F(4) = 3
// F(5) = 5
// F(6) = 8
// F(7) = 13
// F(8) = 21
// F(9) = 34
// F(10) = 55
// F(11) = 89
// F(12) = 144
// The 12th term, F(12), is the first term to contain three digits.
// 
// What is the first term in the Fibonacci sequence to contain 1000 digits?
//
package main

import "fmt"
import "math/big"

func main() {
    fibNMinus2 := big.NewInt(1)
    fibNMinus1 := big.NewInt(1)
    fibN       := big.NewInt(2)
    n := 3

    for len(fibN.String()) < 1000 {
        fibNMinus2.Set(fibNMinus1)
        fibNMinus1.Set(fibN)
        fibN.Add(fibNMinus1, fibNMinus2)
        n += 1
        //fmt.Printf("f(%d) = %s\n", n, fibN.String())
    }

    fmt.Println(n)
}
