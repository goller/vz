//go:build !macos10
// +build !macos10

package vz

import (
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
