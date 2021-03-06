package test

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/query/gizmo"
	"log"
	"testing"
	"sort"
	"github.com/davecgh/go-spew/spew"
	"github.com/cayleygraph/cayley/query"
	"golang.org/x/net/context"
	"fmt"
)

func Test_Cayley3(t *testing.T) {
	// Create a brand new graph
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}

	store.AddQuad(quad.Make("13568984989", "useqq", "59135282", "1"))
	store.AddQuad(quad.Make("13568984989", "useaddress", "成都", "2"))
	store.AddQuad(quad.Make("13568984989", "useemail", "59135282@qq.com", "3"))

	store.AddQuad(quad.Make("59135282", "useaddress", "宜昌", "4"))

	qus := []string{
		`g.Emit({"tel":"13568984989"});g.V("13568984989").Tag('source').Out(['useqq', 'useaddress', 'useemail'], 'type').Tag("dest1").Labels().All()`,

		`g.V("13568984989").Tag('source').Out(['useqq', 'useaddress', 'useemail'], 'type1').Tag("dest1").Out(['useqq', 'useaddress', 'useemail'], 'type2').Tag("dest2").All()`,

		`var p1=g.M().Out(['useqq', 'useaddress', 'useemail'], 'type');
		 g.V("13568984989").Tag("source").FollowRecursive(p1).Tag("dest").All()
		`,
		`var p1=g.M().Both(['useqq', 'useaddress', 'useemail'], 'type');
		 g.V("13568984989").Tag("source").Follow(p1).Tag("dest").All()
		`,
		`var p1=g.M().Both(['useqq', 'useaddress', 'useemail']).Both(['useqq', 'useaddress', 'useemail']);
		 g.V("13568984989").Tag("source").Follow(p1).Tag("dest").All()
		`,
	}

	for _, qu := range qus {
		fmt.Println(qu)
		ses := gizmo.NewSession(store)
		c := make(chan query.Result, 5)
		ses.Execute(context.TODO(), qu, c, 100)
		var got []string
		for res := range c {
			func() {
				fmt.Println(ses.FormatREPL(res),"......")
				got = append(got, ses.FormatREPL(res))
			}()
		}
		sort.Strings(got)
		spew.Dump(got)
	}

}
