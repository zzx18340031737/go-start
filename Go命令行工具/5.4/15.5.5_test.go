package main

import "testing"

func TestLevel(t *testing.T) {
	inputs := []int{20, 65, 75, 95, 180}
	expected := []string{"不及格", "及格", "良", "优秀", "数据出错"}
	for i := 0; i < len(inputs); i++ {
		ret := Level(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %s,the actual %s", inputs[i], expected[i], ret)
		}
	}
}
