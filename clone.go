// Copyright 2017 The go-apfs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

/*
#include <fcntl.h> // for AT_FDCWD
#include <stdint.h> // for uint32_t
#include <sys/attr.h>
#include <sys/clonefile.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"path/filepath"
	"syscall"
)

type CLONEFILE_FALG C.uint32_t

var (
	CLONE_NOFOLLOW CLONEFILE_FALG = 0x0001
)

func CloneFile(src, dst string, flag CLONEFILE_FALG) error {
	if err := C.clonefile(C.CString(src), C.CString(dst), C.uint32_t(flag)); err != 0 {
		return errors.New(fmt.Sprintf("error: C.clonefile: %v", syscall.Errno(err)))
	}

	return nil
}

func CloneFileAt(src, dst string, flag CLONEFILE_FALG) error {
	var srcDirFd C.int
	if !filepath.IsAbs(src) {
		srcDirFd = C.AT_FDCWD
	}
	var dstDirFd C.int
	if !filepath.IsAbs(dst) {
		dstDirFd = C.AT_FDCWD
	}
	if err := C.clonefileat(srcDirFd, C.CString(src), dstDirFd, C.CString(dst), C.uint32_t(flag)); err != 0 {
		return errors.New(fmt.Sprintf("error: C.clonefileat: %v", syscall.Errno(err)))
	}

	return nil
}

func FcloneFileAt(srcFd uintptr, dst string, flag CLONEFILE_FALG) error {
	var dstDirFd C.int
	if !filepath.IsAbs(dst) {
		dstDirFd = C.AT_FDCWD
	}
	if err := C.fclonefileat(C.int(srcFd), dstDirFd, C.CString(dst), C.uint32_t(flag)); err != 0 {
		return errors.New(fmt.Sprintf("error: C.fclonefileat: %v", syscall.Errno(err)))
	}

	return nil
}
