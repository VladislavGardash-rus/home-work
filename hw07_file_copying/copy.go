package main

import (
	"errors"                    //nolint:all
	"github.com/cheggaaa/pb/v3" //nolint:all
	"io"                        //nolint:all
	"os"                        //nolint:all
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	resourceFileInfo, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !resourceFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > resourceFileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 || offset+limit > resourceFileInfo.Size() {
		limit = resourceFileInfo.Size() - offset
	}

	resourceFile, err := openResourceFile(fromPath, offset)
	if err != nil {
		return err
	}
	defer resourceFile.Close()

	targetFile, err := openTargetFile(toPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	progressBar := pb.Default.Start64(limit)
	barReader := progressBar.NewProxyReader(resourceFile)
	defer progressBar.Finish()

	_, err = io.CopyN(targetFile, barReader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}

func openResourceFile(name string, offset int64) (*os.File, error) {
	resourceFileInfo, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	if offset > 0 {
		_, err = resourceFileInfo.Seek(offset, io.SeekStart)
		if err != nil {
			resourceFileInfo.Close()
			return nil, err
		}
	}

	return resourceFileInfo, nil
}

func openTargetFile(name string) (*os.File, error) {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		err = os.Remove(name)
		if err != nil {
			return nil, err
		}
	}

	targetFile, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return targetFile, nil
}
