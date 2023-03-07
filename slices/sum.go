package main

func Sum(input []int) int {
	var sum int
	for _, number := range input {
		sum += number
	}
	return sum
}

func SumAll(numberslices ...[]int) []int {
	var total []int
	for _, numberslice := range numberslices {
		total = append(total, Sum(numberslice))
	}
	return total
}

func SumAllTails(numberslices ...[]int) []int {
	var total []int
	var slicetotal int
	for _, numberslice := range numberslices {
		if len(numberslice) == 0 {
			slicetotal = 0
		} else {
			slicetotal = Sum(numberslice[1:])
		}
		total = append(total, slicetotal)
	}
	return total
}
