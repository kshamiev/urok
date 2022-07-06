package workerdatabus

import (
	"sync"
)

type Task interface {
	Execute()
}

type Pool struct {
	tasks       chan Task      // Канал задач - буф., чтобы родительская программа не блокировалась
	workerLimit int            // Количество одновременно работающих обработчиков
	kill        chan struct{}  // Канал отмены, для завершения работы обработчика
	wg          sync.WaitGroup // WaitGroup для контроля полного завершения работ всех задач и каналов
}

func NewPool(sizeChanel, workerLimit int) *Pool {
	p := &Pool{
		tasks:       make(chan Task, sizeChanel),
		workerLimit: workerLimit,
		kill:        make(chan struct{}),
	}
	for i := 0; i < p.workerLimit; i++ {
		p.wg.Add(1)
		go p.worker()
	}
	return p
}

func (p *Pool) TaskAdd(task Task) {
	p.tasks <- task
}

func (p *Pool) Wait() {
	close(p.tasks)
	p.wg.Wait()
}

// Жизненный цикл обработчика
func (p *Pool) worker() {
	defer func() {
		// TODO доработать обработку паники
		p.wg.Done()
	}()

	for task := range p.tasks {
		task.Execute()
		select {
		default:
		case <-p.kill:
			return
		}
	}
}
