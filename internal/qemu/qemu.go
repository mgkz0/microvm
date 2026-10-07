package qemu

import (
	"context"
	"fmt"
	"github.com/mgkz0/microvm/internal/vm"
	"os"
	"os/exec"
	"strconv"
)

func newQemuDisk(ctx context.Context, diskPath string, diskSize int) {
	args := []string{
		"create",
		"-f", "qcow2",
		diskPath,
		strconv.Itoa(diskSize) + "G",
	}

	cmd := exec.CommandContext(ctx, "qemu-img", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}

func NewQemuVM(ctx context.Context, c vm.VMConfig) {
	_, err := os.Stat(c.Dir)
	if err != nil {
		os.Mkdir(c.Dir, 0755)
	}
	diskPath := c.Dir + c.Disk
	newQemuDisk(ctx, diskPath, c.DiskSize)
	print("=====================================================")
	args := []string{
		"-enable-kvm",
		"-machine", "q35",
		"-m", "2048",
		"-smp", "2",

		"-drive",
		"file=" + diskPath + ",if=virtio,format=qcow2",

		"-cdrom", c.ISO,

		"-nic", "user,model=virtio",

		"-display", "none",

		"-qmp",
		"unix:" + c.QMPSocket + ",server=on,wait=off",

		"-pidfile", c.PIDFile,
	}

	cmd := exec.Command("qemu-system-x86_64", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}
