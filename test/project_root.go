package test

import (
	"os"
	"path/filepath"

	"github.com/jsightapi/jsight-schema-core/errs"
	"github.com/jsightapi/jsight-schema-core/internal/sync"
)

var projectRootOnce sync.ErrOnceWithValue[string]

func GetProjectRoot() string {
	v, _ := projectRootOnce.Do(func() (string, error) {
		return determineProjectRoot(), nil
	})
	return v
}

func determineProjectRoot() string {
	path, err := os.Getwd()
	if err != nil {
		panic(errs.ErrInTheTest.F(err))
	}
	for {
		if path == "/" || path == "" {
			panic(errs.ErrInTheTest.F("Project root not found"))
		}
		if isExists(filepath.Join(path, "go.mod")) {
			break
		}
		path = filepath.Dir(path)
	}
	return path
}

func isExists(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}
