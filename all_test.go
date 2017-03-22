// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	sparseImage    = filepath.Join("testdata", "apfs.sparseimage")
	mountPoint     = filepath.Join("testdata", "apfs")
	testFileGolden = filepath.Join("testdata", "testfile.txt")
	testFile       = filepath.Join(mountPoint, "testfile.txt")
)

func setupTest() {
	log.SetFlags(log.Lshortfile)

	if err := createAPFSSparseImage(); err != nil {
		log.Fatal(err)
	}

	if err := mountAPFSSparseImage(); err != nil {
		log.Fatal(err)
	}

	if err := copyFile(testFile, testFileGolden); err != nil {
		log.Fatal(err)
	}
}

func cleanupTest() {
	if err := unmountAPFSSparseImage(); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(sparseImage); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setupTest()
	err := m.Run()
	cleanupTest()
	os.Exit(err)
}

var hdiutl = filepath.Join("/usr", "bin", "hdiutil")

func createAPFSSparseImage() error {
	cmd := exec.Command(hdiutl, []string{"create", "-fs", "-quiet", "APFS", "-size", "1GB", sparseImage}...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func mountAPFSSparseImage() error {
	cmd := exec.Command(hdiutl, []string{"mount", "-mountpoint", mountPoint, sparseImage}...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func unmountAPFSSparseImage() error {
	cmd := exec.Command(hdiutl, []string{"unmount", mountPoint}...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// copyFile copies a file from src to dst.
func copyFile(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		e := s.Close()
		if err == nil {
			err = e
		}
	}()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		e := d.Close()
		if err == nil {
			err = e
		}
	}()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}

	i, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, i.Mode())
}
