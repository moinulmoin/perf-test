package main

func main() {
	var array [10000]int
	for i := 0; i < 10000; i++ {
		for j := 0; j < 100000; j++ {
			array[i] = array[i] + j
		}
	}
}
