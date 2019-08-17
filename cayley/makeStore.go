package main

import (
	"time"
	"github.com/cayleygraph/cayley"
	"os"
	"github.com/cayleygraph/cayley/graph"
	"path/filepath"
	"log"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"fmt"
	"github.com/robfig/cron"
)

var IpgraphPaths = "D:/ipgraph/test"

func main() {
	C:=cron.New()
	C.Start()
	C.AddFunc("0/2 * * * *", func() {
		getGraphs(true).Close()
		getGraphs(false).Close()
	})
	time.Sleep(time.Hour)
}

func getIpgraphPaths(sure bool, t ... time.Time) string {
	tt := time.Now()
	if len(t) > 0 {
		tt = t[0]
	}
	if sure {
		return filepath.Join(IpgraphPaths, tt.Format("20060102150405")+"ipgraph_sure.db")
	} else {
		return filepath.Join(IpgraphPaths, tt.Format("20060102150405")+"ipgraph.db")
	}
}

func getGraphs(sure bool, t ...time.Time) *cayley.Handle {
	fp := getIpgraphPaths(sure, t...)
	fmt.Println(fp)
	_, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	os.Chmod(fp, os.ModePerm)
	// Initialize the database
	graph.InitQuadStore("bolt", fp, nil)
	// Open and use the database
	store, err := cayley.NewGraph("bolt", fp, nil)
	//MainGraph, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatal(err)
	}
	return store
}
