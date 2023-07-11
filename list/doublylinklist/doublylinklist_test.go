package doublylinklist_test

import (
	"testing"

	"github.com/454270186/GoDazzle/list/doublylinklist"
)

func benchmarkAdd(l *doublylinklist.DoublyLinkList, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			l.Add(j)
		}
	}
}

func benchmarkGet(l *doublylinklist.DoublyLinkList, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			l.Get(j)
		}
	}
}

func benchmarkRemove(l *doublylinklist.DoublyLinkList, size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			l.Remove(j)
		}
	}
}

func BenchmarkAdd100(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100
	b.StartTimer()
	benchmarkAdd(l, size, b)
}

func BenchmarkAdd1000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 1000
	b.StartTimer()
	benchmarkAdd(l, size, b)
}

func BenchmarkAdd10000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 10000
	b.StartTimer()
	benchmarkAdd(l, size, b)
}

func BenchmarkAdd100000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100000
	b.StartTimer()
	benchmarkAdd(l, size, b)
}

func BenchmarkGet100(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkGet(l, size, b)
}

func BenchmarkGet1000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 1000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkGet(l, size, b)
}

func BenchmarkGet10000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 10000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkGet(l, size, b)
}

func BenchmarkGet100000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkGet(l, size, b)
}

func BenchmarkRemove100(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkRemove(l, size, b)
}

func BenchmarkRemove1000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 1000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkRemove(l, size, b)
}

func BenchmarkRemove10000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 10000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkRemove(l, size, b)
}

func BenchmarkRemove100000(b *testing.B) {
	b.StopTimer()
	l := doublylinklist.New()
	size := 100000
	for i := 0; i < size; i++ {
		l.Add(i)
	}
	b.StartTimer()
	benchmarkRemove(l, size, b)
}