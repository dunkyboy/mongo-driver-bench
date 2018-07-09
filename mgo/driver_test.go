package mgo

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"github.com/dunkyboy/mongo-driver-bench/shared"
)

// save the test output somewhere global to force the compiler not to optimize it all away
var marshalResult []byte
var unmarshalResult interface{}

func BenchmarkMarshalSmallStructDepth1(b *testing.B) {
	var bytes []byte
	var er error

	for n := 0; n < b.N; n++ {
		bytes, er = bson.Marshal(shared.ASmallStructDepth1)
		if er != nil {
			panic(er)
		}
	}

	marshalResult = bytes
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
