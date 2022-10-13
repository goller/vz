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
	ok, err := config.Validate()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("failed to validate config")
	}
	m, err := NewVirtualMachine(config)
	if err != nil {
		t.Fatal(err)
	}
	canStart := m.CanStart()
	if !canStart {
		t.Fatal("cannot start")
	}
	m.Start(func(err error) {
		t.Error(err)
	})
}
