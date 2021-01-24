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
			Expected: idiom_string.Something(""),
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
			Expected: idiom_string.Something(""),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
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
				idiom_string.Something(""),
			},
			Expected: idiom_string.Something(""),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
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
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
				idiom_string.Something(""),
			},
			Expected: idiom_string.Something(""),
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
			Expected: idiom_string.Something("ap p"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("ap"),
				idiom_string.Something("p"),
				idiom_string.Something("le"),
			},
			Expected: idiom_string.Something("ap p le"),
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
			Expected: idiom_string.Something("B ANA"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("B"),
				idiom_string.Something("ANA"),
				idiom_string.Something("NA"),
			},
			Expected: idiom_string.Something("B ANA NA"),
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
			Expected: idiom_string.Something("d A"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
			},
			Expected: idiom_string.Something("d A t"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("d"),
				idiom_string.Something("A"),
				idiom_string.Something("t"),
				idiom_string.Something("E"),
			},
			Expected: idiom_string.Something("d A t E"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("A line on its own\n"),
				idiom_string.Something("ONE"),
				idiom_string.Something("TWO THREE"),
				idiom_string.Something("FOUR FIVE"),
			},
			Expected: idiom_string.Something("A line on its own\nONE TWO THREE FOUR FIVE"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("APPLE BANANA  CHERRY   DATE    GRAPE"),
				idiom_string.Something("ONE    TWO   THREE  FOUR FIVE\n"),
				idiom_string.Something("hello world!\u2028"),
				idiom_string.Something("2nd verse same as the first"),
			},
			Expected: idiom_string.Something("APPLE BANANA  CHERRY   DATE    GRAPE ONE    TWO   THREE  FOUR FIVE\nhello world!\u20282nd verse same as the first"),
		},
	}

	for testNumber, test := range tests {

		actual := Run(test.Values...)

		if actual.IsNothing() {
			t.Errorf("For test #%d, did not expect the actual value to be nothing.", testNumber)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}
		if actual.IsError() {
			t.Errorf("For test #%d, did not expect the actual value to be an error.", testNumber)
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
