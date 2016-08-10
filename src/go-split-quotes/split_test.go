package split

import (
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestSplit(t *testing.T) {
	withQuotes, err := shellwords.Parse(`a test "with this" 'and this' with final`)
	if err != nil {
		t.Fatal(err)
	}
	withoutQuotes, err := shellwords.Parse("a test")
	if err != nil {
		t.Fatal(err)
	}
	withEscapedQuotes, err := shellwords.Parse(`a "test \" with quote" test`)
	if err != nil {
		t.Fatal(err)
	}
	withOnlyEscapedQuotes, err := shellwords.Parse(`a \" test`)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct{ Has, Want interface{} }{
		{len(withQuotes), 6},
		{withQuotes[2], `with this`},
		{withQuotes[3], `and this`},
		{len(withoutQuotes), 2},
		{withoutQuotes[1], "test"},
		{len(withEscapedQuotes), 3},
		{withEscapedQuotes[0], "a"},
		{withEscapedQuotes[1], `test " with quote`},
		{withEscapedQuotes[2], "test"},
		{withOnlyEscapedQuotes[0], "a"},
		{withOnlyEscapedQuotes[1], `"`},
		{withOnlyEscapedQuotes[2], "test"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i+1, tc.Want, tc.Has)
		}
	}
}
