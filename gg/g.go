package main

import "fmt"

type ChipType int

const (
	None ChipType = iota
	CPU // Z中央处理器
	GPU //图形处理器
)


func (c ChipType) String() string{
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"

	}

	return "N/A"
}

func main(){
	fmt.Printf("%s %d", CPU, CPU)
}