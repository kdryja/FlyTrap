package flytrap

import (
	"testing"
)

func TestCountry(t *testing.T) {
	want := [2]byte{'G', 'B'}
	in := "137.50.144.133"
	if got := country(in); got != want {
		t.Errorf("country(%s) = got %s, want %s", in, got, want)
	}
}
