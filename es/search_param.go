package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	searchJson := simplejson.New()
	searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, "person_phone:13242503108")
	searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, "AND")
	searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
	searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)
	byt, err := searchJson.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byt))
}

