package main

import (
	"fmt"
	"github.com/cch123/elasticsql"
)

var select_sql = `
select * from aaa
where a=1 and (x = '三个男人'
or create_time between '2015-01-01T00:00:00+0800' and '2016-01-01T00:00:00+0800')
and process_id > 1 or b like '%aaa%' order by id desc limit 100,10
`
var update = `update aaa set a=2,x='一个男人'`
var insert = `insert into aaa  (a,x) values(2,'一个男人')`

var group = `select 	person_phone ,count(*) as ct from gaxz_phone_bill group by person_phone`

func main() {
	fmt.Println(elasticsql.Convert(select_sql))
	fmt.Println(elasticsql.Convert(update))
	fmt.Println(elasticsql.Convert(insert))
	//fmt.Println(elasticsql.Convert(group))

	var test = `select * from t where a like '%as%'`
	fmt.Println(elasticsql.Convert(test))
	fmt.Println(elasticsql.Convert(`select id from sss where id like 'as%' and id not like 'asd-%'`))

	dsl, _, _ := elasticsql.Convert(`select * from t where starttime >= 0 and starttime <=1`)
	fmt.Println(dsl)
}
