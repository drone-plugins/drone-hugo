package main

import (
	"testing"
)

func TestVersionEqual(t *testing.T) {
	want := true
	if got := versionsEqual("1.0", "1.0"); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}

	want = false
	if got := versionsEqual("1.5", "1.0"); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}
}
