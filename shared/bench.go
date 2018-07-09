package shared

import "testing"

// save the test output somewhere global to force the compiler not to optimize it all away
var marshalResult []byte
var unmarshalResult interface{}

func BenchMarshal(b *testing.B, marshaler func() []byte) {
	var bytes []byte

	for n := 0; n < b.N; n++ {
		bytes = marshaler()
	}

	marshalResult = bytes
}

func BenchUnmarshal(b *testing.B, unmarshaler func() interface{}) {
	var output interface{}

	for n := 0; n < b.N; n++ {
		output = unmarshaler()
	}

	unmarshalResult = output
}
