package idiom_tokenizer

import (
	"../token"

	"bufio"
	"strings"

	"testing"
)

func TestReadtoken(t *testing.T) {

	tests := []struct{
		Text string
		Expected idiom_token.Type
	}{
		{
			Text: "",
			Expected: idiom_token.EOF(),
		},



		{
			Text: "\n",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO\n",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO\nTHREE",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO\nTHREE\n",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO\nTHREE\nFOUR",
			Expected: idiom_token.EOL(),
		},
		{
			Text: "\nONE TWO\nTHREE\n FOUR  FIVE",
			Expected: idiom_token.EOL(),
		},



		{
			Text:                        "apple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE\n",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:                        "apple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\n",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\nONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:                        "apple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\n",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                        "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:                       " apple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\n",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:                      "\tapple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       "\tapple BANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry dATE GrApE\n",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       "\tapple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                       "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                      "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:                  "\t\t  apple",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR",
			Expected: idiom_token.String("apple"),
		},
		{
			Text:                  "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR FIVE",
			Expected: idiom_token.String("apple"),
		},



		{
			Text:           "  \t \t \t\t O-N-E T-W-O T-H-R-E-E F-O-U-R F-I-V-E\n\n\t aaaaa    bbbb   ccc  dd e",
			Expected: idiom_token.String("O-N-E"),
		},
	}

	for testNumber, test := range tests {

		reader := strings.NewReader(test.Text)
		scanner := bufio.NewReader(reader)

		actual, err := readtoken(scanner)
		if nil != err {
			t.Errorf("For test #%d, actually got an error but did not expect one.", testNumber)
			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %s", err)
			continue
		}
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("TEXT: %q", test.Text)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
