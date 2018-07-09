package mgo

import (
	"github.com/dunkyboy/mongo-driver-bench/shared"
	"gopkg.in/mgo.v2/bson"
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

func BenchmarkMarshalLargeStructDepth1(b *testing.B) {
	shared.BenchMarshal(b, &shared.ALargerStructDepth1, marshaler)
}

func BenchmarkUnmarshalSmallStructDepth1(b *testing.B) {
	shared.BenchUnmarshal(b, marshaler(&shared.ASmallStructDepth1), unmarshaler)
}

func BenchmarkUnmarshalLargeStructDepth1(b *testing.B) {
	shared.BenchUnmarshal(b, marshaler(&shared.ALargerStructDepth1), unmarshaler)
}
