package ch1

import "testing"

var sampleArguments = []string{"Brian", "Kernighan", "was", "in", "the", "Computing", "Science", "Research", "center", "at", "Bell", "Labs", "until", "2000,", "where", "he", "worked", "on", "languages", "and", "tools", "for", "Unix.", "He", "is", "now", "a", "professor", "in", "the", "Computer", "Science", "Department", "at", "Princeton.", "He", "is", "the", "co-author", "of", "several", "books,", "including", "The", "C", "Programming", "Language", "and", "The", "Practice", "of", "Programming."}

func BenchmarkLoop(b *testing.B) {
	b.Run("benchmark for loop based.", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UsingLoop(sampleArguments)
		}
	})
}

func BenchmarkJoin(b *testing.B) {
	b.Run("benchmark using join.", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UsingJoin(sampleArguments)
		}
	})
}

/*

Following were the benchmark results.

Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkJoin$ github.com/hkbiet/v2/chapter1
140942          9935 ns/op         320 B/op          1 allocs/op
PASS
ok      github.com/hkbiet/v2/chapter1   1.490s


Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkJoin$ github.com/hkbiet/v2/chapter1
167667          7777 ns/op         320 B/op          1 allocs/op
PASS
ok      github.com/hkbiet/v2/chapter1   1.383s


Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkJoin$ github.com/hkbiet/v2/chapter1
158581          9297 ns/op         320 B/op          1 allocs/op
PASS
ok      github.com/hkbiet/v2/chapter1   1.571s


Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkLoop$ github.com/hkbiet/v2/chapter1

signal: terminated
FAIL    github.com/hkbiet/v2/chapter1   2.945s
FAIL
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkLoop$ github.com/hkbiet/v2/chapter1

signal: terminated
FAIL    github.com/hkbiet/v2/chapter1   1.739s
FAIL



------------------------------------------------------------------------------------------------------------

As is evident, BenchmarkLoop was intensive than BenchmarkJoin.

*/