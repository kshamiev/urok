// Наблюдатель
// Тип реализующий наблюдателя вызывается (его метод) при выполнении какой-либо бизнес функции.
// Наблюдение за выполнением которой нужно установить.
// Возможно использование типа агрегатора (TestSubject) для управления и вызова используемых обсерверов.
// Этот тип управления встраивается в нужную бизнес модель
//
// Назначение обсервера определяется названием метода интерфейса
package main

import "fmt"

func main() {
	observer1 := TextObserver{"IObserver #1"}
	observer2 := TextObserver{"IObserver #2"}

	textEdit := TextEdit{}
	textEdit.Attach(observer1)
	textEdit.Attach(observer2)

	textEdit.SetText("test text")
	// printed:
	// IObserver #1: test text
	// IObserver #2: test text
}

type observer interface {
	update(state string)
}

// TextObserver is ConcreteObserver
type TextObserver struct {
	_name string
}

func (t TextObserver) update(state string) {
	fmt.Println(t._name + ": " + state)
}

// TestSubject is Subject
type TestSubject struct {
	_observers []observer
}

// Attach adds an observer
func (ts *TestSubject) Attach(observer observer) {
	ts._observers = append(ts._observers, observer)
}

// Detach removes an observer
func (ts *TestSubject) Detach(observer observer) {
	index := 0
	for i := range ts._observers {
		if ts._observers[i] == observer {
			index = i
			break
		}
	}
	ts._observers = append(ts._observers[0:index], ts._observers[index+1:]...)
}

func (ts TestSubject) notify(state string) {
	for _, observer := range ts._observers {
		observer.update(state)
	}
}

// TextEdit is ConcreteSubject
type TextEdit struct {
	TestSubject
	Text string
}

// SetText changes the text and informs observers
func (te TextEdit) SetText(text string) {
	te.Text = text
	te.notify(text)
	// te.TestSubject.notify(text)
}
