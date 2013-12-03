// simple cal funcs test
package cal

import "testing"

func TestAdd(t *testing.T) {
	const a, b, result = 1, 2, 3
	if value := Add(a, b); value != result {
		t.Errorf("error!")
	}
}

func TestMinus(t *testing.T) {
	const a, b, result = 1, 2, -1
	if value := Minus(a, b); value != result {
		t.Errorf("error!")
	}
}

func TestTimes(t *testing.T) {
	const a, b, result = 1, 2, 2
	if value := Times(a, b); value != result {
		t.Errorf("error!")
	}
}

func TestDivide(t *testing.T) {
	const a, b, result = 4, 2, 2
	if value := Divide(a, b); value != result {
		t.Errorf("error!")
	}
}
