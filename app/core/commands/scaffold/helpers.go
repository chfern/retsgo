package scaffold

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReplaceStrInFile(path string, fi os.FileInfo, old string, new string) error {
	if !!fi.IsDir() {
		return nil
	}
	matched, err := filepath.Match("*.go", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), old, new, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
