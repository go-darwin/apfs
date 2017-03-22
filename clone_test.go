// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCloneFile(t *testing.T) {
	type args struct {
		src  string
		dst  string
		flag CLONEFILE_FLAG
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				src:  testFile,
				dst:  filepath.Join(mountPoint, "clonefile.txt"),
				flag: CLONEFILE_FLAG(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CloneFile(tt.args.src, tt.args.dst, tt.args.flag)
			_, statErr := os.Stat(tt.args.dst)
			if (err != nil && statErr != nil) != tt.wantErr {
				t.Errorf("CloneFile(%v, %v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.flag, err, tt.wantErr)
			}
		})
	}
}

func TestCloneFileAt(t *testing.T) {
	type args struct {
		src  string
		dst  string
		flag CLONEFILE_FLAG
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				src:  testFile,
				dst:  filepath.Join(mountPoint, "clonefileat.txt"),
				flag: CLONEFILE_FLAG(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloneFileAt(tt.args.src, tt.args.dst, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("CloneFileAt(%v, %v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.flag, err, tt.wantErr)
			}
		})
	}
}

func TestFcloneFileAt(t *testing.T) {
	type args struct {
		srcFd uintptr
		dst   string
		flag  CLONEFILE_FLAG
	}
	tests := []struct {
		name    string
		file    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				dst:  filepath.Join(mountPoint, "fclonefileat.txt"),
				flag: CLONEFILE_FLAG(0),
			},
			file:    testFile,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		fi, err := os.Open(tt.file)
		if err != nil {
			t.Fatal(err)
		}
		defer fi.Close()

		tt.args.srcFd = fi.Fd()
		t.Run(tt.name, func(t *testing.T) {
			if err := FcloneFileAt(tt.args.srcFd, tt.args.dst, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("FcloneFileAt(%v, %v, %v) error = %v, wantErr %v", tt.args.srcFd, tt.args.dst, tt.args.flag, err, tt.wantErr)
			}
		})
	}
}
