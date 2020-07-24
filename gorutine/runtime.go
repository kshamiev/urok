// примеры работы с горутинами
package main

import (
	"fmt"
	"time"
)

func main() {
	Sample()
	// SamplechanelBlock1()
	// SamplechanelBlock2()
	// SampleRuntimeSelect()
}

func Sample() {
	fmt.Println("старт")
	// можно запустить функцию
	go process1(0)
	// можно запустить анонимную функцию
	go func() {
		fmt.Println("Анонимный запуск")
	}()

	// Можем запустить много горутин
	for i := 0; i < 1000; i++ {
		go process1(i)
	}

	var ch = make(chan int)
	close(ch)
	// num := <-ch
	// fmt.Println(num, "ТАК ЛУЧШЕ НЕ ДЕЛАТЬ")
	if num1, ok := <-ch; ok {
		fmt.Println(num1)
	} else {
		fmt.Println("close")
	}
	// Нужно дождаться заверешния выполнения
	fmt.Scanln()

}

func process1(i int) {
	fmt.Println("обработка: ", i)
}

// //////////////////////////
// Пример конструкции управление работой горутин с помощью стандартного закрытия канала
func SamplechanelBlock1() {

	// буферизированные
	ch := make(chan int64, 500)

	go func() {
		for {
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
	ch <- 123
	close(ch)
	fmt.Scanln()
}

// Пример конструкции управление работой горутин с помощью стандартного закрытия канала
func SamplechanelBlock2() {

	// буферизированные
	ch := make(chan int64, 500)

	go func() {
		for {
			select {
			// Блокируется пока канал не будет закрыт, либо не поступит новая задача
			case num, ok := <-ch:
				if ok {
					time.Sleep(time.Second * 1)
					fmt.Println("FOR ITER ", num)
				} else {
					fmt.Println("CLOSE")
					return
				}
				// ... следующие кейсы
			}
		}
	}()
	ch <- 45
	ch <- 34
	ch <- 76
	ch <- 123
	close(ch)
	fmt.Scanln()
}

// //////////////////////////////////////////////

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
				time.Sleep(time.Second * 1)
				fmt.Println("FOR ITERATION END ", cnt)
				if cnt >= 30 {
					controlCH <- true
					return
				}
			}
			// использовать в тяжелых горутинах в начале или конце бесконечного цикла!
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
