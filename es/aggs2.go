package main

import (
	"github.com/mattbaird/elastigo/lib"
	"github.com/farmerx/elasticsql"
	"fmt"
	"strings"
	"github.com/bitly/go-simplejson"
)

var es *elastigo.Conn
var essql *elasticsql.ElasticSQL

func init() {
	es = elastigo.NewConn()
	es.Hosts = []string{"192.168.130.201"}
	es.Port = "9200"

	essql = elasticsql.NewElasticSQL(elasticsql.InitOptions{})
}

var ss = `{
  "query": {
    "bool": {
      "must": [
        {
          "match_all": {}
        }
      ]
    }
  },
  "from": 0,
  "size": 0,
  "aggregations": {
    "colors": {
      "terms": {
        "field": "color"
      },
      "aggregations": {
        "avg_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}`

func main() {
	sql := "select avg(price) from cars  group by color "
	table, esql, err := essql.SQLConvert(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(table, esql)
	data, _ := simplejson.NewJson([]byte(esql))
	uri := fmt.Sprintf("/%s/_search", table)
	byt, err := doCommand("post", uri, nil, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byt))
}
func doCommand(method string, url string, args map[string]interface{}, data interface{}) ([]byte, error) {
	if strings.Index(url, "?") == -1 {
		url = url + "?ignore_unavailable=true"
	} else {
		url = url + "&ignore_unavailable=true"
	}
	body, err := es.DoCommand(method, url, args, data)
	return body, err
}
