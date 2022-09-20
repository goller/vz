//go:build !macos10
// +build !macos10

package vz

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootLoader(t *testing.T) {
	bootloader, err := NewLinuxBootLoader("dummy")
	assert.NoError(t, err)
	assert.NotNil(t, bootloader)
}

func TestConfiguration(t *testing.T) {
	config, err := NewVirtualMachineConfiguration(&LinuxBootLoader{}, 1, 64*1024*1024)
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestEntropy(t *testing.T) {
	config, err := NewVirtioEntropyDeviceConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestBalloon(t *testing.T) {
	config, err := NewVirtioTraditionalMemoryBalloonDeviceConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestNetwork(t *testing.T) {
	natAttachment, err := NewNATNetworkDeviceAttachment()
	assert.NoError(t, err)
	assert.NotNil(t, natAttachment)
	//unimplemented
	//NewBridgedNetworkDeviceAttachment(networkInterface BridgedNetwork)
	udp := openUDPConn(t)
	defer udp.Close()
	fileAttachment, err := NewFileHandleNetworkDeviceAttachment(udp)
	assert.NoError(t, err)
	assert.NotNil(t, fileAttachment)

	hwaddr, err := net.ParseMAC("52:54:00:70:2b:71")
	assert.NoError(t, err)
	mac, err := NewMACAddress(hwaddr)
	assert.NoError(t, err)
	assert.NotNil(t, mac)
	mac, err = NewRandomLocallyAdministeredMACAddress()
	assert.NoError(t, err)
	assert.NotNil(t, mac)
}

func TestSerial(t *testing.T) {
	inoutFile := testFile(t, "serial-inout", []byte{})
	defer inoutFile.Close()
	handleAttachment, err := NewFileHandleSerialPortAttachment(inoutFile, inoutFile)
	assert.NoError(t, err)
	assert.NotNil(t, handleAttachment)
	fileAttachment, err := NewFileSerialPortAttachment(inoutFile.Name(), true)
	assert.NoError(t, err)
	assert.NotNil(t, fileAttachment)
	console, err := NewVirtioConsoleDeviceSerialPortConfiguration(fileAttachment)
	assert.NoError(t, err)
	assert.NotNil(t, console)
}

func TestStorage(t *testing.T) {
	f := testFile(t, "storage", []byte{})
	defer f.Close()
	attachment, err := NewDiskImageStorageDeviceAttachment(f.Name(), true)
	assert.NoError(t, err)
	assert.NotNil(t, attachment)
	storage, err := NewVirtioBlockDeviceConfiguration(attachment)
	assert.NoError(t, err)
	assert.NotNil(t, storage)
}

func TestVM(t *testing.T) {
	vm := newTestVM(t)
	assert.NotNil(t, vm)
}
