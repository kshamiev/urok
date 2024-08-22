package main

import (
	"time"
)

var ballast []byte

// ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --to=iec --field 4-5 | numfmt --header --from-unit=1024 --to=iec --field 6-7 | column -t | egrep "ballast.*"
// VSZ выделенная виртуальная память
// RSS выделенная физическая память
func main() {

	// Create a large heap allocation of 10 GiB
	ballast = make([]byte, 10<<30)

	// for i := 0; i < len(ballast); i++ {
	// 	ballast[i] = byte('A')
	// }

	// блокировка (процессоры отдыхают)
	<-time.After(time.Hour)

	// печка (печём блины на процессоре)
	// for {
	// }
}
