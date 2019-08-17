package main

import (
	"github.com/mattbaird/elastigo/lib"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strings"
	"encoding/json"
	"github.com/go-xweb/log"
)

var Es *elastigo.Conn

func init() {
	Es = elastigo.NewConn()
	Es.Hosts = []string{"192.168.130.201"}
	Es.Port = "9200"
}

func main() {
	countType := "terms"
	rs, err := AggCount("sas_cdrcb_qy_ads_02_2018_02_09", "", "", "dst_port", countType, "", 10, false)
	if err != nil {
		log.Fatal(err)
	}
	j1, err := simplejson.NewJson([]byte(rs.Aggregations))
	if err != nil {
		log.Fatal(err)
	}

	field2 := "count1"
	j2, ok := j1.CheckGet(field2)
	if !ok {
		log.Fatal(err)
	}
	value, ok := j2.CheckGet("value")
	if ok {
		jj := []*simplejson.Json{}
		jj2 := simplejson.New()
		jj2.Set("key", countType)
		jj2.Set("value", value)
		jj = append(jj, jj2)
		log.Fatalf("%v", jj)
	}

	buckets, ok := j2.CheckGet("buckets")
	if ok {
		log.Fatalf("%v", buckets)
	}
	buckets, ok = j2.Get("count2").CheckGet("buckets")
	if ok {
		log.Fatalf("%v", buckets)
	}

	log.Fatalf("%v", ConvertJsonToArray(countType, j2))
}

func AggCount(index, Keywords, groupby, field, countType, Operator string, size int, NoTimeFilter bool) (elastigo.SearchResult, error) {
	//fmt.Printf("%v\n", pretty.Formatter()

	// countType is: stats, min, max, avg, percentiles
	if countType == "" {
		countType = "terms"
	}

	if Operator == "" {
		Operator = "AND"
	}

	if size == 0 {
		size = 1000
	} else if size > 1000 {
		size = 1000
	}

	action := "_search"
	var timeFilter *simplejson.Json
	var filters []*simplejson.Json
	if !NoTimeFilter {
		timeFilter = BuildTimeRangeFilter()
		filters = []*simplejson.Json{timeFilter}
	}
	filter := BuildAndFilters(filters)

	searchJson := simplejson.New()
	if Keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, simplejson.New())

	} else {
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, Keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)
	}

	aggs := simplejson.New()

	if filter != nil {
		searchJson.SetPath([]string{"query", "bool", "must"}, filter)
		if groupby != "" {
			aggs.SetPath([]string{"count1", "filter"}, filter)
			aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "field"}, groupby)
			aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "size"}, size)
			aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", countType, "field"}, field)
		} else {
			aggs.SetPath([]string{"count1", countType, "field"}, field)
			if countType == "terms" {
				aggs.SetPath([]string{"count1", countType, "size"}, size)
			}
		}

	} else if groupby != "" {
		aggs.SetPath([]string{"count1", "terms", "field"}, groupby)
		aggs.SetPath([]string{"count1", "terms", "size"}, size)
		aggs.SetPath([]string{"count1", "aggs", "count2", countType, "field"}, field)

	} else {

		aggs.SetPath([]string{"count1", countType, "field"}, field)
		if countType == "terms" {
			aggs.SetPath([]string{"count1", countType, "size"}, size)
		}

	}

	searchJson.SetPath([]string{"aggs"}, aggs)
	byt,_:=searchJson.EncodePretty()
	fmt.Println(string(byt))
	var uriVal = fmt.Sprintf("/%s/%s", index, action)
	//增加routing
	var retval elastigo.SearchResult

	body, err := DoCommand("POST", uriVal, nil, searchJson)
	//fmt.Println("searchJson_aggcount...............", string(body))
	if err != nil {
		return retval, err
	}

	// marshall into json
	jsonErr := json.Unmarshal([]byte(body), &retval)
	if jsonErr != nil {
		return retval, jsonErr
	}
	retval.RawJSON = body

	return retval, nil
}
func DoCommand(method string, url string, args map[string]interface{}, data interface{}) ([]byte, error) {
	if strings.Index(url, "?") == -1 {
		url = url + "?ignore_unavailable=true"
	} else {
		url = url + "&ignore_unavailable=true"
	}
	body, err := Es.DoCommand(method, url, args, data)
	return body, err
}

func BuildAndFilters(filters []*simplejson.Json) *simplejson.Json {
	ln := len(filters)
	if ln < 1 {
		return nil
	}
	if ln == 1 {
		return filters[0]
	}
	filterJson := simplejson.New()
	filterJson.SetPath([]string{"and"}, filters)
	return filterJson
}

func BuildTimeRangeFilter() *simplejson.Json {

	filterJson := simplejson.New()
	filterJson.SetPath([]string{"range", "__time", "gte"}, 0)
	endTime := 0
	if endTime == 0 {
		endTime = 1735689600000 // "2025-01-01"
	}
	filterJson.SetPath([]string{"range", "__time", "lte"}, endTime)

	return filterJson
}
func ConvertJsonToArray(countType string, jj *simplejson.Json) []*simplejson.Json {
	// if countType is: stats, min, max, avg, percentiles
	var result []*simplejson.Json
	if countType == "stats" {
		keys := []string{"sum", "min", "max", "avg", "count"}
		for _, key := range keys {
			j := simplejson.New()
			j.Set("key", key)
			j.Set("value", jj.Get(key).MustFloat64())
			result = append(result, j)
		}
		return result

	} else {
		j := simplejson.New()
		j.Set("key", countType)
		j.Set("value", jj.Get("value").MustFloat64())
		result = append(result, j)
		return result
	}
}
