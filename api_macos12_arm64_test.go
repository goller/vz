//go:build !macos11
// +build !macos11

package vz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacOS(t *testing.T) {
	dummyFile := testFile(t, "dummy", []byte("test"))
	defer dummyFile.Close()

	bootloader, err := NewMacOSBootLoader()
	assert.NoError(t, err)
	assert.NotNil(t, bootloader)

	deviceConfig, err := NewMacGraphicsDeviceConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, deviceConfig)

	displayConfig, err := NewMacGraphicsDisplayConfiguration(1024, 768, 75)
	assert.NoError(t, err)
	assert.NotNil(t, displayConfig)

	platformConfig, err := NewMacPlatformConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, platformConfig)

	hardwareModel, err := NewMacHardwareModelWithDataPath(dummyFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, hardwareModel)

	hardwareModel, err = NewMacHardwareModelWithData([]byte("test"))
	assert.NoError(t, err)
	assert.NotNil(t, hardwareModel)

	identifier, err := NewMacMachineIdentifierWithDataPath(dummyFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, identifier)

	identifier, err = NewMacMachineIdentifierWithData([]byte("test"))
	assert.NoError(t, err)
	assert.NotNil(t, identifier)

	auxiliaryStorage, err := NewMacAuxiliaryStorage(dummyFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, auxiliaryStorage)

	/*
		_, err = FetchLatestSupportedMacOSRestoreImage(context.Background(), dummyFilePath)
		assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

		restoreImage, err := LoadMacOSRestoreImageFromPath(dummyFile.Name())
		assert.NoError(t, err)
		assert.NotNil(t, restoreImage)

		vm := newTestVM(t)
		installer, err := NewMacOSInstaller(vm.VirtualMachine, dummyFile.Name())
		assert.NoError(t, err)
		assert.NotNil(t, installer)
	*/
}
