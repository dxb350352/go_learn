package main
import "fmt"

type A struct {
	AName string
}

func (a A)Run() {
	fmt.Println(a.AName, "a.Run()")
}

type Job interface {
	Run()
}

type B struct {
	Job   Job
	BName string
}

func main() {
	a := A{AName:"aname"}
	b := B{BName:"bname", Job:&a}
	fmt.Println(b.Job)
}
