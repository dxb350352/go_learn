package main

import (
	"fmt"
	"log"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	"os"
	"github.com/cayleygraph/cayley/query/gizmo"
	"sort"
	"github.com/davecgh/go-spew/spew"
	"github.com/cayleygraph/cayley/query"
	"golang.org/x/net/context"
)

func main() {
	//// File for your new BoltDB. Use path to regular file and not temporary in the real world
	//tmpfile, err := ioutil.TempFile("", "example")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer os.Remove(tmpfile.Name()) // clean up
	tmpfile, err := os.OpenFile("d:/bolt.db", os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	// Initialize the database
	graph.InitQuadStore("bolt", tmpfile.Name(), nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", tmpfile.Name(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	store.AddQuad(quad.Make("phrase of the day", "is of course", "Hello BoltDB!", "demo graph"))

	// Now we create the path, to get to our data

	qu := ` g.V().Tag("source").Out().Tag("dest").Labels().All()`
	fmt.Println(qu)
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 5)
	ses.Execute(context.TODO(), qu, c, 100)
	var got []string
	for res := range c {
		func() {
			fmt.Println(ses.FormatREPL(res))
			got = append(got, ses.FormatREPL(res))
		}()
	}
	fmt.Println("...................")
	sort.Strings(got)
	spew.Dump(got)
}
