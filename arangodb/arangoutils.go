package models

import (
	"github.com/bitly/go-simplejson"
	"sync"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"context"
	"github.com/sas/utils"
	"strings"
	"io"
	"fmt"
	"crypto/md5"
)
//所有字段
var allFields = []string{
	"person_id",
	"person_phone",
	"person_email",
	"person_qq",
	"person_bankcard_id",
	"person_loginname",
	"person_realname",
	"school_name",
	"company_name",
	"company_fax",
	"person_name",
	"person_nick",
	"person_ip",
}
//所有字段对应的所有表
var allCollections = []string{
	"ent_person_id",
	"ent_person_phone",
	"ent_person_email",
	"ent_person_qq",
	"ent_person_bankcard_id",
	"ent_person_loginname",
	"ent_person_realname",
	"ent_school_name",
	"ent_company_name",
	"ent_company_fax",
	"ent_person_name",
	"ent_person_nick",
	"ent_person_ip",
}
var Arangodb driver.Database
var Graph driver.Graph
var GraphName string = "graph"
var EdgeCollectionName string = "rel_all_collection"

func initArangodb() error {

	arahosts := utils.ConfString(Config, "arango", "ara.hosts", "http://localhost:8529")
	arauser := utils.ConfString(Config, "arango", "ara.user", "root")
	arapassword := utils.ConfString(Config, "arango", "ara.password", "123456")
	aradbname := utils.ConfString(Config, "arango", "ara.dbname", "test")

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: strings.Split(arahosts, ","),
	})
	if err != nil {
		return err
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication(arauser, arapassword),
	})
	if err != nil {
		return err
	}
	ctx := context.Background()
	Arangodb, err = client.Database(ctx, aradbname)
	if err != nil {
		Arangodb, err = client.CreateDatabase(ctx, aradbname, nil)
		if err != nil {
			return err
		}
	}
	Graph, err = Arangodb.Graph(ctx, GraphName)
	if err != nil {
		Graph, err = Arangodb.CreateGraph(ctx, GraphName, nil)
		if err != nil {
			return err
		}
	}
	return err
	////创建表

	//CreateCollection("ent_person_id")
	//CreateCollection("ent_person_phone")
	//CreateCollection("ent_person_email")
	//CreateCollection("ent_person_qq")
	//CreateCollection("ent_person_bankcard_id")
	//CreateCollection("ent_person_loginname")
	//CreateCollection("ent_person_realname")
	//CreateCollection("ent_school_name")
	//CreateCollection("ent_company_name")
	//CreateCollection("ent_company_fax")
	//
	//CreateCollection("ent_person_name")
	//CreateCollection("ent_person_nick")
	////建关系
}
//创建图
func ensureGraph(ctx context.Context, name string) driver.Graph {
	g, err := Arangodb.Graph(ctx, name)
	if driver.IsNotFound(err) {
		g, err = Arangodb.CreateGraph(ctx, name, nil)
		if err != nil {
			Logger.Errorf("ensureGraph err:%v", err)
		}
	} else if err != nil {
		Logger.Errorf("ensureGraph err:%v", err)
	}
	return g
}
func GetEdgeCollection(ctx context.Context, name string) (driver.Collection, error) {
	ec, _, err := Graph.EdgeCollection(ctx, name)
	if err != nil {
		ec, err = Graph.CreateEdgeCollection(ctx, name, driver.VertexConstraints{From: allCollections, To: allCollections})
	}
	return ec, err
}
func SaveEdges(name string, arr []*simplejson.Json) error {
	ctx := context.Background()
	if len(arr) == 0 {
		Logger.Info("SaveEdges no data to save")
		return nil
	}
	ec, err := GetEdgeCollection(ctx, name)
	if err != nil {
		return err
	}
	_, _, err = ec.CreateDocuments(driver.WithSilent(ctx), arr)
	return err
}
//创建表
func CreateCollection(name string) error {
	ctx := context.Background()
	b, err := Arangodb.CollectionExists(ctx, name);
	if err != nil {
		return err
	}
	if !b {
		_, err = Arangodb.CreateCollection(ctx, name, nil)
	}
	return err
}
//获取表
func GetCollection(name string) (driver.Collection, error) {
	collection, err := Arangodb.Collection(nil, name)
	if err != nil {
		collection, err = Arangodb.CreateCollection(nil, name, nil)
	}
	return collection, err
}

func SaveVertexs(name string, arr []*simplejson.Json) error {
	if len(arr) == 0 {
		return nil
	}
	col, err := GetCollection(name)
	if err != nil {
		return err
	}
	_, _, err = col.CreateDocuments(nil, arr)
	return err
}

func DoSearchHit(j *simplejson.Json, wg *sync.WaitGroup) {
	defer wg.Done()
	ent := map[string][]*simplejson.Json{}
	rel := map[string][]*simplejson.Json{}
	Logger.Info(len(j.MustArray()), j)
	for i := 0; ; i++ {
		hit := j.GetIndex(i)
		if hit.Interface() == nil {
			break
		}
		getVerticesAndEdges(hit, &ent, &rel)
	}
	for k, v := range ent {
		err := SaveVertexs(k, v)
		if err != nil {
			Logger.Errorf("SaveVertexs err:%v", err)
		}
	}
	for k, v := range rel {
		err := SaveEdges(k, v)
		if err != nil {
			Logger.Errorf("SaveVertexs err:%v", err)
		}
	}
}

func getVerticesAndEdges(j *simplejson.Json, ent, relmap *map[string][]*simplejson.Json) {
	var mainField string
	var mainFieldKey string
	//事件
	__id := j.Get("_id").MustString()
	_, b := (*ent)["ent_dm"]
	if !b {
		(*ent)["ent_dm"] = []*simplejson.Json{}
	}
	j = j.Get("_source")
	j.Set("_key", __id)
	(*ent)["ent_dm"] = append((*ent)["ent_dm"], j)
	//顶点
	for _, field := range allFields {
		tv := j.Get(field).MustString()
		if tv != "" {
			field = "ent_" + field
			_, b := (*ent)[field]
			if !b {
				(*ent)[field] = []*simplejson.Json{}
			}
			fieldjson := simplejson.New()
			fieldjson.Set("_key", tv)
			fieldjson.Set("value", tv)
			(*ent)[field] = append((*ent)[field], fieldjson)
			//所有关系以此为准
			if mainField == "" {
				mainField = field
				mainFieldKey = fieldjson.Get("_key").MustString()
				continue
			}
			//建关系
			rel := simplejson.New()
			rel.Set("_key", fmt.Sprintf("%s_%s_%s", j.Get("_key").MustString(), fieldjson.Get("_key").MustString(), mainFieldKey))
			rel.Set("_from", mainField + "/" + mainFieldKey)
			rel.Set("_to", field + "/" + fieldjson.Get("_key").MustString())
			rel.Set("type", field)
			rel.Set("event", "ent_dm/" + j.Get("_key").MustString())
			//暂时所有都存在一起
			_, b = (*relmap)[EdgeCollectionName]
			if !b {
				(*relmap)[EdgeCollectionName] = []*simplejson.Json{}
			}
			(*relmap)[EdgeCollectionName] = append((*relmap)[EdgeCollectionName], rel)
		}
	}
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%X", h.Sum(nil))
}