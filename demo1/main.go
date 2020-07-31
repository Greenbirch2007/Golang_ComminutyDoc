package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type MyC struct {
 	
 }

func (m MyC) Generate(rand *rand.Rand, size int) reflect.Value {
	panic("implement me")
}

func main(){
	var  n = 30
	if n>24{
		fmt.Println("成年人")
		goto xxx

	}
	fmt.Println("aaa")
	fmt.Println("bbb")
	xxx:
	fmt.Println("ccc")
	fmt.Println("ddd")


}