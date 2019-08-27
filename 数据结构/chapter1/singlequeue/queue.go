package main
import (
	"fmt"
	"os"
	"errors"
)
//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array [5] int //数组--》模拟队列
	front int //表示指向队列首部
	rear int //表示指向队列的尾部
}
//添加数据到队列的思路
func (this *Queue) AddQueue(val int)(err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize - 1 {//重要重要：rear是队列尾部含队尾，
return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val
	return
}
//显示队列，找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是")
	//队首是不包含对首的元素的
for i:= this.front + 1; i <= this.rear; i++{
	fmt.Printf("array[%d]=%d\t",i,this.array[i])
}
//编写一个主函数测试
func main() {

	//先创建一个队列
	queue := &Queue {
		maxSize : 5
		front : -1
		rear : -1
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示结束队列")
		fmt.Scanln(&key)
		switch key {
		case "add" :
			fmt.Scanln("输入你要输入的队列数")
			fmt.Scanln(&val)
			err := Queue.AddQueue(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("加入队列OK")
				
			}
		case "get":
			fmt.Println("get")
		case "show":
			queue.showQueue()
		case "exit":
			os.Exit(0)
			
		}

}



