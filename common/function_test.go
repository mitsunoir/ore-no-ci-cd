package common

import "testing"

func TestAdd(t *testing.T) {
    want := 5
    if got := Add(2, 3); got != want {
        t.Errorf("Add(2, 3) != %d, want %d", got, want)
    }
}

func TestSub(t *testing.T) {
    want := -1
    if got := Sub(2, 3); got != want {
        t.Errorf("Sub(2, 3) != %d, want %d", got, want)
    }
}
