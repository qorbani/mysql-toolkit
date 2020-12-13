package cmd

import (
	"fmt"
	"os"

	"github.com/qorbani/mysql-toolkit/drivers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var sql *drivers.MySQL

var err error

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mysql-toolkit",
	Short: "MySQL Toolkit",
	Long:  `MySQL Toolkit is a light-weight CLI to help you run queries via the command line.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.mysql-toolkit.yaml)")

	rootCmd.PersistentFlags().StringP("server", "s", "127.0.0.1", "MySQL server host name")
	rootCmd.PersistentFlags().StringP("user", "u", "root", "username for MySQL server")
	rootCmd.PersistentFlags().StringP("pass", "p", "password", "password for MySQL server")
	rootCmd.PersistentFlags().StringP("db", "d", "", "database name to select")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pass", rootCmd.PersistentFlags().Lookup("pass"))
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".mysql-toolkit")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/")
	}

	viper.SetEnvPrefix("MYSQL_TOOLKIT")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
