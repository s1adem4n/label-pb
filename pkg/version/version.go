package version

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"
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
