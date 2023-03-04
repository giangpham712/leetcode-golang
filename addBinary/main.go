package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := addBinary("1010", "1011")

	fmt.Println(result)
}

func addBinary(a string, b string) string {

	i := len(a) - 1
	j := len(b) - 1

	result := ""

	var numA, numB int
	carry := 0
	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			if a[i] == '1' {
				numA = 1
			} else {
				numA = 0
			}
			sum += numA
		}

		if j >= 0 {
			if b[j] == '1' {
				numB = 1
			} else {
				numB = 0
			}
			sum += numB
		}

		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}

		result = strconv.Itoa(sum%2) + result

		i--
		j--

	}

	if carry > 0 {
		result = strconv.Itoa(1) + result
	}

	return result
}
