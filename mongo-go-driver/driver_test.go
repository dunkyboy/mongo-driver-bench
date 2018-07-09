package driver

import (
	"testing"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/dunkyboy/mongo-driver-bench/shared"
)

// save the test output somewhere global to force the compiler not to optimize it all away
var marshalResult []byte
var unmarshalResult interface{}

func BenchmarkMarshalSmallStructDepth1(b *testing.B) {
	in := shared.ASmallStructDepth1

	shared.BenchMarshal(
		b,
		func() []byte {
			out, err := bson.Marshal(&in)
			if err != nil {
				panic(err)
			}
			return out
		},
	)
}

func BenchmarkUnmarshalSmallStructDepth1(b *testing.B) {
	var output shared.SmallStructDepth1
	var er error

	input, err := bson.Marshal(&shared.ASmallStructDepth1)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		er = bson.Unmarshal(input, &output)
		if er != nil {
			panic(er)
		}
	}

	unmarshalResult = output
}
