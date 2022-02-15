package main

type X struct{ v int }

func foo(x *X) {
	// fmt.Println(x.v)
}

// go build -gcflags=-m memory/analiz/main.go
func main() {
	x := &X{1}
	foo(x)
}
