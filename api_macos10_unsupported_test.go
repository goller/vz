//go:build macos10
// +build macos10

package vz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootLoader(t *testing.T) {
	_, err := NewLinuxBootLoader("dummy")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
