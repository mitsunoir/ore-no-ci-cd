package common

import "testing"

func TestAdd(t *testing.T) {
    want := 5
    if got := Add(2, 3); got != want {
        t.Errorf("Add(2, 3) != %q, want %q", got, want)
    }
}
