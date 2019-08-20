/*接口类型是由一组方法定义的集合，接口类型的值可以存放实现这些方法的任何值
go 语言提供了另外一种数据类型即接口，它把所有具有共性的方法定义在一起，任何
其他类型只要是实现了这些方法就是实现了这个接口*/
package main

import (
	"fmt"
	"math"
)

/* Abser is ...（定义接口Abser，Abs是要实现的方法，返回值是float64类型的，Abs就是
所谓的所有的具有共性的方法）*/
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v

	//这里只返回一个值是因为函数Abs只有一个返回值
}

//MyFloat is...给float64定义一个方法
type MyFloat float64

//Abs is ...
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Vertex is... 定义一个结构体型的方法接收者
type Vertex struct {
	X, Y float64
}

// Abs is...
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//运行结果是1.4142135623730951
