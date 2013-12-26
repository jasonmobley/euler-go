// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
// 
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
// how many letters would be used?
// 
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two)
// contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use
// of "and" when writing out numbers is in compliance with British usage.
package main

import "fmt"
import "strings"

func main() {
    count := 0

    for i := 1; i <= 1000; i += 1 {
        count += len(strings.Join(strings.Split(numberAsWords(i), " "), ""))
    }

    fmt.Println(count)
}

var ones []string = []string{
    "",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen",
}

var tens []string = []string{
    "",
    "ten",
    "twenty",
    "thirty",
    "forty",
    "fifty",
    "sixty",
    "seventy",
    "eighty",
    "ninety",
}

func numberAsWords(n int) string {
    var result string

    switch {
    case n > 999:
        result = numberAsWords(n/1000) + " thousand " + numberAsWords(n%1000)
    case n > 99:
        result = numberAsWords(n/100) + " hundred"
        if (n % 100) > 0 {
            result += " and " + numberAsWords(n % 100)
        }
        break
    case n > 19:
        result = tens[n/10] + " " + numberAsWords(n%10)
        break
    default:
        result = ones[n]
        break
    }

    return strings.TrimSpace(result)
}
