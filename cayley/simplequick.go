package main

import (
	"fmt"
	"log"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/query/gizmo"
	"github.com/cayleygraph/cayley/query"
	"github.com/bitly/go-simplejson"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/context"
)

var (
	Predicate_Children = "children"
	Predicate_Pointto = "pointto"
	Src = "src"
	Dst = "dst"
)

func getobject(ip, area string) interface{} {
	if ip == "" {
		return area
	}
	j := simplejson.New()
	j.Set("ip", ip)
	j.Set("area", area)
	byt, _ := j.MarshalJSON()
	return string(byt)
}
func main() {
	// Create a brand new graph
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}
	store.AddQuad(quad.Make("a", Predicate_Pointto, "b", nil))
	store.AddQuad(quad.Make("b", Predicate_Pointto, "c", nil))
	store.AddQuad(quad.Make("c", Predicate_Pointto, "d", nil))

	ses := gizmo.NewSession(store)

	path("c", ses)

}

func path(begain string, ses *gizmo.Session) {
	dd:for i := 0; ; i++ {
		var qu = fmt.Sprintf(`g.V("%s")`, begain)
		for ii := 0; ii < i; ii++ {
			qu += fmt.Sprintf(`.In("%s")`, Predicate_Pointto)
		}
		qu += fmt.Sprintf(`.Tag("%s").In("%s").Tag("%s").TagArray()`, Dst, Predicate_Pointto, Src)
		fmt.Println(qu)
		c := make(chan query.Result, 2)
		ses.Execute(context.TODO(), qu, c, 100)
		for res := range c {
			if res.Result() == nil {
				break dd
			}
			result := res.Result().([]map[string]interface{})
			if len(result) == 0 {
				break dd
			}
			fmt.Println(result)
		}
	}
}

func search(parent, has  string, level int, ses *gizmo.Session) {
	var out, in string
	for ; level < 4; level++ {
		out += fmt.Sprintf(`.Out("%s")`, Predicate_Children)
		in += fmt.Sprintf(`.In("%s")`, Predicate_Children)
	}
	var qu string
	if has != "" {
		qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")%s.Is("%s")`, parent, out, Src, Predicate_Pointto, Dst, in, has)
		qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")%s.Is("%s")).Unique().TagArray()`, parent, out, Dst, Predicate_Pointto, Src, in, has)
	} else {
		qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")`, parent, out, Src, Predicate_Pointto, Dst)
		qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")).Unique().TagArray()`, parent, out, Dst, Predicate_Pointto, Src)
	}


	//qu := fmt.Sprintf(`g.V("%s").Out("%s").Out("%s").Tag("source").Out("%s").Tag("dest").Or(g.V("%s").Out("%s").Tag("dest").In("%s").Tag("source")).All()`, key, Predicate_Children, Predicate_Children, Predicate_Pointto, key, Predicate_Children, Predicate_Pointto)
	//qu := ` g.V().Tag("source").Out().Tag("dest").Labels().All()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").Count()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").All()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").ToArray()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").ToValue()`
	//qu := ` g.V().Tag("source").Out().Tag("dest").TagArray()`
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
