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
	"sync"
	"github.com/davecgh/go-spew/spew"
)

type Loggerer struct {
}

func (l Loggerer) Error(i... interface{}) {
	spew.Dump(i)
}

var (
	Logger Loggerer
	Predicate_Pointto = "pointto"
	Predicate_Children = "children"
	SearchUrl = "https://192.168.130.240:9005/search/search"
	IpgraphPath string
	IpgraphMain string
	IpgraphSureMain string
	Src = "src"
	Dst = "dst"
	PageSize = 100
	MainGraph *cayley.Handle
	MainGraphSure *cayley.Handle
	Lock = new(sync.Mutex)
	LockFlag bool
	Cond = sync.NewCond(Lock)
	ScoreMax = 0.1
	AppType = []string{
		//"cdrcb|qy_ids|src_addr|dst_addr|",
		"cdrcb|score_netflow|sip|dip|score",
	}
)

func init() {
	IpgraphPath, _ = os.Getwd()
	IpgraphPath = filepath.Join(IpgraphPath, "ipgraph")
	err := os.MkdirAll(IpgraphPath, 0777 | os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	IpgraphMain = filepath.Join(IpgraphPath, "ipgraph.db")
	IpgraphSureMain = filepath.Join(IpgraphPath, "ipgraph_sure.db")
	fmt.Println(IpgraphPath)
	fmt.Println(IpgraphMain)
	fmt.Println(IpgraphSureMain)
}
func count(path string) {
	if path == "" {
		MainGraph.Close()
		path = IpgraphMain
	}
	// Open and use the database
	store, err := cayley.NewGraph("bolt", path, nil)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer store.Close()
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 2)
	qu := `g.V().Tag("source").Out().Tag("dest").Count()`
	ses.Execute(context.TODO(), qu, c, 10)
	for res := range c {
		//spew.Dump(res.Result())
		fmt.Println(res.Result())
	}

}
func main() {
	//Ipgraph()
	//fmt.Println("Ipgraph")
	//Combine30DataFile()
	//fmt.Println("Combine30DataFile")
	spew.Dump(Map("中国", "", 10000, 0, false), ".............")
	//parent := strings.Replace(getNode("192.168.130.240", "四川省,成都市,武侯区", false).(string), `"`, `\"`, -1)
	//spew.Dump(Map(parent, "", 10000, 4, false), ".............")
	//spew.Dump(Map("四川省", "上海省", 10000, 1, true), ".............")
	//spew.Dump(Map("五华区", "武侯区", 10000, 3, false), ".............")
	//spew.Dump(Path("192.168.128.104", "上海,上海市,金山区"))
	//spew.Dump(Path("ip上海", "上海省,上海市,金山区"))
	//count(filepath.Join(IpgraphPath, "20171101ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "20171117ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "20171118ipgraph.db"))
	//count(filepath.Join(IpgraphPath, "20171121ipgraph.db"))
	count("")
	//关闭graph
	if MainGraph != nil {
		MainGraph.Close()
		MainGraph = nil
	}
	if MainGraphSure != nil {
		MainGraphSure.Close()
		MainGraphSure = nil
	}
}

type Node struct {
	Ip   string `json:"ip"`
	Area string `json:"area"`
	Sure bool `json:"sure"`
}
//每天一个文件
func getIpgraphPath(sure bool, t... time.Time) string {
	tt := time.Now()
	if len(t) > 0 {
		tt = t[0]
	}
	if sure {
		return filepath.Join(IpgraphPath, tt.Format("20060102") + "ipgraph_sure.db")
	} else {
		return filepath.Join(IpgraphPath, tt.Format("20060102") + "ipgraph.db")
	}
}

func GetGraph(sure bool, t... time.Time) *cayley.Handle {
	fp := IpgraphMain
	if sure {
		fp = IpgraphSureMain
	}
	if len(t) > 0 {
		fp = getIpgraphPath(sure, t...)
	}
	_, err := os.OpenFile(fp, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	// Initialize the database
	graph.InitQuadStore("bolt", fp, nil)
	// Open and use the database
	store, err := cayley.NewGraph("bolt", fp, nil)
	//MainGraph, err := cayley.NewMemoryGraph()
	if err != nil {
		Logger.Error(err)
		return nil
	}
	return store
}
//
func GetMainGraph(sure bool) *cayley.Handle {
	var store *cayley.Handle
	if sure {
		if MainGraphSure == nil {
			MainGraphSure = GetGraph(sure)
		}
		store = MainGraphSure
	} else {
		if MainGraph == nil {
			MainGraph = GetGraph(sure)
		}
		store = MainGraph
	}
	return store
}
//查询到这个节点的所有路径
func Path(ip, area string, sure bool) []*simplejson.Json {
	var all []*simplejson.Json
	store := GetMainGraph(sure)
	//begain
	node := strings.Replace(getNode(ip, area, sure).(string), `"`, `\"`, -1)
	m := map[string]bool{ip:true}

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
	return all
}
//查询
func Map(parent, has string, limit, level int, sure bool) []*simplejson.Json {
	store := GetMainGraph(sure)
	//begain
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
	return getResult(store, qu)
}
//获取查询结果
func getResult(store *cayley.Handle, qu string) []*simplejson.Json {
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 2)
	ses.Execute(context.TODO(), qu, c, 10)
	var got []*simplejson.Json
	for res := range c {
		fmt.Println(res.Result(),"..................1")
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
//把节点中数据抽取出来
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
//添加各种数据到图中
func Ipgraph() {
	Lock.Lock()
	defer Lock.Unlock()
	if LockFlag {
		Cond.Wait()
	}
	LockFlag = true
	defer func() {
		LockFlag = false
		Cond.Signal()
	}()
	//不确定的
	store := GetGraph(false, time.Now())
	defer store.Close()

	mainstore := GetMainGraph(false)
	//确定的
	store_sure := GetGraph(true, time.Now())
	defer store_sure.Close()

	mainstore_sure := GetMainGraph(true)

	//startTime := "0"
	startTime := fmt.Sprint((time.Now().Unix() - 300) * 1000)
	endTime := fmt.Sprint(time.Now().Unix() * 1000)

	for _, v := range AppType {
		arr := strings.Split(v, "|")
		if len(arr) != 5 {
			continue
		}
		if arr[4] == "" {
			GainData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], store_sure, mainstore_sure)
		} else {
			GainData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], store, mainstore)
		}
	}
}
//查询并添加到图中
func GainData(__app, __type, startTime, endTime, sip, dip, scoreField string, store, mainstore *cayley.Handle) {
	request := gorequest.New()
	if strings.HasPrefix(SearchUrl, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	request.Get(SearchUrl)

	total := 1

	request.Param("app", __app)
	request.Param("_type", __type)
	request.Param("size", fmt.Sprint(PageSize))
	request.Param("startTime", startTime)
	request.Param("endTime", endTime)
	for start := 0; start < total; {
		request.QueryData.Set("from", fmt.Sprint(start))
		start += PageSize
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
			if vsip == "" || src_province == "" || src_city == "" || src_district == "" || vdip == "" || dst_province == "" || dst_city == "" || dst_district == "" {
				continue
			}
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
			err = mainstore.AddQuadSet(quadarr)
			if err != nil {
				Logger.Error(err)
			}
		}
	}
}
//生成一条路径
func AddQuad(sip, dip, subject, predicate, object, label string, quadarr  []quad.Quad, score float64, scoreField string) []quad.Quad {
	if subject != "" && object != "" && strings.Trim(subject, ",") != "" && strings.Trim(object, ",") != "" {
		subnode := getNode(sip, subject, scoreField == "")
		objnode := getNode(dip, object, scoreField == "")
		if scoreField == "" {
			quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
		} else {
			if score > 0 && score - ScoreMax < 0 {
				quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
			}
		}
	}
	return quadarr
}
//生成节点
func getNode(ip, area string, sure bool) interface{} {
	if ip == "" {
		return area
	}
	node := Node{Ip:ip, Area:area, Sure:sure}
	byt, _ := json.Marshal(node)
	return string(byt)
}
//合并30天数据
func Combine30DataFile() {
	Lock.Lock()
	defer Lock.Unlock()
	if LockFlag {
		Cond.Wait()
	}
	LockFlag = true
	defer func() {
		LockFlag = false
		Cond.Signal()
	}()
	//关闭graph
	if MainGraph != nil {
		MainGraph.Close()
		MainGraph = nil
	}
	if MainGraphSure != nil {
		MainGraphSure.Close()
		MainGraphSure = nil
	}
	//删除之前文件
	err := os.Remove(IpgraphMain)
	if err != nil {
		Logger.Error(err)
	}
	err = os.Remove(IpgraphSureMain)
	if err != nil {
		Logger.Error(err)
	}
	//30天文件
	MainGraph = GetGraph(false)
	MainGraphSure = GetGraph(true)
	var data []quad.Quad
	t := time.Now()
	for _, sure := range []bool{true, false} {
		store := MainGraph
		if sure {
			store = MainGraphSure
		}
		for i := 0; i < 30; i++ {
			//30天文件
			tt := t.Add(time.Duration(-24 * i) * time.Hour)
			if _, err := os.Stat(getIpgraphPath(sure, tt)); err != nil {
				Logger.Error(err)
				continue
			}
			fstore := GetGraph(sure, tt)
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
						data = append(data, quad.Make(src, Predicate_Pointto, dst, ""))
					} else {
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
			fstore.Close()
		}
	}
}