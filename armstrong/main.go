package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	digits := decodeStringIntoDigits(input)
	if len(digits) == 0 {
		fmt.Printf("YES")
	} else {
		sum := 0
		for _, digit := range digits {
			sum += digit
		}
		isArmstrong := checkArmstrong(sum)
		if isArmstrong {
			fmt.Printf("YES")
		} else {
			fmt.Printf("NO")
		}
	}
}

func decodeStringIntoDigits(input string) []int {
	digits := make([]int, 0)
	var d0 string
	for index, char := range input {
		d, err := strconv.Atoi(string(char))
		if err == nil {
			d0 += strconv.Itoa(d)
		} else if d0 != "" {
			digit, _ := strconv.Atoi(d0)
			digits = append(digits, digit)
			d0 = ""
		}
		if d0 != "" && index == len(input)-1 {
			digit, _ := strconv.Atoi(d0)
			digits = append(digits, digit)
		}
	}
	return digits
}

func checkArmstrong(number int) bool {
	numberString := strconv.Itoa(number)
	length := len(numberString)
	power := 0
	for _, char := range numberString {
		digit, _ := strconv.Atoi(string(char))
		power += int(math.Pow(float64(digit), float64(length)))
	}

	return power == number
}
