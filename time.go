package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"github.com/sas/utils"
)

func main() {
	gogogo(`now-1d`)
}

func gogogo(st string) {
	b, err := regexp.MatchString(`\d+`, st)
	fmt.Println(b);
	fmt.Println(err);
	ab, errr := strconv.Atoi(st)
	fmt.Println(ab);
	fmt.Println(errr);
	t := time.Unix(int64(ab) / 1000, 0)
	fmt.Println(t);
	fmt.Println(utils.ParseFloat64("116.28835088721563"))
	fmt.Println(time.Now().Day(), time.Now().Month(), time.Now().Year(),time.Now().Format("2006-01-02"))
	var tttt time.Time
	fmt.Println(tttt.Format("2006-01-02"))
}
