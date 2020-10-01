// create by d1y<chenhonzhou@gmail.com>
// idea by https://github.com/MusicalCreeper01/Linux-Application-List

package applist

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"
)

const (
	// scan darwin dir
	darwinDir = "/Applications"
)

var (
	// darwin ignore file
	darwinIgnores = []string{
		".DS_Store",
		".localized",
	}
)

// Exists check file/dir exists
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// array filter
//
// https://medium.com/@habibridho/here-is-why-no-one-write-generic-slice-filter-in-go-8b3d1063674e
func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

// scan darwin
func scanDarwin() ([]string, error) {
	if exists(darwinDir) {
		lists, _ := ioutil.ReadDir(darwinDir)
		var results = []string{}
		for _, item := range lists {
			var filename = item.Name()
			var isIgnore = false
			for _, ignoreFile := range darwinIgnores {
				if filename == ignoreFile {
					isIgnore = true
				}
			}
			if !isIgnore {
				results = append(results, filename)
			}
		}
		return results, nil
	}
	return []string{}, errors.New("scan darwin app is error")
}

// GetApps get apps
//
// only supports darwin/linux
//
// darwin output example:
// [ yoxi.app, qq.app ]
func GetApps() ([]string, error) {
	switch runtime.GOOS {
	case "darwin":
		return scanDarwin()
	case "linux":
		// TODO
		return []string{}, nil
	}
	return []string{}, errors.New("unsupported the os")
}
