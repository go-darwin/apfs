// Copyright 2017 Koichi Shiraishi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	apfs "github.com/zchee/go-apfs"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <src> <dst>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	state := apfs.CopyFileStateAlloc()
	defer func() {
		if err := apfs.CopyFileStateFree(state); err != nil {
			log.Fatal(err)
		}
	}()

	src, dst := flag.Arg(0), flag.Arg(1)
	cloned, err := apfs.CopyFile(src, dst, state, apfs.COPYFILE_CLONE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("isCloned: %v", cloned)
}
