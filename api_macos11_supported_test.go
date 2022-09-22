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
