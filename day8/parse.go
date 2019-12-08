package day8

import (
	"io"
	"io/ioutil"
	"strings"
)

func Parse(r io.Reader) ([]rune, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := string(data)
	s = strings.TrimSpace(s)
	return []rune(s), nil
}
