package gohttp

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {
	u, err := url.Parse("http://usr:pwd@127.0.0.1:8080/some/path?with=args")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		Name     string
		Expected interface{}
		Value    interface{}
	}{
		{"Fragment", "", u.Fragment},
		{"Host", "127.0.0.1:8080", u.Host},
		{"Opaque", "", u.Opaque},
		{"Path", "/some/path", u.Path},
		{"RawQuery", "with=args", u.RawQuery},
		{"RequestURI", "/some/path?with=args", u.RequestURI()},
		{"Scheme", "http", u.Scheme},
		{"String", "http://usr:pwd@127.0.0.1:8080/some/path?with=args", u.String()},
		{"User", "&url.Userinfo{username:\"usr\", password:\"pwd\", passwordSet:true}", fmt.Sprintf("%#v", u.User)},
	}

	for _, tst := range tests {
		if tst.Expected != tst.Value {
			t.Errorf("expected %s to be %#v, was %#v", tst.Name, tst.Expected, tst.Value)
		}
	}

}
