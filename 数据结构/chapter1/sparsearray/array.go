package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子
	//输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println(v)
	}
	//这里遍历了两次，因为是二位数组，有v和v2
	//转成稀疏数组，思路：遍历chessMap，有效元素的值放到对应的切片中

	var sparseArr []ValNode
	//标准的稀疏数组应该还有记录元素二维数组的规模（行和列）
	//创建一个valnode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个valNode值结点
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("当前的稀疏数组是")
	for i, valNode := range sparseArr {
		fmt.Printf("%d %d %d\n", i, valNode.row, valNode.col)
	}
}
