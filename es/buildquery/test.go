package main

import (
	"github.com/tj/es"
	"fmt"
)

func main() {
	//a()
	aa()
}

func a() {
	sum := es.Agg("duration_sum", es.Sum("duration"))

	labels := es.Agg("labels",
		es.Terms("issue.labels.keyword", 100),
		es.Aggs(sum))

	days := es.Agg("days",
		es.DateHistogram("1d"),
		es.Aggs(sum, labels))

	query := es.Query(es.Aggs(sum, labels, days))
	fmt.Println(query)
}

func aa() {
	labels := es.Aggs(
		es.Agg("labels",
			es.Terms("issue.labels.keyword", 100),
			es.Aggs(
				es.Agg("duration_sum",
					es.Sum("duration")))))

	repos := es.Aggs(
		es.Agg("repos",
			es.Terms("repository.name.keyword", 100),
			labels))

	filter := es.Filter(
		es.Term("user.login", "tj"),
		es.Range("now-7d", "now"))

	results := es.Aggs(
		es.Agg("results",
			filter(repos)))

	query := es.Pretty(es.Query(results))
	fmt.Println(query)
}
