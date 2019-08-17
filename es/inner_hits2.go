package main

import (
	"encoding/json"
	"github.com/mattbaird/elastigo/lib"
	"github.com/bitly/go-simplejson"
	"fmt"
)

type SearchResult struct {
	RawJSON      []byte
	Took         int                  `json:"took"`
	TimedOut     bool                 `json:"timed_out"`
	ShardStatus  elastigo.Status      `json:"_shards"`
	Hits         Hits                 `json:"hits"`
	Facets       json.RawMessage      `json:"facets,omitempty"` // structure varies on query
	ScrollId     string               `json:"_scroll_id,omitempty"`
	Aggregations json.RawMessage      `json:"aggregations,omitempty"` // structure varies on query
	Suggestions  elastigo.Suggestions `json:"suggest,omitempty"`
}

type Hits struct {
	Total int `json:"total"`
	//	MaxScore float32 `json:"max_score"`
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Index       string                   `json:"_index"`
	Type        string                   `json:"_type,omitempty"`
	Id          string                   `json:"_id"`
	Score       elastigo.Float32Nullable `json:"_score,omitempty"` // Filters (no query) dont have score, so is null
	Source      *json.RawMessage         `json:"_source"`          // marshalling left to consumer
	TTL         *elastigo.TTL            `json:"_ttl,omitempty"`
	Fields      *json.RawMessage         `json:"fields"` // when a field arg is passed to ES, instead of _source it returns fields
	Explanation *elastigo.Explanation    `json:"_explanation,omitempty"`
	Highlight   *elastigo.Highlight      `json:"highlight,omitempty"`
	Sort        []interface{}            `json:"sort,omitempty"`
	InnerHits   string                   `json:"inner_hits"`
}

var datas = `{
	"_shards": {
		"failed": 0,
		"successful": 5,
		"total": 5
	},
	"hits": {
		"hits": [{
			"_id": "AVoESHYF_OA-dG63Txsd",
			"_index": "recipes",
			"_score": null,
			"_source": {
				"name": "鲫鱼汤（变态辣）",
				"rating": 5,
				"type": "湘菜"
			},
			"_type": "type",
			"fields": {
				"type": ["湘菜"]
			},
			"inner_hits": {
				"hits": {
					"hits": [{
						"_id": "AVoESHYF_OA-dG63Txsd",
						"_index": "recipes",
						"_score": null,
						"_source": {
							"name": "鲫鱼汤（变态辣）",
							"rating": 5,
							"type": "湘菜"
						},
						"_type": "type",
						"sort": [5]
					}, {
						"_id": "AVoESHX7_OA-dG63Txsc",
						"_index": "recipes",
						"_score": null,
						"_source": {
							"name": "鲫鱼汤（微辣）",
							"rating": 4,
							"type": "湘菜"
						},
						"_type": "type",
						"sort": [4]
					}],
					"max_score": null,
					"total": 6
				}
			},
			"sort": [5]
		}, {
			"_id": "AVoESHYW_OA-dG63Txse",
			"_index": "recipes",
			"_score": null,
			"_source": {
				"name": "广式鲫鱼汤",
				"rating": 5,
				"type": "粤菜"
			},
			"_type": "type",
			"fields": {
				"type": ["粤菜"]
			},
			"inner_hits": {
				"hits": {
					"hits": [{
						"_id": "AVoESHYW_OA-dG63Txse",
						"_index": "recipes",
						"_score": null,
						"_source": {
							"name": "广式鲫鱼汤",
							"rating": 5,
							"type": "粤菜"
						},
						"_type": "type",
						"sort": [5]
					}],
					"max_score": null,
					"total": 1
				}
			},
			"sort": [5]
		}],
		"max_score": null,
		"total": 9
	},
	"timed_out": false,
	"took": 1
}`

func Convert2MySearchResult(bys []byte) (*SearchResult, error) {
	j, err := simplejson.NewJson(bys)
	if err != nil {
		return nil, err
	}
	paths := []string{"hits", "hits"}
	hits := j.GetPath(paths...)
	var hitsnew []*simplejson.Json
	for i := 0; ; i++ {
		hit := hits.GetIndex(i)
		if hit.Interface() == nil {
			break
		}
		if a, exist := hit.CheckGet("inner_hits"); exist {
			if b, exist := a.CheckGet("hits"); exist {
				if c, exist := b.CheckGet("hits"); exist {
					var cnew []*simplejson.Json
					for j := 0; ; j++ {
						inhit := c.GetIndex(j)
						if inhit.Interface() == nil {
							break
						}
						inhit.SetPath([]string{"_source", "_id"}, inhit.Get("_id"))
						fmt.Println(inhit.Get("_id").MustString(), "...1")
						fmt.Println(inhit.Get("_source").Get("_id").MustString(), "...2")
						cnew = append(cnew, inhit.Get("_source"))
					}
					cbys, err := json.Marshal(cnew)
					if err != nil {
						continue
					}
					fmt.Println(string(cbys))
					hit.Set("inner_hits", string(cbys))
					hitsnew = append(hitsnew, hit)
				}
			}
		}
	}
	j.SetPath(paths, hitsnew)
	bys, err = j.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(bys, &result)
	if err != nil {
		return nil, err
	}
	fmt.Println(result.Hits.Hits[0].InnerHits)
	return &result, nil
}

func main() {
	Convert2MySearchResult([]byte(datas))
}
