package main

import (
	"fmt"
	"sort"
)

func main() {
	//长度是线段，容量是射线
	//var sliceA = make([]int,4,8)
	//sliceA[0]=10
	//sliceA[1]=10
	//sliceA[2]=10
	//fmt.Println(sliceA)
	//s1 := "big"
	//bytestr := []byte(s1)
	//bytestr[0] ='p'
	//fmt.Println(string(bytestr))
	// string----byte----string
	intL := []int{2,6,9,1,6,66}

	sort.Sort(sort.Reverse(sort.IntSlice(intL)))
	fmt.Println(intL)
}

