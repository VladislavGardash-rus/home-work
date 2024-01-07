//nolint:all
package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environments := make(Environment)
	for _, file := range files {
		if strings.Contains(file.Name(), "=") || !file.Mode().IsRegular() {
			continue
		}

		if file.Size() == 0 {
			environments[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		envValue, err := readEnvironmentValue(dir, file.Name())
		if err != nil {
			return nil, err
		}

		environments[file.Name()] = EnvValue{Value: envValue}
	}

	return environments, nil
}

func readEnvironmentValue(dir, fileName string) (string, error) {
	file, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", nil
	}

	str := scanner.Text()
	str = strings.TrimRightFunc(str, unicode.IsSpace)
	str = strings.ReplaceAll(str, "\x00", "\n")

	return str, nil
}
