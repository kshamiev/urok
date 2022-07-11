package workerbus

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type Task interface {
	Execute()
}

type WorkerBus struct {
	muConsumer     sync.Mutex       // для безопасного конкурентного доступа
	muWait         sync.Mutex       // для безопасного конкурентного доступа
	wg             sync.WaitGroup   // WaitGroup для контроля полного завершения работ всех задач и каналов
	chTask         chan Task        // Канал задач - буф., чтобы родительская программа не блокировалась
	workerLimit    int              // Количество одновременно работающих обработчиков
	chData         chan interface{} // Канал обмена данными - буф., чтобы родительская программа не блокировалась
	storeSubscribe map[string]map[chan interface{}]struct{}
}

func NewWorkerBus(sizeBufferChanel, workerLimit int, debug bool) *WorkerBus {
	p := &WorkerBus{
		chTask:         make(chan Task, sizeBufferChanel),
		chData:         make(chan interface{}, sizeBufferChanel),
		workerLimit:    workerLimit,
		storeSubscribe: make(map[string]map[chan interface{}]struct{}),
	}
	for i := 0; i < p.workerLimit; i++ {
		p.wg.Add(1)
		go p.workerTask()
	}
	p.wg.Add(1)
	go p.workerData()
	if debug {
		go p.debugReport()
	}
	return p
}

func (p *WorkerBus) Wait() {
	p.muWait.Lock()
	p.workerLimit = 0
	p.muWait.Unlock()

	close(p.chTask)
	close(p.chData)

	p.wg.Wait()
}

func (p *WorkerBus) SendTask(task Task) {
	p.muWait.Lock()
	if p.workerLimit > 0 {
		p.chTask <- task
	}
	p.muWait.Unlock()
}

func (p *WorkerBus) SendData(obj interface{}) {
	p.muWait.Lock()
	if p.workerLimit > 0 {
		p.chData <- obj
	}
	p.muWait.Unlock()
}

func (p *WorkerBus) Subscribe(typ interface{}) chan interface{} {
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

func (p *WorkerBus) workerData() {
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

func (p *WorkerBus) workerTask() {
	defer func() {
		p.wg.Done()
		if rvr := recover(); rvr != nil {
			log.Println(fmt.Errorf("%+v", rvr))
			p.wg.Add(1)
			go p.workerTask()
		}
	}()
	var task Task

	for task = range p.chTask {
		task.Execute()
	}
}

func (p *WorkerBus) debugReport() {
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = f.Close()
	}()
	for {
		s := "count chanel Data: " + strconv.Itoa(len(p.chData)) + "\n"
		for i := range p.storeSubscribe {
			s += i + ": " + strconv.Itoa(len(p.storeSubscribe[i])) + "\n"
		}
		s += "\n"
		buf := &bytes.Buffer{}
		buf.WriteString(s)
		_, _ = buf.WriteTo(f)
		time.Sleep(time.Minute)
	}
}
