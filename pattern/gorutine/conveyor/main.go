package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := make(chan int)
	// собственно сам конвейер (функции можно добавлять и вкладывать далее друг в друга
	// при закрытии входящего канала все последующие (каналы и горутины)
	// закрываются по завершении обработки последнего элемента автоматически
	out := work(square(in))
	done := make(chan bool)
	go func() {
		for val := range out {
			fmt.Println("ответ конвейера: ", val)
		}
		fmt.Println("Концерт окончен")
		done <- true
	}()

	fmt.Println("Введите целое число либо слово exit (ctrl+D) для выхода")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text() // получаем строку
		if line == "exit" {
			break
		} else {
			number, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				continue
			}
			in <- number
		}
	}
	close(in)
	<-done
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var val int
		for val = range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func work(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var val int
		for val = range in {
			out <- val * 2
		}
		close(out)
	}()
	return out
}
