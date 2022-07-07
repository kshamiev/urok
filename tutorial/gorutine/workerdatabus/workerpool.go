package workerdatabus

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"reflect"
	"sync"

	"github.com/kshamiev/urok/sample/excel/typs"
)

type Task interface {
	Execute()
}

type Pool struct {
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
	return p
}

func (p *Pool) TaskAdd(task Task) {
	p.chTask <- task
}

func (p *Pool) Wait() {
	close(p.chTask)
	p.wg.Wait()
}

func (p *Pool) workerTask() {
	defer func() {
		// TODO доработать обработку паники
		p.wg.Done()
	}()

	for task := range p.chTask {
		task.Execute()
	}
}

// //// DATA BUS

func (p *Pool) workerData() {
	defer func() {
		// TODO доработать обработку паники
		p.wg.Done()
	}()

	for task := range p.chTask {
		task.Execute()
	}
}

func (p *Pool) DataSubscribe(typ interface{}) chan interface{} {
	t := reflect.TypeOf(typ)
	i := t.String()
	ch := make(chan interface{})
	if _, ok := p.storeSubscribe[i]; !ok {
		p.storeSubscribe[i] = make(map[chan interface{}]struct{})
	}
	p.storeSubscribe[i][ch] = struct{}{}
	return ch
}

func (p *Pool) DataUnsubscribeData(typ interface{}, ch chan interface{}) {
	t := reflect.TypeOf(typ)
	i := t.String()
	if _, ok := p.storeSubscribe[i]; ok {
		delete(p.storeSubscribe[i], ch)
	}
	close(ch)
}

func (p *Pool) DataSend(typ interface{}) {
	t := reflect.TypeOf(typ)
	i := t.String()
	for ch := range p.storeSubscribe[i] {
		ch <- typ
		fmt.Println("SEND: " + typ.(*typs.Cargo).Name)
	}
	for ch := range p.storeSubscribe[i] {
		<-ch
		fmt.Println("FINISH: ")
	}

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
