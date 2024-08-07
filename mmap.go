package mmap

import (
	"errors"
	"runtime"
	"unsafe"
)

var (
	ErrUnsupportedPlatform = errors.New("unsupported platform")
	ErrMmapFailed          = errors.New("mmap failed")
	ErrMunmapFailed        = errors.New("munmap failed")
)

// MmapOpts represents options for memory mapping
type MmapOpts struct {
	Length int
}

// Mmap maps a file into memory
func Mmap(opts MmapOpts) (unsafe.Pointer, error) {
	switch runtime.GOOS {
	case "windows":
		return mmapWindows(opts)
	case "linux", "darwin", "freebsd":
		return mmapUnix(opts)
	default:
		return nil, ErrUnsupportedPlatform
	}
}

// Munmap unmaps a memory-mapped file
func Munmap(addr unsafe.Pointer, length int) error {
	switch runtime.GOOS {
	case "windows":
		return munmapWindows(addr, length)
	case "linux", "darwin", "freebsd":
		return munmapUnix(addr, length)
	default:
		return ErrUnsupportedPlatform
	}
}
