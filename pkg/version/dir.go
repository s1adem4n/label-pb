package version

import (
	"os"
	"path"
)

var GetAppDirectory = func() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	appDir := path.Join(homedir, "label")
	return appDir, nil
}

var CreateAppDirectory = func() error {
	appDir, err := GetAppDirectory()
	if err != nil {
		return err
	}

	_, err = os.Stat(appDir)
	if err == nil {
		return nil
	}

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return err
	}

	return nil
}
