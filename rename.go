// Copyright 2017 The go-apfs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

/*
#include <fcntl.h> // for AT_FDCWD
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"path/filepath"
	"syscall"
)

type RENAME_FALG uint

var (
	RENAME_SECLUDE RENAME_FALG = 0x00000001
	RENAME_SWAP    RENAME_FALG = 0x00000002
	RENAME_EXCL    RENAME_FALG = 0x00000004
)

func RenamexNp(src, dst string, flags RENAME_FALG) error {
	if err := C.renamex_np(C.CString(src), C.CString(dst), C.unsigned(flags)); err != 0 {
		return fmt.Errorf("error: C.renamex_np: %v", (syscall.Errno(err)))
	}

	return nil
}

func RenameatxNp(src, dst string, flags RENAME_FALG) error {
	var srcFd C.int
	if !filepath.IsAbs(src) {
		srcFd = C.AT_FDCWD
	}
	var dstFd C.int
	if !filepath.IsAbs(dst) {
		dstFd = C.AT_FDCWD
	}
	if err := C.renameatx_np(srcFd, C.CString(src), dstFd, C.CString(dst), C.unsigned(flags)); err != 0 {
		return fmt.Errorf("error: C.renameatx_np: %v", syscall.Errno(err))
	}

	return nil
}
