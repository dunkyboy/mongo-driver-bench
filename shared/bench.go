package shared

import "testing"

var marshalResult []byte

func BenchMarshal(b *testing.B, marshaler func() []byte) {
	var bytes []byte

	for n := 0; n < b.N; n++ {
		bytes = marshaler()
	}

	marshalResult = bytes
}
