package download

import (
	"testing"
)

func TestDownloadURL(t *testing.T) {
	want := "https://github.com/gohugoio/hugo/releases/download/v1.0/hugo_1.0_Linux-64bit.tar.gz"
	if got := downloadURL("1.0", "64bit"); got != want {
		t.Errorf("Download url is not correct, got: %s, want: %s", got, want)
	}
}

func TestGet(t *testing.T) {
	if binPath , err := Get("0.42","64bit"); err != nil {
		t.Errorf("Failed to download archive: %s", err)
		if binPath == "" {
			t.Errorf("binPath was nil")
		}
	}
}
