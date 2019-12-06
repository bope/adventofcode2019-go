package day4

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Parse(r io.Reader) (a int, b int, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	s := string(data)
	s = strings.TrimSpace(s)
	parts := strings.Split(s, "-")

	if len(parts) != 2 {
		err = fmt.Errorf("invalid input")
		return
	}

	if a, err = strconv.Atoi(parts[0]); err != nil {
		return
	}

	if b, err = strconv.Atoi(parts[1]); err != nil {
		return
	}

	return
}
