package main

import "fmt"

const (
	SecondPerMinute = 60
	SecondPerHour = SecondPerMinute * 60
	SecondPerDay = SecondPerHour*24

)

func resolveTime(seconds int) (day int,hour int, minute int){
	day = seconds /SecondPerDay
	hour = seconds / SecondPerHour
	minute = seconds / SecondPerMinute

	return
}

func main(){
	fmt.Println(resolveTime(1000))
	_,hour,minute := resolveTime(12000)
	fmt.Println(hour,minute)
	day,_,_ :=resolveTime(909000)
	fmt.Println(day)
}