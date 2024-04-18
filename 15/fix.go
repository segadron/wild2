package main

func createHugeString(length int) string {
	return string(make([]byte, length))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	println(justString) // Используем justString в функции main()
}
