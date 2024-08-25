package main

import "fmt"

func main() {
	const value int = 100
	var intVal int = 20
	var floatVal float32 = float32(intVal)
	fmt.Println(intVal, floatVal)
	intVal = value
	floatVal = float32(value)
	fmt.Println(intVal, floatVal)

	var b int = 1
	fmt.Println(b)
	b = value
	fmt.Println(b)
}
