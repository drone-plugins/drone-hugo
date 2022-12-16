package download

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var (
	url = "https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_%s-%s.tar.gz"
)

// Get will download the specified hugo verion
func Get(version string, extended bool) (string, error) {
	resp, err := http.Get(download(version, extended))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	gz, err := gzip.NewReader(resp.Body)

	if err != nil {
		return "", err
	}

	defer gz.Close()
	targz := tar.NewReader(gz)

	hugoPath, hugoBin, err := tempfile()

	if err != nil {
		return "", err
	}

	defer hugoBin.Close()

	for {
		h, err := targz.Next()

		if err == io.EOF {
			return "", fmt.Errorf("no hugo binary found")
		}

		if err != nil {
			return "", err
		}

		if strings.HasSuffix(h.Name, "hugo") {
			_, _ = io.Copy(hugoBin, targz)

			if err := os.Chmod(hugoPath, 0755); err != nil {
				return "", err
			}

			return hugoPath, nil
		}
	}
}

func download(version string, extended bool) string {
	var (
		binary   string
		osName   string
		archType string
	)

	if extended {
		binary = "hugo_extended"
	} else {
		binary = "hugo"
	}

	switch runtime.GOOS {
	case "linux":
		osName = "Linux"
	case "windows":
		osName = "Windows"
	default:
		osName = "unsupported"
	}

	switch runtime.GOARCH {
	case "amd64":
		archType = "64bit"
	case "arm64":
		archType = "arm64"
	case "arm":
		archType = "arm"
	case "386":
		archType = "32bit"
	default:
		archType = "unsupported"
	}

	return fmt.Sprintf(url, version, binary, version, osName, archType)
}

func tempfile() (string, io.WriteCloser, error) {
	dName, err := os.MkdirTemp("", "hugo")

	if err != nil {
		return "", nil, err
	}

	f, err := os.CreateTemp(dName, "")
	return f.Name(), f, err
}
