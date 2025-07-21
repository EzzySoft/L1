package l1

import (
	"fmt"
)

// Структура Human
type Human struct {
	Name string
	Age  int
}

// Метод Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s!\n", h.Name)
}

// Встраивание Human в Action (композиция)
type Action struct {
	Human
	Type string
}

// Метод только для Action
func (a *Action) DoSomething() {
	fmt.Println("Выполняю действие:", a.Type)
}

func main() {
	// Создаём экземпляр Action
	a := Action{
		Human: Human{Name: "Данила", Age: 21},
		Type:  "Учить Golang",
	}
	// Вызываем метод Human через Action
	a.SayHello()
	// Вызываем метод Action
	a.DoSomething()
}
