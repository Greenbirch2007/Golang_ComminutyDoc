package main

import "fmt"

/*

指针（指针的核心是通过内存指针的操作，把值类型改成了引用类型）

&(取地址符)
*(根据地址取值，也即是*&)
*/

func fn1(x int){
	x = 10

}
func fn2(x *int){
	*x = 40
}
func main(){
	//必须用make来分配存储空间
	//var userinfo = make(map[string]string)
	//
	//userinfo["username"] ="ss"
	//fmt.Println(userinfo)
	//var a = make([]int,4)
	//a[0] =6
	//fmt.Println()

		//用new函数分配你村 ,new主要是真毒指针乐行
	//var a *int //指针也是引用类型
	//*a = 100
	//fmt.Println()
	//a := new(int) //a是一个指针变量  类型是*int的指针类型 值是0
	//b := new(bool) //b是一个指针变量 类型是*bool的指针类型 值为false
	//// new 不太常用，使用new之后得到一个指针类型，并且该指针对应的值为该类型的零值
	//
	//fmt.Printf("%T %T ",a,b)
	//fmt.Println(*a,*b)
	//var a *int
	//a  = new(int)
	//*a= 1000
	//fmt.Println(*a)

	/*
	make,区别于New，make	只用于slice,map,channel的内存创建
	而且它返回的类型就是这三个类型本身，而不是它们的指针类型，
	因为这三种类型就是引用类型，所以没有必要返回它们的指针了

	func make(t Type, size ...IntegerType) Type

	make函数是无可替代的，在使用slice,map,channel是，都需
	使用make进行初始化，然后才可以对他们进行操作，

	 */
	var f = new(bool)
	fmt.Println(*f)
}