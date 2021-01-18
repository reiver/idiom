package main

import (
	"./string"

	"bufio"
	"strings"

	"testing"
)

func TestReadtoken(t *testing.T) {

	tests := []struct{
		Text string
		Expected idiom_string.Type
	}{
		{
			Text:                        "",
			Expected: idiom_string.Error("end of file"),
		},



		{
			Text:                        "\n",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO\n",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO\nTHREE",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO\nTHREE\n",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO\nTHREE\nFOUR",
			Expected: idiom_string.Error("end of line"),
		},
		{
			Text:                        "\nONE TWO\nTHREE\n FOUR  FIVE",
			Expected: idiom_string.Error("end of line"),
		},



		{
			Text:                            "apple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\n",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            "apple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\n",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\nONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple\tBANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            "apple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\n",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "apple  BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            " apple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\n",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            " apple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            "\tapple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\n",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\nONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\nONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\tapple BANANA Cherry dATE GrApE\nONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            "\t\t  apple",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR",
			Expected: idiom_string.Something("apple"),
		},
		{
			Text:                            "\t\t  apple BANANA Cherry dATE GrApE\u2028ONE TWO THREE FOUR FIVE",
			Expected: idiom_string.Something("apple"),
		},



		{
			Text:                            "  \t \t \t\t O-N-E T-W-O T-H-R-E-E F-O-U-R F-I-V-E\n\n\t aaaaa    bbbb   ccc  dd e",
			Expected: idiom_string.Something("O-N-E"),
		},
	}

	for testNumber, test := range tests {

		reader := strings.NewReader(test.Text)
		scanner := bufio.NewReader(reader)

		actual := readtoken(scanner)
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("TEXT: %q", test.Text)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
