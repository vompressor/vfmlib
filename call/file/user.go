package file

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// POSIX ONLY
func GetUser(fi os.FileInfo) (*user.User, error) {
	stat := fi.Sys().(*syscall.Stat_t)

	return user.LookupId(strconv.FormatInt(int64(stat.Uid), 10))
}

// POSIX ONLY
func GetGroup(fi os.FileInfo) (*user.Group, error) {
	stat := fi.Sys().(*syscall.Stat_t)
	return user.LookupGroupId(strconv.FormatInt(int64(stat.Gid), 10))
}

// POSIX ONLY
func GetUID(fi os.FileInfo) uint32 {
	stat := fi.Sys().(*syscall.Stat_t)
	return stat.Uid
}

// POSIX ONLY
func GetGID(fi os.FileInfo) uint32 {
	stat := fi.Sys().(*syscall.Stat_t)
	return stat.Gid
}
