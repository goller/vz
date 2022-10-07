package vz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNonExistingFileSerialPortAttachment(t *testing.T) {
	_, err := NewFileSerialPortAttachment("/non/existing/path", false)
	require.Error(t, err)
}

func TestNonExistingImageStorageDeviceAttachment(t *testing.T) {
	_, err := NewDiskImageStorageDeviceAttachment("/non/existing/path", true)
	require.Error(t, err)
}
