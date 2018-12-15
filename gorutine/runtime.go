// примеры работы с горутинами
// gorutine - облегченный отдельный поток (обычно это функция запускаемая как отедльная программа)
// gorutine - не зависимо выполняемая логичаская часть программы (обычно функция) запущенная оператором go
// runtime занимается распределением горутин по тредам
// go run -race test.go (аргумент race) дает информацию о взаимодействии горутин с общими объектами
// с его помощью можно тестировать отдельный участки программы работающие с горутинами
// deadlock это когда мы бесконечно ждем чтения из канала и не можем прочитать или наоборот бесконечно записать в канал
package main

import (
	"fmt"
	//	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Пример конструкции управление работой горутин с помощью стандартного закрытие канала
func SamplechanelBlock() {

	// буферизированные
	ch := make(chan int64, 500)

	go func() {
		for {
			// такая конструкция будет ловить закрытие канала
			// блокируется пока канал не будет закрыт либо пока не поступят новые данные
			if num, ok := <-ch; ok {
				time.Sleep(time.Second * 1)
				fmt.Println("FOR ITER ", num)
			} else {
				fmt.Println("CLOSE")
				return
			}
		}
	}()
	ch <- 45
	ch <- 34
	ch <- 76
	close(ch)

	// selec с блоком default
	// case ch1, ok := <-ch:
	// ...
	fmt.Scanln()
}

// пример использование Once & atomic
// выполенине содержимого только один раз (даже если в разных горутинах)
func SampleRuntimeOnce() {

	var onc sync.Once

	var T = int64(345)
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&T, 10)
			fmt.Println(atomic.LoadInt64(&T))
			onc.Do(func() {
				fmt.Println("sync.Once")
			})
		}()
	}
	fmt.Scanln()
	fmt.Print("res: ", T)
}

// Удобно работать с горутинами через структуру
// задаем каналы в свойства структуры после этого запускаем метод в горутине в конструкторе

/////////////////////////////////////////////////////////////
// Пример работы Mutex WaitGroup Atomic

type account1Mutex struct {
	sync.Mutex
	balance float64
}

// блокирует для других горутин любую работу с shared обьектом
// для этого такой блок прописывается во всех необходимых методах
func (a *account1Mutex) accountCalc() {
	a.Lock()
	defer a.Unlock()
	// здесь какая-то работа
	// другие горутины ждут (полная блокировка)
}

type account2RWMutex struct {
	sync.RWMutex
	balance float64
}

// блокировка на чтение другие горутины могут читать
// блокировка на запись будет ждать прочтения
func (a *account2RWMutex) balanceLoad() float64 {
	a.RLock()
	defer a.RUnlock()
	return a.balance
}

// блокировка на запись
// будет ждать пока все прочитают
// когда будет изменять все будут ждать
func (a *account2RWMutex) balanceSet(x float64) {
	a.Lock()
	defer a.Unlock()
	a.balance += x
}

type atomicCnt struct {
	val int64
}

// атомарно будет добавлять значение в горутинах
func (a *atomicCnt) Add(x int64) {
	atomic.AddInt64(&a.val, x)
}

// атомарно будет читать значение в горутинах
func (a *atomicCnt) Value() int64 {
	return atomic.LoadInt64(&a.val)
}

func SampleRuntimeWaitGroupAndMutex() {

	var wg sync.WaitGroup

	// счетчик выполняющийхся задач (горутин) (ставится в контролирующей функции перед запуском горутины)
	wg.Add(3)

	// ставится в порождаемой горутине в блок завершения ее работы убавляет счетчик на 1 (передается по ссылке)
	wg.Done()
	wg.Done()
	wg.Done()

	// ждет пока счетчик не будет равено 0 (ставится в контролирующей функции после запуска горутин)
	wg.Wait()

	//////

	acc1 := account1Mutex{}
	acc1.accountCalc()

	acc2 := account2RWMutex{}
	acc2.balanceLoad()
	acc2.balanceSet(34.56)
	acc2.balanceLoad()

	////

	atm := &atomicCnt{}
	atm.Add(64)

}

////////////////////////////////////////////////

// пример работы свича select для горутин
func SampleRuntimeSelect() {

	fmt.Println("START OK")

	controlCH := make(chan bool)
	ch1 := make(chan int64, 10)
	ch2 := make(chan int64, 10)
	ch3 := make(chan int64, 10)

	// заполняем каналы после задержки
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("END TIMEOUT SET CHANEL")
		for i := int64(0); i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
		fmt.Println("END SEND TO CHANEL")
		controlCH <- true
	}()

	// читаем из каналов
	go func() {
		var cnt uint16
		for {
			cnt++
			select {
			case num := <-ch1:
				fmt.Println("CH 1: ", num)
			case num := <-ch2:
				fmt.Println("CH 2: ", num)
			case num := <-ch3:
				fmt.Println("CH 3: ", num)
			// если блок default не указан то программа заблокируется в ожидании ответа от каналов
			// (рекомендуется поскольку ресурсы не тратятся и runtime может выполнять другую горутину)
			// если блок default указан программа будет проваливаться в него и блокировки ожидания не будет
			// (плохо поскольку греем космос и тратим ресурсы)
			default:
				//				fmt.Println("SELECT DEFAULT")
			}
			time.Sleep(time.Second * 1)
			fmt.Println("FOR ITERATION END ", cnt)
			if cnt >= 30 {
				controlCH <- true
				return
			}

			// использовать в тяжелых горутинах в начале или конце бесконечного цикла
			// говорит runtime что можно прерваться и преключиться на другие задачи
			// runtime.Gosched()
		}
	}()

	// ждем завершения работы подпрограмм
	flag := <-controlCH
	flag = <-controlCH

	close(ch1)
	close(ch2)
	close(ch3)

	if flag == true {
		fmt.Sprintln("FINISH OK")
	} else {
		fmt.Sprintln("FINISH OK")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// особенность буферизированного канала
// особенность: после закрытия канала из него можно читать.
// особенность: после закрытия буферизированного канала в нем сохраняются ранее добавленные и не обработанные данные
func SampleRuntime2() {
	stuff := make(chan int64, 7)
	for i := int64(0); i < 19; i = i + 3 {
		stuff <- i
	}
	close(stuff)
	fmt.Println("Res: ", process(stuff))

}

func process(input <-chan int64) (res int64) {
	for r := range input {
		res += r
	}
	return
}

/////////////////////////////////////////////////////////////////////////////////////////////////

// родительская функция верхнего уровня
// то есть какя-то программа которая использует горутины (или которой нужны горутины) для убыстрения
func SampleRuntimeKanal2() {
	chSet := make(chan int64)
	chGet := make(chan int64)
	go samplewrapper2(chSet, chGet)
	for i := int64(1); i < 20; i = i + 3 {
		chSet <- i
		// тут что то можно делать паралельно тяжелое
		res := <-chGet
		fmt.Println(res)
	}
}

// оберточная функция, посредник, реализующая сам механизм взаимодействия с родительской программой
func samplewrapper2(chIn chan int64, chOut chan int64) {
	for v := range chIn {
		chOut <- sampleWork2(v)
	}
}

// непосредственно рабочая функция которая работает в отдельной программе
func sampleWork2(i int64) int64 {
	return i + 10
}

/////////////////////////////////////////////////////////////////////////////////////////////////

// родительская функция верхнего уровня
// то есть какя-то программа которая использует горутины (или которой нужны горутины) для убыстрения
func SampleRuntimeKanal1() {
	chSend := make(chan int64)
	chInput := samplewrapper1(chSend)
	for i := int64(1); i < 20; i = i + 3 {
		chSend <- i
		// тут что то можно делать паралельно тяжелое
		res := <-chInput
		fmt.Println(res)
	}
}

// оберточная функция посредник реализующая сам механизм взаимодействия с родительской программой
// принимает и отдает канал, запускает внутри себя горутину используя замыкание
func samplewrapper1(chIn <-chan int64) <-chan int64 {
	var chOut = make(chan int64)
	go func() {
		for v := range chIn {
			chOut <- sampleWork1(v)
		}
	}()
	return chOut
}

// непосредственно рабочая функция которая работает в отдельной программе
func sampleWork1(i int64) int64 {
	return i + 10
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////

// базовый синтаксис и механизм запуска горутин
func SampleRuntime() {
	// создание не буфиризированного канала
	// ch := make(chan int64)
	// создание буфиризированного канала
	// ch := make(chan int64, 7)

	// запускаем функцию
	go sampleCalc(999)

	// запускаем анонимную функцию
	go func() {
		fmt.Println("anonim")
	}()

	// запускаем много функций
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range arr {
		go sampleCalc(v)
		// захватываемые переменные нужно явно передавать
		// поскольку в момент выполнения и вывода значения
		// они могут быть уже изменены родительской программой
		go func(i int64) {
			fmt.Println(i)
		}(v)

	}
	fmt.Scanln()
}

func sampleCalc(num int64) {
	fmt.Println(num)
}
