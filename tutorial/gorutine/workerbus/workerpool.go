package workerbus

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"log"
	"math/big"
	"reflect"
	"sync"
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

func NewWorkerBus(sizeChanel, workerLimit int) *WorkerBus {
	p := &WorkerBus{
		chTask:         make(chan Task, sizeChanel),
		chData:         make(chan interface{}, sizeChanel),
		workerLimit:    workerLimit,
		storeSubscribe: make(map[string]map[chan interface{}]struct{}),
	}
	for i := 0; i < p.workerLimit; i++ {
		p.wg.Add(1)
		go p.workerTask()
	}
	p.wg.Add(1)
	go p.workerData()
	return p
}

func (p *WorkerBus) Wait() {
	p.muWait.Lock()
	p.workerLimit = 0
	p.muWait.Unlock()

	p.chData <- nil

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
	for obj := range p.chData {
		p.muConsumer.Lock()
		if obj == nil {
			for i := range p.storeSubscribe {
				for ch := range p.storeSubscribe[i] {
					ch <- nil
				}
			}
			for i := range p.storeSubscribe {
				for ch := range p.storeSubscribe[i] {
					<-ch
					delete(p.storeSubscribe[i], ch)
				}
			}
			continue
		}

		i := reflect.TypeOf(obj).String()
		for ch := range p.storeSubscribe[i] {
			ch <- obj
		}
		for ch := range p.storeSubscribe[i] {
			if _, ok := <-ch; !ok {
				delete(p.storeSubscribe[i], ch)
			}
		}
		p.muConsumer.Unlock()
	}
	p.wg.Done()
}

func (p *WorkerBus) workerTask() {
	defer func() {
		// TODO доработать обработку паники
		p.wg.Done()
		if rvr := recover(); rvr != nil {
			log.Println(fmt.Errorf("%+v", rvr))
		}
	}()

	for task := range p.chTask {
		task.Execute()
	}
}

// //// FOR TEST

func GenInt(x int64) int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}

// Dumper all variables to STDOUT
// From local debug
func Dumper(idl ...interface{}) string {
	ret := dump(idl...)
	fmt.Print(ret.String())

	return ret.String()
}

// dump all variables to bytes.Buffer
func dump(idl ...interface{}) bytes.Buffer {
	var buf bytes.Buffer

	var wr = io.MultiWriter(&buf)

	for _, field := range idl {
		fset := token.NewFileSet()
		_ = ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}

	return buf
}
