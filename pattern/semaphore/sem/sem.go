package sem

import (
	"sync"
)

type dataChIn struct {
	key string
	val []int64
}

type dataChOut struct {
	key string
	val int64
}

type Pool struct {
	chanelIn  chan dataChIn
	chanelOut chan dataChOut
	wgWork    sync.WaitGroup
	wgRes     sync.WaitGroup
	result    map[string]int64
}

func NewPool(cnt int) *Pool {
	s := &Pool{
		chanelIn:  make(chan dataChIn),
		chanelOut: make(chan dataChOut),
		result:    map[string]int64{},
	}
	go func() {
		s.wgRes.Add(1)
		defer s.wgRes.Done()
		for d := range s.chanelOut {
			s.result[d.key] = d.val
		}
	}()
	for i := 0; i < cnt; i++ {
		go func() {
			s.wgWork.Add(1)
			defer s.wgWork.Done()
			for d := range s.chanelIn {
				s.chanelOut <- s.work(d)
			}
		}()
	}
	return s
}

func (s *Pool) work(d dataChIn) dataChOut {
	r := dataChOut{key: d.key}
	for i := range d.val {
		r.val += d.val[i]
	}
	return r
}

func (s *Pool) Work(src map[string][]int64) map[string]int64 {
	for key := range src {
		s.chanelIn <- dataChIn{
			key: key,
			val: src[key],
		}
	}
	close(s.chanelIn)
	s.wgWork.Wait()
	close(s.chanelOut)
	s.wgRes.Wait()
	return s.result
}

// ////

// Семофор
// type dataCh struct {
// 	key string
// 	val int64
// }
//
// type Sem struct {
// 	chanelReq chan string
// 	chanelRes chan dataCh
// 	wg        sync.WaitGroup
// 	src       map[string][]int64
// 	dist      map[string]int64
// }
//
// func NewSem(cnt int, src map[string][]int64) *Sem {
// 	s := &Sem{
// 		chanelReq: make(chan string),
// 		chanelRes: make(chan dataCh),
// 		src:       src,
// 		dist:      map[string]int64{},
// 	}
// 	for i := 0; i < cnt; i++ {
// 		go s.worker()
// 	}
// 	for key := range src {
// 		s.chanelReq <- key
// 	}
// 	return s
// }
//
// func (s *Sem) worker() {
// 	fmt.Println("start")
// 	s.wg.Add(1)
// 	defer s.wg.Done()
// 	for key := range s.chanelReq {
// 		s.chanelRes <- s.work(key)
// 	}
// 	fmt.Println("stop")
// 	return
// }
//
// func (s *Sem) work(key string) dataCh {
// 	d := dataCh{key: key}
// 	for i := range s.src[key] {
// 		d.val += s.src[key][i]
// 	}
// 	return d
// }
//
// func (s *Sem) Wait() {
// 	fmt.Println("start close")
// 	close(s.chanelReq)
// 	s.wg.Wait()
// 	fmt.Println("stop close")
// }
