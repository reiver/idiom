package idiom_string

import (
	"math/rand"
	"time"

	"testing"
)

func TestError(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
	runes := []rune(".0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz          ")
	lenRunes := len(runes)

	randomString := func() string {
		length := 1 + randomness.Intn(59)
		rs := make([]rune, length)
		for i := range rs {
			rs[i] = runes[randomness.Intn(lenRunes)]
		}

		return string(rs)
	}

	const limit = 20
	for testNumber:=0; testNumber<limit; testNumber++ {
		msg := randomString()
		str := Error(msg)

		if nothing := Nothing(); nothing == str {
			t.Error("An errored idiom_string.Type is equal to idiom_string.Nothing()")
			t.Logf("msg -> %q", msg)
			t.Logf("                 str -> %#v", str)
			t.Logf("idiom_string.Nothing -> %#v", nothing)
			continue
		}

		value, err := str.Unwrap()

		if expected, actual := "", value; expected != actual {
			t.Error("The actual returned string is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

		if expected, actual := msg, err.Error(); expected != actual {
			t.Error("The actual error message is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}


	}
}
