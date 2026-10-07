package vm

type VMConfig struct {
	Name string

	MemoryMiB int
	VCPUs     int

	Dir       string
	Disk      string
	DiskSize  int
	QMPSocket string
	PIDFile   string

	ISO string
}

func NewVMConfig(name, dir, disk, iso string, memory, vcpus, diskSize int) *VMConfig {
	return &VMConfig{
		Name:      name,
		MemoryMiB: memory,
		VCPUs:     vcpus,
		Dir:       dir,
		Disk:      disk,
		DiskSize:  diskSize,
		QMPSocket: "qmp.sock",
		PIDFile:   "qemu.pid",
		ISO:       iso,
	}
}
