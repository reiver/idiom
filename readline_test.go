package main

import (
	"./string"

	"bufio"
	"reflect"
	"strings"

	"testing"
)

func TestReadline(t *testing.T) {

	tests := []struct{
		Text string
		Expected []idiom_string.Type
	}{
		{
			Text: "",
			Expected: []idiom_string.Type{
				idiom_string.Error("end of file"),
			},
		},



		{
			Text: "\n",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO\n",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\n",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\nFOUR",
			Expected: []idiom_string.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\n FOUR  FIVE",
			Expected: []idiom_string.Type(nil),
		},



		{
			Text: "apple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: "apple BANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: "apple BANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: "apple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: "apple\tBANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: "apple\tBANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\n",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: "apple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: "apple  BANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: "apple  BANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\n",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: " apple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: " apple BANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: " apple BANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: "\tapple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: "\tapple BANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: "\tapple BANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: "\t\t  apple",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
			},
		},
		{
			Text: "\t\t  apple BANANA",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR FIVE",
			Expected: []idiom_string.Type{
				idiom_string.Something("apple"),
				idiom_string.Something("BANANA"),
				idiom_string.Something("Cherry"),
				idiom_string.Something("dATE"),
				idiom_string.Something("GrApE"),
			},
		},



		{
			Text: "  \t \t \t\t O-N-E T-W-O T-H-R-E-E F-O-U-R F-I-V-E\n\n\t aaaaa    bbbb   ccc  dd e",
			Expected: []idiom_string.Type{
				idiom_string.Something("O-N-E"),
				idiom_string.Something("T-W-O"),
				idiom_string.Something("T-H-R-E-E"),
				idiom_string.Something("F-O-U-R"),
				idiom_string.Something("F-I-V-E"),
			},
		},
	}

	for testNumber, test := range tests {

		reader := strings.NewReader(test.Text)
		scanner := bufio.NewReader(reader)

		actual := readline(scanner)
		if expected := test.Expected; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("TEXT: %q", test.Text)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}

func TestReadline_many(t *testing.T) {

	var text string = "ONE TWO THREE FOUR FIVE"                                  + "\u2028" +
	                  "\tapple\t\tBANANA\t\t\tCherry\t\t\t\tdATE\t\t\t\t\tGrApE" + "\u2028" +
	                  ""                                                         + "\u2028" +
	                  "wow"                                                      + "\n"

	reader := strings.NewReader(text)
	scanner := bufio.NewReader(reader)

	actual1 := readline(scanner)
	actual2 := readline(scanner)
	actual3 := readline(scanner)
	actual4 := readline(scanner)
	actual5 := readline(scanner)
	actual6 := readline(scanner)

	if expected, actual := []idiom_string.Type{
		idiom_string.Something("ONE"),
		idiom_string.Something("TWO"),
		idiom_string.Something("THREE"),
		idiom_string.Something("FOUR"),
		idiom_string.Something("FIVE"),
	}, actual1; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 1st read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_string.Type{
		idiom_string.Something("apple"),
		idiom_string.Something("BANANA"),
		idiom_string.Something("Cherry"),
		idiom_string.Something("dATE"),
		idiom_string.Something("GrApE"),
	}, actual2; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 2nd read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_string.Type(nil), actual3; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 3rd read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_string.Type{idiom_string.Something("wow")}, actual4; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 4th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_string.Type{idiom_string.Error("end of file")}, actual5; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 5th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_string.Type{idiom_string.Error("end of file")}, actual6; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 6th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
}
