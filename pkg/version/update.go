package version

import (
	"fmt"
	"os"
	"path"
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

	fmt.Println("Getting correct asset for current os and arch...")
	asset, err := GetCorrectAsset(assets)
	if err != nil {
		return err
	}

	fmt.Println("Downloading asset...")
	data, err := DownloadAsset(asset)
	if err != nil {
		return err
	}

	appDir, err := GetAppDirectory()
	if err != nil {
		return err
	}
	fmt.Println("Writing asset to disk at \"" + appDir + "\"...")

	installPath := path.Join(appDir, "label")
	err = os.WriteFile(installPath, data, 0755)
	if err != nil {
		return err
	}

	return nil
}
