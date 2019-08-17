package module

import (
	"regexp"
	"github.com/jackiedong168/sequence"
	"net/url"
	"strings"
)

func Prepare(app, _type, by, keywords, sortField, order string, startTime, endTime int64) (*EsApi, error) {
	//spew.Dump(keywords)
	var index, nestedPath, nestedQuery, notimefilter string
	esapi := EsApi{}
	timeField := ""
	scan := sequence.NewScanner()
	lparser := sequence.NewLucParser()
	seq1, _ := scan.Scan(keywords)
	seq2, _ := lparser.Parse(seq1)
	isfirst := true
	m := map[string]bool{"AND": true, "NOT": true, "OR": true}

	var l sequence.Sequence
	var nestedSequence sequence.Sequence
	for _, token := range seq2 {
		if isfirst && m[token.Value] {
			continue
		}
		if isfirst && m[token.Op] {
			token.Op = ""
		}

		isfirst = false

		if token.FieldS == "" {
			l = append(l, token)
			isfirst = false
			continue
		}
		switch token.FieldS {
		case "__app":
			app = token.Value
		case "__type":
			_type = token.Value
		case "__index":
			index = token.Value
		case "_index":
			index = token.Value
		case "__notimefilter":
			notimefilter = token.Value
		case "__nestedPath":
			nestedPath = token.Value
		case "joinField":
			esapi.JoinField = token.Value
		case "joinIndices":
			esapi.JoinIndices = token.Value
		case "joinTypes":
			esapi.JoinTypes = token.Value
		case "joinPath":
			esapi.JoinPath = token.Value
		case "joinQuery":
			q, _ := url.QueryUnescape(token.Value)
			esapi.JoinQueryString = q
		case "timeField":
			timeField = token.Value
		default:
			if nestedPath != "" {
				if strings.HasPrefix(token.FieldS, nestedPath+".") {
					nestedSequence = append(nestedSequence, token)
				}
			} else {
				l = append(l, token)
			}
		}
	}
	nestedQuery = lparser.String(nestedSequence, ":")

	keywords = lparser.String(l, ":")

	if index != "" && _type == "" {
		_type = "*"
	} else if index == "" && app == "" {
		index = "*"
		_type = "*"
		app = "*"
		timeField = "__time"

	}

	if app == "jdzh" {
		by = "month"
	}

	if index == "" {
		index, keywords = HandleIndex("", _type, by, keywords, startTime, endTime)
	}
	if timeField == "" {
		timeField = "__time"
	}

	keywords = strings.TrimSpace(keywords)
	if strings.HasPrefix(keywords, "NOT") && index != "" {
		keywords = index + " " + keywords
	}

	reg1 := regexp.MustCompile(`^[\s|NOT|AND|OR]*(\w+)`)
	keywords = reg1.ReplaceAllString(keywords, "$1")

	esapi.Index = index
	esapi.IndexPrefix = "*"
	esapi.NestedPath = nestedPath
	esapi.NestedQuery = nestedQuery
	esapi.Type = _type
	esapi.By = by
	esapi.Keywords = keywords
	esapi.TimeField = timeField
	esapi.StartTime = startTime
	esapi.EndTime = endTime
	esapi.SortField = sortField
	esapi.Order = order
	if notimefilter == "true" {
		esapi.NoTimeFilter = true
	}

	//spew.Dump(esapi)
	return &esapi, nil

}
