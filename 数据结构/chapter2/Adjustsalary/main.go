package main

import "fmt"

func add(n int, salary int) (sum int) {
    if n >= 0 && n < 1 {
        sum = salary + 200
    } else if n >= 1 && n < 3 {
        sum = salary + 500
    } else if n >= 3 && n < 5 {
        sum = salary + 1000
    } else if n >= 5 && n < 10 {
        sum = salary + 2500
    } else if n >= 10 && n < 15 {
        sum = salary + 5000
    }
    fmt.Printf("您目前工作了%d年，基本工资为%d元,应涨工资%d元,涨后工资%d元", n, salary, sum-salary, sum)
    return sum
}
func main() {
    add(10, 3000)
}