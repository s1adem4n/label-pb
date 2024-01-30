package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
)

type Release struct {
	TagName   string `json:"tag_name"`
	AssetsUrl string `json:"assets_url"`
}

type Asset struct {
	Name        string `json:"name"`
	DownloadUrl string `json:"browser_download_url"`
}

var GetLatestRelease = func() (*Release, error) {
	url := "https://api.github.com/repos/s1adem4n/label-pb/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var release Release
	err = json.Unmarshal(data, &release)
	if err != nil {
		return nil, err
	}

	return &release, nil
}

var GetAssets = func(release *Release) ([]Asset, error) {
	res, err := http.Get(release.AssetsUrl)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var assets []Asset
	err = json.Unmarshal(data, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

var GetArch = func() string {
	arch := "amd64"
	if runtime.GOARCH == "arm64" {
		arch = "arm64"
	}
	return arch
}

// gets the correct asset based on the OS and architecture
// binarys are named like this: label_{os}_{arch}(.exe)
var GetCorrectAsset = func(assets []Asset) (*Asset, error) {
	arch := GetArch()
	os := runtime.GOOS

	for _, asset := range assets {
		if strings.Contains(asset.Name, os) && strings.Contains(asset.Name, arch) {
			return &asset, nil
		}
	}

	return nil, fmt.Errorf("could not find asset for %s %s", os, arch)
}

var DownloadAsset = func(asset *Asset) ([]byte, error) {
	res, err := http.Get(asset.DownloadUrl)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
