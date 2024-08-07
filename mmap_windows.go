//go:build windows

package mmap

/*
#include <windows.h>

void* my_mmap(size_t length) {
    HANDLE hMapFile;
    LPVOID lpMapAddress;

    hMapFile = CreateFileMapping(INVALID_HANDLE_VALUE, NULL, PAGE_READWRITE, 0, (DWORD)length, NULL);
    if (hMapFile == NULL) {
        return NULL;
    }

    lpMapAddress = MapViewOfFile(hMapFile, FILE_MAP_ALL_ACCESS, 0, 0, length);
    CloseHandle(hMapFile);

    if (lpMapAddress == NULL) {
        return NULL;
    }

    return lpMapAddress;
}

int my_munmap(void* addr, size_t length) {
    if (!UnmapViewOfFile(addr)) {
        return GetLastError();
    }
    return 0;
}
*/
import "C"
import "unsafe"

func mmapWindows(opts MmapOpts) (unsafe.Pointer, error) {
	addr := C.my_mmap(C.size_t(opts.Length))
	if addr == nil {
		return nil, ErrMmapFailed
	}
	return addr, nil
}

func munmapWindows(addr unsafe.Pointer, length int) error {
	err := C.my_munmap(addr, C.size_t(length))
	if err != 0 {
		return ErrMunmapFailed
	}
	return nil
}
