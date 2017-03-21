// Copyright 2017 The go-apfs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

import (
	"fmt"
	"log"
	"os"
)

func ExampleCopyFile() {
	src, dst := os.Args[1], os.Args[2]

	state := CopyFileStateAlloc()
	defer func() {
		if err := CopyFileStateFree(state); err != nil {
			log.Fatal(err)
		}
	}()

	cloned, err := CopyFile(src, dst, state, COPYFILE_CLONE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("isCloned: %v", cloned)

	// Outpt: true // or false
}

func ExampleFcopyFile()          {}
func ExampleCopyFileStateAlloc() {}
func ExampleCopyFileStateFree()  {}
func ExampleCopyFileStateGet()   {}

func ExampleCloneFile() {
	src, dst := os.Args[1], os.Args[2]
	err := CloneFile(src, dst, CLONEFILE_FLAG(0))
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleCloneFileAt()  {}
func ExampleFcloneFileAt() {}
