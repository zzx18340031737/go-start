package main

func main() {
	GoMap := make(map[int]int)
	for i := 0; i < 10000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)

	}
}

func readMap(Gomap map[int]int, key int) int {
	return Gomap[key]
}

func writeMap(Gomap map[int]int, key int, value int) {
	Gomap[key] = value
}
