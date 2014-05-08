package deephash

import (
	"fmt"
	"testing"
)

type testStruct struct {
	S   string
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64
	F32 float32
	F64 float64
}

var differentTestCases = []interface{}{
	// simple types
	"dave",
	"foo",
	"foobar",
	" foo",
	1,
	1.0,

	// structs
	testStruct{S: "foo"},
	testStruct{S: "bar"},

	// pointers to structs
	&testStruct{S: "foo1"},
	&testStruct{S: "bar1"},

	// structs with different types of ints
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 43, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 44, I8: 44, I16: 44, I32: 44, I64: 44, U8: 44, U16: 44, U32: 44, U64: 44},
	&testStruct{I: 11, I8: 43, I16: 43, I32: 43, I64: 43, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 11, I16: 43, I32: 43, I64: 43, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 11, I32: 43, I64: 43, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 11, I64: 43, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 11, U8: 43, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 43, U8: 11, U16: 43, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 43, U8: 43, U16: 11, U32: 43, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 43, U8: 43, U16: 43, U32: 11, U64: 43},
	&testStruct{I: 43, I8: 43, I16: 43, I32: 43, I64: 43, U8: 43, U16: 43, U32: 43, U64: 11},

	// structs with different types of floats
	&testStruct{F32: 43.0, F64: 43.0},
	&testStruct{F32: 44.0, F64: 44.0},
	&testStruct{F32: 11.0, F64: 43.0},
	&testStruct{F32: 43.0, F64: 11.0},

	// string maps
	map[string]testStruct{
		"foo": testStruct{S: "baz"},
		"bar": testStruct{S: "baz"},
	},
	map[string]testStruct{
		"foo": testStruct{S: "BAZZER"},
		"bar": testStruct{S: "BAZZER"},
	},
}

var sameCases = [][]interface{}{
	// simple stuff
	[]interface{}{
		"foo",
		"foo",
	},

	// hash order shouldn't matter
	[]interface{}{
		map[string]testStruct{
			"foo": testStruct{S: "baz"},
			"bar": testStruct{S: "baz"},
		},
		map[string]testStruct{
			"bar": testStruct{S: "baz"},
			"foo": testStruct{S: "baz"},
		},
	},
}

func TestDifferentCases(t *testing.T) {
	seen := make(map[string]bool)
	for _, tc := range differentTestCases {
		h := Hash(tc)
		hs := fmt.Sprintf("%x", h)
		if len(h) == 0 {
			t.Errorf("Test case %v yields zero length hash", tc)
			continue
		}
		if seen[hs] {
			t.Errorf("Test case %v hashes to %v which has already been seen", tc, hs)
		}
		seen[hs] = true
	}
}

func TestSameCases(t *testing.T) {
	for _, tcs := range sameCases {
		hash := ""
		for _, tc := range tcs {
			h := Hash(tc)
			hs := fmt.Sprintf("%x", h)
			if len(h) == 0 {
				t.Errorf("Test case %v yields zero length hash", tc)
				continue
			}

			if hash == "" {
				hash = hs
			} else if hash != hs {
				t.Errorf("Test case %v hashes to '%v' which is different to previous '%v'", tc, hs, hash)
			}
		}
	}
}