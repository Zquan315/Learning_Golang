package test

import (
	"testing"
)

// func TestMultiple(t *testing.T) {
// 	exp := 55
// 	a := 11
// 	b := 5
// 	res := multiple(a, b)
// 	if res != exp {
// 		t.Errorf("%d * %d != %d", a, b, exp)
// 	}
// }

func TestMultipleWithTable(t *testing.T) {
	multiples := []struct {
		a, b int
		exp  int
	}{
		{2, 1, 2},
		{10, 5, 50},
		{6, 5, 30},
	}

	for _, v := range multiples {
		res := multiple(v.a, v.b)
		if res != v.exp {
			t.Errorf("%d * %d != %d", v.a, v.b, v.exp)
		}
	}
}
