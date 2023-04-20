
package vz

/*
#cgo darwin CFLAGS: -mmacosx-version-min=12 -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization -framework Cocoa
# include "private_12.h"
*/
import "C"
import (
	"github.com/Code-Hex/vz/v3/internal/objc"
)

// UART16550SerialPortConfiguration represents Virtio Console Serial Port Device.
//
// The device creates a console which enables communication between the host and the guest through the Virtio interface.
// The device sets up a single port on the Virtio console device.
// see: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledeviceserialportconfiguration?language=objc
type UART16550SerialPortConfiguration struct {
	*pointer
}

// New16550SerialPortConfiguration creates a new New16550SerialPortConfiguration.
//
// This is only supported on macOS 11 and newer, error will
// be returned on older versions.
func New16550SerialPortConfiguration(attachment SerialPortAttachment) (*UART16550SerialPortConfiguration, error) {
	if err := macOSAvailable(12); err != nil {
		return nil, err
	}

	config := &UART16550SerialPortConfiguration{
		pointer: objc.NewPointer(
			C.newVZ16550SerialPortConfiguration(
				objc.Ptr(attachment),
			),
		),
	}
	objc.SetFinalizer(config, func(self *UART16550SerialPortConfiguration) {
		objc.Release(self)
	})
	return config, nil
}
