// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
package main

import "fmt"
import "math"
import "os"
import "strconv"
import "strings"

const MAX_FACTORS uint64 = 100

func main() {
    var arg uint64
    var arg_error error
    var factors [MAX_FACTORS]uint64

    if len(os.Args) < 2 || len(strings.TrimSpace(os.Args[1])) == 0 {
        fmt.Println("No argument (or argument was just whitespace).");
        os.Exit(1);
    }

    arg, arg_error = strconv.ParseUint(strings.TrimSpace(os.Args[1]), 10, 64)

    if arg_error != nil {
        fmt.Println("Argument must be a non-negative integer.");
        os.Exit(1);
    }

   var work_arg uint64 = arg;
   var root_arg uint64 = uint64(math.Sqrt(float64(arg))) + 1;
   var i uint64;
   var factor_index uint64 = 0;

   fmt.Printf("Factors: ");

   for i = 2; i < root_arg; i++ {
      if work_arg % i == 0 && isPrime(i) {
         factors[factor_index] = work_arg;
         work_arg = work_arg / i;
         factor_index++;

         fmt.Printf("%d ", i);
      }

      if factor_index == MAX_FACTORS {
         fmt.Printf("Number of factors exceeded %d. Quitting.\n", MAX_FACTORS);
         os.Exit(1);
      }
   }

   fmt.Printf("\nThe largest factor of %d is %d.\n", arg, factors[factor_index-1])
}

func isPrime(value uint64) bool {
    var i uint64 = 2
    var rootPlusOne uint64 = uint64(math.Sqrt(float64(value))) + 1

    for i < rootPlusOne {
        if value % i == 0 {
            return false
        }
        i += 1
    }
    return true
}

