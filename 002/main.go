// Each new term in the Fibonacci sequence is generated
// by adding the previous two terms. By starting with
// 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// 
// By considering the terms in the Fibonacci sequence
// whose values do not exceed four million, find the
// sum of the even-valued terms.
package main

import "fmt"

func main() {
    sum := 0
    fib_curr := 2
    fib_prev := 1

    for fib_curr <= 4000000 {
        if fib_curr % 2 == 0 {
            sum += fib_curr
        }

        fib_curr, fib_prev = fib_curr + fib_prev, fib_curr
    }

    fmt.Println(sum);
}

