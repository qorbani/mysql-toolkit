package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/qorbani/mysql-toolkit/drivers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug internal information",
	Long:  `Debug CLI command to expose internal settings`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config File:", viper.ConfigFileUsed())
		if jsonData, err := json.Marshal(viper.AllSettings()); err != nil {
			fmt.Println("Unable to read settings.")
		} else {
			fmt.Println("Configuration:", string(jsonData))
		}

		if _, err := drivers.NewMySQL(); err != nil {
			fmt.Println("Error connection to MySQL:", err.Error())
		} else {
			fmt.Println("MySQL connection stablished.")
		}
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
