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

func TestRenamexNp(t *testing.T) {
	type args struct {
		src  string
		dst  string
		flag RENAME_FALG
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				src:  filepath.Join(mountPoint, "test-renamexnp.txt"),
				dst:  filepath.Join(mountPoint, "renamexnp.txt"),
				flag: RENAME_FALG(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fi, err := os.Create(tt.args.src)
			if err != nil {
				t.Fatal(err)
			}
			defer fi.Close()

			err = RenamexNp(tt.args.src, tt.args.dst, tt.args.flag)
			if (err != nil && fi.Name() == tt.args.dst) != tt.wantErr {
				t.Errorf("RenamexNp(%v, %v, %v) error = %v, gotPath %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.flag, err, fi.Name(), tt.wantErr)
			}
		})
	}
}

func TestRenameatxNp(t *testing.T) {
	type args struct {
		src  string
		dst  string
		flag RENAME_FALG
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				src:  filepath.Join(mountPoint, "test-renameatxnp.txt"),
				dst:  filepath.Join(mountPoint, "renameatxnp.txt"),
				flag: RENAME_FALG(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fi, err := os.Create(tt.args.src)
			if err != nil {
				t.Fatal(err)
			}
			defer fi.Close()

			err = RenameatxNp(tt.args.src, tt.args.dst, tt.args.flag)
			if (err != nil && fi.Name() == tt.args.dst) != tt.wantErr {
				t.Errorf("RenameatxNp(%v, %v, %v) error = %v, gotPath %v,, wantErr %v", tt.args.src, tt.args.dst, tt.args.flag, err, fi.Name(), tt.wantErr)
			}
		})
	}
}
