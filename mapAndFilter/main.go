package main

import "math"

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	return math.Sqrt(float64(x)) == float64(int(math.Sqrt(float64(x))))
}

func IsPalindrome(x int) bool {
	if x < 0 {
		x = -x
	}
	if x < 10 {
		return true
	}
	
	num := x
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return x == reversed
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	filtered := []int{}
	for _, num := range input {
		if f(num) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

func Map(input []int, m MapperFunc) []int {
	mapped := []int{}
	for _, num := range input {
		mapped = append(mapped, m(num))
	}
	return mapped
}
