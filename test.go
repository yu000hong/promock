package main

import (
	"fmt"
	"reflect"
)

type person struct {
	age int
	name string
}

func test2() {
	var yuhong = person{
		age: 18,
		name: "yuhong",
	}
	var x = "ss"
	const c = 5 + 4i
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(yuhong))
	elements := []string{
		0:     "zero",
		1:     "one",
		4: 	"two",
		2:     "also two",
	}
	fmt.Println(reflect.TypeOf(elements))
	fmt.Println(elements[3])

	xx := person{age:18}
	fmt.Println(reflect.TypeOf(xx))
	yy := person{name:"yh"}
	fmt.Println(reflect.TypeOf(yy))
}

func test1(){
	// var (
	// 	x =10
	// 	y = ""
	// )
	// var x=10, y=""
	// var (
	// 	x int = 10
	// 	y string = ""
	// )
	var (
		x = 10
		y = "hello"
		z int64
		age int = 18
		f func(int, string) int
	)
	age, name := 20, "yuhong"
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(f))
}

func test3(){
	var s map[string]int
	s = make(map[string]int)
	s["yu000hong"] = 99
	fmt.Println(len(s))
	
	var p *int
	p = new(int)
	var sli *[3]int
	sli = new([3]int)
	fmt.Println(*p)
	fmt.Println(sli[0])
	fmt.Println(sli[1])
	fmt.Println(sli[2])
}

func test4(){
	s := make([]int, 0)
	fmt.Println(reflect.TypeOf(s))
	m := make(map[string]int, 16)
	fmt.Println(reflect.TypeOf(m))

	var f func(int, string) int
	f = func(a int, b string) int {
		return 10
	}
	fmt.Println(reflect.TypeOf(f))
}

func test5(){
	c := make(chan int)
	go func(){
		c <- 10
		c <- 0
		close(c)
	}()
	t1 := <- c
	println(t1)
	t2 := <- c
	println(t2)
	t3 := <- c
	println(t3)
	t4 := <- c
	println(t4)

}

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	<-c
}

func test6(){
	// i := 0
	// var i int
	while := 3
	println(while)
	i := 0
	for  ; ;  {
		i++
		if i > 3 { break }
		println(i)
	}
}

func main() {
	test6()
}