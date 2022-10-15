package bloom

import (
	"fmt"
	"hash"
	"hash/fnv"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	filter, err := New(128, []hash.Hash64{fnv.New64(), fnv.New64a()})
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		filter.Add([]byte("test"))
	}
}

func BenchmarkNotExists(b *testing.B) {
	filter, err := New(128, []hash.Hash64{fnv.New64(), fnv.New64a()})
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		filter.Exists([]byte("test"))
	}
}

func BenchmarkExists(b *testing.B) {
	filter, err := New(128, []hash.Hash64{fnv.New64(), fnv.New64a()})
	if err != nil {
		b.Error(err)
	}

	filter.Add([]byte("test"))

	for n := 0; n < b.N; n++ {
		filter.Exists([]byte("test"))
	}
}

func BenchmarkMany(b *testing.B) {
	filter, err := New(128, []hash.Hash64{fnv.New64(), fnv.New64a()})
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		item := []byte(fmt.Sprintf("test%d", n))
		filter.Add(item)
		filter.Exists(item)
	}
}
