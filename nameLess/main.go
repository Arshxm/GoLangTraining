package main

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	min := (*numbers)[0]
	for _, number := range *numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func ReverseSlice(numbers *[]int) {
	for i := 0; i < len(*numbers)/2; i++ {
		(*numbers)[i], (*numbers)[len(*numbers)-i-1] = (*numbers)[len(*numbers)-i-1], (*numbers)[i]
	}
}

func SwapElements(numbers *[]int, i, j int) {
	if i < 0 || j < 0 || i >= len(*numbers) || j >= len(*numbers) {
		return
	}
	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
