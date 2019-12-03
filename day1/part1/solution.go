package part1

func CalcMass(mass int) int {
	return (mass / 3) - 2
}

func Solution(input []int) int {
	total := 0
	for _, mass := range input {
		total += CalcMass(mass)
	}
	return total
}
