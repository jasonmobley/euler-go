// Using names.txt, a 46K text file containing over five-thousand first names,
// begin by sorting it into alphabetical order. Then working out the alphabetical
// value for each name, multiply this value by its alphabetical position in the
// list to obtain a name score.
// 
// For example, when the list is sorted into alphabetical order, COLIN, which is
// worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN
// would obtain a score of 938 × 53 = 49714.
// 
// What is the total of all the name scores in the file?
package main

import "fmt"
import "io/ioutil"
import "os"
import "sort"
import "strconv"
import "strings"

func main() {
    var filename string
    sumOfScores := 0

    if len(os.Args) > 1 {
        filename = os.Args[1]
    } else {
        fmt.Println("No filename given.")
        os.Exit(1)
    }

    input_file, read_err := ioutil.ReadFile(filename)

    if read_err != nil {
        fmt.Println(read_err.Error())
        os.Exit(1)
    }

    names := strings.Split(strings.TrimSpace(string(input_file)), ",")

    // Trim quotes off each name
    for index, name := range names {
        names[index], _ = strconv.Unquote(name)
    }

    sort.Strings(names)

    for index, name := range names {
        nameValue := 0
        namePosition := index + 1

        for _, char := range name {
            nameValue += (int(char) - 64)
        }

        nameScore := nameValue * namePosition

        // fmt.Printf("%s: %d * %d = %d\n", name, nameValue, namePosition, nameScore)

        sumOfScores += nameScore
    }

    fmt.Println(sumOfScores)
}
