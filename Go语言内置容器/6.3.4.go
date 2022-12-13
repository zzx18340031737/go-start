package main

import "fmt"

func main() {
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)
	studentScoreMap["Tom"] = 80
	studentScoreMap["Ben"] = 85
	studentScoreMap["Peter"] = 90
	//for k,v :=range studentScoreMap{
	//	fmt.Println(k,v)
	//}
	//for k :=range studentScoreMap{
	//	fmt.Println(k)
	//}
	for _, v := range studentScoreMap {
		fmt.Println(v)
	}
}
