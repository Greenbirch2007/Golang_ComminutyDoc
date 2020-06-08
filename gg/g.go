package main

import (
	"container/list"
	"fmt"
)

func main()  {
	l := list.New()
	l.PushBack("cannon")
	l.PushFront("99")

	for i := l.Front(); i!= nil; i=i.Next(){
		fmt.Println(i.Value)
	}
}