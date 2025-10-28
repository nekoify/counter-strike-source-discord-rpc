package utils

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetValueFromAddress(hProcess windows.Handle, baseAddress uintptr) (string, error) {
	const bufferSize = 1024
	buffer := make([]byte, bufferSize)
	var bytesRead uintptr

	err := windows.ReadProcessMemory(hProcess, baseAddress, &buffer[0], uintptr(bufferSize), &bytesRead)
	if err != nil {
		return "", fmt.Errorf("reading memory failed: %w", err)
	}

	n := bytes.IndexByte(buffer[:bytesRead], 0)
	if n < 0 {
		return "", fmt.Errorf("unterminated string or too long")
	}

	return string(buffer[:n]), nil
}

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

func GetProcessID(processName string) (uint32, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return 0, err
	}
	defer windows.CloseHandle(snapshot)

	var pe32 windows.ProcessEntry32
	pe32.Size = uint32(unsafe.Sizeof(pe32))
	err = windows.Process32First(snapshot, &pe32)
	if err != nil {
		return 0, err
	}

	for {
		if syscall.UTF16ToString(pe32.ExeFile[:]) == processName {
			return pe32.ProcessID, nil
		}
		err = windows.Process32Next(snapshot, &pe32)
		if err != nil {
			break
		}
	}

	return 0, fmt.Errorf("process not found")
}
