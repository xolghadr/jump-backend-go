package nasa

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
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())
		label := inputs[0]
		var numbers []int = []int{}
		for _, score := range inputs[1:] {
			num, _ := strconv.Atoi(score)
			numbers = append(numbers, num)
		}
		length := len(numbers)
		if length < 3 {
			println(label, "0")
			break
		}

		count := 0
		for i := 0; i <= length-3; i++ {
			j := length
			for j >= i+3 {
				if IsArithmetic(numbers[i:j]) {
					count++
				}
				j--
			}
		}

		fmt.Println(label, float64(count))
	}
}

func IsArithmetic(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	diff := arr[0] - arr[1]
	for i := 2; i < len(arr); i++ {
		if arr[i-1]-arr[i] != diff {
			return false
		}
	}
	return true
}
