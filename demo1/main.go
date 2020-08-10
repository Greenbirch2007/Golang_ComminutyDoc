package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do"
	var strSlice= strings.Split(str," ")
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)

 }


