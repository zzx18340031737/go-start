package main

import "fmt"

func main() {
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)
	studentScoreMap["Tom"] = 80
	studentScoreMap["Ben"] = 85
	studentScoreMap["Peter"] = 90
	//fmt.Println(cap(studentScoreMap)
	fmt.Println("map长度为：", len(studentScoreMap))
	fmt.Println(studentScoreMap)

}
