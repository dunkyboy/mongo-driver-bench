package driver

import (
	"github.com/dunkyboy/mongo-driver-bench/shared"
	"github.com/mongodb/mongo-go-driver/bson"
)

var Marshaler = func(in interface{}) []byte {
	out, err := bson.Marshal(in)
	if err != nil {
		panic(err)
	}
	return out
}

var Unmarshaler = func(input []byte) interface{} {
	var output shared.SmallStruct
	err := bson.Unmarshal(input, &output)
	if err != nil {
		panic(err)
	}
	return output
}
