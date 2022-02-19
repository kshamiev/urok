package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// go test ./sample/imgtobase/. -run=^# -bench=Home -benchtime=1000x -count 1
func BenchmarkHome(b *testing.B) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		Home(w, r)
		if w.Code != 200 {
			greeting, _ := ioutil.ReadAll(w.Body)
			b.Fatal(w.Code, string(greeting))
		}
	}
}
