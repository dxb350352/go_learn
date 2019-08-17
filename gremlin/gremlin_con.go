package main

import (
	"github.com/go-gremlin/gremlin"
	"fmt"
	"github.com/sas/utils"
	"errors"
	"github.com/bitly/go-simplejson"
)

func main() {
	data, err := gremlin.Query(`graph = JanusGraphFactory.open('conf/janusgraph-cassandra-es.properties')
	g = graph.traversal()
	id1 = g.V().has("person_qq", "4586243325").next()`).Exec()
	//data, err := gremlin.Query(`g.V(obj).outE()`).Session(session).SetProcessor("session").Exec()
	fmt.Println(string(data), err,"...")
}

func init() {
	fmt.Println("init..............")
	if err := gremlin.NewCluster("ws://192.168.131.114:8182/gremlin","ws://192.168.131.115:8182/gremlin","ws://192.168.131.116:8182/gremlin"); err != nil {
		println(err.Error(),"..")
		return
	}
	//data, err := Query(`
	//m = graph.openManagement()
	//g = graph.traversal()`).Exec()
	//if err != nil {
	//	println(err.Error(), "...............init")
	//	return
	//}
	//fmt.Println(string(data))
}

func testSession() {
	fmt.Println("with out session......................")
	data, err := Query(`
	id1 = g.V().hasLabel("eveCables").getAt(1)
	id2 = g.V(id1).out("relCable").next()
	g.V(id2).out("relCable").out("relKuaidi").valueMap()
	`).Exec()
	if err != nil {
		println(err.Error(), "...............0")
		return
	}
	fmt.Println(string(data))
	fmt.Println("with session......................")
	session := utils.UUIDName()
	data, err = QueryBySession(`id1 = g.V().hasLabel("eveCables").getAt(1)`, session).Exec()
	if err != nil {
		println(err.Error(), "...............1")
		return
	}
	fmt.Println(string(data))
	jsondata, err := QueryExecBySession(`id2 = g.V(id1).out("relCable").next()`, session)
	if err != nil {
		println(err.Error(), "...............2")
		return
	}
	fmt.Println(jsondata.MarshalJSON())
	jsondata, err = QueryExecBySession(`g.V(id2).out("relCable").out("relKuaidi").valueMap()`, session)
	if err != nil {
		println(err.Error(), "...............3")
		return
	}
	data, err = jsondata.Bytes()
	fmt.Println(string(data))
}

func Query(query string) *gremlin.Request {
	return gremlin.Query(query)
}

func QueryExec(query string) (*simplejson.Json, error) {
	data, err := Query(query).Exec()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(data)
}

func QueryBySession(query, session string) *gremlin.Request {
	return gremlin.Query(query).Session(session).ManageTransaction(true).SetProcessor("session")
}

func QueryExecBySession(query, session string) (*simplejson.Json, error) {
	data, err := QueryBySession(query, session).Exec()
	if err != nil {
		return nil, err
	}
	//fmt.Println(query)
	//fmt.Println(string(data))
	return simplejson.NewJson(data)
}
/**
新增顶点--lable
 */
func makeVertexLabel(lable string) error {
	_, err := QueryExec(`
	m = graph.openManagement()
	` + lable + ` = m.makeVertexLabel("` + lable + `").make()
	m.commit()
	`)
	return err
}
/**
新增边--lable
 */
func makeEdgeLabel(lable string) error {
	_, err := QueryExec(`
	m = graph.openManagement()
	` + lable + ` = m.makeEdgeLabel("` + lable + `").make()
	m.commit()
	`)
	return err
}
/**
新增属性
 */
func makePropertyKey(key, cls string) error {
	_, err := QueryExec(`
	m = graph.openManagement()
	` + key + ` = m.makePropertyKey("` + key + `").dataType(` + cls + `.class).make()
	m.commit()
	`)
	return err
}
/**
查询
T.gt - 大于
T.gte - 大于或等于
T.eq - 等于
T.neq - 不等于
T.lte - 小于或等于
T.lt - 小于
T.in - 包含在列表中
T.notin - 不包含在列表中
 */
func findByLable(lable string, limit, skip int64, m map[string]string) (int64, *simplejson.Json, error) {
	if utils.IsBlankString(lable) {
		return 0, nil, errors.New("lable is nil")
	}
	gql := lable + `=g.V().hasLabel("` + lable + `")`
	if limit > 0 {
		gql += fmt.Sprintf(`.range(%d,%d)`, skip, skip + limit)
	}
	data, err := QueryExec(gql)
	return 0, data, err
}

//func findRelations(label, key string) (*simplejson.Json, error) {
//	if utils.IsBlankString(key) || utils.IsBlankString(label) {
//		return nil, errors.New("label or key can't be nil")
//	}
//	gql:`a1=g.V().has("` + label + `","_key","` + key + `").next()`
//}


func getById(id string) (*simplejson.Json, error) {
	gql := fmt.Sprintf(`g.V().hasId("%s").next()`, id)
	return QueryExec(gql)
}

func getRelationsById(id string) (*simplejson.Json, error) {
	session := utils.UUIDName()
	gql := fmt.Sprintf(`obj = g.V().hasId("%s").next()
	g.V(obj).outE()`, id)
	oute, err := QueryExecBySession(gql, session)
	if err != nil {
		return nil, err
	}
	out, err := QueryExecBySession("g.V(obj).out()", session)
	if err != nil {
		return nil, err
	}
	//有多少边就有多少顶点
	length := len(oute.MustArray())
	for i := 0; i < length; i++ {
		ve := oute.GetIndex(i)
		vo := out.GetIndex(i)
		vo.Set("__edge", *ve)
	}
	return out, err
}
//添加一行数据
func AddVertex(label string, j *simplejson.Json) (*simplejson.Json, error) {
	m, err := j.Map()
	if err != nil {
		return nil, err
	}
	gql := fmt.Sprintf(`graph.addVertex(label, "%s"`, label)
	for k, v := range m {
		gql += fmt.Sprintf(`, "%s", "%s"`, k, v)
	}
	gql += `)`
	fmt.Println(gql)
	return QueryExec(gql)
}