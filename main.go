package main

func main() {
	todos := Todos{}
	todos.add("Get harue Nightreign")
	todos.add("Convince kippie linux >>>> windows")
	todos.toggle(0)
	todos.print()
}