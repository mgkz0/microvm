package cmd

import (
	"github.com/mgkz0/microvm/internal/qemu"
	"github.com/mgkz0/microvm/internal/vm"
	"github.com/spf13/cobra"
)

var (
	memoryMiB int
	vcpus     int
	diskSize  int
	iso       string
	dir       string
	disk      string
)

var rootCmd = &cobra.Command{
	Use:   "microvm",
	Short: "Minimalistic VM manager",
}

var cmdNew = &cobra.Command{
	Use:   "new [vm-name]",
	Short: "Create a new VM",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		vmName := args[0]
		vmDir := dir + vmName + "/"
		vm := vm.NewVMConfig(vmName, vmDir, disk, iso, memoryMiB, vcpus, diskSize)
		qemu.NewQemuVM(cmd.Context(), *vm)
		return nil
	},
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

	cmdNew.Flags().IntVar(
		&diskSize,
		"disksize",
		20,
		"virtual disk size",
	)

	cmdNew.Flags().StringVar(
		&iso,
		"iso",
		"",
		"installer ISO",
	)

	cmdNew.Flags().StringVar(
		&dir,
		"dir",
		"./",
		"VM directory",
	)

	cmdNew.Flags().StringVar(
		&disk,
		"disk",
		"disk.qcow2",
		"Virtual disk name",
	)

	rootCmd.AddCommand(cmdNew)
	rootCmd.AddCommand(cmdStop)
	rootCmd.AddCommand(cmdRemove)
}

func Execute() error {
	return rootCmd.Execute()
}
