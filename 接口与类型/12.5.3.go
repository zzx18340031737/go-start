package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main(){
	s := make([]interface{},3)
	s[0] = 1
	s[1] = "str"
	s[2] = Person{"张三",20}
	for index,data := range s {
		if value,ok := data.(int); ok == true {
			fmt.Printf("s[%d] Type = int,Value = %d\n",index,value)
		}
		if value,ok := data.(string); ok == true {
			fmt.Printf("s[%d] Type = string,Value = %s\n",index,value)
		}
		if value,ok := data.(Person); ok == true {
			fmt.Printf("s[%d] Type = Person,Person.Name = %v,Person.Age = %d\n",index,value.Name,value.Age)
		}
	}
}