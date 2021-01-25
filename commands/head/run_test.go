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
			Values: []idiom_string.Type{
				idiom_string.Something(""),
			},
			Expected: idiom_string.Something(""),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 2"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 3"),
			},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 4"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 5"),
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
				idiom_string.Something("A line on its own"),
			},
			Expected: idiom_string.Something("A"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("A line on its own\n"),
			},
			Expected: idiom_string.Something("A"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("A line on its own\n"),
				idiom_string.Something("ONE"),
				idiom_string.Something("TWO THREE"),
				idiom_string.Something("FOUR FIVE"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 4"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("APPLE    BANANA   CHERRY  DATE GRAPE"),
			},
			Expected: idiom_string.Something("APPLE"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("APPLE    BANANA   CHERRY  DATE GRAPE"),
				idiom_string.Something("ONE    TWO   THREE  FOUR FIVE\n"),
				idiom_string.Something("hello world!\u2028"),
				idiom_string.Something("2nd verse same as the first"),
			},
			Expected: idiom_string.Error("expected 1 parameter(s), but actually got 4"),
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
