package pageloop_test

import (
	"reflect"
	"testing"

	pageloop "github.com/nasa9084/go-pageloop"
)

func TestPageLoop(t *testing.T) {
	pager := pageloop.NewPager(10, 45)

	got := [][2]int{}
	for pager.Next() {
		begin, end := pager.Index()
		got = append(got, [2]int{begin, end})
	}

	want := [][2]int{
		{0, 10},
		{10, 20},
		{20, 30},
		{30, 40},
		{40, 45},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected:\n  got:  %#v\n  want: %#v", got, want)
		return
	}
}

func TestZeroLength(t *testing.T) {
	pager := pageloop.NewPager(10, 0)

	if pager.Next() {
		t.Error("unexpected cond: true")
		return
	}
}

func TestZeroPerPage(t *testing.T) {
	pager := pageloop.NewPager(0, 45)

	if !pager.Next() {
		t.Errorf("unexpected cond: false")
		return
	}
	begin, end := pager.Index()

	if begin != 0 {
		t.Errorf("unexpected begin at first time: %d != 0", begin)
		return
	}
	if end != 45 {
		t.Errorf("unexpected end at first time: %d != 45", end)
		return
	}

	if pager.Next() {
		t.Errorf("unexpected cond: true")
		return
	}
}
