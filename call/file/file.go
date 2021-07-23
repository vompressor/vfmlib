package file

import (
	"os"
	"syscall"
	"time"
)

func GetDetailSize(fi os.FileInfo) (size, blockSize, blocks int64) {
	stat := fi.Sys().(*syscall.Stat_t)
	size = stat.Size
	blockSize = stat.Blksize
	blocks = stat.Blocks
	return
}

// GetFileTimes return times of file
// a - time of last access
// m - time of last modification
// c - time of last status change
// Windows, POSIX
func GetFileTimes(fi os.FileInfo) (a, m, c time.Time) {
	m = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	a = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	c = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}

// GetDileINode return inode of file
// Windows, POSIX
func GetFileINode(fi os.FileInfo) uint64 {
	return fi.Sys().(*syscall.Stat_t).Ino

}

// GetHardLink return number of hard links
// Windows, POSIX
func GetHardLinks(fi os.FileInfo) uint64 {
	return fi.Sys().(*syscall.Stat_t).Nlink
}

// GetFileDev return inode of device containing file
// POSIX Only
func GetFileDev(fi os.FileInfo) uint64 {
	return fi.Sys().(*syscall.Stat_t).Dev
}

// GetFileRDev return inode of device containing file
// POSIX Only
func GetFileRDev(fi os.FileInfo) uint64 {
	return fi.Sys().(*syscall.Stat_t).Rdev
}
