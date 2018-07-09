package driver

import (
	"github.com/mongodb/mongo-go-driver/bson"
)

var Marshaler = func(in interface{}) []byte {
	out, err := bson.Marshal(in)
	if err != nil {
		panic(err)
	}
	return out
}

var Unmarshaler = func(input []byte, output interface{}) (success bool) {
	err := bson.Unmarshal(input, output)
	if err != nil {
		panic(err)
	}
	return err != nil
}
