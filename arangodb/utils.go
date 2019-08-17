package main

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"context"
	"strings"
	"fmt"
	"log"
	"github.com/bitly/go-simplejson"
	"crypto/md5"
	"io"
)

var Arangodb driver.Database
var EdgeConnection driver.Collection

type MyDocument struct {
	Key     string `json:"_key"`
	Name    string `json:"name"`
	Counter int64 `json:"counter"`
	Value   string `json:"value"`
}

var from = []string{
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

func main() {
	email := "dxbfdsfs@df.com"
	phone := "13564789548"
	key := "001_100_1"
	initArangodb()
	ctx := context.Background()
	EdgeConnection = ensureEdgeCollection(ctx, "graph", from, from)

	ent_md, _ := GetCollection("ent_md")
	md := createDocument(ctx, ent_md, map[string]interface{}{"_key":Md5(key), "email": email, "phone": phone})
	ent_person_email, _ := GetCollection("ent_person_email")
	person_email := createDocument(ctx, ent_person_email, map[string]interface{}{"_key":Md5(email), "value": email})
	ent_person_phone, _ := GetCollection("ent_person_phone")
	person_phone := createDocument(ctx, ent_person_phone, map[string]interface{}{"_key":Md5(phone), "value": phone})

	fmt.Println(md.ID, person_email.ID, person_phone.ID, "...............")
	var arr[]*simplejson.Json
	rel := simplejson.New()
	rel.Set("_from", "ent_person_email/" + Md5(email))
	rel.Set("_to", "ent_person_loginname/" + Md5(phone))
	rel.Set("type", "ent_person_loginname")
	rel.Set("event", "ent_dm/" + Md5(key))
	arr = append(arr, rel)
	fmt.Println(EdgeConnection.CreateDocuments(driver.WithSilent(ctx), arr))
}

func createDocument(ctx context.Context, col driver.Collection, document interface{}) driver.DocumentMeta {
	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		fmt.Println(err)
	}
	return meta
}

func ensureGraph(ctx context.Context, name string) driver.Graph {
	g, err := Arangodb.Graph(ctx, name)
	if driver.IsNotFound(err) {
		g, err = Arangodb.CreateGraph(ctx, name, nil)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}
	return g
}
func ensureEdgeCollection(ctx context.Context, name string, from, to []string) driver.Collection {
	g := ensureGraph(ctx, name + "__graph")
	ec, _, err := g.EdgeCollection(ctx, name)
	if driver.IsNotFound(err) {
		ec, err := g.CreateEdgeCollection(ctx, name, driver.VertexConstraints{From: from, To: to})
		if err != nil {
			log.Fatal(err)
		}
		return ec
	} else if err != nil {
		log.Fatal(err)
	}
	return ec
}

func testDocument() {
	ctx := context.Background()
	err := initArangodb()
	if err != nil {
		log.Fatal(err)
	}
	table, err := GetCollection("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(table.Count(ctx))
	doc := MyDocument{
		Key:"001_100_1",
		Name: "jan",
		Counter: 23,
		Value: "dddd",
	}
	meta, err := table.CreateDocument(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)
}

func initArangodb() error {
	arahosts := "http://192.168.130.201:8529"
	arauser := "root"
	arapassword := "123456"

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
	Arangodb, err = client.Database(ctx, "test")
	if err != nil {
		Arangodb, err = client.CreateDatabase(ctx, "test", nil)
	}
	return err
}
//获取表
func GetCollection(name string) (driver.Collection, error) {
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
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%X", h.Sum(nil))
}