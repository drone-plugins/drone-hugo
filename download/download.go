package download

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

var (
	_downloadURL = "https://github.com/gohugoio/hugo/releases/download/v%s/%s_%s_Linux-%s.tar.gz"
)

func downloadURL(version string, extended bool) string {
	var archType string
	var binary string
	
	if extended {
		binary = "hugo_extended"
	} else {
		binary = "hugo"
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
	return fmt.Sprintf(_downloadURL, version, binary, version, archType)
}

func getTempFile() (string, io.WriteCloser, error) {
	d, err := ioutil.TempDir("", "")
	if err != nil {
		return "", nil, errors.Wrap(err, "")
	}
	f, err := ioutil.TempFile(d, "")
	return f.Name(), f, err
}

// Get will download the specified hugo verion
func Get(version string, extended bool) (string, error) {
	resp, err := http.Get(downloadURL(version, extended))
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	defer resp.Body.Close()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	defer gz.Close()
	targz := tar.NewReader(gz)

	hugoPath, hugoBin, err := getTempFile()
	if err != nil {
		log.Printf("ERROR: %s", err)
		return "", errors.Wrap(err, "")
	}
	defer hugoBin.Close()

	for {
		h, err := targz.Next()
		if err == io.EOF {
			return "", errors.New("no hugo binary found")
		}
		if err != nil {
			return "", errors.Wrap(err, "")
		}
		if strings.HasSuffix(h.Name, "hugo") {
			io.Copy(hugoBin, targz)

			if err := os.Chmod(hugoPath, 0755); err != nil {
				log.Fatal(err)
			}

			return hugoPath, nil
		}
	}
}
