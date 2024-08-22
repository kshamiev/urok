package main

import (
	"time"
)

// ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --from-unit=1024 --to=iec --field 6-7 | column -t | egrep "ballast.*|PID"
// VSZ выделенная виртуальная память
// RSS выделенная физическая память
func main() {
	// блокировка (процессоры отдыхают)
	<-time.After(time.Hour)
}
