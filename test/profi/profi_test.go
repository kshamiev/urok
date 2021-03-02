// go test -bench=.
// GOGC=off go test -bench=BenchmarkRegex -cpuprofile cpu.out -memprofile mem.out
// go tool pprof profi.test cpu.out
// go tool pprof -alloc_space profi.test mem.out
// go tool pprof -inuse_space memtest profi.test mem.out
// (pprof) svg
//
// https://github.com/pkg/profile

package profi

import (
	"regexp"
	"strings"
	"testing"
)

var haystack = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras accumsan nisl et iaculis fringilla. Integer sapien orci, facilisis ut venenatis nec, suscipit at massa. Cras suscipit lectus non neque molestie, et imperdiet sem ultricies. Donec sit amet mattis nisi, efficitur posuere enim. Aliquam erat volutpat. Curabitur mattis nunc nisi, eu maximus dui facilisis in. Quisque vel tortor mauris. Praesent tellus sapien, vestibulum nec purus ut, luctus egestas odio. Ut ac ipsum non ipsum elementum pretium in id enim. Aenean eu augue fringilla, molestie orci et, tincidunt ipsum.
Nullam maximus odio vitae augue fermentum laoreet eget scelerisque ligula. Praesent pretium eu lacus in ornare. Maecenas fermentum id sapien non faucibus. Donec est tellus, auctor eu iaculis quis, accumsan vitae ligula. Fusce dolor nisl, pharetra eu facilisis non, hendrerit ac turpis. Pellentesque imperdiet aliquam quam in luctus. Curabitur ut orci sodales, faucibus nunc ac, maximus odio. Vivamus vitae nulla posuere, pellentesque quam posuere`

var pattern = regexp.MustCompile("auctor")

func BenchmarkSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains(haystack, "auctor")
	}
}

func BenchmarkRegex(b *testing.B) {
	// pprof.StartCPUProfile(...)
	// pprof.StopCPUProfile()
	// runtime.MemProfile(...)

	for i := 0; i < b.N; i++ {
		// _, _ = regexp.MatchString("auctor", haystack)
		pattern.MatchString(haystack)
	}
}
