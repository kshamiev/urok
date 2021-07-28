// Более сложный пример, с использованием пула обработчиков для типовых задач
package main

import (
	"fmt"
	"sync"
)

// Task - описание интерфейса работы
type Task interface {
	Execute()
}

// Pool
type Pool struct {
	size  int            // количество одновременно работающих воркеров (лимит)
	tasks chan Task      // Канал задач - буферизированный, чтобы основная программа не блокировалась при постановке задач
	kill  chan struct{}  // канал отмены, для завершения работы воркеров
	wg    sync.WaitGroup // WaitGroup для контроля полного завершения работ всех задач и каналов
	mu    sync.Mutex     // мутекс для безопасного изменения
}

// Скроем внутреннее устройство за конструктором, пользователь может влиять только на размер пула
func NewPool(size int) *Pool {
	pool := &Pool{
		tasks: make(chan Task, 128),
		kill:  make(chan struct{}),
	}
	// Вызовем метод resize, чтобы установить соответствующий размер пула
	pool.Resize(size)
	return pool
}

func (p *Pool) Resize(n int) {
	// Захватываем лок, чтобы избежать одновременного изменения состояния
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < n {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
	for p.size > n {
		p.kill <- struct{}{}
	}
}

func (p *Pool) Wait() {
	close(p.tasks)
	p.wg.Wait()
}

func (p *Pool) TaskAdd(task Task) {
	p.tasks <- task
}

// Жизненный цикл воркера
func (p *Pool) worker() {
	defer func() {
		p.size--
		p.wg.Done()
	}()
	for {
		select {
		// Если есть задача, то ее нужно обработать
		// Блокируется пока канал не будет закрыт, либо не поступит новая задача
		case task, ok := <-p.tasks:
			if ok {
				task.Execute()
			} else {
				return
			}
			// Если пришел сигнал умирать, выходим
		case <-p.kill:
			return
		}
	}
}

type ExampleTask string

func (e ExampleTask) Execute() {
	fmt.Println("executing:", string(e))
}

func main() {
	pool := NewPool(5)
	pool.TaskAdd(ExampleTask("foo"))
	pool.TaskAdd(ExampleTask("bar"))
	pool.Resize(3)
	pool.Resize(6)
	for i := 0; i < 20; i++ {
		pool.TaskAdd(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}
	pool.Wait()
	fmt.Println(pool.size)
}
