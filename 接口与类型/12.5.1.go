package main

import "fmt"

func checkTpye(t interface{},ok bool){
	if ok {
		fmt.Println("断言成功！")
	} else {
		fmt.Println("断言失败！")

	}
	fmt.Printf("变量t的类型 = %T,值 = %v \n",t,t)
}

func main(){
	var X interface{} = 1

	fmt.Println("第一次断言：")
	t0,ok := X.(string)
	checkTpye(t0,ok)

	fmt.Println("第二次断言：")
	t1,ok := X.(float64)
	checkTpye(t1,ok)

	fmt.Println("第三次断言：")
	t2,ok := X.(int)
	checkTpye(t2,ok)
}

