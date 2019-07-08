package download

import (
	"testing"
)

func TestDownloadURL(t *testing.T) {
	want := "https://github.com/gohugoio/hugo/releases/download/v1.0/hugo_1.0_Linux-64bit.tar.gz"
	if got := downloadURL("1.0", false); got != want {
		t.Errorf("Download url is not correct, got: %s, want: %s", got, want)
	}
}

func TestDownloadURLExtended(t *testing.T) {
	want := "https://github.com/gohugoio/hugo/releases/download/v0.55.4/hugo_extended_0.55.4_Linux-64bit.tar.gz"
	if got := downloadURL("0.55.4", true); got != want {
		t.Errorf("Download url is not correct, got: %s, want: %s", got, want)
	}
}

func TestGet(t *testing.T) {
	if binPath , err := Get("0.42", false); err != nil {
		t.Errorf("Failed to download archive: %s", err)
		if binPath == "" {
			t.Errorf("binPath was nil")
		}
	}
}
