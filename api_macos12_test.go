//go:build !macos11
// +build !macos11

package vz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSharing(t *testing.T) {
	config, err := NewVirtioFileSystemDeviceConfiguration("vz-test")
	assert.NoError(t, err)
	assert.NotNil(t, config)

	share, err := NewSharedDirectory("/", false)
	assert.NoError(t, err)
	assert.NotNil(t, share)
}

func TestPlatform(t *testing.T) {
	config, err := NewGenericPlatformConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestKeyboardAndPointingDevice(t *testing.T) {
	keyboardConfig, err := NewUSBKeyboardConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, keyboardConfig)

	pointingDeviceConfig, err := NewUSBScreenCoordinatePointingDeviceConfiguration()
	assert.NoError(t, err)
	assert.NotNil(t, pointingDeviceConfig)

}
