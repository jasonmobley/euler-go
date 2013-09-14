// A Pythagorean triplet is a set of three natural numbers,
// a < b < c, for which, a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which
// a + b + c = 1000.
//
// Find the product a*b*c.

package main

import "fmt"
import "math"

func main() {
    var a, b, c float64

    for a = 2; a < 1000; a++ {
        for b = a; b < 1000; b++ {
            c = math.Sqrt(float64(a*a + b*b))
            if c == math.Trunc(c) && a + b + c == 1000 {
                fmt.Printf("%d\n", int(a*b*c))
            }
        }
    }
}
