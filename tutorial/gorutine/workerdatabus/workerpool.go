package workerdatabus

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

type Pool struct {
	muConsumer     sync.Mutex       // для безопасного конкурентного доступа
	muWait         sync.Mutex       // для безопасного конкурентного доступа
	wg             sync.WaitGroup   // WaitGroup для контроля полного завершения работ всех задач и каналов
	chTask         chan Task        // Канал задач - буф., чтобы родительская программа не блокировалась
	workerLimit    int              // Количество одновременно работающих обработчиков
	chData         chan interface{} // Канал обмена данными - буф., чтобы родительская программа не блокировалась
	storeSubscribe map[string]map[chan interface{}]struct{}
}

func NewPool(sizeChanel, workerLimit int) *Pool {
	p := &Pool{
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

func (p *Pool) Wait() {
	p.muWait.Lock()
	p.workerLimit = 0
	p.muWait.Unlock()
	// p.TaskSend(nil)
	// p.DataSend(nil)
	close(p.chTask)
	close(p.chData)
	p.wg.Wait()
}

func (p *Pool) TaskSend(task Task) {
	p.muWait.Lock()
	if p.workerLimit > 0 {
		p.chTask <- task
	}
	p.muWait.Unlock()
}

func (p *Pool) DataSend(obj interface{}) {
	p.muWait.Lock()
	if p.workerLimit > 0 {
		p.chData <- obj
	}
	p.muWait.Unlock()
}

func (p *Pool) workerTask() {
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

func (p *Pool) workerData() {
	for obj := range p.chData {
		// отправка данных подписчику
		t := reflect.TypeOf(obj)
		i := t.String()
		for ch := range p.storeSubscribe[i] {
			ch <- obj
		}
		// подтверждение обработки полученных данных и отписка
		for ch := range p.storeSubscribe[i] {
			if f, ok := (<-ch).(bool); ok && !f {
				close(ch)
				p.muConsumer.Lock()
				delete(p.storeSubscribe[i], ch)
				p.muConsumer.Unlock()
				fmt.Println("CLOSE:")
			}
		}
	}
	p.wg.Done()
}

func (p *Pool) DataSubscribe(typ interface{}) chan interface{} {
	t := reflect.TypeOf(typ)
	i := t.String()
	ch := make(chan interface{})
	p.muConsumer.Lock()
	if _, ok := p.storeSubscribe[i]; !ok {
		p.storeSubscribe[i] = make(map[chan interface{}]struct{})
	}
	p.storeSubscribe[i][ch] = struct{}{}
	p.muConsumer.Unlock()
	return ch
}

// //// для тестирования

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

func GenInt(x int64) int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
