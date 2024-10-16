package benchmarks

import (
	"goji/internal/pkg/storage"
	"strconv"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	s, err := storage.NewStorage()
	if err != nil {
		b.Errorf("new storage: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), strconv.Itoa(i))
	}
}

func BenchmarkGet(b *testing.B) {
	s, err := storage.NewStorage()
	if err != nil {
		b.Errorf("new storage: %v", err)
	}

	s.Set("key", "value")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Get("key")
	}
}

func BenchmarkSetGet(b *testing.B) {
	s, err := storage.NewStorage()
	if err != nil {
		b.Errorf("new storage: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), strconv.Itoa(i))
		s.Get(strconv.Itoa(i))
	}
}
