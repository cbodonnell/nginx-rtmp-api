package main

import (
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

// RemoveContents -  Removes contents from a directory
func RemoveContents(dir string) error {
	if config.Filesystem {
		d, err := os.Open(dir)
		if err != nil {
			return err
		}
		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			return err
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CopyDirectory - Copies a directory
func CopyDirectory(fromDir string, toDir string) error {
	if config.Filesystem {
		return copy.Copy(fromDir, toDir)
	}
	return nil
}
