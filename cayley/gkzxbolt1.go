package main

import (
	"os"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"path/filepath"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"fmt"
	"crypto/tls"
	"github.com/jackiedong168/gorequest"
	"strings"
	"time"
	"github.com/bitly/go-simplejson"
	"github.com/cayleygraph/cayley/query/gizmo"
	"golang.org/x/net/context"
	"github.com/cayleygraph/cayley/query"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

type Loggerer struct {
}

func (l Loggerer) Error(i interface{}) {
	fmt.Println(i)
}

var (
	Logger Loggerer
	Predicate_Pointto = "pointto"
	Predicate_Children = "children"
	SearchUrl = "https://192.168.130.201:9005/search/search"
	IpgraphPath = "d:/ipgraph"
	IpgraphMain = "d:/ipgraph/ipgraph.db"
	Src = "src"
	Dst = "dst"
	PageSize = 100
	MainGraph *cayley.Handle
	CurrentGraph *cayley.Handle
)

func init() {

	dir, _ := filepath.Split(IpgraphPath)
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//Ipgraph()
	//Combine30DataFile()
	spew.Dump(Map("中国", "", 10000, 0), ".............")
	//spew.Dump(Map("四川省", "上海", 10000, 1), ".............")
	//Path("192.168.128.104", "上海,上海市,金山区")
	//count(filepath.Join(IpgraphPath, "20171101ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "20171117ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "20171119ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "ipgraph.db"))
}

func count(path string) {
	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer store.Close()
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 2)
	qu := ` g.V().Tag("source").Out().Tag("dest").Count()`
	ses.Execute(context.TODO(), qu, c, 10)
	for res := range c {
		fmt.Println(res.Result())
	}

}
func getMainGraph() *cayley.Handle {
	if MainGraph != nil {
		return MainGraph
	}
	_, err := os.OpenFile(IpgraphMain, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	// Initialize the database
	graph.InitQuadStore("bolt", IpgraphMain, nil)
	// Open and use the database
	MainGraph, err = cayley.NewGraph("bolt", IpgraphMain, nil)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	return MainGraph
}
func getCurrentGraph() *cayley.Handle {
	if CurrentGraph != nil {
		return CurrentGraph
	}
	fpath := getIpgraphPath()
	_, err := os.OpenFile(fpath, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	// Initialize the database
	graph.InitQuadStore("bolt", fpath, nil)
	// Open and use the database
	CurrentGraph, err = cayley.NewGraph("bolt", fpath, nil)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	return CurrentGraph
}
func Combine30DataFile() {
	//删除之前文件
	err := os.Remove(IpgraphMain)
	if err != nil {
		Logger.Error(err)
	}
	//30天文件
	MainGraph = nil
	store := getMainGraph()

	var data []quad.Quad
	t := time.Now()
	for i := 0; i < 30; i++ {
		fstore := getCurrentGraph()
		daypath := getIpgraphPath(t.Add(time.Duration(-24 * i) * time.Hour))
		if i != 0 {
			//30天文件
			// Open and use the database
			if _, err := os.Stat(daypath); err != nil {
				Logger.Error(err)
				continue
			}
			fstore, err = cayley.NewGraph("bolt", daypath, nil)
			if err != nil {
				Logger.Error(err)
				continue
			}
		}
		qu := fmt.Sprintf(`g.V().Tag("%s").Out().Tag("%s").TagArray()`, Src, Dst)
		c := make(chan query.Result, 2)
		fses := gizmo.NewSession(fstore)
		fses.Execute(context.TODO(), qu, c, 2)
		for res := range c {
			if res.Result() == nil {
				break
			}
			byt, err := json.Marshal(res.Result())
			if err != nil {
				Logger.Error(err)
				break
			}
			j, err := simplejson.NewJson(byt)
			if err != nil {
				Logger.Error(err)
				break
			}
			var i int
			for {
				tj := j.GetIndex(i)
				i++
				if tj.Interface() == nil {
					break
				}
				src := tj.Get(Src).MustString()
				dst := tj.Get(Dst).MustString()
				if strings.Contains(tj.Get(Src).MustString(), "{") {
					fmt.Println(src, Predicate_Pointto, dst, "......................", daypath)
					data = append(data, quad.Make(src, Predicate_Pointto, dst, ""))
				} else {
					fmt.Println(src, Predicate_Children, dst, "......................", daypath)
					data = append(data, quad.Make(src, Predicate_Children, dst, ""))
				}
				if len(data) >= PageSize {
					err = store.AddQuadSet(data)
					data = []quad.Quad{}
					if err != nil {
						Logger.Error(err)
					}
				}
			}
		}
		if len(data) > 0 {
			store.AddQuadSet(data)
			data = []quad.Quad{}
			if err != nil {
				Logger.Error(err)
			}
		}
	}
}

func Path(ip, area string) {
	// Open and use the database
	store, err := cayley.NewGraph("bolt", IpgraphPath, nil)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer store.Close()
	//begain
	node := strings.Replace(getNode(ip, area).(string), `"`, `\"`, -1)
	m := map[string]bool{ip:true}
	var all []*simplejson.Json
	for i := 0; ; i++ {
		var qu = fmt.Sprintf(`g.V("%s")`, node)
		for ii := 0; ii < i; ii++ {
			qu += fmt.Sprintf(`.In("%s")`, Predicate_Pointto)
		}
		qu += fmt.Sprintf(`.Tag("%s").In("%s").Tag("%s").TagArray()`, Dst, Predicate_Pointto, Src)
		result := getResult(store, qu)
		if len(result) == 0 {
			break
		}

		var usefull bool
		for ii := 0; ii < len(result); ii++ {
			if m[result[ii].Get(Src + "_ip").MustString()] {
				continue
			}
			usefull = true
			m[result[ii].Get(Src + "_ip").MustString()] = true
			all = append(all, result[ii])
		}
		if !usefull {
			break
		}
	}
	//组建所有路径

}

func Map(parent, has string, limit, level int) []*simplejson.Json {
	var got []*simplejson.Json
	// Open and use the database
	store, err := cayley.NewGraph("bolt", IpgraphMain, nil)
	if err != nil {
		fmt.Println(err, ".....ddd")
		return got
	}
	defer store.Close()
	//qu := ` g.V().Tag("source").Out().Tag("dest").Count()`
	//qu := fmt.Sprintf(`g.V("%s").Out("%s").Out("%s").Tag("source").Out("%s").Tag("dest").Or(g.V("%s").Out("%s").Tag("dest").In("%s").Tag("source")).All()`, parent, Predicate_Children, Predicate_Children, Predicate_Pointto, parent, Predicate_Children, Predicate_Pointto)
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
	fmt.Println(qu)
	//qu = `g.V("四川省").Out("children").Out("children").Out("children").Tag("src").Out("pointto").Tag("dst").In("children").In("children").In("children").Is("上海省").Or(g.V("四川省").Out("children").Out("children").Out("children").Tag("dst").In("pointto").Tag("src").In("children").In("children").In("children").Is("上海")).Unique().TagArray()`
	//fmt.Println(qu)

	return getResult(store, qu)
}
//获取查询结果
func getResult(store *cayley.Handle, qu string) []*simplejson.Json {
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 2)
	ses.Execute(context.TODO(), qu, c, 10)
	var got []*simplejson.Json
	for res := range c {
		if res.Result() == nil {
			return got
		}
		byt, err := json.Marshal(res.Result())
		if err != nil {
			Logger.Error(err)
			return got
		}
		j, err := simplejson.NewJson(byt)
		if err != nil {
			Logger.Error(err)
			return got
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
	}
	return got
}

func buildJson(j, tj *simplejson.Json, prefix string) {
	ttjj, err := simplejson.NewJson([]byte(tj.Get(prefix).MustString()))
	if err != nil {
		Logger.Error(err)
		return
	}
	m := ttjj.MustMap()
	for k, v := range m {
		j.Set(prefix + "_" + k, v)
	}
}

func Ipgraph() {
	tmpfile, err := os.OpenFile(IpgraphPath, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		Logger.Error(err)
	}
	defer tmpfile.Close()
	// Initialize the database
	graph.InitQuadStore("bolt", tmpfile.Name(), nil)

	// Open and use the database
	store, err := cayley.NewGraph("bolt", tmpfile.Name(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer store.Close()
	//
	startTime := "0"
	endTime := fmt.Sprint(time.Now().Unix() * 1000)
	//Searchback("sas", "netflow_score", startTime, endTime, "sip", "dip", "score", store)
	//Searchback("cdrcb", "cdrcb", startTime, endTime, "src_addr", "dst_addr", "", store)
	//GainData("cdrcb", "cdrcb", startTime, endTime, "", store)
	GainData("sas", "netflow", startTime, endTime, "sip", "dip", "", store)
}

func GainData(__app, __type, startTime, endTime, sip, dip, scoreField string, store *cayley.Handle) {
	request := gorequest.New()
	if strings.HasPrefix(SearchUrl, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	request.Get(SearchUrl)

	total := 1
	size := 100
	request.Param("app", __app)
	request.Param("_type", __type)
	request.Param("size", fmt.Sprint(size))
	request.Param("startTime", startTime)
	request.Param("endTime", endTime)
	for start := 0; start < total; {
		request.QueryData.Set("from", fmt.Sprint(start))
		start += size
		_, body, errs := request.EndBytes()
		if len(errs) != 0 {
			Logger.Error(errs)
			continue
		}
		j, err := simplejson.NewJson(body)
		if err != nil {
			Logger.Error(err)
			continue
		}
		total = j.Get("total").MustInt()
		if total > 10000 {
			total = 10000
		}
		j = j.Get("hits")
		var quadarr []quad.Quad
		for i := 0; ; i++ {
			tj := j.GetIndex(i)
			if tj == nil || tj.Interface() == nil {
				break
			}
			vsip := tj.GetPath("_source", sip).MustString()
			vdip := tj.GetPath("_source", dip).MustString()
			src_province := strings.TrimSuffix(tj.GetPath("_source", "src_province").MustString(), "市")
			src_city := tj.GetPath("_source", "src_city").MustString()
			src_district := tj.GetPath("_source", "src_district").MustString()
			src_pcd := fmt.Sprintf(`%s,%s,%s`, src_province, src_city, src_district)
			dst_province := strings.TrimSuffix(tj.GetPath("_source", "dst_province").MustString(), "市")
			dst_city := tj.GetPath("_source", "dst_city").MustString()
			dst_district := tj.GetPath("_source", "dst_district").MustString()
			dst_pcd := fmt.Sprintf(`%s,%s,%s`, dst_province, dst_city, dst_district)
			_id := ""
			//_id := tj.Get("_id").MustString()
			score := tj.GetPath("_source", scoreField).MustFloat64()
			//省
			quadarr = AddQuad("", "", "中国", Predicate_Children, src_province, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", "中国", Predicate_Children, dst_province, _id, quadarr, score, scoreField)
			//市
			quadarr = AddQuad("", "", src_province, Predicate_Children, src_city, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", dst_province, Predicate_Children, dst_city, _id, quadarr, score, scoreField)
			//县
			quadarr = AddQuad("", "", src_city, Predicate_Children, src_pcd, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", dst_city, Predicate_Children, dst_pcd, _id, quadarr, score, scoreField)
			//县ip
			quadarr = AddQuad("", vsip, src_pcd, Predicate_Children, src_pcd, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", vdip, dst_pcd, Predicate_Children, dst_pcd, _id, quadarr, score, scoreField)
			//真数据
			quadarr = AddQuad(vsip, vdip, src_pcd, Predicate_Pointto, dst_pcd, _id, quadarr, score, scoreField)
		}
		if len(quadarr) > 0 {
			err = store.AddQuadSet(quadarr)
			if err != nil {
				Logger.Error(err)
			}
		}
	}
}
//生成一条路径
func AddQuad(sip, dip, subject, predicate, object, label string, quadarr  []quad.Quad, score float64, scoreField string) []quad.Quad {
	if subject != "" && object != "" {
		subnode := getNode(sip, subject)
		objnode := getNode(dip, object)
		if scoreField == "" {
			quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
		} else {
			if score > 0 && score < 1 / 1e3 {
				quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
			}
		}
	}
	return quadarr
}
//生成节点
func getNode(ip, area string) interface{} {
	fmt.Println(ip, area)
	if ip == "" {
		return area
	}
	return fmt.Sprintf(`{"ip":"%s","area":"%s"}`, ip, area)
}
func getIpgraphPath(t... time.Time) string {
	tt := time.Now()
	if len(t) > 0 {
		tt = t[0]
	}
	return filepath.Join(IpgraphPath, tt.Format("20060102") + "ipgraph.db")
}

//func Searchback(__app, __type, startTime, endTime, sip, dip, scoreField string, store *cayley.Handle) {
//	request := gorequest.New()
//	if strings.HasPrefix(SearchUrl, "https") {
//		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	}
//	request.Get(SearchUrl)
//
//	total := 1
//	size := 100
//	request.Param("app", __app)
//	request.Param("_type", __type)
//	request.Param("size", fmt.Sprint(size))
//	request.Param("startTime", startTime)
//	request.Param("endTime", endTime)
//	for start := 0; start < total; {
//		request.QueryData.Set("from", fmt.Sprint(start))
//		start += size
//		_, body, errs := request.EndBytes()
//		if len(errs) != 0 {
//			Logger.Error(errs)
//			continue
//		}
//		j, err := simplejson.NewJson(body)
//		if err != nil {
//			Logger.Error(err)
//			continue
//		}
//		//total = j.Get("total").MustInt()
//		if total > 10000 {
//			total = 10000
//		}
//		j = j.Get("hits")
//		for i := 0; ; i++ {
//			tj := j.GetIndex(i)
//			if tj == nil || tj.Interface() == nil {
//				break
//			}
//			vsip := tj.GetPath("_source", sip).MustString()
//			vdip := tj.GetPath("_source", dip).MustString()
//			//_id := ""
//			_id := tj.Get("_id").MustString()
//			if vsip != "" && vdip != "" {
//				fmt.Println(i + start, vsip, vdip, _id)
//				if scoreField == "" {
//					store.AddQuad(quad.Make(vsip, Predicate, vdip, _id))
//				} else {
//					score := tj.GetPath("_source", scoreField).MustFloat64()
//					if score > 0 && score < 1 / 1e3 {
//						store.AddQuad(quad.Make(vsip, Predicate, vdip, _id))
//					}
//				}
//			}
//		}
//	}
//}