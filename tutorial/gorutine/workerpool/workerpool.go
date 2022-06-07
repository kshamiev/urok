package workerpool

import (
	"sync"
	"time"
)

type Task interface {
	Execute()
}

type Pool struct {
	tasks      chan Task      // Канал задач - буферизированный, чтобы порождающая программа не блокировалась при постановке задач
	workerMin  int            // Минимальное количество одновременно работающих обработчиков
	workerMax  int            // Максимальное количество одновременно работающих обработчиков
	workerSelf int            // Текущее количество одновременно работающих обработчиков
	kill       chan struct{}  // Канал отмены, для завершения работы обработчика
	wg         sync.WaitGroup // WaitGroup для контроля полного завершения работ всех задач и каналов
}

func NewPool(sizeChanel, workerMin, workerMax int) *Pool {
	p := &Pool{
		tasks:     make(chan Task, sizeChanel),
		workerMin: workerMin,
		workerMax: workerMax,
		kill:      make(chan struct{}),
	}
	for p.workerSelf < p.workerMin {
		p.workerSelf++
		p.wg.Add(1)
		go p.worker()
	}
	p.wg.Add(1)
	go p.control()
	return p
}

func (p *Pool) TaskAdd(task Task) {
	p.tasks <- task
}

func (p *Pool) Wait() {
	p.workerMin = 0
	close(p.tasks)
	p.wg.Wait()
}

func (p *Pool) control() {
	for {
		select {
		case <-time.After(time.Millisecond):
			check := len(p.tasks)
			if check == 0 {
				break
			}
			check = cap(p.tasks) / check
			switch {
			case check < 10: // канал заполнен задачами больше чем на 10%
				if p.workerSelf < p.workerMax {
					p.workerSelf++
					p.wg.Add(1)
					go p.worker()
				}
				if check == 1 { // канал заполнен задачами больше чем на половину // TODO метрика предупреждения
				}
			case check > 50: // канал заполнен меньше чем на 30%
				if p.workerSelf > p.workerMin {
					p.kill <- struct{}{}
				}
			}
		}
		if p.workerMin == 0 || p.workerMax == 0 {
			break
		}
	}
	p.wg.Done()
}

// Жизненный цикл обработчика
func (p *Pool) worker() {
	defer func() {
		// TODO доработать обработку паники
		p.workerSelf--
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
