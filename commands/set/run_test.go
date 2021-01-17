package verboten

import (
	"../../string"

	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct{
		Values []idiom_string.Type
		Expected idiom_string.Type
	}{
		{
			Values: []idiom_string.Type{},
			Expected: idiom_string.Error("too few parameters"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("too few parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Something(""),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("too many parameters"),
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
			Expected: idiom_string.Error("too many parameters"),
		},


		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
			},
			Expected: idiom_string.Error("too few parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
			},
			Expected: idiom_string.Something("p"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
				idiom_string.Something("le"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
			},
			Expected: idiom_string.Error("too few parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
			},
			Expected: idiom_string.Something("ANA"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
				idiom_string.Something("NA"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("Cherry"),
			},
			Expected: idiom_string.Error("too few parameters"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
			},
			Expected: idiom_string.Error("too few parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
			},
			Expected: idiom_string.Something("A"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
				idiom_string.Something("E"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
	}

	for testNumber, test := range tests {

		actual := Run(test.Values...)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
