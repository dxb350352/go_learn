package main
import (
	"bytes"
	"github.com/BurntSushi/toml"
	"fmt"
	"io/ioutil"
)

type Name struct {
	Name    []string
	Version string
	Inners  []Inner
}

type Inner struct {
	Id int
	Go string
}
func main() {
	c := Name{Name:[]string{"ad", "db", "ee"}, Version:"2.2"}
	var in []Inner
	in = append(in, Inner{Id:1, Go:"go1"})
	in = append(in, Inner{Id:2, Go:"go2"})
	in = append(in, Inner{Id:3, Go:"go3"})
	c.Inners = in
	WriteFile("e://config.toml", c)
	ReadFile("e://config.toml")
}

func ReadFile(path string) Name {
	var n Name
	metaData, err := toml.DecodeFile(path, &n)
	if err != nil {
		fmt.Println(err)
		return n
	}
	fmt.Println(n, n.Name, n.Version, metaData.Keys())
	fmt.Println(n.Inners[0].Go)
	return n
}
func WriteFile(filename string, config  interface{}) error {
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	err := enc.Encode(config)
	if err != nil {
		return fmt.Errorf("encode toml config error")
	}
	s := buf.String()
	fmt.Println(s, buf, "............", config)
	err = ioutil.WriteFile(filename, []byte(s), 0644)
	if err != nil {
		return fmt.Errorf("write %s error", filename)
	}
	return nil
}