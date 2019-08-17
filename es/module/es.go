// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package module

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"regexp"
	"time"

	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"

	js "github.com/bitly/go-simplejson"
	"github.com/mattbaird/elastigo/lib"
	"github.com/sas/utils"
)

var (
	StatsIncludeFields []string
	StatsExcludeFields []string
)

type EsApi struct {
	IndexPrefix string
	By          string
	Index       string
	Type        string
	StartTime   int64
	EndTime     int64
	TimeField   string
	Interval    string
	Keywords    string
	Operator    string
	SortField   string
	Order       string

	GeoField    string
	Lon         float64
	Lat         float64
	Radius      string
	NestedPath  string
	NestedQuery string

	JoinField       string
	JoinIndices     string
	JoinTypes       string
	JoinPath        string
	JoinQueryString string
	NoTimeFilter    bool

	Opts map[string]interface{}

	buf bytes.Buffer

	Routing string
}

var (
	Es           *elastigo.Conn
	MultiCluster string //多集群
	Logger       FmtLogger
)

func init() {
	EsInit([]string{"192.168.130.201"}, "9200")
}

func EsInit(hosts []string, port string) {
	Es = elastigo.NewConn()
	Es.Hosts = hosts
	Es.Port = port

}

func CreateRepo(repoType, name string, path string) error {
	if repoType == "" || name == "" || path == "" {
		return errors.New("parameter error")
	}

	esapi := &EsApi{}
	url := "/_snapshot/" + name
	request := js.New()
	request.Set("type", repoType)
	request.SetPath([]string{"settings", "location"}, path)
	result, err := esapi.DoCommand("PUT", url, nil, request)
	if err != nil {
		return err
	}
	if result.Get("acknowledged").MustBool() != true {
		return errors.New("resp error")
	}
	return nil
}

func (self *EsApi) AllIndices() string {
	prefix := self.IndexPrefix + self.Type + "*"
	return prefix
}

func (self *EsApi) CalculateIndex(t int64) string {
	prefix := self.IndexPrefix + self.Type
	format := "2006_01_02"
	if self.By == "month" {
		format = "2006_01"
	} else if self.By == "year" {
		format = "2006"
	} else if self.By == "all" {
		return prefix
	}
	str_time := time.Unix(int64(t/1000), 0).Format(format)
	return prefix + "_" + str_time
}

func (self *EsApi) CalculateIndices(startTime int64, endTime int64) string {
	if self.By == "all" {
		return self.AllIndices()
	}

	if startTime == 0 && endTime == 0 {
		return self.AllIndices()

	}

	t1 := time.Unix(int64(startTime/1000), 0)
	y1, m1, day1 := t1.Date()
	xt1 := time.Date(y1, m1, 1, 0, 0, 0, 0, time.Local)

	t2 := time.Unix(int64(endTime/1000), 0)
	y2, m2, day2 := t2.Date()
	xt2 := time.Date(y2, m2, 2, 0, 0, 0, 0, time.Local)

	if (y2-y1)*12*30+(int(m2)-int(m1))*30+day2-day1 > 12*30 {
		return self.AllIndices()
	}

	l := []string{}
	prefix := self.IndexPrefix + self.Type + "_"

	t := xt1

	if y1 == y2 && m1 == m2 && day1 == day2 {
		var s string
		if self.By == "month" {
			s = t1.Format("2006_01")
		} else if self.By == "year" {
			s = t1.Format("2006")
		} else {
			s = t1.Format("2006_01_02")
		}
		return prefix + s

	}

	format := "2006_01_*"
	if self.By == "month" {
		format = "2006_01"
	}
	for !t.After(xt2) {
		s := prefix + t.Format(format)
		l = append(l, s)
		t = t.AddDate(0, 1, 0)
	}
	return strings.Join(l, ",")
}

func (self *EsApi) Indexing(v interface{}) error {
	_, err := Es.Index(self.Index, self.Type, "", nil, v)
	if err != nil {
		fmt.Println("%#v", err)
		return err
	}
	return nil
}

func (self *EsApi) Indexing2(id string, args map[string]interface{}, v interface{}) error {
	_, err := Es.Index(self.Index, self.Type, id, args, v)
	if err != nil {
		fmt.Println("%#v", err)

		return err
	}
	return nil
}

func (self *EsApi) Get(id string) (*json.RawMessage, error) {
	res, err := Es.Get(self.Index, self.Type, id, nil)
	if err != nil {
		return nil, err
	}
	return res.Source, nil
}

func (self *EsApi) DeleteByQuery(indices, types, keywords, routing_field string) error {
	esapi := EsApi{}
	esapi.Index = indices
	esapi.Type = types
	esapi.Keywords = keywords
	esapi.NoTimeFilter = true
	total := 1
	per := 10000
	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		rs, err := esapi.Search("", "", 0, per)
		if err != nil {
			Logger.Errorf("err:%# v", err)
			return err
		}
		if rs.Hits.Total <= 0 {
			return nil
		}
		total = rs.Hits.Total/per + 1
		hits := rs.Hits.Hits
		n := int(math.Ceil(float64(len(hits)) / float64(100)))
		j := 0
		k := 0
		for i := 0; i < n; i++ {
			if j+100 > len(hits) {
				k = len(hits)
			} else {
				k = j + 100
			}
			wg.Add(1)
			go dodelete(routing_field, hits[j:k], &wg)
			j = j + 100
		}
		wg.Wait()
	}
	return nil

}

func dodelete(routing_field string, hits []elastigo.Hit, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, hit := range hits {
		b, err := hit.Source.MarshalJSON()
		if err != nil {
			Logger.Errorf("err:%v\n", err)
			return
		}
		j, _ := js.NewJson(b)
		Logger.Debugf("j:%+v", j)
		if routing_field == "" {
			_, err := Es.DoCommand("DELETE", "/"+hit.Index+"/"+hit.Type+"/"+hit.Id, nil, nil)
			if err != nil {
				Logger.Errorf("err:%# v", err)
			}
			continue
		}
		routing, ok := j.CheckGet(routing_field)
		if !ok {
			_, err := Es.DoCommand("DELETE", "/"+hit.Index+"/"+hit.Type+"/"+hit.Id, nil, nil)
			if err != nil {
				Logger.Errorf("err:%# v", err)
			}
			continue
		}
		v, err := routing.String()
		if err == nil {
			if v != "" {
				Logger.Debugf("v:%+v", v)
				_, err := Es.DoCommand("DELETE", "/"+hit.Index+"/"+hit.Type+"/"+hit.Id+"?routing="+v, nil, nil)
				if err != nil {
					Logger.Errorf("err:%# v", err)
				}
			} else {
				_, err := Es.DoCommand("DELETE", "/"+hit.Index+"/"+hit.Type+"/"+hit.Id, nil, nil)
				if err != nil {
					Logger.Errorf("err:%# v", err)
				}

			}
			continue
		}
		vs, err := routing.StringArray()
		if err != nil {
			Logger.Errorf("err:%# v", err)
			continue
		}
		for _, v := range vs {
			_, err := Es.DoCommand("DELETE", "/"+hit.Index+"/"+hit.Type+"/"+hit.Id+"?routing="+v, nil, nil)
			if err != nil {
				Logger.Errorf("err:%# v", err)
			}
		}
	}

}

func (self *EsApi) DeleteDoc(id string) (bool, error) {
	body, err := DoCommand("DELETE", "/"+self.Index+"/"+self.Type+"/"+id, nil, nil)
	if err != nil {
		fmt.Println("%#v", err)

		return false, err
	}
	result, err := js.NewJson(body)
	if err != nil {
		fmt.Println("%#v", err)

		return false, err
	}

	if result.Get("acknowledged").MustBool() == true {
		return true, nil
	}

	e := result.Get("error").MustString()

	if e == "" {
		return true, nil
	} else {
		return false, errors.New(e)
	}
}

func (self *EsApi) GeoPointField(field string, lon, lat float64) *js.Json {
	j := js.New()
	j.SetPath([]string{field, "type"}, "point")
	j.SetPath([]string{field, "coordinates"}, []float64{lon, lat})
	return j
}

func (self *EsApi) GeoFilter(field, radius string, lon, lat float64) *js.Json {
	j := js.New()
	j.SetPath([]string{"geo_shape", field, "shape", "type"}, "circle")
	j.SetPath([]string{"geo_shape", field, "shape", "coordinates"}, []float64{lon, lat})
	j.SetPath([]string{"geo_shape", field, "shape", "radius"}, radius)

	return j
}

func (self *EsApi) DeleteIndex(index string) (bool, error) {
	body, err := DoCommand("DELETE", "/"+index, nil, nil)
	if err != nil {
		fmt.Println("%#v", err)

		return false, err
	}
	result, err := js.NewJson(body)
	if err != nil {
		fmt.Println("%#v", err)

		return false, err
	}

	if result.Get("acknowledged").MustBool() == true {
		return true, nil
	}

	e := result.Get("error").MustString()
	if e == "IndexmissingException" {
		return true, nil
	} else {
		return false, errors.New(e)
	}
}

func (self *EsApi) GetStats(isall bool) (*js.Json, error) {
	type Res struct {
		Index string  `json:"index"`
		Count int64   `json:"count"`
		Size  float64 `json:"size"`
	}

	var result *js.Json
	body, err := DoCommand("GET", "/_stats", nil, nil)
	if err != nil {
		fmt.Println("%#v", err)

		return nil, err
	}
	result, err = js.NewJson(body)
	if err != nil {
		fmt.Println("%#v", err)

		return nil, err
	}

	res := js.New()
	var total_doc int64
	var total_size float64

	var res_indices []Res

	indices := result.Get("indices")
	m1, err := indices.Map()

	if err != nil {
		return nil, err
	}
	for k1, _ := range m1 {
		if !isall {
			reg := "^" + self.IndexPrefix
			match, _ := regexp.MatchString(reg, k1)
			if !match {
				continue
			}
		}
		pri := indices.GetPath(k1, "primaries")
		total := pri.Get("docs").Get("count").MustInt64()
		byte_size := pri.Get("store").Get("size_in_bytes").MustInt64()
		sf_size := fmt.Sprintf("%.2f", float64(byte_size)/1024/1024)
		f_size, _ := strconv.ParseFloat(sf_size, 64)
		res_indices = append(res_indices, Res{Index: k1, Count: total, Size: f_size})
		total_doc += total
		total_size += f_size
	}

	res.Set("total", total_doc)
	sf_total := fmt.Sprintf("%.2f", total_size)
	f_total, _ := strconv.ParseFloat(sf_total, 64)
	res.Set("size", f_total)
	res.Set("indices", res_indices)

	return res, nil
}

func (self *EsApi) GetMapping() (*js.Json, error) {
	var result *js.Json
	body, err := DoCommand("GET", fmt.Sprintf("/%s/%s/_mapping", self.Index, self.Type), nil, nil)
	if err != nil {
		fmt.Println("%#v", err)

		return nil, err
	}
	result, err = js.NewJson(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (self *EsApi) CalculateFields() ([]string, error) {
	j, err := self.GetMapping()
	if err != nil {
		return nil, err
	}
	m1, err := j.Map()
	if err != nil {
		return nil, err
	}
	fields := []string{}

	excludes := StatsExcludeFields
	for k1, _ := range m1 {
		mappings := j.GetPath(k1, "mappings")
		m2, _ := mappings.Map()
		for k2, _ := range m2 {
			properties := mappings.GetPath(k2, "properties")
			m3, _ := properties.Map()
			for k3, _ := range m3 {
				field := properties.Get(k3)
				if field.Get("type").MustString() == "string" && field.Get("index").MustString() != "not_analyzed" {
					excludes = append(excludes, k3)
				} else {
					fields = append(fields, k3)
				}
			}

		}
	}
	fields2 := UniqStr(fields)
	fields3 := []string{}
	for _, v := range fields2 {
		if IsMatch(v, excludes) {
			continue
		} else {
			fields3 = append(fields3, v)
		}
	}

	return fields3, nil
}

func (self *EsApi) BuildTimeRangeFilter() *js.Json {

	filterJson := js.New()
	filterJson.SetPath([]string{"range", self.TimeField, "gte"}, self.StartTime)
	endTime := self.EndTime
	if endTime == 0 {
		endTime = 1735689600000 // "2025-01-01"
	}
	filterJson.SetPath([]string{"range", self.TimeField, "lte"}, endTime)

	return filterJson
}

func (self *EsApi) BuildGeoBoundingBoxFilter(field string, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64) *js.Json {
	filterJson := js.New()
	filterJson.SetPath([]string{"geo_bounding_box", field, "top_left", "lon"}, topLeftLon)
	filterJson.SetPath([]string{"geo_bounding_box", field, "top_left", "lat"}, topLeftLat)
	filterJson.SetPath([]string{"geo_bounding_box", field, "bottom_right", "lon"}, bottomRightLon)
	filterJson.SetPath([]string{"geo_bounding_box", field, "bottom_right", "lat"}, bottomRightLat)

	return filterJson
}

func (self *EsApi) BuildGeoEnvelopeFilter(field string, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64) *js.Json {
	filterJson := js.New()
	filterJson.SetPath([]string{"geo_shape", field, "shape", "type"}, "envelope")
	filterJson.SetPath([]string{"geo_shape", field, "shape", "coordinates"}, [][]float64{{topLeftLon, topLeftLat}, {bottomRightLon, bottomRightLat}})
	return filterJson
}

func (self *EsApi) BuildGeoContainFilter(field string, lon, lat float64) *js.Json {
	filterJson := js.New()
	filterJson.SetPath([]string{"geo_shape", field, "shape", "type"}, "point")
	filterJson.SetPath([]string{"geo_shape", field, "shape", "coordinates"}, []float64{lon, lat})
	filterJson.SetPath([]string{"geo_shape", field, "relation"}, "contains")

	return filterJson
}

func (self *EsApi) BuildGeoDistanceFilter(field, distance string, lon, lat float64) *js.Json {
	filterJson := js.New()
	filterJson.SetPath([]string{"geo_distance", "distance"}, distance)
	filterJson.SetPath([]string{"geo_distance", field, "lon"}, lon)
	filterJson.SetPath([]string{"geo_distance", field, "lat"}, lat)

	return filterJson
}

func (self *EsApi) BuildGeoDistanceRangeFilter(field, from, to string, lon, lat float64) *js.Json {
	filterJson := js.New()
	filterJson.SetPath([]string{"geo_distance_range", "gte"}, from)
	filterJson.SetPath([]string{"geo_distance_range", "lt"}, to)
	filterJson.SetPath([]string{"geo_distance_range", field, "lon"}, lon)
	filterJson.SetPath([]string{"geo_distance_range", field, "lat"}, lat)

	return filterJson
}

func (self *EsApi) BuildAndFilters(filters []*js.Json) *js.Json {
	ln := len(filters)
	if ln < 1 {
		return nil
	}
	if ln == 1 {
		return filters[0]
	}
	filterJson := js.New()
	filterJson.SetPath([]string{"and"}, filters)
	return filterJson
}

func (self *EsApi) BuildSort() *js.Json {
	sortJson := js.New()
	j1 := js.New()

	if self.SortField == "" {
		self.SortField = self.TimeField
	}
	if self.Order == "" {
		j1.SetPath([]string{self.SortField, "order"}, "desc")
	} else {
		j1.SetPath([]string{self.SortField, "order"}, self.Order)
	}
	sortJson.SetPath([]string{}, j1)
	return sortJson
}

func (self *EsApi) BuildDistanceSort(field string, lon, lat float64) *js.Json {
	sortJson := js.New()
	j1 := js.New()
	j1.SetPath([]string{"_geo_distance", field, "lat"}, lat)
	j1.SetPath([]string{"_geo_distance", field, "lon"}, lon)
	j1.SetPath([]string{"_geo_distance", "order"}, "asc")
	j1.SetPath([]string{"_geo_distance", "unit"}, "km")
	sortJson.SetPath([]string{}, j1)
	return sortJson
}

func (self *EsApi) BuildJoinFilter(field, to_indices, to_types, path, keywords string) *js.Json {
	var j = js.New()
	j.SetPath([]string{"filterjoin", field, "indices"}, strings.Split(to_indices, ","))
	j.SetPath([]string{"filterjoin", field, "types"}, strings.Split(to_types, ","))
	j.SetPath([]string{"filterjoin", field, "path"}, path)

	//j.SetPath([]string{"filterjoin", field, "query", "terms", field}, keywords)

	if keywords == "" {
		j.SetPath([]string{"filterjoin", field, "query"}, "match_all")
	} else {
		j.SetPath([]string{"filterjoin", field, "query", "query_string", "query"}, keywords)
		j.SetPath([]string{"filterjoin", field, "query", "query_string", "auto_generate_phrase_queries"}, true)
		j.SetPath([]string{"filterjoin", field, "query", "query_string", "analyze_wildcard"}, true)
	}

	return j
}

func (self *EsApi) BuildStringQuery(keywords string) *js.Json {
	var j = js.New()
	if keywords == "" {
		j.SetPath([]string{"match_all"}, js.New())
	} else {
		j.SetPath([]string{"query_string", "query"}, keywords)
		j.SetPath([]string{"query_string", "auto_generate_phrase_queries"}, true)
		j.SetPath([]string{"query_string", "analyze_wildcard"}, true)
	}
	return j
}

func (self *EsApi) BuildQuery(operator string) *js.Json {
	if operator == "" {
		if self.Operator == "" {
			self.Operator = "AND"
		}
	} else {
		self.Operator = operator
	}
	var filter *js.Json
	if !self.NoTimeFilter {

		filter = self.BuildTimeRangeFilter()
	}
	fileteredQuery := js.New()

	if self.Keywords == "" {
		fileteredQuery.SetPath([]string{"filtered", "query", "match_all"}, js.New())
	} else {
		fileteredQuery.SetPath([]string{"filtered", "query", "query_string", "query"}, self.Keywords)
		fileteredQuery.SetPath([]string{"filtered", "query", "query_string", "default_operator"}, self.Operator)
		fileteredQuery.SetPath([]string{"filtered", "query", "query_string", "auto_generate_phrase_queries"}, true)
		fileteredQuery.SetPath([]string{"filtered", "query", "query_string", "analyze_wildcard"}, true)

	}

	if self.Radius != "" && filter != nil {
		geoFilter := self.GeoFilter(self.GeoField, self.Radius, self.Lon, self.Lat)
		filters := []interface{}{filter, geoFilter}
		fileteredQuery.SetPath([]string{"filtered", "filter", "and"}, filters)

	} else if filter != nil {
		fileteredQuery.SetPath([]string{"filtered", "filter"}, filter)
	} else {
		geoFilter := self.GeoFilter(self.GeoField, self.Radius, self.Lon, self.Lat)
		fileteredQuery.SetPath([]string{"filtered", "filter"}, geoFilter)

	}

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))
	if self.NestedPath != "" && self.NestedQuery != "" {
		nestedQuery := js.New()
		nestedQuery.SetPath([]string{"nested", "path"}, self.NestedPath)
		nestedQuery.SetPath([]string{"nested", "query", "query_string", "query"}, self.NestedQuery)
		var mustQueryList []*js.Json
		mustQueryList = append(mustQueryList, fileteredQuery, nestedQuery)
		boolQuery := js.New()
		boolQuery.SetPath([]string{"bool", "filter"}, mustQueryList)
		return boolQuery
	} else {
		return fileteredQuery
	}

}

func (self *EsApi) JoinSearch(from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
		sort := self.BuildSort()
		searchJson.Set("sort", sort)

	}
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	if self.Keywords == "" {
		searchJson.SetPath([]string{"query", "filtered", "query", "match_all"}, js.New())
	} else {
		searchJson.SetPath([]string{"query", "filtered", "query", "query_string", "query"}, self.Keywords)
		searchJson.SetPath([]string{"query", "filtered", "query", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "filtered", "query", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "filtered", "query", "query_string", "analyze_wildcard"}, true)
	}

	filters := []*js.Json{timeFilter}

	if self.JoinIndices != "" {
		joinFilter := self.BuildJoinFilter(self.JoinField, self.JoinIndices, self.JoinTypes, self.JoinPath, self.JoinQueryString)
		filters = append(filters, joinFilter)
	}
	andFilter := self.BuildAndFilters(filters)
	searchJson.SetPath([]string{"query", "filtered", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))
	var uriVal string
	if len(self.Type) > 0 && self.Type != "*" {
		uriVal = fmt.Sprintf("/%s/%s/_coordinate_search", self.Index, self.Type)
	} else {
		uriVal = fmt.Sprintf("/%s/_coordinate_search", self.Index)
	}
	var retval elastigo.SearchResult

	body, err := DoCommand("POST", uriVal, nil, searchJson)
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

func (self *EsApi) ScrollSearch(scroll, scroll_id string) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	searchJson.Set("sort", []string{"_doc"})
	searchJson.Set("scroll", scroll)
	searchJson.Set("scroll_id", scroll_id)
	action := "_search/scroll"

	var uriVal string
	uriVal = fmt.Sprintf("/%s", action)
	var retval elastigo.SearchResult

	body, err := DoCommand("POST", uriVal, self.Opts, searchJson)
	if err != nil {
		return retval, err
	}

	jsonErr := json.Unmarshal([]byte(body), &retval)
	if jsonErr != nil {
		return retval, jsonErr
	}
	retval.RawJSON = body

	return retval, nil
}

func (self *EsApi) Search(search_type string, operator string, from int, size int) (elastigo.SearchResult, error) {
	if search_type == "" {
		search_type = "query_then_fetch"
	}

	if operator == "" {
		if self.Operator == "" {
			self.Operator = "AND"
		}
	} else {
		self.Operator = operator
	}
	var searchJson = js.New()

	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	action := "_search"

	var timeFilter *js.Json
	var filters []*js.Json
	if !self.NoTimeFilter {
		timeFilter = self.BuildTimeRangeFilter()
		sort := self.BuildSort()
		searchJson.Set("sort", sort)
	} else {
		if self.SortField != "" {
			sort := self.BuildSort()
			searchJson.Set("sort", sort)
		}
	}
	if timeFilter != nil {
		filters = []*js.Json{timeFilter}
	}

	if self.Radius != "" && timeFilter != nil {
		geoFilter := self.GeoFilter(self.GeoField, self.Radius, self.Lon, self.Lat)
		filters = append(filters, geoFilter)
	}
	if self.JoinIndices != "" {
		action = "_coordinate_search"
		joinFilter := self.BuildJoinFilter(self.JoinField, self.JoinIndices, self.JoinTypes, self.JoinPath, self.JoinQueryString)
		filters = append(filters, joinFilter)
	}
	Logger.Debugf("self:%+v", self)
	Logger.Debugf("timeFilter:%+v", timeFilter)
	Logger.Debugf("filters:%+v", filters)

	if filters != nil {
		q := js.New()
		if self.Keywords == "" {
			q.SetPath([]string{"match_all"}, js.New())
		} else {
			q.SetPath([]string{"query_string", "query"}, self.Keywords)
			q.SetPath([]string{"query_string", "default_operator"}, self.Operator)
			q.SetPath([]string{"query_string", "auto_generate_phrase_queries"}, true)
			q.SetPath([]string{"query_string", "analyze_wildcard"}, true)

		}
		filters = append(filters, q)
		searchJson.SetPath([]string{"query", "bool", "filter"}, filters)

	} else {
		if self.Keywords == "" {
			searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
		} else {
			searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, self.Keywords)
			searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
			searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
			searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

		}

	}
	//spew.Dump(searchJson)
	var uriVal string
	if len(self.Type) > 0 && self.Type != "*" {
		uriVal = fmt.Sprintf("/%s/%s/%s", self.Index, self.Type, action)
	} else {
		uriVal = fmt.Sprintf("/%s/%s", self.Index, action)
	}
	//增加routing
	if self.Routing != "" {
		if strings.Index(uriVal, "?") == -1 {
			uriVal = uriVal + "?routing=" + self.Routing
		} else {
			uriVal = uriVal + "&routing=" + self.Routing
		}
	}
	var retval elastigo.SearchResult

	Logger.Debugf("searchJson:%+v", searchJson)
	body, err := DoCommand("POST", uriVal, self.Opts, searchJson)
	//fmt.Println(string(body))
	if err != nil {
		Logger.Errorf("err:%+v", err)
		return retval, err
	}
	j, err := js.NewJson(body)
	if err != nil {
		return retval, err
	}
	j.Del("_shards")
	body, err = j.MarshalJSON()
	if err != nil {
		return retval, err
	}
	// marshall into json
	jsonErr := json.Unmarshal([]byte(body), &retval)
	if jsonErr != nil {
		return retval, jsonErr
	}
	retval.RawJSON = body

	// out, err := Es.Search(self.Index, self.Type, map[string]interface{}{"search_type": search_type}, searchJson)
	// if err != nil && err != elastigo.RecordNotFound {
	// 	fmt.Println("%#v", err)

	// 	return elastigo.SearchResult{}, err
	// }
	return retval, nil
}

func (self *EsApi) NestedSearch(search_type string, operator string, from int, size int) (elastigo.SearchResult, error) {
	if search_type == "" {
		search_type = "query_then_fetch"
	}
	var searchJson = js.New()
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}
	sort := self.BuildSort()
	searchJson.Set("sort", sort)
	query := self.BuildQuery(operator)
	searchJson.Set("query", query)

	//spew.Dump(searchJson)
	out, err := Es.Search(self.Index, self.Type, map[string]interface{}{"search_type": search_type}, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) GeoSearch(keywords, field, distance string, lon, lat float64, from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoDistanceFilter := self.BuildGeoDistanceFilter(field, distance, lon, lat)
	andFilter := self.BuildAndFilters([]*js.Json{timeFilter, geoDistanceFilter})
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	sort := self.BuildSort()
	searchJson.Set("sort", sort)

	if keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
	} else {
		if self.Operator == "" {
			self.Operator = "AND"
		}
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	searchJson.SetPath([]string{"query", "bool", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) DistanceRangeSearch(keywords, field, gte, lt string, lon, lat float64, from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoDistanceFilter := self.BuildGeoDistanceRangeFilter(field, gte, lt, lon, lat)
	andFilter := self.BuildAndFilters([]*js.Json{timeFilter, geoDistanceFilter})
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	//sort1 := self.BuildSort()
	sort2 := self.BuildDistanceSort(field, lon, lat)
	sort := []*js.Json{sort2}
	searchJson.Set("sort", sort)

	if keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
	} else {
		if self.Operator == "" {
			self.Operator = "AND"
		}
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	searchJson.SetPath([]string{"query", "bol", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) GeoBoundingBoxSearch(keywords, field string, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64, from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoFilter := self.BuildGeoBoundingBoxFilter(field, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat)
	andFilter := self.BuildAndFilters([]*js.Json{timeFilter, geoFilter})
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	sort := self.BuildSort()
	searchJson.Set("sort", sort)

	if keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
	} else {
		if self.Operator == "" {
			self.Operator = "AND"
		}
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	searchJson.SetPath([]string{"query", "bool", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) GeoEnvelopeSearch(keywords, field string, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64, from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoFilter := self.BuildGeoEnvelopeFilter(field, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat)
	andFilter := self.BuildAndFilters([]*js.Json{timeFilter, geoFilter})
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	sort := self.BuildSort()
	searchJson.Set("sort", sort)

	if keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
	} else {
		if self.Operator == "" {
			self.Operator = "AND"
		}
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	searchJson.SetPath([]string{"query", "bool", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) GeoContainSearch(keywords, field string, lon, lat float64, from int, size int) (elastigo.SearchResult, error) {
	var searchJson = js.New()
	var timeFilter *js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoFilter := self.BuildGeoContainFilter(field, lon, lat)
	andFilter := self.BuildAndFilters([]*js.Json{timeFilter, geoFilter})
	if from != 0 {
		searchJson.Set("from", from)
	}
	if size != 0 {
		if size > 10000 {
			size = 10000
		}
		searchJson.Set("size", size)
	}

	sort := self.BuildSort()
	searchJson.Set("sort", sort)

	if keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())
	} else {
		if self.Operator == "" {
			self.Operator = "AND"
		}
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	searchJson.SetPath([]string{"query", "bool", "filter"}, andFilter)

	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) NestedAggCount(groupby, countType, field string, size int) (elastigo.SearchResult, error) {
	if size == 0 {
		size = 1000
	} else if size > 1000 {
		size = 1000
	}
	//fmt.Printf("self:%v\n", pretty.Formatter(self))

	// countType is: stats, min, max, avg, percentiles
	if countType == "" {
		countType = "terms"
	}

	if self.Operator == "" {
		self.Operator = "AND"
	}

	query := self.BuildQuery("")

	aggs := js.New()
	if groupby != "" {
		aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "field"}, groupby)
		aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "size"}, size)
		aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", countType, "field"}, field)

	} else {
		aggs.SetPath([]string{"count1", "aggs", "count2", countType, "field"}, field)
		if countType == "terms" {
			aggs.SetPath([]string{"count1", "aggs", "count2", countType, "size"}, size)
		}

	}

	if self.NestedPath != "" && strings.HasPrefix(groupby, self.NestedPath+".") || strings.HasPrefix(field, self.NestedPath+".") {
		aggs.SetPath([]string{"count1", "nested", "path"}, self.NestedPath)
	} else if !self.NoTimeFilter {
		filter := self.BuildTimeRangeFilter()
		aggs.SetPath([]string{"count1", "filter"}, filter)

	}

	searchJson := js.New()
	searchJson.Set("query", query)
	searchJson.Set("aggs", aggs)

	// bs, _ := searchJson.EncodePretty()
	// spew.Dump(string(bs))
	out, err := Es.Search(self.Index, self.Type, nil, searchJson)
	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}
	return out, nil
}

func (self *EsApi) AggCount(groupby, countType, field string, size int) (elastigo.SearchResult, error) {
	//fmt.Printf("self:%v\n", pretty.Formatter(self))

	// countType is: stats, min, max, avg, percentiles
	if countType == "" {
		countType = "terms"
	}

	if self.Operator == "" {
		self.Operator = "AND"
	}

	if size == 0 {
		size = 1000
	} else if size > 1000 {
		size = 1000
	}

	action := "_search"
	var timeFilter *js.Json
	var filters []*js.Json
	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
		filters = []*js.Json{timeFilter}
	}
	if self.JoinIndices != "" {
		action = "_coordinate_search"
		joinFilter := self.BuildJoinFilter(self.JoinField, self.JoinIndices, self.JoinTypes, self.JoinPath, self.JoinQueryString)
		filters = append(filters, joinFilter)
	}
	filter := self.BuildAndFilters(filters)

	searchJson := js.New()
	if self.Keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())

	} else {
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, self.Keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)
	}

	aggs := js.New()

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
	//pretty.Println("searchJson_aggcount:", searchJson)
	var uriVal string
	if len(self.Type) > 0 && self.Type != "*" {
		uriVal = fmt.Sprintf("/%s/%s/%s", self.Index, self.Type, action)
	} else {
		uriVal = fmt.Sprintf("/%s/%s", self.Index, action)
	}
	//增加routing
	if self.Routing != "" {
		if strings.Index(uriVal, "?") == -1 {
			uriVal = uriVal + "?routing=" + self.Routing
		} else {
			uriVal = uriVal + "&routing=" + self.Routing
		}
	}
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

func (self *EsApi) AggDateHist() (elastigo.SearchResult, error) {
	if self.Operator == "" {
		self.Operator = "AND"
	}

	action := "_search"
	var timeFilter *js.Json
	var filters []*js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
		filters = []*js.Json{timeFilter}

	}
	if self.JoinIndices != "" {
		action = "_coordinate_search"
		joinFilter := self.BuildJoinFilter(self.JoinField, self.JoinIndices, self.JoinTypes, self.JoinPath, self.JoinQueryString)
		filters = append(filters, joinFilter)
	}
	filter := self.BuildAndFilters(filters)

	searchJson := js.New()
	if self.Keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())

	} else {
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, self.Keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	aggs := js.New()

	if filter != nil {
		searchJson.SetPath([]string{"query", "bool", "must"}, filter)
		aggs.SetPath([]string{"count1", "filter"}, filter)

	}

	datehistAgg, _ := js.NewJson([]byte(fmt.Sprintf(`{"field":"%s", "interval":"%s","offset":"-8h","min_doc_count":1}`, self.TimeField, self.Interval)))

	aggs.SetPath([]string{"count1", "aggs", "count2", "date_histogram"}, datehistAgg)

	searchJson.SetPath([]string{"aggs"}, aggs)
	//pretty.Println("searchJson_AggDateHist:", searchJson)
	var uriVal string
	if len(self.Type) > 0 && self.Type != "*" {
		uriVal = fmt.Sprintf("/%s/%s/%s", self.Index, self.Type, action)
	} else {
		uriVal = fmt.Sprintf("/%s/%s", self.Index, action)
	}
	var retval elastigo.SearchResult
	body, err := DoCommand("POST", uriVal, nil, searchJson)
	//fmt.Println(".........searchJson_AggDateHist", string(body))
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

func (self *EsApi) AggDateHistStat(groupby, countType, field string, size int) (elastigo.SearchResult, error) {
	if countType == "" {
		countType = "terms"
	}
	if self.Operator == "" {
		self.Operator = "AND"
	}

	if size == 0 {
		size = 1000
	} else if size > 1000 {
		size = 1000
	}

	action := "_search"
	var timeFilter *js.Json
	var filters []*js.Json

	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
		filters = []*js.Json{timeFilter}

	}
	if self.JoinIndices != "" {
		action = "_coordinate_search"
		joinFilter := self.BuildJoinFilter(self.JoinField, self.JoinIndices, self.JoinTypes, self.JoinPath, self.JoinQueryString)
		filters = append(filters, joinFilter)
	}
	filter := self.BuildAndFilters(filters)

	searchJson := js.New()
	if self.Keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())

	} else {
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, self.Keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)

	}

	aggs := js.New()

	if filter != nil {
		searchJson.SetPath([]string{"query", "bool", "must"}, filter)
		aggs.SetPath([]string{"count1", "filter"}, filter)

	}
	datehistAgg, _ := js.NewJson([]byte(fmt.Sprintf(`{"field":"%s", "interval":"%s","offset":"-8h","min_doc_count":1}`, self.TimeField, self.Interval)))

	if groupby != "" {
		// first groupby, then datehist
		// aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "field"}, groupby)
		// aggs.SetPath([]string{"count1", "aggs", "count2", "terms", "size"}, size)

		// aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "date_histogram"}, datehistAgg)
		// aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "aggs", "count4", countType, "field"}, field)

		// first datehist, then groupby
		aggs.SetPath([]string{"count1", "aggs", "count2", "date_histogram"}, datehistAgg)
		aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "terms", "field"}, groupby)
		aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "terms", "size"}, size)
		aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "aggs", "count4", countType, "field"}, field)

		if countType == "terms" {
			aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", "aggs", "count4", countType, "size"}, size)
		}

	} else {
		aggs.SetPath([]string{"count1", "aggs", "count2", "date_histogram"}, datehistAgg)
		aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", countType, "field"}, field)
		if countType == "terms" {
			aggs.SetPath([]string{"count1", "aggs", "count2", "aggs", "count3", countType, "size"}, size)
		}

	}

	searchJson.SetPath([]string{"aggs"}, aggs)
	//pretty.Println("searchJson_AggDateHistStat:", searchJson)
	var uriVal string
	if len(self.Type) > 0 && self.Type != "*" {
		uriVal = fmt.Sprintf("/%s/%s/%s", self.Index, self.Type, action)
	} else {
		uriVal = fmt.Sprintf("/%s/%s", self.Index, action)
	}
	var retval elastigo.SearchResult

	body, err := DoCommand("POST", uriVal, nil, searchJson)
	//fmt.Println("..........", string(body))
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

func (self *EsApi) AggGeoHashGrid(field string, precision int, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat float64) (elastigo.SearchResult, error) {
	if self.Operator == "" {
		self.Operator = "AND"
	}
	var timeFilter *js.Json
	if !self.NoTimeFilter {

		timeFilter = self.BuildTimeRangeFilter()
	}
	geoBoundingBoxFilter := self.BuildGeoBoundingBoxFilter(field, topLeftLon, topLeftLat, bottomRightLon, bottomRightLat)

	searchJson := js.New()
	if self.Keywords == "" {
		searchJson.SetPath([]string{"query", "bool", "filter", "match_all"}, js.New())

	} else {
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "query"}, self.Keywords)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "default_operator"}, self.Operator)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "auto_generate_phrase_queries"}, true)
		searchJson.SetPath([]string{"query", "bool", "filter", "query_string", "analyze_wildcard"}, true)
	}
	if timeFilter != nil {
		searchJson.SetPath([]string{"query", "bool", "must"}, timeFilter)
	}

	aggs := js.New()
	aggs.SetPath([]string{"count1", "filter"}, geoBoundingBoxFilter)
	aggs.SetPath([]string{"count1", "aggs", "count2", "geohash_grid", "field"}, field)
	aggs.SetPath([]string{"count1", "aggs", "count2", "geohash_grid", "precision"}, precision)

	searchJson.SetPath([]string{"aggs"}, aggs)
	//fmt.Printf("searchJson:%v\n", pretty.Formatter(searchJson))

	//out, err := Es.Search(self.Index, self.Type, map[string]interface{}{"search_type": "count"}, searchJson)
	out, err := Es.Search(self.Index, self.Type, nil, searchJson)

	if err != nil && err != elastigo.RecordNotFound {
		fmt.Println("%#v", err)

		return elastigo.SearchResult{}, err
	}

	return out, nil
}

func (self *EsApi) HandleIndex() {
	keywords := self.Keywords
	newKeywords := self.Keywords
	index := self.Index
	if keywords != "" {
		var reg = regexp.MustCompile(`__index:(\w+)`)
		r := reg.FindStringSubmatch(keywords)
		if r != nil {
			index = r[1]
			var reg2 = regexp.MustCompile(`__index:\w+\s*[AND|OR]*\s*(.*)`)
			newKeywords = reg2.ReplaceAllString(keywords, "$1")
		} else {
			index = self.CalculateIndices(self.StartTime, self.EndTime)
		}
	} else {
		index = self.CalculateIndices(self.StartTime, self.EndTime)

	}
	self.Index = index
	self.Keywords = newKeywords
	return

}

func (self *EsApi) DoCommand(method, url string, args map[string]interface{}, data interface{}) (*js.Json, error) {
	var result *js.Json
	body, err := DoCommand(method, url, args, data)
	if err != nil {
		fmt.Println("%#v", err)
		return nil, err
	}
	result, err = js.NewJson(body)
	if err != nil {
		fmt.Println("%#v", err)
		j := js.New()
		j.Set("raw", string(body))
		return j, nil
	}
	return result, nil
}

func (self *EsApi) OperateIndices(indices, action string) error {
	var url string
	if action == "close" {
		url = "/" + indices + "/_close"
	} else if action == "open" {
		url = "/" + indices + "/_open"
	} else {
		return errors.New("action only could be close or open")
	}
	result, err := self.DoCommand("POST", url, nil, nil)
	if err != nil {
		return err
	}
	if result.Get("acknowledged").MustBool() != true {
		return errors.New("acknowledge is not true")
	}
	return nil
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

func (c *EsApi) WriteBulkBytes(op string, index string, _type string, id, parent, ttl string, date *time.Time, data interface{}) error {
	// only index and update are currently supported
	if op != "index" && op != "update" {
		return errors.New(fmt.Sprintf("Operation '%s' is not yet supported", op))
	}

	// // First line
	c.buf.WriteString(fmt.Sprintf(`{"%s":{"_index":"`, op))
	c.buf.WriteString(index)
	c.buf.WriteString(`","_type":"`)
	c.buf.WriteString(_type)
	c.buf.WriteString(`"`)
	if len(id) > 0 {
		c.buf.WriteString(`,"_id":"`)
		c.buf.WriteString(id)
		c.buf.WriteString(`"`)
	}

	if len(parent) > 0 {
		c.buf.WriteString(`,"_parent":"`)
		c.buf.WriteString(parent)
		c.buf.WriteString(`"`)
	}

	if op == "update" {
		c.buf.WriteString(`,"_retry_on_conflict":3`)
	}

	if len(ttl) > 0 {
		c.buf.WriteString(`,"ttl":"`)
		c.buf.WriteString(ttl)
		c.buf.WriteString(`"`)
	}
	if date != nil {
		c.buf.WriteString(`,"_timestamp":"`)
		c.buf.WriteString(strconv.FormatInt(date.UnixNano()/1e6, 10))
		c.buf.WriteString(`"`)
	}

	c.buf.WriteString(`}}`)
	c.buf.WriteRune('\n')
	//buf.WriteByte('\n')
	switch v := data.(type) {
	case *bytes.Buffer:
		io.Copy(&c.buf, v)
	case []byte:
		c.buf.Write(v)
	case string:
		c.buf.WriteString(v)
	default:
		body, jsonErr := json.Marshal(data)
		if jsonErr != nil {
			return jsonErr
		}
		c.buf.Write(body)
	}
	c.buf.WriteRune('\n')
	return nil
}

func (c *EsApi) Send(refresh string) error {
	type responseStruct struct {
		Took   int64                    `json:"took"`
		Errors bool                     `json:"errors"`
		Items  []map[string]interface{} `json:"items"`
	}

	response := responseStruct{}

	para := make(map[string]interface{})

	para["refresh"] = refresh

	body, err := DoCommand("POST", fmt.Sprintf("/_bulk"), para, c.buf.Bytes())

	if err != nil {
		return err
	}
	// check for response errors, bulk insert will give 200 OK but then include errors in response
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr == nil {
		if response.Errors {
			return errors.New("insert error")
		}
	}
	return nil
}

func HandleIndex(indexPrefix, _type, by, keywords string, startTime, endTime int64) (index string, newKeywords string) {
	esapi := EsApi{Type: _type, IndexPrefix: indexPrefix, By: by}
	newKeywords = keywords
	if keywords != "" {
		var reg = regexp.MustCompile(`__index:(\w+)`)
		r := reg.FindStringSubmatch(keywords)
		if r != nil {
			index = r[1]
			var reg2 = regexp.MustCompile(`__index:\w+\s*[AND|OR]*\s*(.*)`)
			newKeywords = reg2.ReplaceAllString(keywords, "$1")
		} else {
			index = esapi.CalculateIndices(startTime, endTime)
		}
	} else {
		index = esapi.CalculateIndices(startTime, endTime)

	}

	return index, newKeywords

}

// UniqStr returns a copy if the passed slice with only unique string results.
func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))
	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func IsMatch(field string, coll []string) bool {
	for _, v := range coll {
		if v == "" {
			return false
		}
		match, _ := regexp.MatchString(v, field)
		if match {
			return true
		}

	}
	return false
}

func CacTimePeriod(startTime, endTime int64) int {
	t1 := time.Unix(int64(startTime/1000), 0)
	y1, m1, day1 := t1.Date()

	t2 := time.Unix(int64(endTime/1000), 0)
	y2, m2, day2 := t2.Date()

	return (y2-y1)*12*30 + (int(m2)-int(m1))*30 + day2 - day1
}

//根据DebugShowlog，打印日志
func DoCommand(method string, url string, args map[string]interface{}, data interface{}) ([]byte, error) {
	if strings.Index(url, "?") == -1 {
		url = url + "?ignore_unavailable=true"
	} else {
		url = url + "&ignore_unavailable=true"
	}
	//多集群
	if MultiCluster != "" {
		urlarr := strings.SplitN(url, "/", 3)
		if len(urlarr) == 3 {
			//index没有冒号
			if strings.Index(url, ":") == -1 {
				clu_idx := strings.Replace(MultiCluster, ",", fmt.Sprintf(":%s,", urlarr[1]), -1)
				//最后还有一个没有加index，再加个本地查询
				urlarr[1] = fmt.Sprintf("%s:%s,%s", clu_idx, urlarr[1], urlarr[1])
				url = strings.Join(urlarr, "/")
			}
		}
	}
	//if DebugShowlog {
	var request string
	switch v := data.(type) {
	case string:
		request = v
	case []byte:
		request = string(v)
	default:
		byt, err := json.Marshal(v)
		if err == nil {
			request = string(byt)
		}
	}
	Logger.Infof(`>>>>>>No=[%s],status=[request],url=[%s]`, utils.GoroutineID(), url)
	Logger.Debugf(`>>>>>>No=[%s],status=[request],data=%s`, utils.GoroutineID(), request)
	//}
	body, err := Es.DoCommand(method, url, args, data)
	//if DebugShowlog {
	var response string
	if len(body) > 10000 {
		response = string(body[:10000]) + "......"
	} else {
		response = string(body)
	}
	Logger.Infof(`>>>>>>No=[%s],status=[response]`, utils.GoroutineID())
	Logger.Debugf(`>>>>>>No=[%s],status=[response],data=%s`, utils.GoroutineID(), response)
	Logger.Flush()
	//}
	return body, err
}
