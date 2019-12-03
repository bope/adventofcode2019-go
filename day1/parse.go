package day1

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Parse(r io.Reader) ([]int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s := string(b)
	s = strings.TrimRight(s, "\n")
	parts := strings.Split(s, "\n")
	ints := make([]int, len(parts))
	for i, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		ints[i] = v
	}
	return ints, nil
}
