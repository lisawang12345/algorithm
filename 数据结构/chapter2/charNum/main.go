package main //第三题输入一行字符，分别统计出其英文字母、空格、数字和其它字符的个数。

import "fmt"

func main() {
    var a string
    numNum := 0
    letterNum := 0
    spaceNum := 0
    othersNum := 0
    fmt.Println("please input a string")
    fmt.Scanf("%s", &a)
    for _, value := range a {
        if value <= 'Z' && value >= 'A' || value <= 'z' && value >= 'a' {
            letterNum++
        } else if value <= '9' && value >= 0 {
            numNum++
        } else if value == ' ' {
            spaceNum++
        } else {
            othersNum++
        }
    }
    fmt.Println("空格字符的个数为", spaceNum)
    fmt.Println("数字字符的个数为", numNum)
    fmt.Println("字母字符的个数为", letterNum)
    fmt.Println("其他字符的个数为", othersNum)
}
