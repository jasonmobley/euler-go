// A palindromic number reads the same both ways. The largest
// palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import "fmt"

func main() {
    var buffer string;
    var product int;
    var max_palindrome int = 0;

    for i := 100; i < 1000; i++ {
        for j := 100; j < 1000; j++ {
            product = i * j;
            buffer = fmt.Sprintf("%d", product);
            if buffer == reverse(buffer) && product > max_palindrome {
                max_palindrome = product;
            }
        }
    }

    fmt.Printf("The largest palindrome is %d.\n", max_palindrome);
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

