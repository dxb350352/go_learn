package workpool

type Runner interface {
	Run(interface{}) (interface{}, error)
}

//定义一个有趣的函数类型，这个类型实现了接口 代理模式
type RunnerFunc func(interface{}) (interface{}, error)

func (r RunnerFunc) Run(avar interface{}) (interface{}, error) {
	return r(avar)
}
