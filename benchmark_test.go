package benchmark

import (
	"github.com/dunkyboy/mongo-driver-bench/mgo"
	"github.com/dunkyboy/mongo-driver-bench/mongo-go-driver"
	"testing"
	"github.com/dunkyboy/mongo-driver-bench/shared"
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
	BenchMarshal(b, shared.ASmallStruct, mgo.Marshaler)
}

func BenchmarkMgoMarshalSmallNestedStruct(b *testing.B) {
	BenchMarshal(b, shared.ASmallNestedStruct, mgo.Marshaler)
}

func BenchmarkMgoMarshalLargeStruct(b *testing.B) {
	BenchMarshal(b, shared.ALargerStruct, mgo.Marshaler)
}

func BenchmarkMgoUnmarshalSmallStruct(b *testing.B) {
	BenchUnmarshal(b, mgo.Marshaler(shared.ASmallStruct), mgo.Unmarshaler)
}

func BenchmarkMgoUnmarshalSmallNestedStruct(b *testing.B) {
	BenchUnmarshal(b, mgo.Marshaler(shared.ASmallNestedStruct), mgo.Unmarshaler)
}

func BenchmarkMgoUnmarshalLargeStruct(b *testing.B) {
	BenchUnmarshal(b, mgo.Marshaler(shared.ALargerStruct), mgo.Unmarshaler)
}

func BenchmarkDriverMarshalSmallStruct(b *testing.B) {
	BenchMarshal(b, shared.ASmallStruct, driver.Marshaler)
}

func BenchmarkDriverMarshalSmallNestedStruct(b *testing.B) {
	BenchMarshal(b, shared.ASmallNestedStruct, driver.Marshaler)
}

func BenchmarkDriverMarshalLargeStruct(b *testing.B) {
	BenchMarshal(b, shared.ALargerStruct, driver.Marshaler)
}

func BenchmarkDriverUnmarshalSmallStruct(b *testing.B) {
	BenchUnmarshal(b, driver.Marshaler(shared.ASmallStruct), driver.Unmarshaler)
}

func BenchmarkDriverUnmarshalSmallNestedStruct(b *testing.B) {
	BenchUnmarshal(b, driver.Marshaler(shared.ASmallNestedStruct), driver.Unmarshaler)
}

func BenchmarkDriverUnmarshalLargeStruct(b *testing.B) {
	BenchUnmarshal(b, driver.Marshaler(shared.ALargerStruct), driver.Unmarshaler)
}


