package part1

func Count(input []rune, val rune) int {
	ret := 0
	for _, v := range input {
		if v == val {
			ret++
		}
	}
	return ret
}

func Solution(input []rune, width, height int) int {
	layerSize := (width * height)
	layerCount := len(input) / layerSize

	var leastZerosLayer int
	leastZeros := layerSize + 1

	for i := 0; i < layerCount; i++ {
		zeros := Count(input[layerSize*i:layerSize*(i+1)], '0')
		if zeros < leastZeros {
			leastZeros = zeros
			leastZerosLayer = i
		}
	}

	ones := Count(input[layerSize*leastZerosLayer:layerSize*(leastZerosLayer+1)], '1')
	twos := Count(input[layerSize*leastZerosLayer:layerSize*(leastZerosLayer+1)], '2')
	return ones * twos

}
