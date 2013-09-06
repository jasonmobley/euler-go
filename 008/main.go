// Find the greatest product of five consecutive digits
// in the 1000-digit number listed in number.txt.
package main

import "fmt"
import "io/ioutil"
import "os"
import "strconv"

func main() {
   var filename, numbers string

   if len(os.Args) > 1 {
      filename = os.Args[1]
   } else {
      fmt.Println("No filename given.")
      os.Exit(1)
   }

   number_file, read_err := ioutil.ReadFile(filename)

   if read_err != nil {
      fmt.Println(read_err.Error());
      os.Exit(1)
   }

   numbers = string(number_file)

   var d1, d2, d3, d4, d5 int
   var currentProduct int = 0
   var maxProduct int = 0
   var numDigits = len(numbers)

   for index := 4; index < numDigits; index++ {
       d1, _ = strconv.Atoi(string(numbers[index-4]))
       d2, _ = strconv.Atoi(string(numbers[index-3]))
       d3, _ = strconv.Atoi(string(numbers[index-2]))
       d4, _ = strconv.Atoi(string(numbers[index-1]))
       d5, _ = strconv.Atoi(string(numbers[index]))

       if d1 == 0 || d2 == 0 || d3 == 0 || d4 == 0 || d5 == 0 {
           continue
       }

       currentProduct = d1 * d2 * d3 * d4 * d5

       if currentProduct > maxProduct {
           maxProduct = currentProduct
       }
   }

   fmt.Printf("Max product of 5 consecutive digits: %d\n", maxProduct);
}

