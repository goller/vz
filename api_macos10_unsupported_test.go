//go:build macos10
// +build macos10

package vz

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootLoader(t *testing.T) {
	_, err := NewLinuxBootLoader("dummy")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestConfiguration(t *testing.T) {
	_, err := NewVirtualMachineConfiguration(&LinuxBootLoader{}, 1, 64*1024*1024)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestEntropy(t *testing.T) {
	_, err := NewVirtioEntropyDeviceConfiguration()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestBalloon(t *testing.T) {
	_, err := NewVirtioTraditionalMemoryBalloonDeviceConfiguration()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestNetwork(t *testing.T) {
	_, err := NewNATNetworkDeviceAttachment()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
	//unimplemented
	//NewBridgedNetworkDeviceAttachment(networkInterface BridgedNetwork)
	udp := openUDPConn(t)
	defer udp.Close()
	_, err := NewFileHandleNetworkDeviceAttachment(udp)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	hwaddr, err := net.ParseMAC("52:54:00:70:2b:71")
	assert.NoError(t, err)
	_, err = NewMACAddress(hwaddr)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
	_, err = NewRandomLocallyAdministeredMACAddress()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestSerial(t *testing.T) {
	inoutFile := testFile(t, "serial-inout", []byte{})
	defer inoutFile.Close()
	_, err := NewFileHandleSerialPortAttachment(inoutFile, inoutFile)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
	_, err = NewFileSerialPortAttachment(inoutFile.Name(), true)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
	_, err = NewVirtioConsoleDeviceSerialPortConfiguration(&FileSerialPortAttachment{})
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
