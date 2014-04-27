// By starting at the top of the triangle below and moving to adjacent
// numbers on the row below, the maximum total from top to bottom is 23.
// 
//       [3]
//     [7]  4
//    2  [4]  6
//  8   5  [9]  3
// 
// That is, 3 + 7 + 4 + 9 = 23.
// 
// Find the maximum total from top to bottom of the triangle below:
// 
// 75
// 95 64
// 17 47 82
// 18 35 87 10
// 20 04 82 47 65
// 19 01 23 75 03 34
// 88 02 77 73 07 63 67
// 99 65 04 28 06 16 70 92
// 41 41 26 56 83 40 80 70 33
// 41 48 72 33 47 32 37 16 94 29
// 53 71 44 65 25 43 91 52 97 51 14
// 70 11 33 28 77 73 17 78 39 68 17 57
// 91 71 52 38 17 14 91 43 58 50 27 29 48
// 63 66 04 68 89 53 67 30 73 16 69 87 40 31
// 04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
// 
// NOTE: As there are only 16384 routes, it is possible to solve this
// problem by trying every route. However, Problem 67, is the same
// challenge with a triangle containing one-hundred rows; it cannot be
// solved by brute force, and requires a clever method! ;o)
package main

import "fmt"
import "io/ioutil"
import "math"
import "os"
import "strconv"
import "strings"

func main() {
    var filename string

    if len(os.Args) > 1 {
        filename = os.Args[1]
    } else {
        fmt.Println("No filename given.")
        os.Exit(1)
    }

    number_file, read_err := ioutil.ReadFile(filename)

    if read_err != nil {
        fmt.Println(read_err.Error())
        os.Exit(1)
    }

    filecontent := strings.TrimSpace(string(number_file))
    filelines := strings.Split(filecontent, "\n")
    numbers := make([][]int, len(filelines))

    // Populate the numbers slice with the content of the file
    for line_index, line := range filelines {
        line_items := strings.Split(line, " ")
        numbers[line_index] = make([]int, len(line_items))

        for item_index, item := range line_items {
            numbers[line_index][item_index], _ = strconv.Atoi(string(item))
        }
    }

    for row_index := 1; row_index < len(numbers); row_index++ {
        for col_index, number := range numbers[row_index] {
            if col_index == 0 {
                numbers[row_index][col_index] = number + numbers[row_index-1][0]
            } else if col_index == (len(numbers[row_index]) - 1) {
                numbers[row_index][col_index] = number + numbers[row_index-1][len(numbers[row_index-1])-1]
            } else {
                numbers[row_index][col_index] = int(math.Max(float64(number + numbers[row_index-1][col_index-1]), float64(number + numbers[row_index-1][col_index])))
            }
        }
    }

    max := 0

    for _, path_sum := range numbers[len(numbers)-1] {
        if path_sum > max {
            max = path_sum
        }
    }

    fmt.Println(max)
}

