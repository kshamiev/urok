// go test function/* -v -run Test_add_one

// go test -coverprofile=cover.out ./function - генерим отчет о coverage
// go tool cover -html=cover.out -o cover.html - переводим отчет в html

// go tool cover -func=cover.out - покрытие по функциям
package function

import (
	"fmt"
	"testing"
)

func Test_Closure(t *testing.T) {

	Closure()

}

func Test_Parameter(t *testing.T) {

	ParameterSum1(45, 55, Parameter)

	res, flag := ParameterSum2(5, 7, func(x, y int) int {
		return x * y
	})
	fmt.Println(res, flag)
}

func Test_Add(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{args: []int{1}}, want: 1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.args...); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
