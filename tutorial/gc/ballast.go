package main

import (
	"time"
)

var ballast []byte

// ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --from-unit=1024 --to=iec --field 6-7 | column -t | egrep "ballast.*"
// VSZ выделенная виртуальная память
// RSS выделенная физическая память
func main() {

	// Create a large heap allocation of 10 GiB
	ballast = make([]byte, 10<<30)

	// блокировка (процессоры отдыхают)
	<-time.After(time.Hour)

}
