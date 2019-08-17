package main
import "fmt"

func main() {
	fmt.Println("mytest", "--------------------------------")
	mytest()
	fmt.Println(pointer_test, "--------------------------------")
	pointer_test()
	fmt.Println(memery_test, "--------------------------------")
	memery_test()
}
func mytest() {
	var p *int
	fmt.Println(p, &p)
	var i int
	p = &i
	fmt.Println(p, &p, *p)
	*p = 8
	fmt.Println(p, &p, *p)
	m := [3]int{2, 3, 4}
	x := [3]*int{&m[0], &m[1], &m[2]}
	fmt.Println(x[0], &x[0], *x[0])
	var n [3]*int
	n = x
	fmt.Println(n[0], &n[0], *n[0])

}
func pointer_test() {
	//空指针，输出为nil
	var p *int
	fmt.Printf("p: %v\n", p)
	//指向局部变量，变量值初始为0
	var i int
	p = &i
	fmt.Printf("p: %v,%v\n", p, *p)
	//通过指针修改变量数值
	*p = 8
	fmt.Printf("p: %v,%v\n", p, *p)
	//数组的初始化及输出
	m := [3]int{3, 4, 5}
	fmt.Printf("m:%v--%v,%v,%v\n", m, m[0], m[1], m[2])
	//指针数组的初始化及输出
	//j, k, l := 3, 4, 5
	//x := [3]*int{&j, &k, &l}
	x := [3]*int{&m[0], &m[1], &m[2]}
	fmt.Printf("x:%v,%v,%v\n", x[0], x[1], x[2])
	fmt.Printf("*x:%v,%v,%v\n", *x[0], *x[1], *x[2])

	var n [3]*int
	n = x
	fmt.Printf("n:%v,%v,%v\n", n[0], n[1], n[2])
	//指向数组的指针，也即二级指针的使用
	y := []*[3]*int{&x}
	fmt.Printf("y:%v,%v\n", y, y[0])
	fmt.Printf("*y[0]:%v\n", *y[0])
	fmt.Printf("*y[][]:%v,%v,%v\n", *y[0][0], *y[0][1], *y[0][2])
	/*
	*X            ->    v
	[3]X        ->    [v0][v1][v2]
	[]X            ->    [v0]...[vi]
	[3]*X        ->    [p0][p1][p2]
					  |   |   |
					  j   k   l
	[]*[3]*X    ->    [n0]...[ni]
					  | ... |
					 [p0]->j
					 [p1]->k
					 [p2]->l
*/
}

type Student struct {
	name  string
	id    int
	score uint
}

func memery_test() {
	//new分配出来的数据是指针形式
	p := new(Student)
	p.name = "China"
	p.id = 63333
	p.score = 99
	fmt.Println(p, *p)
	//var定义的变量是数值形式
	var st Student
	st.name = "Chinese"
	st.id = 666333
	st.score = 100
	fmt.Println(st)
	//make分配slice、map和channel的空间，并且返回的不是指针
	var ptr *[]Student
	fmt.Println(ptr)     //ptr == nil
	ptr = new([]Student) //指向一个空的slice
	fmt.Println(ptr)
	*ptr = make([]Student, 3, 100)
	fmt.Println(ptr)
	stu := []Student{{"China", 3333, 66}, {"Chinese", 4444, 77}, {"Chince", 5555, 88}}
	fmt.Println(stu)
}
