//go:build linux || darwin || freebsd

package mmap

/*
#include <sys/mman.h>
#include <unistd.h>
#include <errno.h>

void* my_mmap(size_t length) {
    void* addr = mmap(NULL, length, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    if (addr == MAP_FAILED) {
        return NULL;
    }
    return addr;
}

int my_munmap(void* addr, size_t length) {
    if (munmap(addr, length) == -1) {
        return errno;
    }
    return 0;
}
*/
import "C"
import (
	"unsafe"
)

func mmapUnix(opts MmapOpts) (unsafe.Pointer, error) {
	addr := C.my_mmap(C.size_t(opts.Length))
	if addr == nil {
		return nil, ErrMmapFailed
	}
	return addr, nil
}

func munmapUnix(addr unsafe.Pointer, length int) error {
	err := C.my_munmap(addr, C.size_t(length))
	if err != 0 {
		return ErrMunmapFailed
	}
	return nil
}
