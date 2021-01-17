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
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},

		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Nothing(),
		},


		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
				idiom_string.Something("le"),
			},
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
				idiom_string.Something("NA"),
			},
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("Cherry"),
			},
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
				idiom_string.Something("E"),
			},
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Nothing(),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dAtE"),
				idiom_string.Nothing(),
				idiom_string.Error("something happened"),
			},
			Expected: idiom_string.Nothing(),
		},
	}

	for testNumber, test := range tests {

		actual := Run(test.Values...)
		if !actual.IsNothing() {
			t.Errorf("For test #%d, expectd the actual value to be nothing.", testNumber)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
