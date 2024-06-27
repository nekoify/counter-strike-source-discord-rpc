package utils

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func GetValueFromAddress(hProcess windows.Handle, baseAddress uintptr) string {
	const bufferSize = 1024
	buffer := make([]byte, bufferSize)
	var bytesRead uintptr

	for i := 0; i < bufferSize; i++ {
		err := windows.ReadProcessMemory(hProcess, baseAddress+uintptr(i), &buffer[i], 1, &bytesRead)
		if err != nil {
			fmt.Println(err)
		}
		if buffer[i] == 0 {
			return string(buffer[:i])
		}
	}
	fmt.Println("string too long to fit in buffer")
	return ""
}
