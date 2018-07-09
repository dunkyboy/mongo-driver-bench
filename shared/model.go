package shared

import "time"

type SmallStruct struct {
	aNumber int64
	aTime time.Time
	someText string
}

type LargerStruct struct {
	aNumber1 int64
	aTime1 time.Time
	someText1 string
	aNumber2 int64
	aTime2 time.Time
	someText2 string
	aBool bool
	aTime3 time.Time
	someText3 string
	aNumber4 int64
	aTime4 time.Time
	someText4 string
	aNumber5 int64
	aTime5 time.Time
	someText5 string
	aNumber6 int64
	aTime6 time.Time
	someText6 string
	aFloat float64
	aTime7 time.Time
	someText7 string
	bytes [16]byte
	aTime8 time.Time
	someText8 string
	aNumber9 int64
	aTime9 time.Time
	someText9 string
	aByte byte
	aTime10 time.Time
	someText10 string
}

type SmallStructDepth9 struct {
	someText string
	nested  smallNestedStruct2
}

type smallNestedStruct2 struct {
	aTime time.Time
	nested smallNestedStruct3
}

type smallNestedStruct3 struct {
	aBool bool
	nested smallNestedStruct4
}

type smallNestedStruct4 struct {
	bytes [16]byte
	nested smallNestedStruct5
}

type smallNestedStruct5 struct {
	aByte byte
	nested smallNestedStruct6
}

type smallNestedStruct6 struct {
	aTime time.Time
	nested smallNestedStruct7
}

type smallNestedStruct7 struct {
	aFloat float64
	nested smallNestedStruct8
}

type smallNestedStruct8 struct {
	aTime time.Time
	nested smallNestedStruct9
}

type smallNestedStruct9 struct {
	aTime time.Time
	nested SmallStruct
}
