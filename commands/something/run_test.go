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
			Values: []idiom_string.Type(nil),
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 0"),
		},
		{
			Values: []idiom_string.Type{},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 0"),
		},



		{
			Values: []idiom_string.Type{idiom_string.Nothing()},
			Expected:                   idiom_string.Error("not something"),
		},



		{
			Values: []idiom_string.Type{idiom_string.Error("just some error")},
			Expected:                   idiom_string.Error("not something"),
		},



		{
			Values: []idiom_string.Type{idiom_string.Something("")},
			Expected:                   idiom_string.Something(""),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 2"),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 3"),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 4"),
		},
		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 5"),
		},

		{
			Values: []idiom_string.Type{idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something(""),idiom_string.Something("")},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 10"),
		},


		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
			},
			Expected: idiom_string.Something("ap"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 2"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
				idiom_string.Something("le"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 3"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
			},
			Expected: idiom_string.Something("B"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 2"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
				idiom_string.Something("NA"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 3"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("Cherry"),
			},
			Expected: idiom_string.Something("Cherry"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
			},
			Expected: idiom_string.Something("d"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 2"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 3"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
				idiom_string.Something("E"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 4"),
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
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 7"),
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
