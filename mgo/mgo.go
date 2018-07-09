package mgo

import (
	"github.com/dunkyboy/mongo-driver-bench/benchmark"
	"gopkg.in/mgo.v2/bson"
)

var Marshaler = func(in interface{}) []byte {
	out, err := bson.Marshal(
		in)
	if err != nil {
		panic(err)
	}
	return out
}

var Unmarshaler = func(input []byte) interface{} {
	var output benchmark.SmallStructDepth1
	err := bson.Unmarshal(input, &output)
	if err != nil {
		panic(err)
	}
	return output
}
