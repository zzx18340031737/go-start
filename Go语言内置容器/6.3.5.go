package main

import "fmt"

func main() {
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)
	studentScoreMap["Tom"] = 80
	studentScoreMap["Ben"] = 85
	studentScoreMap["Peter"] = 90
	delete(studentScoreMap, "Tom")
	fmt.Println(studentScoreMap)
}
