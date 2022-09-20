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
