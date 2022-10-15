package bloom

import (
	"errors"
	"hash"
)

func New(size int, k []hash.Hash64) (*Filter, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	if len(k) < 2 {
		return nil, errors.New("must provide at least 2 hash functions")
	}

	filter := &Filter{
		set:     make([]bool, size),
		hashers: k,
		size:    uint64(size),
	}

	return filter, nil
}

type Filter struct {
	hashers []hash.Hash64
	size    uint64
	set     []bool
}

func (f *Filter) Add(item []byte) error {
	res := hashAll(f.hashers, item)

	for _, val := range res {
		position := val % uint64(len(f.set))
		f.set[position] = true
	}

	return nil
}

func (f *Filter) Exists(item []byte) bool {
	res := hashAll(f.hashers, item)

	for _, val := range res {
		position := val % f.size
		if !f.set[position] {
			return false
		}
	}

	return true
}

func hashAll(hashers []hash.Hash64, item []byte) (res []uint64) {
	for _, hasher := range hashers {
		hasher.Write(item)
		res = append(res, hasher.Sum64())
		hasher.Reset()
	}

	return res
}
