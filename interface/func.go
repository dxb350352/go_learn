package main

import (
	"fmt"
)

type Status interface {
	GetStatus() bool
	SetStatus(b bool)
}

type Stop struct {
	Status bool
}

func (s *Stop)GetStatus() bool {
	return s.Status
}
func (s *Stop)SetStatus(b bool) {
	s.Status = b
}

func main() {
	var stop Stop
	TTest(&stop)
	stop.SetStatus(true)
	TTest(&stop)
}

func TTest(s Status) {
	fmt.Println(s.GetStatus())
}