package verboten

import (
	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct{
		Values []string
		Expected string
	}{
		{
			Values: []string{},
			Expected: "",
		},



		{
			Values: []string{""},
			Expected: "",
		},
		{
			Values: []string{"",""},
			Expected: "",
		},
		{
			Values: []string{"","",""},
			Expected: "",
		},
		{
			Values: []string{"","","",""},
			Expected: "",
		},
		{
			Values: []string{"","","","",""},
			Expected: "",
		},

		{
			Values: []string{"","","","","","","","","",""},
			Expected: "",
		},


		{
			Values: []string{"ap"},
			Expected: "ap",
		},
		{
			Values: []string{"ap","p"},
			Expected: "app",
		},
		{
			Values: []string{"ap","p","le"},
			Expected: "apple",
		},



		{
			Values: []string{"B"},
			Expected: "B",
		},
		{
			Values: []string{"B","ANA"},
			Expected: "BANA",
		},
		{
			Values: []string{"B","ANA","NA"},
			Expected: "BANANA",
		},



		{
			Values: []string{"Cherry"},
			Expected: "Cherry",
		},



		{
			Values: []string{"d"},
			Expected: "d",
		},
		{
			Values: []string{"d","A"},
			Expected: "dA",
		},
		{
			Values: []string{"d","A","t"},
			Expected: "dAt",
		},
		{
			Values: []string{"d","A","t","E"},
			Expected: "dAtE",
		},
	}

	for testNumber, test := range tests {

		actual, err := Run(test.Values...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %s", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
