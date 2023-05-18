package main

import (
	"fmt"
	"strconv"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.

Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.

Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

func possibleTimes(digits []int) int {
	// Your code here
	fmt.Println("Input: " + strconv.Itoa(digits[0]) + ", " + strconv.Itoa(digits[1]) + ", " + strconv.Itoa(digits[2]) + ", " + strconv.Itoa(digits[3]))
	if len(digits) > 4 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(digits); i++ {
		time := [4]int{digits[i], 0, 0, 0}
		for j := 0; j < len(digits); j++ {
			if i == j {
				continue
			}
			time[1] = digits[j]
			if time[0] == time[1] {
				return 0
			}
			for k := 0; k < len(digits); k++ {
				if i == k || j == k {
					continue
				}
				time[2] = digits[k]
				if time[0] == time[2] || time[1] == time[2] {
					return 0
				}
				for l := 0; l < len(digits); l++ {
					if i == l || j == l || k == l {
						continue
					}
					time[3] = digits[l]
					if time[3] == time[0] || time[3] == time[1] || time[3] == time[2] {
						return 0
					}
					if time[0]*10+time[1] > 23 || time[2] > 6 {
						break
					}
					if cnt == 0 {
						fmt.Print("Valid times: ")
					}
					fmt.Print(strconv.Itoa(time[0]) + "" + strconv.Itoa(time[1]) + ":" + strconv.Itoa(time[2]) + "" + strconv.Itoa(time[3]) + " ")
					cnt++
					break
				}
			}
		}
	}
	fmt.Println()
	return cnt
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}
