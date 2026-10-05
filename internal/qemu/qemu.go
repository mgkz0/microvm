package qemu

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func newQemuDisk(ctx context.Context, vmName string, diskSize int) string {
	dir_path := "./" + vmName
	_, err := os.Stat(dir_path)
	if err != nil {
		os.Mkdir(dir_path, 0755)
	}

	full_path := dir_path + "/disk.qcow2"
	args := []string{
		"create",
		"-f", "qcow2",
		full_path,
		strconv.Itoa(diskSize) + "GB",
	}

	cmd := exec.CommandContext(ctx, "qemu-img", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
	return full_path
}

func NewQemuVM(ctx context.Context, iso, qmpSocket, pidFile, vmName string, diskSize int) {
	disk := newQemuDisk(ctx, vmName, diskSize)
	print("=====================================================")
	args := []string{
		"-enable-kvm",
		"-machine", "q35",
		"-m", "2048",
		"-smp", "2",

		"-drive",
		"file=" + disk + ",if=virtio,format=qcow2",

		"-cdrom", iso,

		"-nic", "user,model=virtio",

		"-display", "none",

		"-qmp",
		"unix:" + qmpSocket + ",server=on,wait=off",

		"-pidfile", pidFile,
	}

	cmd := exec.Command("qemu-system-x86_64", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}
