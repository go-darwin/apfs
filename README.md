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


## API

```go
import "github.com/zchee/go-apfs"
```

## <a name="pkg-index">Index</a>
* [func CloneFile(src, dst string, flag CLONEFILE_FLAG) error](#CloneFile)
* [func CloneFileAt(src, dst string, flag CLONEFILE_FLAG) error](#CloneFileAt)
* [func CopyFile(src, dst string, state COPYFILE_STATE, flag COPYFILE_FLAG) (bool, error)](#CopyFile)
* [func CopyFileStateFree(state COPYFILE_STATE) error](#CopyFileStateFree)
* [func CopyFileStateGet(state COPYFILE_STATE, flag COPYFILE_STATE_FLAG, result *int) error](#CopyFileStateGet)
* [func CopyFileStateSet(state COPYFILE_STATE, flag COPYFILE_STATE_FLAG, result *int) error](#CopyFileStateSet)
* [func FcloneFileAt(srcFd uintptr, dst string, flag CLONEFILE_FLAG) error](#FcloneFileAt)
* [func FcopyFile(src, dst uintptr, state COPYFILE_STATE, flag COPYFILE_FLAG) error](#FcopyFile)
* [func RenameatxNp(src, dst string, flag RENAME_FALG) error](#RenameatxNp)
* [func RenamexNp(src, dst string, flag RENAME_FALG) error](#RenamexNp)
* [type CLONEFILE_FLAG](#CLONEFILE_FLAG)
* [type COPYFILE_FLAG](#COPYFILE_FLAG)
* [type COPYFILE_RECURSE_CALLBACK](#COPYFILE_RECURSE_CALLBACK)
* [type COPYFILE_STATE](#COPYFILE_STATE)
  * [func CopyFileStateAlloc() COPYFILE_STATE](#CopyFileStateAlloc)
* [type COPYFILE_STATE_FLAG](#COPYFILE_STATE_FLAG)
* [type RENAME_FALG](#RENAME_FALG)

#### <a name="pkg-files">Package files</a>
[apfs.go](/src/github.com/zchee/go-apfs/apfs.go) [clone.go](/src/github.com/zchee/go-apfs/clone.go) [copy.go](/src/github.com/zchee/go-apfs/copy.go) [rename.go](/src/github.com/zchee/go-apfs/rename.go) 

## <a name="CloneFile">func</a> [CloneFile](/src/target/clone.go?s=1953:2011#L44)
``` go
func CloneFile(src, dst string, flag CLONEFILE_FLAG) error
```
CloneFile function causes the named file src to be cloned to the named file dst.

The cloned file dst shares its data blocks with the src file but has its own copy of attributes,
extended attributes and ACL's which are identical to those of the named file src with the exceptions listed below

1. ownership information is set as it would be if dst was created by openat(2) or mkdirat(2) or symlinkat(2) if the current


	user does not have privileges to change ownership.

2. setuid and setgid bits are turned off in the mode bits for regular files.

Subsequent writes to either the original or cloned file are private to the file being modified (copy-on-write).

The named file dst must not exist for the call to be successful. Since the clonefile() system call might not allocate new storage for
data blocks, it is possible for a subsequent overwrite of an existing data block to return ENOSPC.

If src names a directory, the directory hierarchy is cloned as if each item was cloned individually.

However, the use of copyfile(3) is more appropriate for copying large directory hierarchies instead of clonefile(2).

```c
int
clonefile(const char * src, const char * dst, int flags);
```


## <a name="CloneFileAt">func</a> [CloneFileAt](/src/target/clone.go?s=2925:2985#L64)
``` go
func CloneFileAt(src, dst string, flag CLONEFILE_FLAG) error
```
CloneFileAt is equivalent to clonefile() except in the case where either src or dst specifies a relative path.

If src is a relative path, the file to be cloned is located relative to the directory associated with the file descriptor
src_dirfd instead of the current working directory.

If dst is a relative path, the same happens only relative to the directory associated with dst_dirfd.

If clonefileat() is passed the special value AT_FDCWD in either the src_dirfd or dst_dirfd parameters,
the current working directory is used in the determination of the file for the respective path parameter.

```c
int
clonefileat(int src_dirfd, const char * src, int dst_dirfd, const char * dst, int flags);
```


## <a name="CopyFile">func</a> [CopyFile](/src/target/copy.go?s=11965:12051#L204)
``` go
func CopyFile(src, dst string, state COPYFILE_STATE, flag COPYFILE_FLAG) (bool, error)
```
CopyFile function can copy the named from file to the named to file.

If the state parameter is the return value from CopyFileStateAlloc, then CopyFile will use the information from the state object.

If it is NULL, then both functions will work normally, but less control will be available to the caller.

```c
int
copyfile(const char *from, const char *to, copyfile_state_t state, copyfile_flags_t flags);
```


## <a name="CopyFileStateFree">func</a> [CopyFileStateFree](/src/target/copy.go?s=13720:13770#L247)
``` go
func CopyFileStateFree(state COPYFILE_STATE) error
```
CopyFileStateFree function is used to deallocate the object and its contents.

```c
int
copyfile_state_free(copyfile_state_t state);
```



## <a name="CopyFileStateGet">func</a> [CopyFileStateGet](/src/target/copy.go?s=14172:14260#L261)
``` go
func CopyFileStateGet(state COPYFILE_STATE, flag COPYFILE_STATE_FLAG, result *int) error
```
CopyFileStateGet functions can be used to manipulate the COPYFILE_STATE object returned by CopyFileStateAlloc.

The dst parameter's type depends on the flag parameter that is passed in.

```c
int
copyfile_state_get(copyfile_state_t state, uint32_t flag, void * dst);
```



## <a name="CopyFileStateSet">func</a> [CopyFileStateSet](/src/target/copy.go?s=14689:14777#L275)
``` go
func CopyFileStateSet(state COPYFILE_STATE, flag COPYFILE_STATE_FLAG, result *int) error
```
CopyFileStateSet functions can be used to manipulate the COPYFILE_STATE object returned by CopyFileStateAlloc.

The dst parameter's type depends on the flag parameter that is passed in.

```c
int
copyfile_state_set(copyfile_state_t state, uint32_t flag, const void * src);
```


## <a name="FcloneFileAt">func</a> [FcloneFileAt](/src/target/clone.go?s=3669:3740#L87)
``` go
func FcloneFileAt(srcFd uintptr, dst string, flag CLONEFILE_FLAG) error
```
FcloneFileAt function is similar to clonefileat() except that the source is identified by file descriptor srcfd rather
than a path (as in clonefile() or clonefileat())

The flags parameter specifies the options that can be passed.


```c
int
fclonefileat(int srcfd, int dst_dirfd, const char * dst, int flags);
```


## <a name="FcopyFile">func</a> [FcopyFile](/src/target/copy.go?s=12901:12981#L225)
``` go
func FcopyFile(src, dst uintptr, state COPYFILE_STATE, flag COPYFILE_FLAG) error
```
FcopyFile function does the same CopyFile, but using the file descriptors of already-opened files.

If the state parameter is the return value from CopyFileStateAlloc, then FcopyFile will use the information from the state object.

If it is NULL, then both functions will work normally, but less control will be available to the caller.

```c
int
fcopyfile(int from, int to, copyfile_state_t state, copyfile_flags_t flags);
```


## <a name="RenameatxNp">func</a> [RenameatxNp](/src/target/rename.go?s=1696:1753#L40)
``` go
func RenameatxNp(src, dst string, flag RENAME_FALG) error
```
RenamexNp system calls are similar to rename() and renameat() counterparts except that they take a flags argument.

```c
int
renameatx_np(int fromfd, const char *from, int tofd, const char *to, unsigned int flags);
```


## <a name="RenamexNp">func</a> [RenamexNp](/src/target/rename.go?s=1244:1299#L29)
``` go
func RenamexNp(src, dst string, flag RENAME_FALG) error
```
RenamexNp system calls are similar to rename() and renameat() counterparts except that they take a flags argument.

```c
int
renamex_np(const char *from, const char *to, unsigned int flags);
```



## <a name="CLONEFILE_FLAG">type</a> [CLONEFILE_FLAG](/src/target/clone.go?s=442:468#L14)
``` go
type CLONEFILE_FLAG uint32
```
CLONEFILE_FLAG provides the clonefile(2) flag.


``` go
var (
    // CLONE_NOFOLLOW don't follow the src file if it is a symbolic link (applicable only if the source is not a directory).
    //
    // The symbolic link is itself cloned if src names a symbolic link.
    CLONE_NOFOLLOW CLONEFILE_FLAG = 0x0001
)
```


## <a name="COPYFILE_FLAG">type</a> [COPYFILE_FLAG](/src/target/copy.go?s=341:363#L11)
``` go
type COPYFILE_FLAG int
```
COPYFILE_FLAG provides the copyfile flag.


``` go
var (
    // COPYFILE_ACL copy the source file's access control lists.
    COPYFILE_ACL COPYFILE_FLAG = 1 << 0
    // COPYFILE_STAT copy the source file's POSIX information (mode, modification time, etc.).
    COPYFILE_STAT COPYFILE_FLAG = 1 << 1
    // COPYFILE_XATTR copy the source file's extended attributes.
    COPYFILE_XATTR COPYFILE_FLAG = 1 << 2
    // COPYFILE_DATA copy the source file's data.
    COPYFILE_DATA COPYFILE_FLAG = 1 << 3

    // COPYFILE_SECURITY copy the source file's POSIX and ACL information; equivalent to (COPYFILE_STAT|COPYFILE_ACL).
    COPYFILE_SECURITY COPYFILE_FLAG = COPYFILE_STAT | COPYFILE_ACL
    // COPYFILE_METADATA copy the metadata; equivalent to (COPYFILE_SECURITY|COPYFILE_XATTR).
    COPYFILE_METADATA COPYFILE_FLAG = COPYFILE_SECURITY | COPYFILE_XATTR
    // COPYFILE_ALL copy the entire file; equivalent to (COPYFILE_METADATA|COPYFILE_DATA).
    COPYFILE_ALL COPYFILE_FLAG = COPYFILE_METADATA | COPYFILE_DATA

    // This flag is not used by FcopyFile; see below for more information.
    COPYFILE_RECURSIVE COPYFILE_FLAG = 1 << 15 // Descend into hierarchies
    // COPYFILE_CHECK return a bitmask (corresponding to the flags argument) indicating which contents would be copied
    // no data are actually copied.
    //
    // (E.g., if flags was set to COPYFILE_CHECK|COPYFILE_METADATA, and the from
    // file had extended attributes but no ACLs, the return value would be COPYFILE_XATTR .)
    COPYFILE_CHECK COPYFILE_FLAG = 1 << 16 // return flags for xattr or acls if set
    // COPYFILE_EXCL fail if the to file already exists.  (This is only applicable for the CopyFile function.)
    COPYFILE_EXCL COPYFILE_FLAG = 1 << 17 // fail if destination exists
    // COPYFILE_NOFOLLOW_SRC do not follow the from file, if it is a symbolic link.  (This is only applicable for the CopyFile function.)
    COPYFILE_NOFOLLOW_SRC COPYFILE_FLAG = 1 << 18 // don't follow if source is a symlink
    // COPYFILE_NOFOLLOW_DST do not follow the to file, if it is a symbolic link.  (This is only applicable for the CopyFile function.)
    COPYFILE_NOFOLLOW_DST COPYFILE_FLAG = 1 << 19 // don't follow if dst is a symlink
    // COPYFILE_MOVE unlink (using remove(3)) the from file.  (This is only applicable for the copyfile() function.) No error is returned if remove(3) fails.
    //
    // Note that remove(3) removes a symbolic link itself, not the target of the link.
    COPYFILE_MOVE COPYFILE_FLAG = 1 << 20 // unlink src after copy
    // COPYFILE_UNLINK unlink the to file before starting.  (This is only applicable for the copyfile() function.)
    COPYFILE_UNLINK COPYFILE_FLAG = 1 << 21 // unlink dst before copy
    // COPYFILE_NOFOLLOW this is a convenience macro, equivalent to (COPYFILE_NOFOLLOW_DST|COPYFILE_NOFOLLOW_SRC).
    COPYFILE_NOFOLLOW COPYFILE_FLAG = COPYFILE_NOFOLLOW_SRC | COPYFILE_NOFOLLOW_DST

    // COPYFILE_PACK serialize the from file. The to file is an AppleDouble-format file.
    COPYFILE_PACK COPYFILE_FLAG = 1 << 22
    // COPYFILE_UNPACK unserialize the from file.
    //
    // The from file is an AppleDouble-format file; the to file will have the extended attributes, ACLs, resource fork, and FinderInfo data from the to file, regardless of the flags argument passed in.
    COPYFILE_UNPACK COPYFILE_FLAG = 1 << 23

    // COPYFILE_CLONE try to clone the file/directory instead.
    //
    // This is a best try flag i.e. if cloning fails, fallback to copying the file.
    //
    // This flag is equivalent to (COPYFILE_EXCL | COPYFILE_ACL | COPYFILE_STAT | COPYFILE_XATTR | COPYFILE_DATA).
    //
    // Note that if cloning is successful, callbacks will not be invoked.
    COPYFILE_CLONE COPYFILE_FLAG = 1 << 24
    // COPYFILE_CLONE_FORCE clone the file/directory instead.  This is a force flag i.e. if cloning fails, an error is returned.
    //
    // This flag is equivalent to (COPYFILE_EXCL | COPYFILE_ACL | COPYFILE_STAT | COPYFILE_XATTR | COPYFILE_DATA).
    //
    // Note that if cloning is successful, callbacks will not be invoked.
    COPYFILE_CLONE_FORCE COPYFILE_FLAG = 1 << 25

    // COPYFILE_RUN_IN_PLACE if the src file has quarantine information, add the QTN_FLAG_DO_NOT_TRANSLOCATE flag to the quarantine information of the dst file.
    //
    // This allows a bundle to run in place instead of being translocated.
    COPYFILE_RUN_IN_PLACE COPYFILE_FLAG = 1 << 26

    COPYFILE_VERBOSE COPYFILE_FLAG = 1 << 30
)
```


## <a name="COPYFILE_RECURSE_CALLBACK">type</a> [COPYFILE_RECURSE_CALLBACK](/src/target/copy.go?s=8510:8547#L147)
``` go
type COPYFILE_RECURSE_CALLBACK uint32
```
COPYFILE_RECURSE_CALLBACK provides the copyfile callbacks.


``` go
var (
    // COPYFILE_RECURSE_ERROR there was an error in processing an element of the source hierarchy.
    //
    // This happens when fts(3) returns an error or unknown file type.  (Currently, the second argument to the call-back function will always be COPYFILE_ERR in this case.)
    COPYFILE_RECURSE_ERROR COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_ERROR // 0
    // COPYFILE_RECURSE_FILE the object being copied is a file (or, rather, something other than a directory).
    COPYFILE_RECURSE_FILE COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_FILE // 1
    // COPYFILE_RECURSE_DIR the object being copied is a directory, and is being entered. (That is, none of the filesystem objects contained within the directory have been copied yet.)
    COPYFILE_RECURSE_DIR COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR // 2
    // COPYFILE_RECURSE_DIR_CLEANUP the object being copied is a directory, and all of the objects contained have been copied.
    //
    // At this stage, the destination directory being copied will have any extra permissions that were added to allow the copying will be removed.
    COPYFILE_RECURSE_DIR_CLEANUP COPYFILE_RECURSE_CALLBACK = C.COPYFILE_RECURSE_DIR_CLEANUP // 3
    COPYFILE_COPY_DATA           COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_DATA           // 4
    COPYFILE_COPY_XATTR          COPYFILE_RECURSE_CALLBACK = C.COPYFILE_COPY_XATTR          // 5
)
```

``` go
var (
    // COPYFILE_START before copying has begun. The third parameter will be a newly-created copyfile_state_t object with the callback function and context pre-loaded.
    COPYFILE_START COPYFILE_RECURSE_CALLBACK = C.COPYFILE_START // 1
    // COPYFILE_FINISH after copying has successfully finished.
    COPYFILE_FINISH COPYFILE_RECURSE_CALLBACK = C.COPYFILE_FINISH // 2
    // COPYFILE_ERR indicates an error has happened at some stage.
    //
    // If the first argument to the call-back function is COPYFILE_RECURSE_ERROR, then an error occurred while processing the source hierarchy;
    //
    // otherwise, it will indicate what type of object was being copied, and errno will be set to indicate the error.
    COPYFILE_ERR      COPYFILE_RECURSE_CALLBACK = C.COPYFILE_ERR      // 3
    COPYFILE_PROGRESS COPYFILE_RECURSE_CALLBACK = C.COPYFILE_PROGRESS // 4
)
```

``` go
var (
    // COPYFILE_CONTINUE the copy will continue as expected.
    COPYFILE_CONTINUE COPYFILE_RECURSE_CALLBACK = C.COPYFILE_CONTINUE // 0
    // COPYFILE_SKIP this object will be skipped, and the next object will be processed.
    //
    // (Note that, when entering a directory. returning COPYFILE_SKIP from the call-back function will prevent the contents of the directory from being copied.)
    COPYFILE_SKIP COPYFILE_RECURSE_CALLBACK = C.COPYFILE_SKIP // 1
    // COPYFILE_QUIT the entire copy is aborted at this stage.  Any filesystem objects created up to this point will remain.
    //
    // CopyFile will return -1, but errno will be unmodified.
    COPYFILE_QUIT COPYFILE_RECURSE_CALLBACK = C.COPYFILE_QUIT // 2
)
```


## <a name="COPYFILE_STATE">type</a> [COPYFILE_STATE](/src/target/copy.go?s=11499:11537#L194)
``` go
type COPYFILE_STATE C.copyfile_state_t
```
COPYFILE_STATE provides the copyfile state.


### <a name="CopyFileStateAlloc">func</a> [CopyFileStateAlloc](/src/target/copy.go?s=13484:13524#L239)
``` go
func CopyFileStateAlloc() COPYFILE_STATE
```
CopyFileStateAlloc function initializes a copyfile_state_t object (which is an opaque data type).

This object can be passed to CopyFile and FcopyFile; CopyfileStateGet and CopyFileStateSet can be used to manipulate the state (see below).

```c
copyfile_state_t
copyfile_state_alloc(void);
```


## <a name="COPYFILE_STATE_FLAG">type</a> [COPYFILE_STATE_FLAG](/src/target/copy.go?s=4736:4767#L86)
``` go
type COPYFILE_STATE_FLAG uint32
```
COPYFILE_STATE_FLAG provides the copyfile state flag.


``` go
var (
    // COPYFILE_STATE_SRC_FD get or set the file descriptor associated with the source (or destination) file.
    //
    // If this has not been initialized yet, the value will be -2.
    //
    // The dst (for CopyFileStateGet) and src (for CopyFileStateSet) parameters are pointers to int.
    COPYFILE_STATE_SRC_FD COPYFILE_STATE_FLAG = C.COPYFILE_STATE_SRC_FD
    // COPYFILE_STATE_DST_FD get or set the file descriptor associated with the source (or destination) file.
    //
    // If this has not been initialized yet, the value will be -2.
    //
    // The dst (for CopyFileStateGet) and src (for CopyFileStateSet) parameters are pointers to int.
    COPYFILE_STATE_DST_FD COPYFILE_STATE_FLAG = C.COPYFILE_STATE_DST_FD
    // COPYFILE_STATE_SRC_FILENAME get or set the filename associated with the source (or destination) file.
    //
    // If it has not been initialized yet, the value will be NULL.
    //
    // For copyfile_state_set(), the src parameter is a pointer to a C string (i.e., char* ); copyfile_state_set() makes a private copy of this string.
    //
    // For copyfile_state_get() function, the dst parameter is a pointer to a pointer to a C string (i.e., char** ); the returned value is a pointer to the state 's copy, and must not be modified or released.
    COPYFILE_STATE_SRC_FILENAME COPYFILE_STATE_FLAG = C.COPYFILE_STATE_SRC_FILENAME
    // COPYFILE_STATE_DST_FILENAME get or set the filename associated with the source (or destination) file.
    //
    // If it has not been initialized yet, the value will be NULL.
    //
    // For copyfile_state_set(), the src parameter is a pointer to a C string (i.e., char* ); copyfile_state_set() makes a private copy of this string.
    //
    // For copyfile_state_get() function, the dst parameter is a pointer to a pointer to a C string (i.e., char** ); the returned value is a pointer to the state 's copy, and must not be modified or released.
    COPYFILE_STATE_DST_FILENAME COPYFILE_STATE_FLAG = C.COPYFILE_STATE_DST_FILENAME
    // COPYFILE_STATE_STATUS_CB get or set the callback status function (currently only used for recursive copies; see below for details).
    //
    // The src parameter is a pointer to a function of type copyfile_callback_t (see above).
    COPYFILE_STATE_STATUS_CB COPYFILE_STATE_FLAG = C.COPYFILE_STATE_STATUS_CB
    // COPYFILE_STATE_STATUS_CTX get or set the context parameter for the status call-back function (see below for details).
    //
    // The src parameter is a void *.
    COPYFILE_STATE_STATUS_CTX COPYFILE_STATE_FLAG = C.COPYFILE_STATE_STATUS_CTX
    // COPYFILE_STATE_QUARANTINE get or set the quarantine information with the source file.
    //
    // The src parameter is a pointer to an opaque object (type void * ).
    COPYFILE_STATE_QUARANTINE COPYFILE_STATE_FLAG = C.COPYFILE_STATE_QUARANTINE
    // COPYFILE_STATE_COPIED get the number of data bytes copied so far.  (Only valid for copyfile_state_get(); see below for
    // more details about callbacks.)
    //
    // If a COPYFILE_CLONE or COPYFILE_CLONE_FORCE operation successfully cloned the requested objects, then this value will be 0.
    //
    // The dst parameter is a pointer to off_t (type off_t * ).
    COPYFILE_STATE_COPIED COPYFILE_STATE_FLAG = C.COPYFILE_STATE_COPIED
    // COPYFILE_STATE_XATTRNAME get the name of the extended attribute during a callback for COPYFILE_COPY_XATTR (see below for details).
    //
    // This field cannot be set, and may be NULL.
    COPYFILE_STATE_XATTRNAME COPYFILE_STATE_FLAG = C.COPYFILE_STATE_XATTRNAME
    // COPYFILE_STATE_WAS_CLONED true if a COPYFILE_CLONE or COPYFILE_CLONE_FORCE operation successfully cloned the requested objects.
    //
    // The dst parameter is a pointer to bool (type bool * ).
    COPYFILE_STATE_WAS_CLONED COPYFILE_STATE_FLAG = C.COPYFILE_STATE_WAS_CLONED
)
```


## <a name="RENAME_FALG">type</a> [RENAME_FALG](/src/target/rename.go?s=356:377#L11)
``` go
type RENAME_FALG uint
```
RENAME_FALG provides the rename flag.


``` go
var (
    // RENAME_SWAP on file systems that support it (see getattrlist(2) VOL_CAP_INT_RENAME_SWAP), it will cause the source and target to be atomically swapped.
    //
    // Source and target need not be of the same type, i.e. it is possible to swap a file with a directory.
    //
    // EINVAL is returned in case of bitwise-inclusive OR with RENAME_EXCL.
    RENAME_SWAP RENAME_FALG = 0x00000002
    // RENAME_EXCL on file systems that support it (see getattrlist(2) VOL_CAP_INT_RENAME_EXCL), it will cause EEXIST to be returned if the destination already exists.
    //
    // EINVAL is returned in case of bitwise-inclusive OR with RENAME_SWAP.
    RENAME_EXCL RENAME_FALG = 0x00000004
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
