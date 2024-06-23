package main

import (
	"fmt"
	"main/discordRPC"
	"main/utils"
	"time"

	"github.com/hugolgst/rich-go/client"
	"golang.org/x/sys/windows"
)

func main() {
	startTime := time.Now()
	const processName = "hl2.exe"
	const engineDll = "engine.dll"
	const clientDll = "client.dll"

	var mapName = "Placeholder"
	var serverName = "Placeholder"

	const mapNameOffset = 0x4779E8
	const serverNameOffset = 0x50A214

	pid, err := utils.GetProcessID(processName)
	if err != nil {
		fmt.Println("Error finding process:", err)
		return
	}

	hProcess, err := windows.OpenProcess(windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, false, pid)
	if err != nil {
		fmt.Println("Error opening process:", err)
		return
	}
	defer windows.CloseHandle(hProcess)

	engineAddr, err := utils.GetModuleBaseAddress(pid, engineDll)
	if err != nil {
		panic(err)
	}

	clientAddr, err := utils.GetModuleBaseAddress(pid, clientDll)
	if err != nil {
		panic(err)
	}

	mapNameAddress := engineAddr + mapNameOffset
	serverNameAddress := clientAddr + serverNameOffset

	errr := client.Login("1253060133940891789")
	if errr != nil {
		panic(errr)
	}

	for true {
		var shouldUpdate = false
		pid, err := utils.GetProcessID(processName)
		_ = pid
		if err != nil {
			panic(err)
		}
		mapVal := utils.GetValueFromAddress(hProcess, uintptr(mapNameAddress))
		if mapVal != mapName && mapVal != "" {
			mapName = mapVal
			shouldUpdate = true
		}
		serverVal := utils.GetValueFromAddress(hProcess, uintptr(serverNameAddress))
		if serverVal != serverName && serverVal != "" {
			serverName = serverVal
			shouldUpdate = true
		}
		if shouldUpdate {
			discordRPC.Update(mapVal, serverVal, startTime)
		}
		time.Sleep(time.Second)
	}
}
