// Copyright 2017 The go-apfs Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package apfs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	type args struct {
		src   string
		dst   string
		state COPYFILE_STATE
		flag  COPYFILE_FLAG
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "both file in apfs",
			args: args{
				src:   testFile,
				dst:   filepath.Join(mountPoint, "copyfile.txt"),
				state: CopyFileStateAlloc(),
				flag:  COPYFILE_CLONE,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "src file not in apfs",
			args: args{
				src:   testFileGolden,
				dst:   filepath.Join(mountPoint, "copyfile2.txt"),
				state: CopyFileStateAlloc(),
				flag:  COPYFILE_CLONE,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer CopyFileStateFree(tt.args.state)

			got, err := CopyFile(tt.args.src, tt.args.dst, tt.args.state, tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFile(%v, %v, %v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.state, tt.args.flag, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CopyFile(%v, %v, %v, %v) = %v, want %v", tt.args.src, tt.args.dst, tt.args.state, tt.args.flag, got, tt.want)
			}
		})
	}
}

func TestFcopyFile(t *testing.T) {
	type args struct {
		src   uintptr
		dst   uintptr
		state COPYFILE_STATE
		flag  COPYFILE_FLAG
	}
	tests := []struct {
		name    string
		args    args
		srcFile string
		dstFile string
		wantErr bool
	}{
		{
			name: "both file in apfs",
			args: args{
				state: CopyFileStateAlloc(),
				flag:  COPYFILE_DATA,
			},
			srcFile: testFile,
			dstFile: filepath.Join(mountPoint, "fcopyfile.txt"),
			wantErr: false,
		},
		{
			name: "src file not in apfs",
			args: args{
				state: CopyFileStateAlloc(),
				flag:  COPYFILE_DATA,
			},
			srcFile: testFileGolden,
			dstFile: filepath.Join(mountPoint, "fcopyfile2.txt"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer CopyFileStateFree(tt.args.state)

			srcFd, err := os.Open(tt.srcFile)
			if err != nil {
				t.Fatal(err)
			}
			defer srcFd.Close()
			tt.args.src = srcFd.Fd()

			dstFd, err := os.Create(tt.dstFile)
			if err != nil {
				t.Fatal(err)
			}
			defer dstFd.Close()
			tt.args.dst = dstFd.Fd()

			if err := FcopyFile(tt.args.src, tt.args.dst, tt.args.state, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("FcopyFile(%v, %v, %v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.state, tt.args.flag, err, tt.wantErr)
			}
		})
	}
}

// TODO(zchee): implements test
func TestCopyFileStateAlloc(t *testing.T) {}

func TestCopyFileStateFree(t *testing.T) {
	type args struct {
		state COPYFILE_STATE
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{state: CopyFileStateAlloc()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFileStateFree(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("CopyFileStateFree(%v) error = %v, wantErr %v", tt.args.state, err, tt.wantErr)
			}
		})
	}
}

// TODO(zchee): implements test
func TestCopyFileStateGet(t *testing.T) {}
