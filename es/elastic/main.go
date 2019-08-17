package main

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.130.201:9200"))
	if err != nil {
		log.Fatal(err)
	}
	gr, err := client.Get().Index("cdr_e_cdr_*").Type("e_cdr").Id("192.168.130.1_e_cdr_20180315_-1434882631").Do(nil)
	if err != nil {
		log.Fatal(err)
	}
	bys, err := gr.Source.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bys))
}
