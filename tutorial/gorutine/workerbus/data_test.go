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
	"testing"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/typs"
)

const (
	countObject = 1000000
)

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=1000000x -count 10 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	b.ReportAllocs()
	Init(100000, 3)

	// подписчики
	for i := 0; i < 1; i++ {
		ch := Gist().Subscribe(&typs.General{})
		go consumerB(ch)
	}

	b.ResetTimer()
	// отправитель
	for j := 0; j < b.N; j++ {
		Gist().SendData(&typs.General{Name: fmt.Sprintf("additional_%d", j), Amount: 1})
	}

	Gist().Wait()
}

func consumerB(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = Gist().Subscribe(&typs.General{})
			go consumerB(ch)
		}
	}()
	for obj := range ch {
		_, ok := obj.(*typs.General)
		if !ok {
			close(ch)
			break
		}
		// It`s Work
		// ...
		ch <- true
		i++
	}
	// log.Println("full count: ", i)
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(100000, 3)

	// подписчики
	for i := 0; i < 1; i++ {
		ch := pool.Subscribe(&typs.General{})
		go consumerT(pool, ch)
	}

	// отправитель
	for i := 0; i < countObject; i++ {
		pool.SendData(&typs.General{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
	}

	pool.Wait()
}

func consumerT(pool *WorkerBus, ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = pool.Subscribe(&typs.General{})
			go consumerT(pool, ch)
		}
	}()
	for obj := range ch {
		_, ok := obj.(*typs.General)
		if !ok {
			close(ch)
			break
		}
		// It`s Work
		if i == 1000 {
			panic("PANICA")
		}
		// ...
		ch <- true
		i++
	}
	log.Println("full count: ", i)
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
