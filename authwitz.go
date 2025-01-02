package authwitz

import (
	"errors"
	"io"
	"os"

	"github.com/yeka/zip"
)

type Checker struct {
	file      *os.File
	zipReader *zip.Reader
}

func NewChecker() *Checker {
	return &Checker{}
}

func (c *Checker) Open(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	c.file = file
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	zipReader, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return err
	}
	c.zipReader = zipReader
	if len(zipReader.File) == 0 {
		return errors.New("no files found in archive")
	}
	return nil
}

func (c *Checker) Try(password string) (bool, error) {
	for _, file := range c.zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		file.SetPassword(password)
		handle, err := file.Open()
		if err != nil {
			return false, err
		}
		buf := make([]byte, 100)
		_, err = handle.Read(buf)
		if err == nil || err == io.EOF {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("no files found in the archive")
}
func (c *Checker) Close() {
	c.file.Close()
}
