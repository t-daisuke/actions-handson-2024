package handson_test

import (
    "testing"
    "handson"
)

func TestAdd(t *testing.T) {
    actual := handson.Add(3, 6)
    if actual != 9  {
        t.Errorf("3+6=9 but got %d", actual)
    }
}