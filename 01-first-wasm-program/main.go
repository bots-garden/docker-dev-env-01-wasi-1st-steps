package main

// ✋ make this function callable from host
//export add
func add(x int, y int) int {
	return x + y
}

// 👋 this one does not work
//export hello
func hello(name string) string {
	return "hello " + name
}

func main() {}
