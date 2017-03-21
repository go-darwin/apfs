# Apple File System

THIS DOCUMENT COPIED FROM THE [Apple File System Guide](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40016999)

## Introduction

[Introduction](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40016999-CH1-DontLinkElementID_18)

HFS+ and its predecessor HFS are more than 30 years old. These file systems were developed in an era of floppy disks and spinning hard drives, when file sizes were calculated in kilobytes or megabytes.

Today, people commonly store hundreds of gigabytes and access millions of files on high-speed, low-latency flash drives. People carry their data with them, and they demand that sensitive information be secure.

Apple File System is a new, modern file system for iOS, macOS, tvOS, and watchOS. It is optimized for Flash/SSD storage and features strong encryption, copy-on-write metadata, space sharing, cloning for files and directories, snapshots, fast directory sizing, atomic safe-save primitives, and improved file system fundamentals.

A Developer Preview of Apple File System is available in macOS Sierra. Apple plans to release Apple File System as a bootable file system in 2017.

### Prerequisites

To understand how your app interacts with the file system, read [File System Programming Guide](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672).


## Features

- [Fuatures](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_17)
  - [Clones](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_4")
  - [Snapshots](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_5")
  - [Space Sharing](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_6")
  - [Encryption](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_7")
  - [Crash Protection](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_8)
  - [Sparse Files](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_9)
  - [Fast Directory Sizing](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_10)
  - [Atomic Safe-Save](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/Features/Features.html#//apple_ref/doc/uid/TP40016999-CH5-DontLinkElementID_11)

Apple File System is a 64-bit file system supporting over 9 quintillion files on a single volume. This state-of-the-art file system features cloning for files and directories, snapshots, space sharing, fast directory sizing, atomic safe-save primitives, and improved filesystem fundamentals, as well as a unique copy-on-write design that uses I/O coalescing to deliver maximum performance while ensuring data reliability.

### Clones

A clone is a nearly instantaneous copy of a file or directory that occupies no additional space for file data. Clones allow the operating system to make fast, power-efficient file copies on the same volume without occupying additional storage space.

Modifications to the data write the new data elsewhere and continue to share the unmodified blocks. Changes to a file are saved as deltas of the cloned file, reducing storage space required for document revisions and copies.

### Snapshots

A volume snapshot is a point-in-time, read-only instance of the file system.

The operating system uses snapshots to make backups work more efficiently and offer a way to revert changes to a given point in time.

### Space Sharing

Space Sharing allows multiple file systems to share the same underlying free space on a physical volume. Unlike rigid partitioning schemes that pre-allocate a fixed amount of space for each file system, APFS-formatted volumes can grow and shrink without volume repartitioning.

Each volume in an APFS container reports the same available disk space, which is equal to the total available disk space of the container. For example, for an APFS container with a capacity of 100GB that contains volume A (which uses 10GB) and volume B (which uses 20GB), the free space reported for both volumes A and B is 70GB (100GB - 10GB - 20GB).

### Encryption

Security and privacy are fundamental in the design of Apple File System. That's why Apple File System implements strong full-disk encryption, encrypting files and all sensitive metadata.

Which encryption methods are available depends on hardware and operating system support, and can vary for Mac, iPhone, iPad, Apple TV, and Apple Watch.

Apple File System supports the following encryption models for each volume in a container:

- No encryption
- Single-key encryption
- Multi-key encryption with per-file keys for file data and a separate key for sensitive metadata

Multi-key encryption ensures the integrity of user data. Even if someone were to compromise the physical security of the device and gain access to the device key, they still couldn't decrypt the user's files.

Apple File System uses AES-XTS or AES-CBC encryption modes, depending on hardware.

### Crash Protection

Apple File System uses a novel copy-on-write metadata scheme to ensure that updates to the file system are crash protected, without the write-twice overhead of journaling.

### Sparse Files

Apple File System supports sparse files, a more efficient way of representing empty blocks on disk. With sparse files storage is allocated only when actually needed.

### Fast Directory Sizing

Fast Directory Sizing allows Apple File System to quickly compute the total space used by a directory hierarchy, and update it as the hierarchy evolves.

### Atomic Safe-Save

Apple File System introduces a new Atomic Safe-Save primitive for bundles and directories. Atomic Safe-Save performs renames in a single transaction such that, from the user’s perspective, the operation either is completed or does not happen at all.


## Frequently Asked Questions

- [Frequently Asked Questions](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/FAQ/FAQ.html#//apple_ref/doc/uid/TP40016999-CH6-DontLinkElementID_16)
  - [Compatibility](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/FAQ/FAQ.html#//apple_ref/doc/uid/TP40016999-CH6-DontLinkElementID_1)
  - [Upgrading](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/FAQ/FAQ.html#//apple_ref/doc/uid/TP40016999-CH6-DontLinkElementID_2)
  - [Implementation](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/FAQ/FAQ.html#//apple_ref/doc/uid/TP40016999-CH6-DontLinkElementID_3)

### Compatibility

#### What are the limitations of Apple File System in macOS Sierra?

macOS Sierra includes a Developer Preview release of Apple File System. As a Developer Preview, it has several limitations:

- Startup Disk: An APFS-formatted volume cannot be used as a startup disk.
- Case Sensitivity: Filenames are case-sensitive only.
- Time Machine: Time Machine backups are not supported.
- FileVault: APFS-formatted volumes cannot be encrypted using FileVault.
- Fusion Drive: Apple File System cannot use Fusion Drives.

#### Can I use Apple File System with my existing hard disk drive?

Yes. Apple File System is optimized for Flash/SSD storage, but can also be used with traditional hard disk drives (HDD) and external, direct-attached storage.

#### Can I reshare APFS-formatted volumes using a network file-sharing protocol?

Yes, you can share APFS-formatted volumes using the SMB or NFS network file-sharing protocol.

You cannot share APFS-formatted volumes using AFP. The AFP protocol is deprecated.

#### Can I use my third-party disk utilities with an APFS-formatted hard disk?

Existing third-party utilities will need to be updated to support Apple File System. Consult the utility's documentation, or contact the vendor for compatibility information.

#### Can I boot macOS Sierra from an APFS-formatted hard disk?

No. macOS Sierra supports Apple File System for data volumes only. You cannot boot macOS Sierra from a APFS-formatted volume. 

### Upgrading

#### How do I upgrade to Apple File System?

Apple will offer nondestructive in-place upgrades from HFS+ to APFS for all boot volumes when Apple File System ships in 2017. Tools will be available to convert external volumes from HFS+ to APFS format.

#### If I convert a volume to APFS, can I later revert to HFS+?

You can use Disk Utility to erase an APFS-formatted volume and reformat as HFS+. However, your data will not be preserved when you reformat the volume as HFS+.

### Implementation

#### Why did Apple develop APFS?

Apple File System is uniquely designed to meet the needs of Apple’s products and ecosystem. Apple File System provides strong encryption, ultra-low latencies and limited memory overhead. It is optimized for Flash/SSD storage and can be used on everything from an Apple Watch to a Mac Pro.

#### Can RAID be used with Apple File System?

Yes. Apple File System does not directly implement software RAID; however APFS-formatted volumes can be combined with an Apple RAID volume to support Striping (RAID 0), Mirroring (RAID 1), and Concatenation (JBOD). APFS-formatted volumes can also be used with direct-attached hardware RAID solutions.

#### Does Apple File System support directory hard links?

Directory hard links are not supported by Apple File System. All directory hard links are converted to symbolic links or aliases when you convert from HFS+ to APFS volume formats on macOS.

#### Does Apple File System support ditto blocks?

Ditto blocks are primarily used to protect against corrupted sectors in large storage arrays with unreliable hard disk drives. Apple File System takes advantage of modern hardware with strong checksums and error correction in firmware, and does not depend on ditto blocks.

#### Does Apple File System support redundant metadata?

With modern Flash/SSD storage, writing two blocks of data to different locations does not guarantee that the blocks will be written to separate locations. The Flash translation layer typically groups writes together into the same NAND block. Therefore it affords no extra protection to write a second copy at the same time the first copy is written.

#### What has Apple done to ensure the reliability of my data?

Apple products are designed to prevent data corruption and protect against data loss.

To protect data from hardware errors, all Flash/SSD and hard disk drives used in Apple products use Error Correcting Code (ECC). ECC checks for transmission errors, and when necessary, corrects on the fly. Apple File System uses a unique copy-on-write scheme to protect against data loss that can occur during a crash or loss of power. And to further ensure data integrity, Apple File System uses the Fletcher's checksum algorithm for metadata operations.

#### Does Apple File System use journaling?

Apple File System uses copy-on-write to avoid in-place changes to file data, which ensures that file system updates are crash protected without the write-twice overhead of journaling.

#### Does Apple File System support data deduplication?

No. With Apple File System individual extents can be encrypted, making it impossible to examine and deduplicate files and their content. Apple File System uses clone files to minimize data storage and data duplication.

#### Does Apple File System support TRIM operations?

Yes. TRIM operations are issued asynchronously from when files are deleted or free space is reclaimed, which ensures that these operations are performed only after metadata changes are persisted to stable storage.

#### Is APFS open source?

An open source implementation is not available at this time. Apple plans to document and publish the APFS volume format specification when Apple File System is released for macOS in 2017.


## Tools and APIs

- [Tools and APIs](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/ToolsandAPIs/ToolsandAPIs.html#//apple_ref/doc/uid/TP40016999-CH7-DontLinkElementID_20)
  - [Disk Image Tools](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/ToolsandAPIs/ToolsandAPIs.html#//apple_ref/doc/uid/TP40016999-CH7-DontLinkElementID_13)
  - [Enhanced APIs](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/ToolsandAPIs/ToolsandAPIs.html#//apple_ref/doc/uid/TP40016999-CH7-DontLinkElementID_14)

### Disk Image Tools

#### hdiutil

```sh
hdiutil create -fs APFS -size 1GB foo.sparseimage
```

#### diskutil apfs ...

```sh
diskutil apfs createContainer /dev/disk1s1

diskutil apfs addVolume disk1s1 APFS newAPFS
```

#### fsck_apfs

```sh
fsck_apfs
```

### Enhanced APIs

#### Foundation / FileManager

```c
func copyItem(atPath srcPath: String,
              toPath dstPath: String) throws
```

```c
func replaceItem(at originalItemURL: URL,
              withItemAt newItemURL: URL,
      backupItemName backupItemName: String?,
                    options options:
          FileManager.ItemReplacementOptions = [],
      resultingItemURL resultingURL:
          AutoreleasingUnsafeMutablePointer<NSURL?>?) throws
```

#### libcopyfile

```c
#include <copyfile.h>
 
int copyfile(const char *from,
             const char *to,
       copyfile_state_t state,
       copyfile_flags_t flags);
 
int fcopyfile(int from_fd,
             int to_fd,
             copyfile_state_t state,
             copyfile_flags_t flags);
// new flag bit: COPYFILE_CLONE
// equivalent to (COPYFILE_EXCL | COPYFILE_ACL | COPYFILE_STAT | COPYFILE_XATTR | COPYFILE_DATA)
```

#### Safe Save APIs

```c
#include <stdio.h>
 
int renamex_np(const char *, const char *, unsigned int)
 
int renameatx_np(int, const char *, int, const char *, unsigned int)
```

#### Cloning APIs

```c
#include <sys/attr.h>
#include <sys/clonefile.h>
 
int clonefileat(int, const char *, int, const char *, uint32_t);
int fclonefileat(int, int, const char *, uint32_t);
int clonefile(const char *, const char *, uint32_t);
```

#### Snapshot APIs

The Snapshot API is still under development. Contact [Apple Developer Relations](https://developer.apple.com/contact/) for more information.


## Volume Format Comparison

- [Volume Format Comparison](https://developer.apple.com/library/prerelease/content/documentation/FileManagement/Conceptual/APFS_Guide/VolumeFormatComparison/VolumeFormatComparison.html#//apple_ref/doc/uid/TP40016999-CH8-DontLinkElementID_21)

|                             | Mac OS Extended (HFS+) | Apple File System (APFS) |
|-----------------------------|------------------------|--------------------------|
| Number of allocation blocks | 232 (4 billion)        | 263 (9 quintillion)      |
| File IDs                    | 32-bit                 | 64-bit                   |
| Maximum file size           | 263 bytes              | 263 bytes                |
| Time stamp granularity      | 1 second               | 1 nanosecond             |
| Copy-on-write               |                        | ✔                        |
| Crash protected             |                        | ✔                        |
| Journaled                   |                        | ✔                        |
| File and directory clones   |                        | ✔                        |
| Snapshots                   |                        | ✔                        |
| Space sharing               |                        | ✔                        |
| Full disk encryption        | ✔                      | ✔                        |
| Sparse files                |                        | ✔                        |
| Fast directory sizing       |                        | ✔                        |
