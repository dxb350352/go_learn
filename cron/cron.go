package main
import (
	"github.com/jakecoffman/cron"
	"fmt"
	"time"
)

type Job struct {
	Id       string
	Schedule string
}

func (j Job) Run() {
	fmt.Println(j.Id, "is running.....")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	job := Job{Id:"1", Schedule:"0/10 * * * *fd"}
	Cron := cron.New()
	Cron.Start()
	fmt.Println("start.......................")
	Cron.AddJob(job.Schedule, job, job.Id)
	fmt.Println("add.......................")
	time.Sleep(time.Hour)
}
