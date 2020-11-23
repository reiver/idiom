package verboten

import (
	"strings"

	"testing"
)

func TestRun(t *testing.T) {

	tests := []struct{
		Values []string
		Expected string
	}{
		{
			Values: []string{},
			Expected: "«»\n",
		},



		{
			Values: []string{""},
			Expected: "«»\n",
		},
		{
			Values: []string{"",""},
			Expected: "«»\n",
		},
		{
			Values: []string{"","",""},
			Expected: "«»\n",
		},
		{
			Values: []string{"","","",""},
			Expected: "«»\n",
		},
		{
			Values: []string{"","","","",""},
			Expected: "«»\n",
		},

		{
			Values: []string{"","","","","","","","","",""},
			Expected: "«»\n",
		},


		{
			Values: []string{"ap"},
			Expected: "«ap»\n",
		},
		{
			Values: []string{"ap","p"},
			Expected: "«ap p»\n",
		},
		{
			Values: []string{"ap","p","le"},
			Expected: "«ap p le»\n",
		},



		{
			Values: []string{"B"},
			Expected: "«B»\n",
		},
		{
			Values: []string{"B","ANA"},
			Expected: "«B ANA»\n",
		},
		{
			Values: []string{"B","ANA","NA"},
			Expected: "«B ANA NA»\n",
		},



		{
			Values: []string{"Cherry"},
			Expected: "«Cherry»\n",
		},



		{
			Values: []string{"d"},
			Expected: "«d»\n",
		},
		{
			Values: []string{"d","A"},
			Expected: "«d A»\n",
		},
		{
			Values: []string{"d","A","t"},
			Expected: "«d A t»\n",
		},
		{
			Values: []string{"d","A","t","E"},
			Expected: "«d A t E»\n",
		},
	}

	for testNumber, test := range tests {

		var output strings.Builder

		result, err := run(&output, test.Values...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %s", testNumber, err, err)
			continue
		}

		if expected, actual := "", result; expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

		if expected, actual := test.Expected, output.String(); expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
