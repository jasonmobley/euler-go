// The following iterative sequence is defined for the set of positive integers:
// 
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
// 
// Using the rule above and starting with 13, we generate the following sequence:
// 
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// 
// It can be seen that this sequence (starting at 13 and finishing at 1)
// contains 10 terms. Although it has not been proved yet (Collatz Problem),
// it is thought that all starting numbers finish at 1.
// 
// Which starting number, under one million, produces the longest chain?
// 
// NOTE: Once the chain starts the terms are allowed to go above one million.

package main

import "fmt"

func main() {
    var num, current, max int

    for i := 2; i < 1000000; i++ {
        current = collatzLength(i)

        if current > max {
            max = current
            num = i
        }
    }

    fmt.Printf("Longest Collatz sequence starts with %d and has %d terms\n", num, max)
}

func collatzLength(n int) int {
    current := n
    length := 1

    for current != 1 {
        if current % 2 == 0 {
            current = current / 2
        } else {
            current = (3 * current) + 1
        }
        length += 1
    }

    return length
}
