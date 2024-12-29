package authwitz

import (
	"archive/zip"
	"errors"
	"os"
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
}

func (c *Checker) Try() (bool, error) {

}
func (c *Checker) Close() {
	c.file.Close()
}
