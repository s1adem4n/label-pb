package version

import (
	"os"
	"path"
)

func Install() error {
	err := CreateAppDirectory()
	if err != nil {
		return err
	}

	appDir, err := GetAppDirectory()
	if err != nil {
		return err
	}
	binaryPath, err := os.Executable()
	if err != nil {
		return err
	}

	installPath := path.Join(appDir, "label")
	err = os.Rename(binaryPath, installPath)
	if err != nil {
		return err
	}

	return nil
}
