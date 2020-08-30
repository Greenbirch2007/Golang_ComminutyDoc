package main

import "fmt"

/*
golang中空接口和类型断言使用细节
注意，把切片赋值给空接口后，没法直接获取切片里面的值
同理，把结构体赋值给空接口后，也无法直接获取结构体里面的属性


解决办法：类型断言
切片--空接口的处理
	hobby2,_ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[1])

结构体---空接口的处理
	address2,_:= userinfo["address"].(Address)
	fmt.Println(address2.Name)

*/

type Address struct {
	Name string
	Phone int
}

func main(){
	var userinfo=make(map[string]interface{})
	userinfo["username"]="张赞"
	userinfo["age"]=90
	userinfo["hobby"]=[]string{"吃饭","睡觉"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo[""])//空接口不支持索引


	var address = Address{
		"李四",
		13111,
	}
	//fmt.Println(address.Name) //李四
	userinfo["address"] =address

	hobby2,_ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[1])

	address2,_:= userinfo["address"].(Address)
	fmt.Println(address2.Name)


}