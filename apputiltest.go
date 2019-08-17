package main
import (
	"strings"
	"fmt"
	"github.com/sas/utils"
	"math/rand"
)

func main() {
	str := "/public/*,/bmitems/*,/bmtriggers/triggers"
	ss := strings.Split(str, ",")
	fmt.Println(utils.RegexpUrl(ss, "/bmtriggers/triggers//f"))
	fmt.Printf("%014d",rand.Intn(123))
	fmt.Printf("%014d",rand.Intn(123))
	fmt.Println(strings.Contains("dddd{dds}","{"))
}
