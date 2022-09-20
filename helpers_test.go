package vz

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type testVM struct {
	*VirtualMachine
	tempKernelFile    *os.File
	stateHandlerError func(err error)
}

func (vm *testVM) Close() error {
	_ = os.Remove(vm.tempKernelFile.Name())
	return vm.tempKernelFile.Close()
}

func newTestVM(t *testing.T) *testVM {
	// use empty file as dummy kernel as we don't expect the VM to successfully start in our tests
	tempKernelFile := testFile(t, "vz_vmlinuz_dummy", []byte{})
	bootloader, err := NewLinuxBootLoader(tempKernelFile.Name())
	require.NoError(t, err)
	config, err := NewVirtualMachineConfiguration(bootloader, 1, 64*1024*1024)
	require.NoError(t, err)
	//passing the config below to NewVirtualMachine reproduces https://github.com/Code-Hex/vz/issues/43
	//config := NewVirtualMachineConfiguration(&LinuxBootLoader{}, 1, 64*1024*1024)

	stateHandlerError := func(err error) {
		require.Error(t, err)
	}

	return &testVM{
		VirtualMachine:    NewVirtualMachine(config),
		tempKernelFile:    tempKernelFile,
		stateHandlerError: stateHandlerError,
	}
}

func testFile(t *testing.T, relPath string, content []byte) *os.File {
	filePath := filepath.Join(t.TempDir(), relPath)
	file, err := os.Create(filePath)
	require.NoError(t, err)
	if len(content) > 0 {
		_, err = file.Write([]byte("test"))
		require.NoError(t, err)
	}
	return file
}

func openUDPConn(t *testing.T) *os.File {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{Port: 8888})
	require.NoError(t, err)
	udpFile, err := udp.File()
	require.NoError(t, err)

	return udpFile
}
