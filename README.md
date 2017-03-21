# go-apfs

Package apfs implements an Apple File System(apfs) bindings for Go.

| **CI (darwin)**                             | **codecov.io**                          | **godoc.org**                      | **Release**                          |
|:-------------------------------------------:|:---------------------------------------:|:----------------------------------:|:------------------------------------:|
| [![travis-ci.com][travis-badge]][travis]    | [![codecov.io][codecov-badge]][codecov] | [![godoc.org][godoc-badge]][godoc] | [![Release][release-badge]][release] |

[![Analytics][ga-badge]][ga]


## Install

Installing `apfs` package:

```sh
go get -u -v github.com/zchee/go-apfs
```


## Usage

```
```


## API

`import "github.com/zchee/go-apfs"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package apfs implements an Apple File System(apfs) bindings for Go.

## <a name="pkg-index">Index</a>
* [func CloneFile(src, dst string, flag CLONEFILE_FLAG) error](#CloneFile)
* [func CloneFileAt(src, dst string, flag CLONEFILE_FLAG) error](#CloneFileAt)
* [func CopyFile(src, dst string, state state, flags COPYFILE_FLAG) (bool, error)](#CopyFile)
* [func CopyFileStateAlloc() state](#CopyFileStateAlloc)
* [func CopyFileStateFree(state state) error](#CopyFileStateFree)
* [func CopyFileStateGet(state state, flag COPYFILE_STATE, result *int) error](#CopyFileStateGet)
* [func FcloneFileAt(srcFd uintptr, dst string, flag CLONEFILE_FLAG) error](#FcloneFileAt)
* [func FcopyFile(src, dst uintptr, state state, flags COPYFILE_FLAG) error](#FcopyFile)
* [func RenameatxNp(src, dst string, flags RENAME_FALG) error](#RenameatxNp)
* [func RenamexNp(src, dst string, flags RENAME_FALG) error](#RenamexNp)
* [type CLONEFILE_FLAG](#CLONEFILE_FLAG)
* [type COPYFILE_FLAG](#COPYFILE_FLAG)
* [type COPYFILE_RECURSE_CALLBACK](#COPYFILE_RECURSE_CALLBACK)
* [type COPYFILE_STATE](#COPYFILE_STATE)
* [type RENAME_FALG](#RENAME_FALG)

#### <a name="pkg-examples">Examples</a>
* [CloneFileAt](#example_CloneFileAt)
* [CopyFile](#example_CopyFile)
* [CopyFileStateAlloc](#example_CopyFileStateAlloc)
* [CopyFileStateFree](#example_CopyFileStateFree)
* [CopyFileStateGet](#example_CopyFileStateGet)
* [FcloneFileAt](#example_FcloneFileAt)
* [FcopyFile](#example_FcopyFile)

#### <a name="pkg-files">Package files</a>
[apfs.go](/src/github.com/zchee/go-apfs/apfs.go) [clone.go](/src/github.com/zchee/go-apfs/clone.go) [copy.go](/src/github.com/zchee/go-apfs/copy.go) [rename.go](/src/github.com/zchee/go-apfs/rename.go) 


## <a name="CloneFile">func</a> [CloneFile](/src/target/clone.go?s=472:530#L19)

``` go
func CloneFile(src, dst string, flag CLONEFILE_FLAG) error
```

## <a name="CloneFileAt">func</a> [CloneFileAt](/src/target/clone.go?s=716:776#L27)

``` go
func CloneFileAt(src, dst string, flag CLONEFILE_FLAG) error
```

## <a name="CopyFile">func</a> [CopyFile](/src/target/copy.go?s=3528:3606#L83)

``` go
func CopyFile(src, dst string, state state, flags COPYFILE_FLAG) (bool, error)
```

## <a name="CopyFileStateAlloc">func</a> [CopyFileStateAlloc](/src/target/copy.go?s=4288:4319#L104)

``` go
func CopyFileStateAlloc() state
```

## <a name="CopyFileStateFree">func</a> [CopyFileStateFree](/src/target/copy.go?s=4365:4406#L108)

``` go
func CopyFileStateFree(state state) error
```

## <a name="CopyFileStateGet">func</a> [CopyFileStateGet](/src/target/copy.go?s=4528:4602#L116)

``` go
func CopyFileStateGet(state state, flag COPYFILE_STATE, result *int) error
```

## <a name="FcloneFileAt">func</a> [FcloneFileAt](/src/target/clone.go?s=1134:1205#L43)

``` go
func FcloneFileAt(srcFd uintptr, dst string, flag CLONEFILE_FLAG) error
```

## <a name="FcopyFile">func</a> [FcopyFile](/src/target/copy.go?s=4015:4087#L96)

``` go
func FcopyFile(src, dst uintptr, state state, flags COPYFILE_FLAG) error
```

## <a name="RenameatxNp">func</a> [RenameatxNp](/src/target/rename.go?s=704:762#L26)

``` go
func RenameatxNp(src, dst string, flags RENAME_FALG) error
```

## <a name="RenamexNp">func</a> [RenamexNp](/src/target/rename.go?s=470:526#L18)

``` go
func RenamexNp(src, dst string, flags RENAME_FALG) error
```

## <a name="CLONEFILE_FLAG">type</a> [CLONEFILE_FLAG](/src/target/clone.go?s=391:421#L13)

``` go
type CLONEFILE_FLAG C.uint32_t
```

``` go
var (
    CLONE_NOFOLLOW CLONEFILE_FLAG = 0x0001
)
```

## <a name="COPYFILE_FLAG">type</a> [COPYFILE_FLAG](/src/target/copy.go?s=296:318#L10)

``` go
type COPYFILE_FLAG int
```

``` go
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
```

## <a name="COPYFILE_RECURSE_CALLBACK">type</a> [COPYFILE_RECURSE_CALLBACK](/src/target/copy.go?s=2363:2400#L57)

``` go
type COPYFILE_RECURSE_CALLBACK uint32
```

``` go
var (
    COPYFILE_RECURSE_ERROR       COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_ERROR       // 0
    COPYFILE_RECURSE_FILE        COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_FILE        // 1
    COPYFILE_RECURSE_DIR         COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR         // 2
    COPYFILE_RECURSE_DIR_CLEANUP COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR_CLEANUP // 3
    COPYFILE_COPY_DATA           COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_DATA           // 4
    COPYFILE_COPY_XATTR          COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_XATTR          // 5
)
```

``` go
var (
    COPYFILE_START    COPYFILE_RECURSE_CALLBACK = C.COPYFILE_START    // 1
    COPYFILE_FINISH   COPYFILE_RECURSE_CALLBACK = C.COPYFILE_FINISH   // 2
    COPYFILE_ERR      COPYFILE_RECURSE_CALLBACK = C.COPYFILE_ERR      // 3
    COPYFILE_PROGRESS COPYFILE_RECURSE_CALLBACK = C.COPYFILE_PROGRESS // 4
)
```

``` go
var (
    COPYFILE_CONTINUE COPYFILE_RECURSE_CALLBACK = C.COPYFILE_CONTINUE // 0
    COPYFILE_SKIP     COPYFILE_RECURSE_CALLBACK = C.COPYFILE_SKIP     // 1
    COPYFILE_QUIT     COPYFILE_RECURSE_CALLBACK = C.COPYFILE_QUIT     // 2
)
```

## <a name="COPYFILE_STATE">type</a> [COPYFILE_STATE](/src/target/copy.go?s=1596:1622#L42)

``` go
type COPYFILE_STATE uint32
```

``` go
var (
    COPYFILE_STATE_SRC_FD       COPYFILE_STATE = C.COPYFILE_STATE_SRC_FD
    COPYFILE_STATE_DST_FD       COPYFILE_STATE = C.COPYFILE_STATE_DST_FD
    COPYFILE_STATE_SRC_FILENAME COPYFILE_STATE = C.COPYFILE_STATE_SRC_FILENAME
    COPYFILE_STATE_DST_FILENAME COPYFILE_STATE = C.COPYFILE_STATE_DST_FILENAME
    COPYFILE_STATE_STATUS_CB    COPYFILE_STATE = C.COPYFILE_STATE_STATUS_CB
    COPYFILE_STATE_STATUS_CTX   COPYFILE_STATE = C.COPYFILE_STATE_STATUS_CTX
    COPYFILE_STATE_QUARANTINE   COPYFILE_STATE = C.COPYFILE_STATE_QUARANTINE
    COPYFILE_STATE_COPIED       COPYFILE_STATE = C.COPYFILE_STATE_COPIED
    COPYFILE_STATE_XATTRNAME    COPYFILE_STATE = C.COPYFILE_STATE_XATTRNAME
    COPYFILE_STATE_WAS_CLONED   COPYFILE_STATE = C.COPYFILE_STATE_WAS_CLONED
)
```

## <a name="RENAME_FALG">type</a> [RENAME_FALG](/src/target/rename.go?s=315:336#L10)

``` go
type RENAME_FALG uint
```

``` go
var (
    RENAME_SECLUDE RENAME_FALG = 0x00000001
    RENAME_SWAP    RENAME_FALG = 0x00000002
    RENAME_EXCL    RENAME_FALG = 0x00000004
)
```


## Contribute

Not yet.  
~~See [CONTRIBUTING.md](CONTRIBUTING.md)~~.


## License

go-apfs is released under the [BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause).  


[travis]: https://travis-ci.org/zchee/go-apfs
[godoc]: https://godoc.org/github.com/zchee/go-apfs
[codecov]: https://codecov.io/gh/zchee/go-apfs
[release]: https://github.com/zchee/go-apfs/releases
[ga]: https://github.com/zchee/go-apfs

[travis-badge]: https://img.shields.io/travis/zchee/go-apfs.svg?style=flat-square&label=%20Travis%20CI&logo=data%3Aimage%2Fsvg%2Bxml%3Bcharset%3Dutf-8%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI2MCIgaGVpZ2h0PSI2MCIgdmlld0JveD0iNSA0IDI0IDI0Ij48cGF0aCBmaWxsPSIjREREIiBkPSJNMTEuMzkyKzkuMzc0aDQuMDk2djEzLjEyaC0xLjUzNnYyLjI0aDYuMDgwdi0yLjQ5NmgtMS45MnYtMTMuMDU2aDQuMzUydjEuOTJoMS45ODR2LTMuOTA0aC0xNS4yOTZ2My45MDRoMi4yNHpNMjkuMjYzKzIuNzE4aC0yNC44NDhjLTAuNDMzKzAtMC44MzIrMC4zMjEtMC44MzIrMC43NDl2MjQuODQ1YzArMC40MjgrMC4zOTgrMC43NzQrMC44MzIrMC43NzRoMjQuODQ4YzAuNDMzKzArMC43NTMtMC4zNDcrMC43NTMtMC43NzR2LTI0Ljg0NWMwLTAuNDI4LTAuMzE5LTAuNzQ5LTAuNzUzLTAuNzQ5ek0yNS43MjgrMTIuMzgyaC00LjU0NHYtMS45MmgtMS43OTJ2MTAuNDk2aDEuOTJ2NS4wNTZoLTguNjR2LTQuOGgxLjUzNnYtMTAuNTZoLTEuNTM2djEuNzI4aC00Ljh2LTYuNDY0aDE3Ljg1NnY2LjQ2NHoiLz48L3N2Zz4=
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org
[codecov-badge]: https://img.shields.io/codecov/c/github/zchee/go-apfs.svg?style=flat-square&label=%20%20Codecov%2Eio&logo=data%3Aimage%2Fsvg%2Bxml%3Bcharset%3Dutf-8%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI0MCIgaGVpZ2h0PSI0MCIgdmlld0JveD0iMCAwIDI1NiAyODEiPjxwYXRoIGZpbGw9IiNFRUUiIGQ9Ik0yMTguNTUxIDM3LjQxOUMxOTQuNDE2IDEzLjI4OSAxNjIuMzMgMCAxMjguMDk3IDAgNTcuNTM3LjA0Ny4wOTEgNTcuNTI3LjA0IDEyOC4xMjFMMCAxNDkuODEzbDE2Ljg1OS0xMS40OWMxMS40NjgtNy44MTQgMjQuNzUtMTEuOTQ0IDM4LjQxNy0xMS45NDQgNC4wNzkgMCA4LjE5OC4zNzMgMTIuMjQgMS4xMSAxMi43NDIgMi4zMiAyNC4xNjUgOC4wODkgMzMuNDE0IDE2Ljc1OCAyLjEyLTQuNjcgNC42MTQtOS4yMDkgNy41Ni0xMy41MzZhODguMDgxIDg4LjA4MSAwIDAgMSAzLjgwNS01LjE1Yy0xMS42NTItOS44NC0yNS42NDktMTYuNDYzLTQwLjkyNi0xOS4yNDVhOTAuMzUgOTAuMzUgMCAwIDAtMTYuMTItMS40NTkgODguMzc3IDg4LjM3NyAwIDAgMC0zMi4yOSA2LjA3YzguMzYtNTEuMjIyIDUyLjg1LTg5LjM3IDEwNS4yMy04OS40MDggMjguMzkyIDAgNTUuMDc4IDExLjA1MyA3NS4xNDkgMzEuMTE3IDE2LjAxMSAxNi4wMSAyNi4yNTQgMzYuMDMzIDI5Ljc4OCA1OC4xMTctMTAuMzI5LTQuMDM1LTIxLjIxMi02LjEtMzIuNDAzLTYuMTQ0bC0xLjU2OC0uMDA3YTkwLjk1NyA5MC45NTcgMCAwIDAtMy40MDEuMTExYy0xLjk1NS4xLTMuODk4LjI3Ny01LjgyMS41LS41NzQuMDYzLTEuMTM5LjE1My0xLjcwNy4yMzEtMS4zNzguMTg2LTIuNzUuMzk1LTQuMTA5LjYzOS0uNjAzLjExLTEuMjAzLjIzMS0xLjguMzUxYTkwLjUxNyA5MC41MTcgMCAwIDAtNC4xMTQuOTM3Yy0uNDkyLjEyNi0uOTgzLjI0My0xLjQ3LjM3NGE5MC4xODMgOTAuMTgzIDAgMCAwLTUuMDkgMS41MzhjLS4xLjAzNS0uMjA0LjA2My0uMzA0LjA5NmE4Ny41MzIgODcuNTMyIDAgMCAwLTExLjA1NyA0LjY0OWMtLjA5Ny4wNS0uMTkzLjEwMS0uMjkzLjE1MWE4Ni43IDg2LjcgMCAwIDAtNC45MTIgMi43MDFsLS4zOTguMjM4YTg2LjA5IDg2LjA5IDAgMCAwLTIyLjMwMiAxOS4yNTNjLS4yNjIuMzE4LS41MjQuNjM1LS43ODQuOTU4LTEuMzc2IDEuNzI1LTIuNzE4IDMuNDktMy45NzYgNS4zMzZhOTEuNDEyIDkxLjQxMiAwIDAgMC0zLjY3MiA1LjkxMyA5MC4yMzUgOTAuMjM1IDAgMCAwLTIuNDk2IDQuNjM4Yy0uMDQ0LjA5LS4wODkuMTc1LS4xMzMuMjY1YTg4Ljc4NiA4OC43ODYgMCAwIDAtNC42MzcgMTEuMjcybC0uMDAyLjAwOXYuMDA0YTg4LjAwNiA4OC4wMDYgMCAwIDAtNC41MDkgMjkuMzEzYy4wMDUuMzk3LjAwNS43OTQuMDE5IDEuMTkyLjAyMS43NzcuMDYgMS41NTcuMTA0IDIuMzM4YTk4LjY2IDk4LjY2IDAgMCAwIC4yODkgMy44MzRjLjA3OC44MDQuMTc0IDEuNjA2LjI3NSAyLjQxLjA2My41MTIuMTE5IDEuMDI2LjE5NSAxLjUzNGE5MC4xMSA5MC4xMSAwIDAgMCAuNjU4IDQuMDFjNC4zMzkgMjIuOTM4IDE3LjI2MSA0Mi45MzcgMzYuMzkgNTYuMzE2bDIuNDQ2IDEuNTY0LjAyLS4wNDhhODguNTcyIDg4LjU3MiAwIDAgMCAzNi4yMzIgMTMuNDVsMS43NDYuMjM2IDEyLjk3NC0yMC44MjItNC42NjQtLjEyN2MtMzUuODk4LS45ODUtNjUuMS0zMS4wMDMtNjUuMS02Ni45MTcgMC0zNS4zNDggMjcuNjI0LTY0LjcwMiA2Mi44NzYtNjYuODI5bDIuMjMtLjA4NWMxNC4yOTItLjM2MiAyOC4zNzIgMy44NTkgNDAuMzI1IDExLjk5N2wxNi43ODEgMTEuNDIxLjAzNi0yMS41OGMuMDI3LTM0LjIxOS0xMy4yNzItNjYuMzc5LTM3LjQ0OS05MC41NTQiLz48L3N2Zz4=
[release-badge]: https://img.shields.io/github/release/zchee/go-apfs.svg?style=flat-square
[ga-badge]: https://ga-beacon.appspot.com/UA-89201129-1/gist-go-template?flat&useReferer&pixel
