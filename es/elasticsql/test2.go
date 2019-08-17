package main

import (
	"fmt"
	"github.com/farmerx/elasticsql"
	"strings"
	"time"
)

var sql = `
select * from test where a=1 and ( b="c" or create_time between '2015-01-01T00:00:00+0800' and '2016-01-01T00:00:00+0800') and process_id > 1 order by id desc limit 100,10
`

var sql2 = `
  select avg(age),min(age),max(age),count(student) from test group by class limit 10
`
var sql3 = `
  select * from test group by class,student limit 10
`
var sql4 = `
  select * from test group by date_histogram(field="changeTime",interval="1h",format="yyyy-MM-dd HH:mm:ss")
`

func main() {
	esql := elasticsql.NewElasticSQL(elasticsql.InitOptions{})
	fmt.Println(esql.SQLConvert(sql))
	fmt.Println(esql.SQLConvert(sql2))
	fmt.Println(esql.SQLConvert(sql3))
	fmt.Println(esql.SQLConvert(sql4))
	var test = `SELECT count(*) as count,avg(price) as avg from t  group by colors`
	fmt.Println(esql.SQLConvert(test))
	fmt.Println(esql.SQLConvert(`select a,b from t where a=1 `))
	fmt.Println(esql.SQLConvert(`select * from t where a=1 `))
	fmt.Println(esql.SQLConvert(`SELECT COUNT(DISTINCT calleee164.keyword) as test from t GROUP BY calleee164.keyword `))
	keywords := []string{"18053519944", "17605900691"}
	__time := fmt.Sprintf(" starttime between %d and %d ", 0, time.Now().Unix()*1000)
	sql := fmt.Sprintf("select COUNT(DISTINCT  %s.keyword) from t where calleee164_phone in ('"+strings.Join(keywords, "','")+"') and "+__time+" group by %s.keyword", "vosip", "vosip")
	fmt.Println(sql)
	fmt.Println(esql.SQLConvert(sql))

	fmt.Println(esql.SQLConvert("select * from t where person_name like '%a%' or person_name like '%b%'"))
}
