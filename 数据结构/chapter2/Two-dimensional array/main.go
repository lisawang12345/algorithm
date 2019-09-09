/*a = [3][4]int{如何翻转这个二维数组
{0, 1, 2, 3} ,
{4, 5, 6, 7} ,
{8, 9, 10, 11},
}*/
package main

import "fmt"

func transpose(a [][]int) [][]int {
    row := len(a)
    if row == 0 {
        return nil
    }

    col := len(a[0])
    arr := make([][]int, col, col)
    for i := 0; i <= row; i++ {
        arr[i] = make([]int, row, row)
    }
    temp := 0
    for k := 0; k < row; k++ {
        for y := 0; y < col; y++ {
            a[k][y] = temp
            for m := 0; m < col; m++ {
                for n := 0; n < row; n++ {
                    arr[m][n] = a[k][y]
                }
            }
        }

    }
    fmt.Println(arr)
    return arr
}
func main() {
    a := [][]int{
        {0, 1, 2, 3},
        {4, 5, 6, 7},
        {8, 9, 10, 11},
    }
    transpose(a)

}
