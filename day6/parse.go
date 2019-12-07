package day6

import (
	"io"
	"io/ioutil"
	"strings"
)

func Parse(r io.Reader) ([][2]string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s := string(data)
	s = strings.TrimSpace(s)

	lines := strings.Split(s, "\n")
	ret := make([][2]string, len(lines))

	for i, part := range lines {
		var d [2]string
		copy(d[:], strings.SplitN(part, ")", 2))
		ret[i] = d

	}
	return ret, nil
}
