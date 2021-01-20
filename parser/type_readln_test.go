package idiom_parser

import (
	"../token"

	"bufio"
	"reflect"
	"strings"

	"testing"
)

func TestType_Readln(t *testing.T) {

	tests := []struct{
		Text string
		Expected []idiom_token.Type
	}{
		{
			Text: "",
			Expected: []idiom_token.Type{
				idiom_token.EOF(),
			},
		},



		{
			Text: "\n",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO\n",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\n",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\nFOUR",
			Expected: []idiom_token.Type(nil),
		},
		{
			Text: "\nONE TWO\nTHREE\n FOUR  FIVE",
			Expected: []idiom_token.Type(nil),
		},



		{
			Text: "apple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: "apple BANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: "apple BANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: "apple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: "apple\tBANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: "apple\tBANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\n",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: "apple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: "apple  BANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: "apple  BANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\n",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: " apple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: " apple BANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: " apple BANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: "\tapple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: "\tapple BANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: "\tapple BANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\n",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: "\t\t  apple",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
			},
		},
		{
			Text: "\t\t  apple BANANA",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},
		{
			Text: "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR FIVE",
			Expected: []idiom_token.Type{
				idiom_token.String("apple"),
				idiom_token.String("BANANA"),
				idiom_token.String("Cherry"),
				idiom_token.String("dATE"),
				idiom_token.String("GrApE"),
			},
		},



		{
			Text: "  \t \t \t\t O-N-E T-W-O T-H-R-E-E F-O-U-R F-I-V-E\n\n\t aaaaa    bbbb   ccc  dd e",
			Expected: []idiom_token.Type{
				idiom_token.String("O-N-E"),
				idiom_token.String("T-W-O"),
				idiom_token.String("T-H-R-E-E"),
				idiom_token.String("F-O-U-R"),
				idiom_token.String("F-I-V-E"),
			},
		},



		{
			Text: "cat file.txt | grep dog",
			Expected: []idiom_token.Type{
				idiom_token.String("cat"),
				idiom_token.String("file.txt"),
			},
		},
	}

	for testNumber, test := range tests {

		reader := strings.NewReader(test.Text)
		scanner := bufio.NewReader(reader)
		parser := New(scanner)

		actual, err := parser.Readln()
		if nil != err {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q",      err)
			continue
		}
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
	                  "wow"                                                      + "\n"     +
	                  "cat file.txt | grep dog"                                  + "\n"     +
	                  "hello-world"                                              + "\n"     +
	                  "| filter aaa bb c"                                        + "\n"

	reader := strings.NewReader(text)
	scanner := bufio.NewReader(reader)
	parser := New(scanner)

	actual1,  err1  := parser.Readln()
	actual2,  err2  := parser.Readln()
	actual3,  err3  := parser.Readln()
	actual4,  err4  := parser.Readln()
	actual5,  err5  := parser.Readln()
	actual6,  err6  := parser.Readln()
	actual7,  err7  := parser.Readln()
	actual8,  err8  := parser.Readln()
	actual9,  err9  := parser.Readln()
	actual10, err10 := parser.Readln()

	if nil != err1 {
		t.Errorf("Actually got unexpected error with 1st read.")
		t.Logf("ERROR TYPE: %T", err1)
		t.Logf("ERROR: %q",      err1)
		return
	}
	if nil != err2 {
		t.Errorf("Actually got unexpected error with 2nd read.")
		t.Logf("ERROR TYPE: %T", err2)
		t.Logf("ERROR: %q",      err2)
		return
	}
	if nil != err3 {
		t.Errorf("Actually got unexpected error with 3rd read.")
		t.Logf("ERROR TYPE: %T", err3)
		t.Logf("ERROR: %q",      err3)
		return
	}
	if nil != err4 {
		t.Errorf("Actually got unexpected error with 4th read.")
		t.Logf("ERROR TYPE: %T", err4)
		t.Logf("ERROR: %q",      err4)
		return
	}
	if nil != err5 {
		t.Errorf("Actually got unexpected error with 5th read.")
		t.Logf("ERROR TYPE: %T", err5)
		t.Logf("ERROR: %q",      err5)
		return
	}
	if nil != err6 {
		t.Errorf("Actually got unexpected error with 6th read.")
		t.Logf("ERROR TYPE: %T", err6)
		t.Logf("ERROR: %q",      err6)
		return
	}
	if nil != err7 {
		t.Errorf("Actually got unexpected error with 7th read.")
		t.Logf("ERROR TYPE: %T", err7)
		t.Logf("ERROR: %q",      err7)
		return
	}
	if nil != err8 {
		t.Errorf("Actually got unexpected error with 8th read.")
		t.Logf("ERROR TYPE: %T", err8)
		t.Logf("ERROR: %q",      err8)
		return
	}
	if nil != err9 {
		t.Errorf("Actually got unexpected error with 9th read.")
		t.Logf("ERROR TYPE: %T", err8)
		t.Logf("ERROR: %q",      err8)
		return
	}
	if nil != err10 {
		t.Errorf("Actually got unexpected error with 10th read.")
		t.Logf("ERROR TYPE: %T", err8)
		t.Logf("ERROR: %q",      err8)
		return
	}

	if expected, actual := []idiom_token.Type{
		idiom_token.String("ONE"),
		idiom_token.String("TWO"),
		idiom_token.String("THREE"),
		idiom_token.String("FOUR"),
		idiom_token.String("FIVE"),
	}, actual1; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 1st read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{
		idiom_token.String("apple"),
		idiom_token.String("BANANA"),
		idiom_token.String("Cherry"),
		idiom_token.String("dATE"),
		idiom_token.String("GrApE"),
	}, actual2; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 2nd read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type(nil), actual3; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 3rd read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{idiom_token.String("wow")}, actual4; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 4th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{
			idiom_token.String("cat"),
			idiom_token.String("file.txt"),
	}, actual5; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 5th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{
			idiom_token.String("|"),
			idiom_token.String("grep"),
			idiom_token.String("dog"),
	}, actual6; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 6th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{
			idiom_token.String("hello-world"),
	}, actual7; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 7th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{
			idiom_token.String("|"),
			idiom_token.String("filter"),
			idiom_token.String("aaa"),
			idiom_token.String("bb"),
			idiom_token.String("c"),
	}, actual8; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 8th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{idiom_token.EOF()}, actual9; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 9th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
	if expected, actual := []idiom_token.Type{idiom_token.EOF()}, actual10; !reflect.DeepEqual(expected,actual) {
		t.Error("Actual value of 10th read is not what was expected.")
		t.Logf("EXPECTED: %#v", expected)
		t.Logf("ACTUAL:   %#v", actual)
	}
}
