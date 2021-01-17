package verboten

import (
	"../../kernel"
	"../../string"

	"testing"
)

func TestRun(t *testing.T) {

	idiom_kernel.Variables.Set("apple",  "one")
	idiom_kernel.Variables.Set("BANANA", "two")
	idiom_kernel.Variables.Set("Cherry", "three")
	idiom_kernel.Variables.Set("dATE",   "four")


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
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
			Expected: idiom_string.Error("too many parameters"),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("abcde"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("qwerty"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("APPLE"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("banana"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("cHERRY"),
			},
			Expected: idiom_string.Nothing(),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("Date"),
			},
			Expected: idiom_string.Nothing(),
		},



		{
			Values: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
			Expected: idiom_string.Something("one"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("BANANA"),
			},
			Expected: idiom_string.Something("two"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("Cherry"),
			},
			Expected: idiom_string.Something("three"),
		},
		{
			Values: []idiom_string.Type{
				idiom_string.Something("dATE"),
			},
			Expected: idiom_string.Something("four"),
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
