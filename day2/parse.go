package day2

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Parse(r io.Reader) ([]int, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := string(data)
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ",")
	ints := make([]int, len(parts))

	for i, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ints[i] = v
	}

	return ints, nil
}
