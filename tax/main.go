package tax

import (
	"fmt"
	"math"
)

func main() {
	var n = 0
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	fmt.Println(math.Floor(CalculateTax(float64(n))))
}

func CalculateTax(income float64) float64 {
	var tax float64 = 0
	switch {
	case income > 1000:
		{
			tax = (income-1000)*20/100 + CalculateTax(1000)
		}
	case income <= 1000 && income > 500:
		{
			tax = (income-500)*15/100 + CalculateTax(500)
		}
	case income <= 500 && income > 100:
		{
			tax = (income-100)*10/100 + CalculateTax(100)
		}
	case income <= 100:
		{
			tax = income * 5 / 100
		}
	}
	return tax
}
