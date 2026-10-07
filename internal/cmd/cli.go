package cmd

import (
	"github.com/mgkz0/microvm/internal/qemu"
	"github.com/mgkz0/microvm/internal/vms"
	"github.com/spf13/cobra"
)

var (
	memoryMiB int
	vcpus     int
	diskSize  string
	isoPath   string
)

var rootCmd = &cobra.Command{
	Use:   "microvm",
	Short: "Minimalistic VM manager",
}

var cmdNew = &cobra.Command{
	Use:   "new [vm-name]",
	Short: "Create a new VM",
	RunE: func(cmd *cobra.Command, args []string) error {
		vmName := args[0]
		vm := vm.NewVMConfig(vmName, )
		qemu.NewQemuVM(cmd.Context(), isoPath, qm)	
	}
}

var cmdRemove = &cobra.Command{
	Use:   "remove [vm-name]",
	Short: "Remove created VM",
}

var cmdStop = &cobra.Command{
	Use:   "stop [vm-name]",
	Short: "Stop running VM",
}

func init() {
	cmdNew.Flags().IntVarP(
		&memoryMiB,
		"memory",
		"m",
		2048,
		"RAM in MiB",
	)

	cmdNew.Flags().IntVarP(
		&vcpus,
		"cpus",
		"c",
		2,
		"number of virtual CPUs",
	)

	cmdNew.Flags().StringVar(
		&diskSize,
		"disk",
		"20G",
		"virtual disk size",
	)

	cmdNew.Flags().StringVar(
		&isoPath,
		"iso",
		"",
		"installer ISO",
	)

	rootCmd.AddCommand(cmdNew)
	rootCmd.AddCommand(cmdStop)
	rootCmd.AddCommand(cmdRemove)
}
