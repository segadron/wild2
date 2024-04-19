package main

import "fmt"

// Интерфейс, который мы хотим адаптировать
type EnglishSpeaker interface {
	SpeakEnglish() string
}

// Структура, имплементирующая интерфейс EnglishSpeaker
type EnglishPerson struct{}

func (e *EnglishPerson) SpeakEnglish() string {
	return "Hello, how are you?"
}

// Интерфейс адаптера
type RussianSpeaker interface {
	SpeakRussian() string
}

// Структура, которую мы адаптируем к интерфейсу RussianSpeaker с помощью адаптера
type Adapter struct {
	englishPerson EnglishSpeaker
}

func (a *Adapter) SpeakRussian() string {
	return "Привет, как дела? (translated: " + a.englishPerson.SpeakEnglish() + ")"
}

func main() {
	english := &EnglishPerson{}

	adapter := &Adapter{
		englishPerson: english,
	}

	fmt.Println(adapter.SpeakRussian())
}
