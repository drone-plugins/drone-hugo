package main

import (
	"github.com/pkg/errors"
	"testing"
)

func TestCommandValidate(t *testing.T) {
	config := Config{Validate: true}
	want := []string{"hugo", "check"}
	got := commandValidate(config)

	if err := argsEqual(want, got.Args); err != nil {
		t.Errorf("%s", err)
	}

	config = Config{Validate: true, Config: "config.toml"}
	want = []string{"hugo", "check", "--config", "config.toml"}
	got = commandValidate(config)

	if err := argsEqual(want, got.Args); err != nil {
		t.Errorf("%s", err)
	}
}

func TestVersionEqual(t *testing.T) {
	want := true
	if got := versionsEqual("1.0", "1.0", false); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}

	want = false
	if got := versionsEqual("1.5", "1.0", false); want != got {
		t.Errorf("want: %t, got: %t", want, got)
	}
}

func argsEqual(want []string, got []string) error {
	for i := range want {
		if want[i] != got[i] {
			return errors.Errorf("Arguments do not match, want: %s, got: %s", want[i], got[i])
		}
	}
	return nil
}
