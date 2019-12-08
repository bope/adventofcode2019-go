package part2

import (
	"bytes"
	"strings"
)

func Solution(input []rune, width, height int) string {
	layerSize := (width * height)
	layerCount := len(input) / layerSize

	var picture bytes.Buffer
	var pixel, layerPixel rune

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel = '2'
			for l := layerCount - 1; l >= 0; l-- {
				layerPixel = input[(layerSize*l)+(y*width)+x]
				if layerPixel != '2' {
					pixel = layerPixel
				}
			}
			picture.WriteRune(pixel)
		}
		picture.WriteString("\n")
	}
	return strings.TrimSpace(picture.String())

}
