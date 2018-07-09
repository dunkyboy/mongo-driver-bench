package mgo

import (
	"github.com/dunkyboy/mongo-driver-bench/shared"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

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

	input, err := bson.Marshal(&shared.ASmallStructDepth1)
	if err != nil {
		panic(err)
	}

	shared.BenchUnmarshal(
		b,
		func() interface{} {
			var output shared.SmallStructDepth1
			err := bson.Unmarshal(input, &output)
			if err != nil {
				panic(err)
			}
			return output
		},
	)
}
