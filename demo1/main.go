package main

import "fmt"

func Test(person string)(work func() string){
	work = func() string {
		return (person+"is working")
	}
	return
}

func main(){
	p := Test("Sss")
	fmt.Println(p)
}