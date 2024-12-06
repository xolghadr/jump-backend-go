package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	var countryCode map[string]string = make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())
		country := inputs[0]
		code := inputs[1]
		countryCode[code] = country
	}

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())
		for _, input := range inputs {
			if len(input) > 3 {
				firstThree := input[:3]
				if country, found := countryCode[firstThree]; found {
					fmt.Println(country)
				} else {
					fmt.Println("Invalid Number")
				}
			} else {
				fmt.Println("Invalid Number")
			}
		}
	}
}
