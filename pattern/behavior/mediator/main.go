// Посредник (реализуется через интерфейс)
// Это тип который является посредником между несколькими объектами
// Его задача производить синхронизацию между ними
// Для этого он содержит все объекты которые нужно синхронизировать
// И имеет метод который принимает объект с которым необходимо синхронизироваться
//
// Для удобства посредник может быть привязан к объекту через свойство

package main

import "fmt"

func main() {
	mediator := &SyncMediator{[]*Switcher{}}
	switcher1 := NewSwitcher(mediator)
	switcher2 := NewSwitcher(mediator)
	switcher3 := NewSwitcher(mediator)

	switcher1.State = true
	fmt.Println(switcher2.State, switcher3.State)

	mediator.Sync(switcher1)
	fmt.Println(switcher2.State, switcher3.State)
}

// Mediator defines an interface for communicating with Colleague objects
type Mediator interface {
	Sync(switcher *Switcher)
	Add(switcher *Switcher)
}

// SyncMediator is ConcreteMediator
type SyncMediator struct {
	switchers []*Switcher
}

// Sync synchronizes the state of all Colleague objects
func (sm *SyncMediator) Sync(switcher *Switcher) {
	for _, curSwitcher := range sm.switchers {
		curSwitcher.State = switcher.State
	}
}

// Add append Colleague to the Mediator list
func (sm *SyncMediator) Add(switcher *Switcher) {
	sm.switchers = append(sm.switchers, switcher)
}

// ////

// Switcher is Colleague
type Switcher struct {
	State bool
}

// NewSwitcher creates a new Switcher
func NewSwitcher(mediator Mediator) *Switcher {
	switcher := &Switcher{false}
	mediator.Add(switcher)
	return switcher
}
