# algo-timegen-test
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

# How to run
- init go module
$ go mod init main

- add module requirements and sums
$ go mod tidy

- How to run project 
$ go run times.go

# How to bulid

- build executable result
$ go build .

# Output

Input: 1, 2, 3, 4
Valid times: 12:34 12:43 13:24 13:42 14:23 14:32 21:34 21:43 23:14 23:41
10
Input: 9, 1, 2, 0
Valid times: 19:20 19:02 12:09 10:29 21:09 20:19 09:12 09:21 01:29 02:19
10
Input: 2, 2, 1, 9
0