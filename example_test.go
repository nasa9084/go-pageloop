package pageloop_test

import (
	"fmt"
	"strings"

	pageloop "github.com/nasa9084/go-pageloop"
)

func ExamplePageLoop() {
	target := []string{
		"foo", "bar", "baz", "qux", "quux",
		"corge", "grault", "garply", "waldo", "fred",
		"plugh", "xyzzy", "thud",
	}
	for pager := pageloop.NewPager(5, len(target)); pager.Next(); {
		begin, end := pager.Index()
		fmt.Printf("%s\n", strings.Join(target[begin:end], ","))
	}

	// Output:
	// foo,bar,baz,qux,quux
	// corge,grault,garply,waldo,fred
	// plugh,xyzzy,thud
}
