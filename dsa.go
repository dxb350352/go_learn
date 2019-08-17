package main
import (
	"fmt"
	"github.com/sas/szga/app/models"
	"time"
)

func main() {
	m := make(map[string]int64)
	fmt.Println(m["fds"])
	fmt.Println(m["f2ds"] + 1)
	termdb := new(models.SzgaTerminalHeartbeat)
	termdb.HeartbeatTime = time.Now()
	his := new(models.SzgaTerminalHeartbeatHis)
	his.NewHis(termdb, true)
	fmt.Println(termdb)
	fmt.Println(his)
}