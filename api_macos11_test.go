//go:build macos11
// +build macos11

package vz

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fetchMajorVersion = func() {
		majorVersion = 11
	}
	os.Exit(m.Run())
}

func TestFileSharing(t *testing.T) {
	_, err := NewVirtioFileSystemDeviceConfiguration("vz-test")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)

	_, err = NewSharedDirectory("/", false)
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
