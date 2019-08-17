package main

import (
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"regexp"
	"strconv"
	"time"
)

func main() {
	ParseTime("1439366133000")
}

func ParseTime(st string) (*time.Time, error) {
	var t time.Time
	ab, errr := strconv.Atoi(st)
	if errr == nil {
		t = time.Unix(int64(ab)/1000, 0)
	} else if st == "now" {
		t = time.Now()
	} else {
		var reg = regexp.MustCompile(`^now(\+|\-)(\d)(\w)`)
		r := reg.FindAllStringSubmatch(st, -1)
		fmt.Printf("r:%# v\n", pretty.Formatter(r))

		if len(r[0]) != 4 {
			err := errors.New("synax error")
			fmt.Printf("err:%# v\n", pretty.Formatter(err))

			return nil, err
		}

		op := r[0][1]

		sn := r[0][2]
		n, err := strconv.Atoi(sn)
		if err != nil {
			return nil, err
		}
		dw := r[0][3]
		switch dw {
		case "ns", "us", "ms", "s", "m", "h":
			i, err := time.ParseDuration(sn + dw)
			if err != nil {
				fmt.Printf("err:%# v\n", pretty.Formatter(err))
				return nil, err
			}
			if op == "-" {
				t = time.Now().Add(-i)
			} else if r[0][1] == "+" {
				t = time.Now().Add(i)
			}

		case "y":
			if op == "-" {

				t = time.Now().AddDate(-n, 0, 0)
			} else {
				t = time.Now().AddDate(n, 0, 0)
			}

		case "M":
			if op == "-" {

				t = time.Now().AddDate(0, -n, 0)
			} else {
				t = time.Now().AddDate(0, n, 0)
			}

		case "d":
			if op == "-" {
				t = time.Now().AddDate(0, 0, -n)
			} else {
				t = time.Now().AddDate(0, 0, n)
			}

		}

	}
	n := t.Unix()
	fmt.Printf("n*1000:%# v\n", pretty.Formatter(n*1000))
	return &t, nil
}
