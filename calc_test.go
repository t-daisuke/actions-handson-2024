package handson_test

import (
	"handson"
	"testing"
)

func TestAdd(t *testing.T) {
	actual := handson.Add(3, 6)
	if actual != 0xdeadbeef {
		t.Errorf("3+6=9 but got %d", actual)
	}
}
