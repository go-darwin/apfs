// Copyright 2017 The go-apfs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

/*
#include <stdlib.h>
#include <copyfile.h>
*/
import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
)

type COPYFILE_FLAG int

var (
	COPYFILE_ACL   COPYFILE_FLAG = 1 << 0
	COPYFILE_STAT  COPYFILE_FLAG = 1 << 1
	COPYFILE_XATTR COPYFILE_FLAG = 1 << 2
	COPYFILE_DATA  COPYFILE_FLAG = 1 << 3

	COPYFILE_SECURITY COPYFILE_FLAG = COPYFILE_STAT | COPYFILE_ACL
	COPYFILE_METADATA COPYFILE_FLAG = COPYFILE_SECURITY | COPYFILE_XATTR
	COPYFILE_ALL      COPYFILE_FLAG = COPYFILE_METADATA | COPYFILE_DATA

	COPYFILE_RECURSIVE    COPYFILE_FLAG = 1 << 15 // Descend into hierarchies
	COPYFILE_CHECK        COPYFILE_FLAG = 1 << 16 // return flags for xattr or acls if set
	COPYFILE_EXCL         COPYFILE_FLAG = 1 << 17 // fail if destination exists
	COPYFILE_NOFOLLOW_SRC COPYFILE_FLAG = 1 << 18 // don't follow if source is a symlink
	COPYFILE_NOFOLLOW_DST COPYFILE_FLAG = 1 << 19 // don't follow if dst is a symlink
	COPYFILE_MOVE         COPYFILE_FLAG = 1 << 20 // unlink src after copy
	COPYFILE_UNLINK       COPYFILE_FLAG = 1 << 21 // unlink dst before copy
	COPYFILE_NOFOLLOW     COPYFILE_FLAG = COPYFILE_NOFOLLOW_SRC | COPYFILE_NOFOLLOW_DST

	COPYFILE_PACK   COPYFILE_FLAG = 1 << 22
	COPYFILE_UNPACK COPYFILE_FLAG = 1 << 23

	COPYFILE_CLONE       COPYFILE_FLAG = 1 << 24
	COPYFILE_CLONE_FORCE COPYFILE_FLAG = 1 << 25

	COPYFILE_RUN_IN_PLACE COPYFILE_FLAG = 1 << 26

	COPYFILE_VERBOSE COPYFILE_FLAG = 1 << 30
)

type COPYFILE_STATE_FLAG uint32

var (
	COPYFILE_STATE_SRC_FD       COPYFILE_STATE_FLAG = C.COPYFILE_STATE_SRC_FD
	COPYFILE_STATE_DST_FD       COPYFILE_STATE_FLAG = C.COPYFILE_STATE_DST_FD
	COPYFILE_STATE_SRC_FILENAME COPYFILE_STATE_FLAG = C.COPYFILE_STATE_SRC_FILENAME
	COPYFILE_STATE_DST_FILENAME COPYFILE_STATE_FLAG = C.COPYFILE_STATE_DST_FILENAME
	COPYFILE_STATE_STATUS_CB    COPYFILE_STATE_FLAG = C.COPYFILE_STATE_STATUS_CB
	COPYFILE_STATE_STATUS_CTX   COPYFILE_STATE_FLAG = C.COPYFILE_STATE_STATUS_CTX
	COPYFILE_STATE_QUARANTINE   COPYFILE_STATE_FLAG = C.COPYFILE_STATE_QUARANTINE
	COPYFILE_STATE_COPIED       COPYFILE_STATE_FLAG = C.COPYFILE_STATE_COPIED
	COPYFILE_STATE_XATTRNAME    COPYFILE_STATE_FLAG = C.COPYFILE_STATE_XATTRNAME
	COPYFILE_STATE_WAS_CLONED   COPYFILE_STATE_FLAG = C.COPYFILE_STATE_WAS_CLONED
)

type COPYFILE_RECURSE_CALLBACK uint32

var (
	COPYFILE_RECURSE_ERROR       COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_ERROR       // 0
	COPYFILE_RECURSE_FILE        COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_FILE        // 1
	COPYFILE_RECURSE_DIR         COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR         // 2
	COPYFILE_RECURSE_DIR_CLEANUP COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR_CLEANUP // 3
	COPYFILE_COPY_DATA           COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_DATA           // 4
	COPYFILE_COPY_XATTR          COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_XATTR          // 5
)

var (
	COPYFILE_START    COPYFILE_RECURSE_CALLBACK = C.COPYFILE_START    // 1
	COPYFILE_FINISH   COPYFILE_RECURSE_CALLBACK = C.COPYFILE_FINISH   // 2
	COPYFILE_ERR      COPYFILE_RECURSE_CALLBACK = C.COPYFILE_ERR      // 3
	COPYFILE_PROGRESS COPYFILE_RECURSE_CALLBACK = C.COPYFILE_PROGRESS // 4
)

var (
	COPYFILE_CONTINUE COPYFILE_RECURSE_CALLBACK = C.COPYFILE_CONTINUE // 0
	COPYFILE_SKIP     COPYFILE_RECURSE_CALLBACK = C.COPYFILE_SKIP     // 1
	COPYFILE_QUIT     COPYFILE_RECURSE_CALLBACK = C.COPYFILE_QUIT     // 2
)

type COPYFILE_STATE C.copyfile_state_t

func CopyFile(src, dst string, state COPYFILE_STATE, flag COPYFILE_FLAG) (bool, error) {
	if err := C.copyfile(C.CString(src), C.CString(dst), state, C.copyfile_flags_t(flag)); err != 0 {
		return false, fmt.Errorf("couldn't copy from %s to %s: %v", src, dst, syscall.Errno(err))
	}

	var isCloned int
	if err := CopyFileStateGet(state, COPYFILE_STATE_WAS_CLONED, &isCloned); err != nil {
		return false, fmt.Errorf("couldn't get copyfile_state_get: %v", err)
	}

	return isCloned != 0, nil
}

func FcopyFile(src, dst uintptr, state COPYFILE_STATE, flag COPYFILE_FLAG) error {
	if err := C.fcopyfile(C.int(src), C.int(dst), state, C.copyfile_flags_t(flag)); err != 0 {
		return fmt.Errorf("couldn't fcopy from %d to %d: %v", src, dst, syscall.Errno(err))
	}

	return nil
}

func CopyFileStateAlloc() COPYFILE_STATE {
	return COPYFILE_STATE(C.copyfile_state_alloc())
}

func CopyFileStateFree(state COPYFILE_STATE) error {
	if err := C.copyfile_state_free(C.copyfile_state_t(state)); err != 0 {
		return syscall.Errno(err)
	}

	return nil
}

func CopyFileStateGet(state COPYFILE_STATE, flag COPYFILE_STATE_FLAG, result *int) error {
	if err := C.copyfile_state_get(state, C.uint32_t(flag), unsafe.Pointer(result)); err != 0 {
		return syscall.Errno(err)
	}

	return nil
}
