package main

import "fmt"

type myInt int

func add (x,y int)int{
	return x+y
}
func sub (x,y int)int{
	return x-y
}

type calcType func(int,int)int

func do(o string) calcType{
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x,y int)int{
			return x*y
		}
	default:
		return nil
	}
}


func main() {
	//匿名函数 匿名自执行函数
	func(){
		fmt.Println("这是匿名函数")
	}()

	f := func(x,y int) int{
		a := x*y
		return a

	}(6,9)
	fmt.Println(f)

	//匿名函数自执行函数接收参数
	func(x,y int){
		fmt.Println(x*y)

	}(9,9)


 }


