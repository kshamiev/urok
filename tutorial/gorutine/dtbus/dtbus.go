package dtbus

import (
	"reflect"
	"sync"
)

type DtBus struct {
	muConsumer     sync.Mutex       // для безопасного конкурентного доступа
	muWait         sync.Mutex       // для безопасного конкурентного доступа
	wg             sync.WaitGroup   // WaitGroup для контроля полного завершения работ всех задач и каналов
	isDone         bool             // Количество одновременно работающих обработчиков
	chData         chan interface{} // Канал обмена данными - буф., чтобы родительская программа не блокировалась
	storeSubscribe map[string]map[chan interface{}]struct{}
}

func NewDtBus(sizeBufferChanel int) *DtBus {
	p := &DtBus{
		chData:         make(chan interface{}, sizeBufferChanel),
		storeSubscribe: make(map[string]map[chan interface{}]struct{}),
	}
	p.wg.Add(1)
	go p.workerData()
	return p
}

func (p *DtBus) Wait() {
	p.muWait.Lock()
	p.isDone = true
	p.muWait.Unlock()

	close(p.chData)

	p.wg.Wait()
}

func (p *DtBus) SendData(obj interface{}) {
	p.muWait.Lock()
	if !p.isDone {
		p.chData <- obj
	}
	p.muWait.Unlock()
}

func (p *DtBus) Subscribe(typ interface{}) chan interface{} {
	i := reflect.TypeOf(typ).String()
	ch := make(chan interface{})

	p.muConsumer.Lock()
	if _, ok := p.storeSubscribe[i]; !ok {
		p.storeSubscribe[i] = make(map[chan interface{}]struct{})
	}
	p.storeSubscribe[i][ch] = struct{}{}
	p.muConsumer.Unlock()

	return ch
}

func (p *DtBus) workerData() {
	var (
		ok  bool
		i   string
		ch  chan interface{}
		obj interface{}
	)
	for obj = range p.chData {
		i = reflect.TypeOf(obj).String()
		p.muConsumer.Lock()
		for ch = range p.storeSubscribe[i] {
			ch <- obj
		}
		for ch = range p.storeSubscribe[i] {
			if _, ok = <-ch; !ok {
				delete(p.storeSubscribe[i], ch)
			}
		}
		p.muConsumer.Unlock()
	}
	p.muConsumer.Lock()
	for i = range p.storeSubscribe {
		for ch = range p.storeSubscribe[i] {
			ch <- nil
		}
	}
	for i = range p.storeSubscribe {
		for ch = range p.storeSubscribe[i] {
			<-ch
			delete(p.storeSubscribe[i], ch)
		}
	}
	p.muConsumer.Unlock()
	p.wg.Done()
}
