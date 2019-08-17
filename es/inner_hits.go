package main

import (
	"encoding/json"
	"github.com/mattbaird/elastigo/lib"
	"github.com/bitly/go-simplejson"
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

func (h *Hits) Len() int {
	return len(h.Hits)
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
	InnerHits   Hits                     `json:"inner_hits"`
}

func Convert2MySearchResult(sr *elastigo.SearchResult) (*SearchResult, error) {
	j, err := simplejson.NewJson(sr.RawJSON)
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
					hit.Set("inner_hits", c)
					hitsnew = append(hitsnew, hit)
				}
			}
		}
	}
	j.SetPath(paths, hitsnew)
	bys, err := j.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(bys, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
