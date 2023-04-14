package cmdutils

import (
	"os"
	"strings"
	"unicode"
)

type PkgFileInfo struct {
	Dir  string
	Path string
}

func LookupDir(dir string) ([]PkgFileInfo, error) {
	paths := []PkgFileInfo{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		entryPath := dir + "/" + entry.Name()
		if entry.IsDir() {
			subDirs, err := LookupDir(entryPath)
			if err != nil {
				return nil, err
			}
			paths = append(paths, subDirs...)
		} else {
			if strings.HasSuffix(entry.Name(), "_grpc.pb.go") {
				paths = append(paths, PkgFileInfo{
					Dir:  dir,
					Path: entryPath,
				})
			}
		}
	}
	return paths, nil
}

func GetModuleName() (string, error) {
	var moduleName string
	b, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			moduleName = line[len("module "):]
		}
	}
	return moduleName, nil
}

func ToSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			res = append(res, '_', unicode.ToLower(r))
		} else {
			res = append(res, unicode.ToLower(r))
		}
	}
	return string(res)
}

func IsStringInArray(arr []string, v string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, s := range arr {
		if s == v {
			return true
		}
	}
	return false
}
