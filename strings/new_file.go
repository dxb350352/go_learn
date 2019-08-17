package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "http://192.168.130.240:9005/search/search"
	fmt.Println(s)
	fmt.Println(strings.HasSuffix(s, "search"))
	fmt.Println(strings.HasPrefix(s, "http"))
	var dd = "3ed"
	fmt.Println(strings.EqualFold(dd, "3eD"))
	fmt.Println("------------------------")
	fmt.Println(s[len("http") + 2:])

	fmt.Println(strings.Fields("e w 32 ew e2  a"))
	fmt.Println(strings.FieldsFunc("0,10", func(r rune) bool {
		return r == ','
	}))

	fmt.Println(strings.Replace("fdsa fdsa fdsa fdsa   ", " ", "", -1))

	fmt.Println("TrimSpace...........",strings.TrimSpace(`         123456782229

fds

	      `))
	fmt.Println(strings.TrimSpace("   123456782229    \r\n"))
	fmt.Println("12345679"[1:5] == "2345")

	fmt.Println(strings.TrimRight("123400000", "4564560"))

	index := strings.Index(s, "/search")
	fmt.Println(s[:index + 1])
	fmt.Println(strings.Index(s, "ddddd"))
	s = "fdsafds/fdsa/fdsa?fd=ds"
	if index := strings.Index(s, "?"); index != -1 {
		s = s[:index]
	}
	fmt.Println(s)
	gql := fmt.Sprintf(`g.V().hasId("%s")`, "fds")
	fmt.Println(gql)
}
