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

func NewVMConfig(name string, dir *string, disk *string, iso string, memory, vcpus, diskSize int) *VMConfig {
	baseDir := "./" + name + "/"
	baseDiskName := "disk.qcow2"
	QMPSocket := "qmp.sock"
	PIDFile := "qemu.pid"
	if dir != nil {
		baseDir = *dir
	}
	if disk != nil {
		baseDiskName = *disk
	}
	return &VMConfig{
		Name:      name,
		MemoryMiB: memory,
		VCPUs:     vcpus,
		Dir:       baseDir,
		Disk:      baseDiskName,
		DiskSize:  diskSize,
		QMPSocket: QMPSocket,
		PIDFile:   PIDFile,
		ISO:       iso,
	}
}
