package vz

/*
#cgo darwin CFLAGS: -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization
# include "virtualization.h"
*/
import "C"
import (
	"io"
	"net"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rs/xid"
)

type NewConnectionFn func(conn *VirtioSocketConnection, err error)
type ShouldAcceptConnectionFn func(conn *VirtioSocketConnection) bool

// SocketDeviceConfiguration for a socket device configuration.
type SocketDeviceConfiguration interface {
	NSObject

	socketDeviceConfiguration()
}

type baseSocketDeviceConfiguration struct{}

func (*baseSocketDeviceConfiguration) socketDeviceConfiguration() {}

var _ SocketDeviceConfiguration = (*VirtioSocketDeviceConfiguration)(nil)

// VirtioSocketDeviceConfiguration is a configuration of the Virtio socket device.
//
// This configuration creates a Virtio socket device for the guest which communicates with the host through the Virtio interface.
// Only one Virtio socket device can be used per virtual machine.
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketdeviceconfiguration?language=objc
type VirtioSocketDeviceConfiguration struct {
	pointer

	*baseSocketDeviceConfiguration
}

// NewVirtioSocketDeviceConfiguration creates a new VirtioSocketDeviceConfiguration.
func NewVirtioSocketDeviceConfiguration() *VirtioSocketDeviceConfiguration {
	config := &VirtioSocketDeviceConfiguration{
		pointer: pointer{
			ptr: C.newVZVirtioSocketDeviceConfiguration(),
		},
	}
	runtime.SetFinalizer(config, func(self *VirtioSocketDeviceConfiguration) {
		self.Release()
	})
	return config
}

// VirtioSocketDevice a device that manages port-based connections between the guest system and the host computer.
//
// Don’t create a VirtioSocketDevice struct directly. Instead, when you request a socket device in your configuration,
// the virtual machine creates it and you can get it via SocketDevices method.
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketdevice?language=objc
type VirtioSocketDevice struct {
	id string

	dispatchQueue unsafe.Pointer
	pointer
}

var connectionHandlers = map[string]NewConnectionFn{}

// most likely needs to be protected by a mutex. Or maybe this belongs in the VM or device struct?
var objcToGoListeners = map[unsafe.Pointer]*VirtioSocketListener{}

func newVirtioSocketDevice(ptr, dispatchQueue unsafe.Pointer) *VirtioSocketDevice {
	id := xid.New().String()
	socketDevice := &VirtioSocketDevice{
		id:            id,
		dispatchQueue: dispatchQueue,
		pointer: pointer{
			ptr: ptr,
		},
	}
	connectionHandlers[id] = func(*VirtioSocketConnection, error) {}

	runtime.SetFinalizer(socketDevice, func(self *VirtioSocketDevice) {
		self.Release()
	})
	return socketDevice
}

// SetSocketListenerForPort configures an object to monitor the specified port for new connections.
//
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketdevice/3656679-setsocketlistener?language=objc
func (v *VirtioSocketDevice) SetSocketListenerForPort(listener *VirtioSocketListener, port uint32, fn ShouldAcceptConnectionFn) {
	objcToGoListeners[listener.Ptr()] = listener
	listener.acceptHandlers[port] = fn
	C.VZVirtioSocketDevice_setSocketListenerForPort(v.Ptr(), v.dispatchQueue, listener.Ptr(), C.uint32_t(port))
}

// RemoveSocketListenerForPort removes the listener object from the specfied port.
//
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketdevice/3656678-removesocketlistenerforport?language=objc
func (v *VirtioSocketDevice) RemoveSocketListenerForPort(listener *VirtioSocketListener, port uint32) {
	delete(objcToGoListeners, listener.Ptr())
	delete(listener.acceptHandlers, port)
	C.VZVirtioSocketDevice_removeSocketListenerForPort(v.Ptr(), v.dispatchQueue, C.uint32_t(port))
}

//export connectionHandler
func connectionHandler(connPtr, errPtr unsafe.Pointer, cid *C.char) {
	id := (*char)(cid).String()
	// see: startHandler
	conn := newVirtioSocketConnection(connPtr)
	if err := newNSError(errPtr); err != nil {
		connectionHandlers[id](conn, err)
	} else {
		connectionHandlers[id](conn, nil)
	}
}

func newConnFromListener(connPtr unsafe.Pointer) (*VirtioSocketConnection, error) {
	conn := newVirtioSocketConnection(connPtr)
	// not fully clear what the lifetime/ownership of the filedescriptor is in the .m code
	// this might be a bad workaround for my lack of understanding
	// after calling this, fileDescriptor needs to be closed in Close(),
	// though golang will garbage collect open file descriptors if needed
	newFd, err := syscall.Dup(int(conn.FileDescriptor()))
	if err != nil {
		return nil, err
	}
	conn.fileDescriptor = uintptr(newFd)

	return conn, nil
}

//export shouldAcceptNewConnectionHandler
func shouldAcceptNewConnectionHandler(connPtr unsafe.Pointer, listenerPtr unsafe.Pointer, devicePtr unsafe.Pointer) C.int {
	listener, hasListener := objcToGoListeners[listenerPtr]
	if !hasListener {
		return 0
	}
	conn, err := newConnFromListener(connPtr)
	if err != nil {
		return 0
	}
	handler, hasHandler := listener.acceptHandlers[conn.DestinationPort()]
	if !hasHandler {
		return 0
	}
	boolResult := handler(conn)
	if boolResult {
		return 1
	}
	return 0
}

// Initiates a connection to the specified port of the guest operating system.
//
// This method initiates the connection asynchronously, and executes the completion handler when the results are available.
// If the guest operating system doesn’t listen for connections to the specifed port, this method does nothing.
//
// For a successful connection, this method sets the sourcePort property of the resulting VZVirtioSocketConnection object to a random port number.
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketdevice/3656677-connecttoport?language=objc
func (v *VirtioSocketDevice) ConnectToPort(port uint32, fn NewConnectionFn) {
	connectionHandlers[v.id] = fn
	cid := charWithGoString(v.id)
	defer cid.Free()
	C.VZVirtioSocketDevice_connectToPort(v.Ptr(), v.dispatchQueue, C.uint32_t(port), cid.CString())
}

// VirtioSocketListener a struct that listens for port-based connection requests from the guest operating system.
//
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketlistener?language=objc
type VirtioSocketListener struct {
	acceptHandlers map[uint32]ShouldAcceptConnectionFn
	pointer
}

// NewVirtioSocketListener creates a new VirtioSocketListener.
func NewVirtioSocketListener() *VirtioSocketListener {
	listener := &VirtioSocketListener{
		acceptHandlers: make(map[uint32]ShouldAcceptConnectionFn),
		pointer: pointer{
			ptr: C.newVZVirtioSocketListener(),
		},
	}
	runtime.SetFinalizer(listener, func(self *VirtioSocketListener) {
		self.Release()
	})
	return listener
}

// VirtioSocketConnection is a port-based connection between the guest operating system and the host computer.
//
// You don’t create connection objects directly. When the guest operating system initiates a connection, the virtual machine creates
// the connection object and passes it to the appropriate VirtioSocketListener struct, which forwards the object to its delegate.
//
// see: https://developer.apple.com/documentation/virtualization/vzvirtiosocketconnection?language=objc
type VirtioSocketConnection struct {
	sourcePort      uint32
	destinationPort uint32
	fileDescriptor  uintptr

	pointer
}

// TODO(codehex): should implement net.Conn?
var _ io.Closer = (*VirtioSocketConnection)(nil)

func newVirtioSocketConnection(ptr unsafe.Pointer) *VirtioSocketConnection {
	vzVirtioSocketConnection := C.convertVZVirtioSocketConnection2Flat(ptr)
	conn := &VirtioSocketConnection{
		sourcePort:      (uint32)(vzVirtioSocketConnection.sourcePort),
		destinationPort: (uint32)(vzVirtioSocketConnection.destinationPort),
		fileDescriptor:  (uintptr)(vzVirtioSocketConnection.fileDescriptor),
		pointer: pointer{
			ptr: ptr,
		},
	}
	runtime.SetFinalizer(conn, func(self *VirtioSocketConnection) {
		self.Release()
	})
	return conn
}

// DestinationPort returns the destination port number of the connection.
func (v *VirtioSocketConnection) DestinationPort() uint32 {
	return v.destinationPort
}

// SourcePort returns the source port number of the connection.
func (v *VirtioSocketConnection) SourcePort() uint32 {
	return v.sourcePort
}

// FileDescriptor returns the file descriptor associated with the socket.
//
// Data is sent by writing to the file descriptor.
// Data is received by reading from the file descriptor.
// A file descriptor of -1 indicates a closed connection.
func (v *VirtioSocketConnection) FileDescriptor() uintptr {
	return v.fileDescriptor
}

func (v *VirtioSocketConnection) Close() error {
	C.VZVirtioSocketConnection_close(v.Ptr())
	return nil
}

// VirtioSocketListener is a low-level type with a single instance, it will handle VM->host connection attempts for all ports
// Listener is a high-level type implementing net.Listener
type Listener struct {
	port          uint32
	incomingConns chan *VirtioSocketConnection
}

func (v *VirtioSocketDevice) Listen(port uint32) *Listener {
	// for a given device, we should only use one instance of *VirtioSocketListener
	listener := &Listener{
		port:          port,
		incomingConns: make(chan *VirtioSocketConnection, 1),
	}
	shouldAcceptConn := func(conn *VirtioSocketConnection) bool {
		listener.incomingConns <- conn
		return true
	}

	virtioSocketListener := NewVirtioSocketListener()
	v.SetSocketListenerForPort(virtioSocketListener, port, shouldAcceptConn)
	return listener
}

func (l *Listener) Accept() (*VirtioSocketConnection, error) {
	conn := <-l.incomingConns
	return conn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *Listener) Close() error {
	return nil
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return nil
}
