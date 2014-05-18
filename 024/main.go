// A permutation is an ordered arrangement of objects. For example,
// 3124 is one possible permutation of the digits 1, 2, 3 and 4.
// If all of the permutations are listed numerically or alphabetically,
// we call it lexicographic order. The lexicographic permutations
// of 0, 1 and 2 are:
//
// 012   021   102   120   201   210
//
// What is the millionth lexicographic permutation of the digits
// 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
//
package main

import "fmt"

func main() {
    const alphabet = "0123456789"
    index := 0

    for str := range permutations(alphabet) {
        index += 1

        // Only care about the millionth one
        if index == 1000000 {
            fmt.Println(str)
        }
    }
}

func permutations(alphabet string) <-chan string {
    c := make(chan string, len(alphabet))

    go func() {
        defer close(c)

        if len(alphabet) == 0 {
            return
        }

        Word(alphabet).Permute(c)
    }()

    return c
}

type Word []rune

// Permute generates all possible combinations of
// the current word. This assumes the runes are sorted.
func (w Word) Permute(out chan<- string) {
    if len(w) <= 1 {
        out <- string(w)
        return
    }

    // Write first result manually.
    out <- string(w)

    // Find and print all remaining permutations.
    for w.next() {
        out <- string(w)
    }
}

// next performs a single permutation by shuffling characters around.
// Returns false if there are no more new permutations.
func (w Word) next() bool {
    var left, right int

    left = len(w) - 2
    for w[left] >= w[left+1] && left >= 1 {
        left--
    }

    if left == 0 && w[left] >= w[left+1] {
        return false
    }

    right = len(w) - 1
    for w[left] >= w[right] {
        right--
    }

    w[left], w[right] = w[right], w[left]

    left++
    right = len(w) - 1

    for left < right {
        w[left], w[right] = w[right], w[left]
        left++
        right--
    }

    return true
}
