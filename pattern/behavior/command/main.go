// Команда
// Интерфейс описывающий выполнение какой-то одной абстрактной команды - операции - функции
// Тип (команда) реализующий данный интерфейс принимает конкретный объект (бизнес модель)
// и производит конкретную работу с бизнес моделью через этот метод.
// Далее этот тип (команда) передается конкретному клиенту для работы.
//
// Назначение команды определяется в имени интерфейса
package main

import "fmt"

func main() {
	bank := Bank{}
	cPut := PutCommand{bank}
	cGet := GetCommand{bank}
	client := BankClient{cPut, cGet}
	client.GetMoney()
	// printed: money to the client
	client.PutMoney()
	// printed: money from the client
}

// Command type
type Command interface {
	Execute()
}

// PutCommand is ConcreteCommand
type PutCommand struct {
	bank Bank
}

// Execute command
func (pc PutCommand) Execute() {
	pc.bank.receiveMoney()
}

// GetCommand is ConcreteCommand
type GetCommand struct {
	bank Bank
}

// Execute command
func (gc GetCommand) Execute() {
	gc.bank.giveMoney()
}

// ////

// Bank is Receiver
type Bank struct{}

func (b Bank) giveMoney() {
	fmt.Println("money to the client")
}

func (b Bank) receiveMoney() {
	fmt.Println("money from the client")
}

// BankClient is Invoker
type BankClient struct {
	putCommand Command
	getCommand Command
}

// PutMoney runs the putCommand
func (bc BankClient) PutMoney() {
	bc.putCommand.Execute()
}

// GetMoney runs the getCommand
func (bc BankClient) GetMoney() {
	bc.getCommand.Execute()
}
