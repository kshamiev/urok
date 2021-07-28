// https://habr.com/ru/post/271789/
package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTickets      = errors.New("не могу захватить семафор")
	ErrIllegalRelease = errors.New("не могу освободить семафор, не захватив его сначала")
)

// Interface содержит поведение семафора,
// который может быть захвачен (Acquire)
// и/или освобожден (Release).
type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) *implementation {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

/*
func main() {

		// Семафор с таймаутом
		tickets, timeout := 1, time.Duration(0)
		s := semaphore.New(tickets, timeout)

		if err := s.Acquire(); err != nil {
			panic(err)
		}

		// Выполняем важную работу

		if err := s.Release(); err != nil {
			panic(err)
		}

		// ////

		// Семафор без таймаутов
		tickets, timeout = 0, 0
		s = semaphore.New(tickets, timeout)

		if err := s.Acquire(); err != nil {
			if err != semaphore.ErrNoTickets {
				panic(err)
			}

			// Билетов не осталось, не могу работать
			os.Exit(1)
		}

	fmt.Println("OK")
}
*/
