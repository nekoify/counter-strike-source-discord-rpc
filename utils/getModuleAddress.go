package utils

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetModuleBaseAddress(pid uint32, moduleName string) (uintptr, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE|windows.TH32CS_SNAPMODULE32, pid)
	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(snapshot)

	var me32 windows.ModuleEntry32
	me32.Size = uint32(unsafe.Sizeof(me32))
	err = windows.Module32First(snapshot, &me32)
	if err != nil {
		return 0, err
	}

	for {
		if syscall.UTF16ToString(me32.Module[:]) == moduleName {
			return uintptr(me32.ModBaseAddr), nil
		}
		err = windows.Module32Next(snapshot, &me32)
		if err != nil {
			break
		}
	}

	return 0, fmt.Errorf("module not found")
}
