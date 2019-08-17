package workpool

type Work struct {
	Runner Runner
	Args   interface{}
}

func NewWork(runner Runner, args interface{}) *Work {
	return &Work{runner, args}
}
