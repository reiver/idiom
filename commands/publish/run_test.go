package verboten

import (
	"../../string"

	"strings"

	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct{
		Values []idiom_string.Type
		Expected string
	}{
		{
			Values: []idiom_string.Type{},
			Expected: "",
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: "",
		},

		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: "",
		},


		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
			},
			Expected: "ap",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
				idiom_string.Something("le"),
			},
			Expected: "",
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
			},
			Expected: "B",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
				idiom_string.Something("NA"),
			},
			Expected: "",
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("Cherry"),
			},
			Expected: "Cherry",
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
			},
			Expected: "d",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
			},
			Expected: "",
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
				idiom_string.Something("E"),
			},
			Expected: "",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder

		result := run(&output, test.Values...)

		if 1 > len(test.Values) {
			if expected, actual := idiom_string.Error("expected 1 parameter(s), but actually got 0"), result; expected != actual {
				t.Errorf("For test #%d, expected a specific error, but did not actually get it.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("EXPECTED")
			}
			continue
		}
		if length := len(test.Values); 1 < length {
			if expected, actual := idiom_string.Errorf("expected 1 parameter(s), but actually got %d", length), result; expected != actual {
				t.Errorf("For test #%d, expected a specific error, but did not actually get it.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("EXPECTED")
			}
			continue
		}

		if expected, actual := test.Values[0], result; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, the actual value of what was outputted is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
