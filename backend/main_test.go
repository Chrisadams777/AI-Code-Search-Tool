package main

import "testing"

func TestSample(t *testing.T) {
    if 1+1 != 2 {
        t.Error("Math is broken!")
    }
}
