// 2520 is the smallest number that can be divided by each
// of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly
// divisible by all of the numbers from 1 to 20?
package main

import "fmt"

func main() {
   var value uint64 = 20
   found := false

   for !found {
      if value % 10 == 0 && value % 11 == 0 && value % 12 == 0 && value % 13 == 0 &&
         value % 14 == 0 && value % 15 == 0 && value % 16 == 0 && value % 17 == 0 &&
         value % 18 == 0 && value % 19 == 0 && value % 20 == 0 {
         found = true;
      } else {
         value++;
      }
   }

   fmt.Printf("The smallest number divisible by 1-20 is %d.\n", value);
}

