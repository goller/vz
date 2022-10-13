package vz

import (
	"os"
	"testing"
)

func TestVM(t *testing.T) {
	f, err := os.CreateTemp("", "vmlinuz")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	bootloader, err := NewLinuxBootLoader(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	config, err := NewVirtualMachineConfiguration(bootloader, 2, 64*1024*1024)
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewVirtualMachine(config)
	if err != nil {
		t.Fatal(err)
	}
	m.Start(func(err error) {
		t.Error(err)
	})
}
