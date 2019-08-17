package main
import (
	"encoding/xml"
	"strings"
	"fmt"
	"github.com/sas/utils"
)

func main() {
	test("")
	fmt.Println(utils.EncryptPassword("123456",""))
}
func test(input string) {
	input = `<Person><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	var t xml.Token
	var err error

	inputReader := strings.NewReader(input)

	// 从文件读取，如可以如下：
	// content, err := ioutil.ReadFile("studygolang.xml")
	// decoder := xml.NewDecoder(bytes.NewBuffer(content))

	decoder := xml.NewDecoder(inputReader)
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		if err != nil {
			fmt.Println(err)
		}
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			fmt.Printf("Token of '%s' end\n", token.Name.Local)
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
		// ...
		}
	}
}
