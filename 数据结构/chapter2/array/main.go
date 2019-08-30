package main

import "fmt"

//Insert expects a slice and value to insert, return the index of inserted valuefunc
func Insert(a []int, value int) []int {
	var index = 1
	var b = make([]int, index)
	copy(b, a[:index])
	b = append(b, value)
	b = append(b, a[index:]...)
	return a
}
func main() {
	var a = []int{1, 3, 5, 6, 7, 8, 9, 11}
	Insert(a, 2)
	fmt.Println(a)
}
