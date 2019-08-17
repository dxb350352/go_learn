package main

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"context"
	"strings"
	"github.com/bitly/go-simplejson"
	"log"
	"fmt"
	"crypto/md5"
	"io"
)

var from = []string{
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
var EdgeCollectionName string = "graph_edge_collection"

func main() {
	initdb()
	email := "dxbfdsfs@df.com"
	phone := "13564789548"
	key := "001_100_1"
	ctx := context.Background()
	ent_md, _ := GetCollection1("ent_md")
	md := createDocument1(ctx, ent_md, map[string]interface{}{"_key":Md51(key), "email": email, "phone": phone})
	ent_person_email, _ := GetCollection1("ent_person_email")
	person_email := createDocument1(ctx, ent_person_email, map[string]interface{}{"_key":Md51(email), "value": email})
	ent_person_phone, _ := GetCollection1("ent_person_phone")
	person_phone := createDocument1(ctx, ent_person_phone, map[string]interface{}{"_key":Md51(phone), "value": phone})

	fmt.Println(md.ID, person_email.ID, person_phone.ID, "...............")
	var arr[]*simplejson.Json
	rel := simplejson.New()
	rel.Set("_from", "ent_person_email/" + Md51(email))
	rel.Set("_to", "ent_person_phone/" + Md51(phone))
	rel.Set("event", "ent_dm/" + Md51(key))
	//rel.Set("_from", person_email.ID)
	//rel.Set("_to", person_phone.ID)
	//rel.Set("event", md.ID)

	rel.Set("type", "ent_person_loginname")
	arr = append(arr, rel)
	fmt.Println(rel)
	SaveEdges(EdgeCollectionName, arr)
}

func initdb() {
	arahosts := "http://192.168.130.201:8529"
	arauser := "root"
	arapassword := "123456"
	aradbname := "test"

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: strings.Split(arahosts, ","),
	})
	if err != nil {
		log.Fatal(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
		Authentication: driver.BasicAuthentication(arauser, arapassword),
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	Arangodb, err = client.Database(ctx, aradbname)
	if err != nil {
		Arangodb, err = client.CreateDatabase(ctx, aradbname, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	Graph, err = Arangodb.Graph(ctx, GraphName)
	if err != nil {
		Graph, err = Arangodb.CreateGraph(ctx, GraphName, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ensureEdgeCollectiona(ctx context.Context, name string) (driver.Collection, error) {
	ec, _, err := Graph.EdgeCollection(ctx, name)
	if err != nil {
		ec, err = Graph.CreateEdgeCollection(ctx, name, driver.VertexConstraints{From: from, To: from})
	}
	return ec, err
}
func SaveEdges(name string, arr []*simplejson.Json) error {
	ctx := context.Background()
	if len(arr) == 0 {
		log.Println("SaveEdges no data to save")
		return nil
	}
	ec, err := ensureEdgeCollectiona(ctx, name)
	if err != nil {
		return err
	}
	dc, errs, err := ec.CreateDocuments(ctx, arr)
	fmt.Println(err)
	fmt.Println(errs)
	fmt.Println(dc)
	return err
}

func GetCollection1(name string) (driver.Collection, error) {
	ctx := context.Background()
	b, err := Arangodb.CollectionExists(ctx, name);
	if err != nil {
		return nil, err
	}
	if !b {
		return Arangodb.CreateCollection(ctx, name, &driver.CreateCollectionOptions{})
	}
	return Arangodb.Collection(ctx, name)
}
func Md51(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%X", h.Sum(nil))
}
func createDocument1(ctx context.Context, col driver.Collection, document interface{}) driver.DocumentMeta {
	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		fmt.Println(err)
	}
	return meta
}