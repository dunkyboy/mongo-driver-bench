package driver

import (
	"github.com/dunkyboy/mongo-driver-bench/shared"
	"github.com/mongodb/mongo-go-driver/bson"
	"testing"
)

var marshaler = func(in interface{}) []byte {
	out, err := bson.Marshal(
		in)
	if err != nil {
		panic(err)
	}
	return out
}

var unmarshaler = func(input []byte) interface{} {
	var output shared.SmallStructDepth1
	err := bson.Unmarshal(input, &output)
	if err != nil {
		panic(err)
	}
	return output
}

func BenchmarkMarshalSmallStructDepth1(b *testing.B) {
	shared.BenchMarshal(b, &shared.ASmallStructDepth1, marshaler)
}

func BenchmarkUnmarshalSmallStructDepth1(b *testing.B) {
	shared.BenchUnmarshal(b, marshaler(&shared.ASmallStructDepth1), unmarshaler)
}
