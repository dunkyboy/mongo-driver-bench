package benchmark

import (
	"github.com/dunkyboy/mongo-driver-bench/mgo"
	"github.com/dunkyboy/mongo-driver-bench/mongo-go-driver"
	"testing"
)

// save the test output somewhere global to force the compiler not to optimize it all away
var marshalResult []byte
var unmarshalResult interface{}

func BenchMarshal(b *testing.B, input interface{}, marshaler func(interface{}) []byte) {
	var bytes []byte

	for n := 0; n < b.N; n++ {
		bytes = marshaler(input)
	}

	marshalResult = bytes
}

func BenchUnmarshal(b *testing.B, input []byte, unmarshaler func([]byte) interface{}) {
	var output interface{}

	for n := 0; n < b.N; n++ {
		output = unmarshaler(input)
	}

	unmarshalResult = output
}

func BenchmarkMgoMarshalSmallStruct(b *testing.B) {
	BenchMarshal(b, &ASmallStruct, mgo.Marshaler)
}

func BenchmarkDriverMarshalSmallStruct(b *testing.B) {
	BenchMarshal(b, &ASmallStruct, driver.Marshaler)
}

func BenchmarkMgoMarshalLargeStruct(b *testing.B) {
	BenchMarshal(b, &ALargerStruct, mgo.Marshaler)
}

func BenchmarkDriverMarshalLargeStruct(b *testing.B) {
	BenchMarshal(b, &ALargerStruct, driver.Marshaler)
}

func BenchmarkMgoUnmarshalSmallStruct(b *testing.B) {
	BenchUnmarshal(b, mgo.Marshaler(&ASmallStruct), mgo.Unmarshaler)
}

func BenchmarkDriverUnmarshalSmallStruct(b *testing.B) {
	BenchUnmarshal(b, driver.Marshaler(&ASmallStruct), driver.Unmarshaler)
}

func BenchmarkMgoUnmarshalLargeStruct(b *testing.B) {
	BenchUnmarshal(b, mgo.Marshaler(&ALargerStruct), mgo.Unmarshaler)
}

func BenchmarkDriverUnmarshalLargeStruct(b *testing.B) {
	BenchUnmarshal(b, driver.Marshaler(&ALargerStruct), driver.Unmarshaler)
}
