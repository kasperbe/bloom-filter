package bloom

import (
	"bytes"
	"encoding/binary"
	"hash"
	"hash/fnv"
	"testing"
)

func TestNewFilterSize(t *testing.T) {
	tt := map[string]struct {
		size int
		err  string
	}{
		"zero value": {
			size: 0,
			err:  "size must be greater than zero",
		},
		"negative value": {
			size: -100,
			err:  "size must be greater than zero",
		},
		"success": {
			size: 100,
			err:  "",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			_, err := New(tc.size, []hash.Hash64{fnv.New64(), fnv.New64a()})
			if err == nil && tc.err == "" {
				return
			}

			if err.Error() != tc.err {
				t.Error(err)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, int32(123))

	tt := map[string]struct {
		input []byte
	}{
		"numbers": {input: buf.Bytes()},
		"strings": {input: []byte("123")},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			filter, err := New(16, []hash.Hash64{fnv.New64(), fnv.New64a()})
			if err != nil {
				t.Error(err)
			}

			err = filter.Add(tc.input)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestExists(t *testing.T) {
	tt := map[string]struct {
		input  [][]byte
		check  []byte
		exists bool
	}{
		"not exists": {input: [][]byte{[]byte("test"), []byte("test1"), []byte("test2"), []byte("test3")}, check: []byte("something else"), exists: false},
		"success_1":  {input: [][]byte{[]byte("test"), []byte("test1"), []byte("test2"), []byte("test3")}, check: []byte("test1"), exists: true},
		"success_2":  {input: [][]byte{[]byte("some"), []byte("thing"), []byte("must"), []byte("exist")}, check: []byte("exist"), exists: true},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			filter, err := New(16, []hash.Hash64{fnv.New64(), fnv.New64a()})
			if err != nil {
				t.Error(err)
			}

			for _, input := range tc.input {
				err := filter.Add(input)
				if err != nil {
					t.Error(err)
				}
			}

			exists := filter.Exists(tc.check)
			if exists != tc.exists {
				t.Errorf("%s: %s - %s: %t", name, tc.input, tc.check, tc.exists)
			}
		})
	}
}
