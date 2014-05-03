// You are given the following information, but you may prefer to do some research for yourself.
// 
// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.

// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
package main

import "fmt"

func main() {
    var dayOfWeek, day, month, year int = 1, 1, 1, 1900
    daysInMonth := []int{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 }
    sundaysOnTheFirstOfTheMonth := 0

    for year < 2001 {
        // Reference dayOfWeek is 1900-01-01, but we don't start counting until
        // 1901-01-01, so add year checks to skip the first year of Sundays (if any)
        if year > 1900 && year < 2001 && dayOfWeek == 0 && day == 1 {
            sundaysOnTheFirstOfTheMonth += 1
        }

        // Move to the next day...

        // Advance the dayOfWeek, wrapping around every 7th day
        dayOfWeek = (dayOfWeek + 1) % 7

        // Special case to give Feb 29 days on leap years
        if year % 4 == 0 && (year % 100 != 0 || (year % 100 == 0 && year % 400 == 0)) && month == 2 && day == 28 {
            day += 1
            continue
        }

        // Advance the day of the month
        day += 1

        // If we've gone past the number of days in the current month,
        // advance month and reset day of month back to 1
        if day > daysInMonth[month-1] {
            day = 1
            month += 1
        }

        // If we've gone past the number of months in a year,
        // advance the year and reset month back to 1
        if month >= 13 {
            month = 1
            year += 1
        }
    }

    fmt.Println(sundaysOnTheFirstOfTheMonth)
}

