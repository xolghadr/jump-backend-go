package noname

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	m := (*numbers)[0]
	for i := 1; i < len(*numbers); i++ {
		if (*numbers)[i] < m {
			m = (*numbers)[i]
		}
	}
	return m
}

func ReverseSlice(numbers *[]int) {
	reversed := make([]int, len(*numbers), cap(*numbers))
	length := len(*numbers)
	for i := 0; i < length; i++ {
		reversed[i] = (*numbers)[length-1-i]
	}
	*numbers = reversed
}

func SwapElements(numbers *[]int, i, j int) {
	if i < 0 || j < 0 || i >= len(*numbers) || j >= len(*numbers) {
		return
	}
	el1, el2 := (*numbers)[i], (*numbers)[j]
	(*numbers)[i], (*numbers)[j] = el2, el1
}
