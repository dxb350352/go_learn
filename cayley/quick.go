package main

import (
	"fmt"
	"log"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/query/gizmo"
	"github.com/cayleygraph/cayley/query"
	"golang.org/x/net/context"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/bitly/go-simplejson"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

var (
	Predicate_Children = "children"
	Predicate_Pointto = "pointto"
	Src = "src"
	Dst = "dst"
)

func getNode(ip, area string) interface{} {
	if ip == "" {
		return area
	}
	return fmt.Sprintf(`{"ip":"%s","area":"%s"}`, ip, area)
}
func main() {
	// Create a brand new graph
	//fp:="d:/ipgraph/20171118ipgraph.db"
	//_, err := os.OpenFile(fp, os.O_RDWR | os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// Initialize the database
	//graph.InitQuadStore("bolt", fp, nil)
	//// Open and use the database
	//store, err := cayley.NewGraph("bolt", fp, nil)
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}
	defer store.Close()
	store.AddQuad(quad.Make(getNode("", "中国"), Predicate_Children, getNode("", "四川省"), nil))
	store.AddQuad(quad.Make(getNode("", "中国"), Predicate_Children, getNode("", "上海省"), nil))
	store.AddQuad(quad.Make(getNode("", "四川省"), Predicate_Children, getNode("", "成都市"), nil))
	store.AddQuad(quad.Make(getNode("", "上海省"), Predicate_Children, getNode("", "上海市"), nil))
	store.AddQuad(quad.Make(getNode("", "成都市"), Predicate_Children, getNode("", "四川省,成都市,武侯区"), nil))
	store.AddQuad(quad.Make(getNode("", "上海市"), Predicate_Children, getNode("", "上海省,上海市,金山区"), nil))
	store.AddQuad(quad.Make(getNode("", "四川省,成都市,武侯区"), Predicate_Children, getNode("ip四川", "四川省,成都市,武侯区"), nil))
	store.AddQuad(quad.Make(getNode("", "上海省,上海市,金山区"), Predicate_Children, getNode("ip上海", "上海省,上海市,金山区"), nil))
	store.AddQuad(quad.Make(getNode("ip四川", "四川省,成都市,武侯区"), Predicate_Pointto, getNode("ip上海", "上海省,上海市,金山区"), nil))
	// Now we create the path, to get to our data
	p := cayley.StartPath(store).Both().Unique()
	for {
		count, err := p.Iterate(nil).Count()
		if err != nil || count == 0 {
			fmt.Println(err, count)
			break
		}
		// Now we iterate over results. Arguments:
		// 1. Optional context used for cancellation.
		// 2. Flag to optimize query before execution.
		// 3. Quad store, but we can omit it because we have already built path with it.
		err = p.Iterate(nil).EachValue(nil, func(value quad.Value) {
			//nativeValue := quad.NativeOf(value) // this converts RDF values to normal Go types
			//fmt.Println(nativeValue)

		})
		if err != nil {
			log.Fatalln(err)
		}
		break
		p = p.Out(quad.String("use")).Unique()
	}
	ses := gizmo.NewSession(store)
	//search("中国", "", 0, ses)
	search("四川省", "上海省", 1, ses)
	//search("上海省", "四川省", 1, ses)
	//search("成都市","上海市",2, ses)
	//search("上海市", "成都市", 2, ses)
}

func search(parent, has  string, level int, ses *gizmo.Session) {
	//var out, in string
	//for ; level < 4; level++ {
	//	out += fmt.Sprintf(`.Out("%s")`, Predicate_Children)
	//	in += fmt.Sprintf(`.In("%s")`, Predicate_Children)
	//}
	//var qu string
	//if has != "" {
	//	qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")%s.Is("%s")`, parent, out, Src, Predicate_Pointto, Dst, in, has)
	//	qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")%s.Is("%s")).Unique().TagArray()`, parent, out, Dst, Predicate_Pointto, Src, in, has)
	//} else {
	//	qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")`, parent, out, Src, Predicate_Pointto, Dst)
	//	qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")).Unique().TagArray()`, parent, out, Dst, Predicate_Pointto, Src)
	//}


	//qu := fmt.Sprintf(`g.V("%s").Out("%s").Out("%s").Tag("source").Out("%s").Tag("dest").Or(g.V("%s").Out("%s").Tag("dest").In("%s").Tag("source")).All()`, key, Predicate_Children, Predicate_Children, Predicate_Pointto, key, Predicate_Children, Predicate_Pointto)
	//qu := ` g.V().Tag("source").Out().Tag("dest").Labels().All()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").Count()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").All()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").ToArray()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").ToValue()`
	qu := ` g.V().Tag("source").Out().Tag("dest").TagArray()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").TagValue()`

	fmt.Println(qu)
	c := make(chan query.Result, 2)
	ses.Execute(context.TODO(), qu, c, 2)

	spew.Dump(getResult(c))
}

func getResult(c chan query.Result) []*simplejson.Json {
	var got []*simplejson.Json
	for res := range c {
		func() {
			defer func() {
				recover()
			}()
			byt, err := json.Marshal(res.Result())
			if err != nil {
				fmt.Println(err)
				return
			}
			j, err := simplejson.NewJson(byt)
			if err != nil {
				fmt.Println(err)
				return
			}
			var i int
			for {
				tj := j.GetIndex(i)
				i++
				if tj.Interface() == nil {
					break
				}
				gt := simplejson.New()
				buildJson(gt, tj, Src)
				buildJson(gt, tj, Dst)
				got = append(got, gt)
			}
		}()
	}
	return got
}

func buildJson(j, tj *simplejson.Json, prefix string) {
	ttjj, err := simplejson.NewJson([]byte(tj.Get(prefix).MustString()))
	if err != nil {
		fmt.Println(err)
		return
	}
	m := ttjj.MustMap()
	for k, v := range m {
		j.Set(prefix + "_" + k, v)
	}
}
