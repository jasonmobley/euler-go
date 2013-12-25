// 2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2**1000?
package main

import "fmt"
import "math/big"
import "strconv"

func main() {
    var sum uint64 = 0
    var n = big.NewInt(0)

    n.Exp(big.NewInt(2), big.NewInt(1000), nil)

    for _, digitRune := range n.String() {
        digit, err := strconv.ParseUint(string(digitRune), 10, 8)

        if err == nil {
            sum += digit
        }
    }
    fmt.Println(sum)
}
