package main

import "fmt"

type Person interface {
	Speak()
}
type Student struct {
	Name string
}
func (s Student) Speak(){
	fmt.Println("My name is",s.Name)

}

func checkTpye(t interface{},ok bool){
	if ok {
		fmt.Println("断言成功！")
	} else {
		fmt.Println("断言失败！")

	}
	fmt.Printf("变量t的类型 = %T,值 = %v \n",t,t)
}

func main(){
	var Student interface{} = Student{"小明"}

	fmt.Println("第一次断言：")
	t0,ok := Student.(string)
	checkTpye(t0,ok)

	fmt.Println("第二次断言：")
	t1,ok := Student.(Person)
	checkTpye(t1,ok)
}
