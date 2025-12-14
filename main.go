package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Eat")
	todos.add("Sleep")
	todos.add("Goon")
	todos.add("Repeat")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(2)
	fmt.Printf("%+v", todos)
}