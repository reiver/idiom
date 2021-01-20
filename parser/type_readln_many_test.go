package idiom_parser

import (
	"../token"

	"reflect"
	"strings"

	"testing"
)

func TestType_Readln_many(t *testing.T) {

	var s string = "ONE two THREE four FIVE" + "\n"     +
	               ""                        + "\n"     +
	               "cat file.txt | grep dog" + "\u2028" +
	               "  hello-world   "        + "\n"     +
	               "  |  grep  target "      + "\n"

	reader := strings.NewReader(s)

	parser := New(reader)

	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 1st read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.String("ONE"),
			idiom_token.String("two"),
			idiom_token.String("THREE"),
			idiom_token.String("four"),
			idiom_token.String("FIVE"),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 1st read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 2nd read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type(nil), tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 2nd read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 3rd read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.String("cat"),
			idiom_token.String("file.txt"),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 3rd read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 4th read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.String("|"),
			idiom_token.String("grep"),
			idiom_token.String("dog"),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 4th read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 5th read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.String("hello-world"),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 5th read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 6th read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.String("|"),
			idiom_token.String("grep"),
			idiom_token.String("target"),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 6th read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 7th read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.EOF(),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 7th read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
	{
		tokens, err := parser.Readln()
		if nil != err {
			t.Error("Did not actually expect an error with the 8th read-line.")
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %#v", err)
			return
		}
		if expected, actual := []idiom_token.Type{
			idiom_token.EOF(),
		}, tokens; !reflect.DeepEqual(expected, actual) {
			t.Error("Actual tokens are not what was expected for 8th read-line.")
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			return
		}
	}
}
