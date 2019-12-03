package day3

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var directions = map[rune]Vec{
	'U': Vec{0, 1},
	'D': Vec{0, -1},
	'L': Vec{-1, 0},
	'R': Vec{1, 0},
}

type Step struct {
	Length    int
	Direction Vec
}

func parsePath(input string) ([]Step, error) {
	var err error
	parts := strings.Split(input, ",")
	steps := make([]Step, len(parts))

	for i, step := range parts {
		steps[i].Direction = directions[rune(step[0])]
		if steps[i].Length, err = strconv.Atoi(step[1:]); err != nil {
			return nil, err
		}
	}

	return steps, nil
}

func Parse(r io.Reader) ([]Step, []Step, error) {
	data, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, nil, err
	}

	s := string(data)
	s = strings.TrimSpace(s)

	parts := strings.Split(s, "\n")

	first, err := parsePath(parts[0])
	if err != nil {
		return nil, nil, err
	}

	second, err := parsePath(parts[1])

	if err != nil {
		return nil, nil, err
	}

	return first, second, nil
}
