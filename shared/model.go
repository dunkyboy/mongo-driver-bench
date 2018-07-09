package shared

import "time"

type SmallStruct struct {
	ANumber int64
	ATime time.Time
	SomeText string
}

type LargerStruct struct {
	ANumber1 int64
	ATime1 time.Time
	SomeText1 string
	ANumber2 int64
	ATime2 time.Time
	SomeText2 string
	ABool bool
	ATime3 time.Time
	SomeText3 string
	ANumber4 int64
	ATime4 time.Time
	SomeText4 string
	ANumber5 int64
	ATime5 time.Time
	SomeText5 string
	ANumber6 int64
	ATime6 time.Time
	SomeText6 string
	AFloat float64
	ATime7 time.Time
	SomeText7 string
	ANumber uint32
	ATime8 time.Time
	SomeText8 string
	ANumber9 int64
	ATime9 time.Time
	SomeText9 string
	AByte byte
	ATime10 time.Time
	SomeText10 string
}

type SmallStructDepth9 struct {
	Other SmallStruct
	Nested smallNestedStruct2
}

type smallNestedStruct2 struct {
	Other SmallStruct
	Nested smallNestedStruct3
}

type smallNestedStruct3 struct {
	Other SmallStruct
	Nested smallNestedStruct4
}

type smallNestedStruct4 struct {
	Other SmallStruct
	Nested smallNestedStruct5
}

type smallNestedStruct5 struct {
	Other SmallStruct
	Nested smallNestedStruct6
}

type smallNestedStruct6 struct {
	Other SmallStruct
	Nested smallNestedStruct7
}

type smallNestedStruct7 struct {
	Other SmallStruct
	Nested smallNestedStruct8
}

type smallNestedStruct8 struct {
	Other SmallStruct
	Nested smallNestedStruct9
}

type smallNestedStruct9 struct {
	Other SmallStruct
	Nested SmallStruct
}

type LargerStructDepth9 struct {
	ANumber1 int64
	ATime1 time.Time
	SomeText1 string
	ANumber2 int64
	ATime2 time.Time
	SomeText2 string
	ABool bool
	ATime3 time.Time
	SomeText3 string
	ANumber4 int64
	ATime4 time.Time
	SomeText4 string
	ANumber5 int64
	ATime5 time.Time
	SomeText5 string
	ANumber6 int64
	ATime6 time.Time
	SomeText6 string
	AFloat float64
	ATime7 time.Time
	SomeText7 string
	ANumber uint32
	ATime8 time.Time
	SomeText8 string
	ANumber9 int64
	ATime9 time.Time
	SomeText9 string
	AByte byte
	ATime10 time.Time
	SomeText10 string
	Nested largerNestedStruct2
}

type largerNestedStruct2 struct {
	Other LargerStruct
	Nested largerNestedStruct3
}

type largerNestedStruct3 struct {
	Other LargerStruct
	Nested largerNestedStruct4
}

type largerNestedStruct4 struct {
	Other LargerStruct
	Nested largerNestedStruct5
}

type largerNestedStruct5 struct {
	Other LargerStruct
	Nested largerNestedStruct6
}

type largerNestedStruct6 struct {
	Other LargerStruct
	Nested largerNestedStruct7
}

type largerNestedStruct7 struct {
	Other LargerStruct
	Nested largerNestedStruct8
}

type largerNestedStruct8 struct {
	Other LargerStruct
	Nested largerNestedStruct9
}

type largerNestedStruct9 struct {
	Other LargerStruct
	Nested LargerStruct
}
