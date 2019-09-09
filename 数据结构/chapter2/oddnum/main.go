//使用for循环,求出1-100之间的奇数之和.
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i <= 100; i++ {
        if getremainder(i) == false {

            sum += i
        }
    }

    fmt.Println("sum=", sum)
}
func getremainder(n int) bool {
    if n%2 == 0 {
        return true
    }
    return false
}
