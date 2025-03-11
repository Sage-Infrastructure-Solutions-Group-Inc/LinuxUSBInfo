//go:build (darwin && cgo) || linux
// +build darwin,cgo linux

package main

import (
	"encoding/json"
	"fmt"
)

// GetUSBDeviceInfoByUDev retrieves detailed USB device information using libudev.
func getUSBDeviceInfoByUDev() {
	u := udev.Udev{}
	e := u.NewEnumerate()
	e.AddMatchSubsystem("usb")
	devices, _ := e.Devices()
	device_list := []map[string]string{}
	for i := range devices {
		device := devices[i]
		device_list = append(device_list, device.Properties())
	}
	data, _ := json.Marshal(device_list)
	data_str := string(data)
	fmt.Print(data_str)
}

func main() {
	getUSBDeviceInfoByUDev()
}
