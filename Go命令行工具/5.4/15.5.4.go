package main

func Level(point int) string {
	switch {
	case point < 60:
		return "不及格"
	case point < 70:
		return "及格"
	case point < 80:
		return "良"
	case point < 100:
		return "优秀"
	default:
		return "数据出错"

	}
}
