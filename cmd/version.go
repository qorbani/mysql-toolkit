package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version will inject by Makefile
	Version = "0.0.0"
	// Build will inject by Makefile
	Build = "dev"
	// BuildTime will inject by Makefile
	BuildTime = "Now"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Toolkit version detail",
	Long:  `Toolkit Version and Build information`,
	Run: func(cmd *cobra.Command, args []string) {
		osArch := runtime.GOOS + "/" + runtime.GOARCH
		VersionString := fmt.Sprintf("MySQL Toolkit v%s (%s) %v BuildTime: %s", Version, Build, osArch, BuildTime)
		fmt.Println(VersionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
