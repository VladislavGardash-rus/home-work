package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const (
		sourceFileName = "testdata/input.txt"
		targetFileName = "out.txt"
	)

	t.Run("negative offset", func(t *testing.T) {
		err := Copy(sourceFileName, targetFileName, -1, 0)
		require.Equal(t, nil, err)

		resourceFileInfo, _ := os.Stat(sourceFileName)
		targetFileInfo, _ := os.Stat(targetFileName)
		require.Equal(t, resourceFileInfo.Size(), targetFileInfo.Size())

		err = os.Remove(targetFileName)
		require.Equal(t, nil, err)
	})

	t.Run("negative limit", func(t *testing.T) {
		err := Copy(sourceFileName, targetFileName, 0, -1)
		require.Equal(t, nil, err)

		resourceFileInfo, _ := os.Stat(sourceFileName)
		targetFileInfo, _ := os.Stat(targetFileName)
		require.Equal(t, resourceFileInfo.Size(), targetFileInfo.Size())

		err = os.Remove(targetFileName)
		require.Equal(t, nil, err)
	})

	t.Run("negative limit offset", func(t *testing.T) {
		err := Copy(sourceFileName, targetFileName, -1, -1)
		require.Equal(t, nil, err)

		resourceFileInfo, _ := os.Stat(sourceFileName)
		targetFileInfo, _ := os.Stat(targetFileName)
		require.Equal(t, resourceFileInfo.Size(), targetFileInfo.Size())

		err = os.Remove(targetFileName)
		require.Equal(t, nil, err)
	})

	t.Run("without source file", func(t *testing.T) {
		err := Copy("", targetFileName, 0, 0)
		require.Error(t, err)
	})

	t.Run("limit 100", func(t *testing.T) {
		err := Copy(sourceFileName, targetFileName, 0, 100)
		require.Equal(t, nil, err)

		targetFileInfo, _ := os.Stat(targetFileName)
		require.Equal(t, int64(100), targetFileInfo.Size())

		err = os.Remove(targetFileName)
		require.Equal(t, nil, err)
	})

	t.Run("offset 50", func(t *testing.T) {
		err := Copy(sourceFileName, targetFileName, 50, 4)
		require.Equal(t, nil, err)

		targetFileInfo, err := os.Stat(targetFileName)
		require.Equal(t, nil, err)
		require.Equal(t, int64(4), targetFileInfo.Size())

		file, err := os.Open(targetFileName)
		require.Equal(t, nil, err)

		data, err := io.ReadAll(file)
		require.Equal(t, nil, err)

		require.Equal(t, "Play", string(data))

		file.Close()
		err = os.Remove(targetFileName)
		require.Equal(t, nil, err)
	})
}
