package main

import "fmt"

type Test struct {
	Cnt int
}

func (obj Test) Add(n int) {
	obj.Cnt += n
}

func (obj Test) Del(n int) {
	obj.Cnt -= n
}

func (obj *Test) String() string {
	return fmt.Sprintf("%v", obj.Cnt)
}

func main() {

	var self Test

	self.Add(5)
	self.Del(10)

	fmt.Println(self)

}
