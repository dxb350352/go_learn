package main

import (
	"errors"
	"fmt"
	"github.com/robfig/config"
	"time"
	"github.com/sas/utils"
	"os"
	"path/filepath"
)

type ConfigItem struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Section string `json:"section"`
}

func (c ConfigItem)Create() error {
	if !Config.HasSection(c.Section) {
		Config.AddSection(c.Section)
	}
	if Config.HasOption(c.Section, c.Key) {
		Config.RemoveOption(c.Section, c.Key)
	}
	bl := Config.AddOption(c.Section, c.Key, c.Value)
	if !bl {
		return errors.New("保存配置失败")
	} else {
		return Write2File()
	}
}

func (c ConfigItem)Delete(section, key string) error {
	if Config.HasOption(section, key) {
		bl := Config.RemoveOption(section, key)
		if !bl {
			return errors.New("删除配置失败")
		} else {
			return Write2File()
		}
	}
	return nil
}

func (c ConfigItem)DeleteSection(section string) error {
	if Config.HasSection(c.Section) {
		bl := Config.RemoveSection(section)
		if !bl {
			return errors.New("删除配置失败")
		} else {
			return Write2File()
		}
	}
	return nil
}
//func (c ConfigItem)Update() error {
//	return nil
//}
//func (c ConfigItem)Get(key string) (*ConfigItem, error) {
//	return nil, nil
//}

func (c ConfigItem)Find(section string) (*[]ConfigItem, error) {
	var items []ConfigItem//结果集
	var defaltItem []interface{}//全局配置名称
	var sections []string//本次查询的所有分组
	//查询出默认项
	options, err := Config.Options("DEFAULT")
	if err != nil {
		return nil, err
	}
	//只查询一个分组的情况
	if section != "" {
		sections = append(sections, section)
	} else {
		sections = Config.Sections()
	}

	for _, o := range options {
		v, err := Config.String("", o)
		if err != nil {
			return nil, err
		}
		//只查询一个分组的情况
		defaltItem = append(defaltItem, o)
		if section == "" {
			items = append(items, ConfigItem{Section:"", Key:o, Value:v})
		}
	}

	for _, s := range sections {
		if s == "DEFAULT" {
			continue
		}
		options, err = Config.Options(s)
		if err != nil {
			return nil, err
		}
		for _, o := range options {
			if utils.Contains(&defaltItem, o) {
				continue
			}
			v, err := Config.String(s, o)
			if err != nil {
				return nil, err
			}
			items = append(items, ConfigItem{Section:s, Key:o, Value:v})
		}
	}
	return &items, nil
}

func Write2File() error {
	defer func() {
		Config, _ = config.ReadDefault("E:/GOPATH_GO/src/testgo/config/my.conf")
	}()
	return Config.WriteFile("E:/GOPATH_GO/src/testgo/config/my.conf", 0644, fmt.Sprintln("Modiffied at:", time.Now()))
}

var (
	Config *config.Config
)

func init() {
	pwd, _:=os.Getwd()
	fp := filepath.Join(pwd, "my.conf")
	c, err := config.Read(fp, config.DEFAULT_COMMENT, config.ALTERNATIVE_SEPARATOR, false, true)
	if err != nil {
		panic(err)
	}
	Config = c
	//show()
}

func main() {
	//	var item ConfigItem = ConfigItem{Section:"123", Key:"kddv", Value:"vv"}
	//	fmt.Println(item.Create())
	//	show()
	fmt.Println(Config.String("", "kdv"))
}

func show() {
	fmt.Println("...................................")
	var item ConfigItem
	items, _ := item.Find("")
	for i, v := range *items {
		fmt.Println(i, v.Section, v.Key, v.Value)
	}
	fmt.Println("...................................")
}