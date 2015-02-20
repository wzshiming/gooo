package gooo

import (
	"os"
	"path/filepath"
	"strings"
)

func JoinPath(path ...string) string {
	return strings.Join(path, "/")
}

func FullPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	} else {
		absPath, _ := filepath.Abs(path)
		return absPath
	}
}

func GetAllFiles(basepath string) (ret []string) {
	filepath.Walk(FullPath(basepath), func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			if basepath != path {
				ret = append(ret, GetAllFiles(path)...)
			}
			return nil
		}
		ret = append(ret, path)
		return nil
	})
	return
}

func RelPaths(basepath string, paths []string) (ret []string) {
	absPath := FullPath(basepath)
	for _, v := range paths {
		relpath, _ := filepath.Rel(absPath, v)
		ret = append(ret, relpath)
	}
	return
}
