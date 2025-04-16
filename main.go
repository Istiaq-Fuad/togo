package main

func main() {
	todos := Todos{}
    storage := NewStorage[Todos]("todos.json")
    storage.Load(&todos)
    todos.Add("Buy Milk")
    todos.Add("Buy Bread")
    todos.Toggle(0)
    todos.Print()
    storage.Save(todos)
}
