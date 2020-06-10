package main

import "fmt"

func main(){
	c := make(chan int)
	go func(){
		c <-1
		c <-2
		close(c)

		c <-6


	}()
	for v:= range c{
		fmt.Println(
			v)
	}

}