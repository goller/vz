//go:build macos11
// +build macos11

package vz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSharing(t *testing.T) {
	_, err := NewVirtioFileSystemDeviceConfiguration("vz-test")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewSharedDirectory("/", false)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}

func TestPlatform(t *testing.T) {
	_, err := NewGenericPlatformConfiguration()
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
