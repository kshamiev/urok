package workerdatabus

import (
	"reflect"
	"sync"
)

type Task interface {
	Execute()
}

type Pool struct {
	tasks       chan Task // Канал задач - буф., чтобы родительская программа не блокировалась
	workerLimit int       // Количество одновременно работающих обработчиков
	store       map[string]map[chan interface{}]struct{}
	wg          sync.WaitGroup // WaitGroup для контроля полного завершения работ всех задач и каналов
}

func NewPool(sizeChanel, workerLimit int) *Pool {
	p := &Pool{
		tasks:       make(chan Task, sizeChanel),
		workerLimit: workerLimit,
		store:       make(map[string]map[chan interface{}]struct{}),
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
	}
}

// //// DATA BUS

func (p *Pool) DataSubscribe(typ interface{}) chan interface{} {
	t := reflect.TypeOf(typ)
	i := t.String()
	ch := make(chan interface{})
	if _, ok := p.store[i]; !ok {
		p.store[i] = make(map[chan interface{}]struct{})
	}
	p.store[i][ch] = struct{}{}
	return ch
}

func (p *Pool) DataUnsubscribeData(typ interface{}, ch chan interface{}) {
	t := reflect.TypeOf(typ)
	i := t.String()
	if _, ok := p.store[i]; ok {
		delete(p.store[i], ch)
	}
	close(ch)
}

func (p *Pool) DataSend(typ interface{}) {
	t := reflect.TypeOf(typ)
	i := t.String()
	for ch := range p.store[i] {
		ch <- typ
	}
}
