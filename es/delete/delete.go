package main

import (
	"github.com/mattbaird/elastigo/lib"
	"errors"
	"github.com/bitly/go-simplejson"
)

var Es     *elastigo.Conn

func main() {

}

func init() {
	Es = elastigo.NewConn()
	Es.Hosts = []string{"192.168.130.201"}
	Es.Port = "9200"
}

func delete(id string)error  {
	url := "/_aliases"
	remove := simplejson.New()
	remove.SetPath([]string{"remove", "indices"}, []string{"*"})
	remove.SetPath([]string{"remove", "alias"}, alias)


	body, err := Es.DoCommand("POST", url, nil, pdata)
	if err != nil {
		return err
	}
	result, err := simplejson.NewJson(body)
	if err != nil {
		return err
	}
	if result.Get("acknowledged").MustBool() != true {
		return errors.New("acknowledge is not true")
	}
	return nil
}