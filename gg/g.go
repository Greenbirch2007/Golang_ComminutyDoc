package main

import (
	"fmt"
)

func Sort(data Interface){
	for i:=1;i< data.Len();i++{
		for j:=i;j>0 && data.Less(j,i-1);j--{
			data.Swap(j,j-1)
		}
	}
}

type Interface interface {
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}

type testType struct {
	a int
	b string
}

func (t *testType)String(){
	return fmt.Sprint(t.a) +" "+t.b
}
func main()  {
	t:= &testType{88,"sunset ss"}
	fmt.Println(t)
}

