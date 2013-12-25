// Starting in the top left corner of a 2×2 grid,
// and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
// 
// How many such routes are there through a 20×20 grid?

package main

import "fmt"

func main() {
    var i, n, sum uint64

    // Looking at a 20x20 grid
    n = 20

    // Using combinations to find the answer, but split the problem up
    // to keep the size of the factorials needed within a uint64. The
    // total path will be 2n steps long, so splitting the problem in
    // half means finding the sum of the number of paths with each
    // possible number of rightward or downward movements in steps 1 to n
    // corresponding with the inverse of movements in steps n+1 to 2n.
    // The sum starts at 2 to account for two paths where the first n
    // movements are all down or right movements (meaning we'd end up
    // trying to calculate nCr(n, 0) which wouldn't work).
    sum = 2

    for i = 1; i < n; i++ {
        sum += (nCr(n, i) * nCr(n, n - i))
    }

    fmt.Println(sum)
}

// Compute the number of combinations of choosing r elements from n options.
// This is the regular factorial-based version of this computation, which
// works for us, since the biggest factorial we need is 20! which can fit
// into a uint64.
func nCr(n, r uint64) uint64 {
    if (n == 0 || r == 0) {
        return 0;
    }
    if n == r {
        return n;
    }
    return fact(n) / (fact(n - r) * fact(r))
}

// Simple factorial function. Shame go doesn't support TCO...
func fact(n uint64) uint64 {
    if (n == 0) {
        return 1
    }
    return n * fact(n - 1)
}
