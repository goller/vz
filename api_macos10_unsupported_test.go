//go:build macos10
// +build macos10

package vz

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fetchMajorVersion = func() {
		majorVersion = 10
	}
	os.Exit(m.Run())
}

func TestBootLoader(t *testing.T) {
	_, err := NewLinuxBootLoader("dummy")
	assert.ErrorIs(t, err, ErrUnsupportedOSVersion)
}
