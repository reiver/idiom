package idiom_string

import (
	"math/rand"
	"time"

	"testing"
)

func TestSomething(t *testing.T) {

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
		s := randomString()
		str := Something(s)

		if nothing := Nothing(); nothing == str {
			t.Error("An errored idiom_string.Type is equal to idiom_string.Nothing()")
			t.Logf("s -> %q", s)
			t.Logf("                 str -> %#v", str)
			t.Logf("idiom_string.Nothing -> %#v", nothing)
			continue
		}

		value, err := str.Return()

		if nil != err {
			t.Error("The actual returned error is not what was expected.")
			t.Logf("EXPECTED: %#v", nil)
			t.Logf("ACTUAL:   %#v", err)
			continue
		}

		if expected, actual := s, value; expected != actual {
			t.Error("The actual string is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
