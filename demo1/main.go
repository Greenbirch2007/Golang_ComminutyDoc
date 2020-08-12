package main

import (
	"errors"
	"fmt"
)
// 模拟一个读取文件的方法

func readFile(fileName string)error{
	if fileName == "main.go"{
		return nil
	}else{
		return errors.New("读取文件失败")
	}
}

func myFn(){
	defer func() {
		err := recover()
		if err!=nil{
			fmt.Println("给管理员发送邮件")

		}
	}()
	err :=readFile("xx.go")
	if  err !=nil{
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("继续执行")

}

