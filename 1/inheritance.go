package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

func (h *Human) Introduce() {
	fmt.Printf("Меня зовут %s, мне %d лет и я %s\n", h.Name, h.Age, h.Gender)
}

type Action struct {
	Human
}

func (a *Action) Jump() {
	fmt.Println("Я прыгаю!")
}

func (a *Action) Run() {
	fmt.Println("Я бегаю!")
}

func main() {
	human := Human{
		Name:   "Андрей",
		Age:    25,
		Gender: "Мужчина",
	}

	action := Action{
		Human: human,
	}

	action.Introduce()
	action.Jump()
	action.Run()
}
