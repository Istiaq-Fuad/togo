package main

func main() {
	todos := Todos{}
	todos.Add("Buy Milk")
	todos.Add("Buy Bread")

	todos.Toggle(0)

	todos.Print()
}
