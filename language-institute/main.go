package language_institute

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
		label := scanner.Text()
		scanner.Scan()
		scores := strings.Fields(scanner.Text())
		var count, sum int
		for _, score := range scores {
			num, _ := strconv.Atoi(score)
			sum += num
			count++
		}
		average := float64(sum) / float64(count)
		fmt.Println(label, " ", NameAverage(average))
	}
}

func NameAverage(average float64) string {
	switch {
	case average >= 80:
		return "Excellent"
	case average < 80 && average >= 60:
		return "Very Good"
	case average < 60 && average >= 40:
		return "Good"
	default:
		return "Fair"
	}
}
