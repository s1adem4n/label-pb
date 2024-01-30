package version

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"github.com/minio/selfupdate"
)

func Update() error {
	fmt.Println("Creating app directory if necessary...")
	err := CreateAppDirectory()
	if err != nil {
		return err
	}

	fmt.Println("Getting latest release...")
	latestRelease, err := GetLatestRelease()
	if err != nil {
		return err
	}

	fmt.Println("Getting assets...")
	assets, err := GetAssets(latestRelease)
	if err != nil {
		return err
	}

	var BinariesAsset Asset

	for _, asset := range assets {
		if strings.Contains(asset.Name, ".zip") {
			BinariesAsset = asset
		}
	}

	if BinariesAsset.Name == "" {
		return fmt.Errorf("no binaries asset found")
	}

	fmt.Println("Downloading asset...")
	res, err := http.Get(BinariesAsset.DownloadUrl)
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	readerAt := bytes.NewReader(bodyBytes)

	var update io.ReadCloser
	read, err := zip.NewReader(readerAt, res.ContentLength)
	if err != nil {
		return err
	}
	for _, file := range read.File {
		platform := runtime.GOOS + "_" + runtime.GOARCH
		if strings.Contains(file.Name, platform) {
			fmt.Println("Found binary for this platform:", file.Name)
			f, err := file.Open()
			if err != nil {
				return err
			}
			update = f
			defer update.Close()

			break
		}
	}

	fmt.Println("Applying update...")
	err = selfupdate.Apply(update, selfupdate.Options{})
	if err != nil {
		return err
	}

	return nil
}
