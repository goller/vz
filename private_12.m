//
//  private_12.m
//
//  Created by Christophe Fergeau.
//

#import "private_12.h"

/*
https://github.com/Code-Hex/vz/wiki/Private-API-on-macOS-12

0x0020f638bf0 VZVirtioConsoleDeviceSerialPortConfiguration : VZSerialPortConfiguration  // class methods
  0x001e661d500 +[VZVirtioConsoleDeviceSerialPortConfiguration serialPortType]

  // instance methods
  0x001e661d4c8 -[VZVirtioConsoleDeviceSerialPortConfiguration init]
  0x001e661d364 -[VZVirtioConsoleDeviceSerialPortConfiguration encodeWithEncoder:]
  0x001e661cf74 -[VZVirtioConsoleDeviceSerialPortConfiguration _serialPort]


0x002169d83b0 _VZ16550SerialPortConfiguration : VZSerialPortConfiguration  // class methods
  0x001e747ffa8 +[_VZ16550SerialPortConfiguration serialPortType]

  // instance methods
  0x001e747ff70 -[_VZ16550SerialPortConfiguration init]
  0x001e747fe34 -[_VZ16550SerialPortConfiguration encodeWithEncoder:]
  0x001e747fa48 -[_VZ16550SerialPortConfiguration _serialPort]
*/

/*!
 @abstract Create a new Virtio Console Serial Port Device configuration
 @param attachment Base class for a serial port attachment.
 @discussion
    The device creates a console which enables communication between the host and the guest through the Virtio interface.

    The device sets up a single port on the Virtio console device.
 */
void *newVZ16550SerialPortConfiguration(void *attachment)
{
    if (@available(macOS 12, *)) {
        _VZ16550SerialPortConfiguration *config = [[_VZ16550SerialPortConfiguration alloc] init];
        [config setAttachment:(VZSerialPortAttachment *)attachment];
        return config;
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
