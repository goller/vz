//go:build !macos10 && macos11
// +build !macos10
// +build macos11

package vz

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacOS(t *testing.T) {
	dummyFile := testFile(t, "dummy", []byte("test"))
	defer dummyFile.Close()

	_, err := NewMacOSBootLoader()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacGraphicsDeviceConfiguration()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacGraphicsDisplayConfiguration(1024, 768, 75)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacPlatformConfiguration()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacHardwareModelWithDataPath(dummyFile.Name())
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacHardwareModelWithData([]byte{})
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacMachineIdentifierWithDataPath(dummyFile.Name())
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacMachineIdentifierWithData([]byte{})
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewMacAuxiliaryStorage(dummyFile.Name())
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = FetchLatestSupportedMacOSRestoreImage(context.Background(), dummyFile.Name())
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = LoadMacOSRestoreImageFromPath(dummyFile.Name())
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	vm := newTestVM(t)
	_, err = NewMacOSInstaller(vm.VirtualMachine, "")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
