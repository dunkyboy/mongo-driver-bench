package shared

var textMediumLength = "The quick brown fox jumps over the lazy dog. Then it did it again. Keep typing a " +
	"little longer so we get a string that's somewhere between a short string (like a hostname perhaps) and a long " +
	"text blob"

// zero values are fine for everything but strings

var ASmallStruct = SmallStruct{
	SomeText: textMediumLength,
}

var ALargerStruct = LargerStruct{
	SomeText1:  textMediumLength,
	SomeText2:  textMediumLength,
	SomeText3:  textMediumLength,
	SomeText4:  textMediumLength,
	SomeText5:  textMediumLength,
	SomeText6:  textMediumLength,
	SomeText7:  textMediumLength,
	SomeText8:  textMediumLength,
	SomeText9:  textMediumLength,
	SomeText10: textMediumLength,
}

var ASmallNestedStruct = SmallStructDepth9{
	Other: ASmallStruct,
	// nested structs will have empty strings, meh
}

var ALargerNestedStruct = LargerStructDepth9{
	SomeText1:  textMediumLength,
	SomeText2:  textMediumLength,
	SomeText3:  textMediumLength,
	SomeText4:  textMediumLength,
	SomeText5:  textMediumLength,
	SomeText6:  textMediumLength,
	SomeText7:  textMediumLength,
	SomeText8:  textMediumLength,
	SomeText9:  textMediumLength,
	SomeText10: textMediumLength,
	// nested structs will have empty strings, meh
}
