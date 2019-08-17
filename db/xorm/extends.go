package main
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Parent struct {
	Id   string
	Name string
}

func (Parent) TableName() string {
	return "t_parent"
}

type Child struct {
	Parent `xorm:"extends"`
	Age int
}

func (Child) TableName() string {
	return "t_child"
}

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "localhost:3306", "szga") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	Engine.ShowSQL(true)
	err=Engine.Sync2(
		new(Parent),
		new(Child),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	var p Parent
	p.Id = "1"
	p.Name = "pn"
	var c Child
	c.Parent = p
	c.Age = 10
	affected, err := Engine.Insert(&p, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(affected)
}