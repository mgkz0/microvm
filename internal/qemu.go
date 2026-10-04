package qemu

import (
	"context"
	"fmt"
	"os/exec"
)

func StartQemu(ctx context.Context, disk, iso, qmpSocket, pidFile string) {
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
