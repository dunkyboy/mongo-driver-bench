package shared

var textMediumLength = "The quick brown fox jumps over the lazy dog. Then it did it again. Keep typing a " +
	"little longer so we get a string that's somewhere between a short string (like a hostname perhaps) and a long " +
	"text blob"

// zero vallues are fine for everything but strings

var ASmallStructDepth1 = SmallStructDepth1{
	someText: textMediumLength,
}