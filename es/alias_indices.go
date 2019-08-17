package main

import (
	js "github.com/bitly/go-simplejson"
	"github.com/mattbaird/elastigo/lib"
	"errors"
	"fmt"
)

var Es     *elastigo.Conn

func main() {
	fmt.Println(AliasIndices("alias11"))
}
func AliasIndices(alias string, indices ...string) error {
	url := "/_aliases"
	remove := js.New()
	remove.SetPath([]string{"remove", "indices"}, []string{"*"})
	remove.SetPath([]string{"remove", "alias"}, alias)

	pdata := js.New()
	if len(indices) > 0 {
		add := js.New()
		add.SetPath([]string{"add", "indices"}, indices)
		add.SetPath([]string{"add", "alias"}, alias)
		pdata.Set("actions", []*js.Json{remove, add})
	} else {
		pdata.Set("actions", []*js.Json{remove})
	}

	body, err := Es.DoCommand("POST", url, nil, pdata)
	if err != nil {
		return err
	}
	result, err := js.NewJson(body)
	if err != nil {
		return err
	}
	if result.Get("acknowledged").MustBool() != true {
		return errors.New("acknowledge is not true")
	}
	return nil
}

func init() {
	Es = elastigo.NewConn()
	Es.Hosts = []string{"192.168.130.201"}
	Es.Port = "9200"
}