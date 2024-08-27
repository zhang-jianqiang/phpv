package pkg

import (
	"errors"
	"os"
	"path/filepath"
)

func ListPhpVersion() ([]string, error) {
	phpPath, err := ReadConfig()
	if err != nil {
		return nil, errors.New("php目录未配置，请先配置, 示例: phpv config D:/phpstudy_pro/Extensions/php")
	}

	items := make([]string, 0)
	dirs, err := os.ReadDir(phpPath)
	if err != nil {
		return items, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		name := d.Name()
		_, err := os.Stat(filepath.Join(phpPath, name, "php.exe"))
		if os.IsNotExist(err) {
			continue
		}
		items = append(items, name)
	}

	return items, nil
}
