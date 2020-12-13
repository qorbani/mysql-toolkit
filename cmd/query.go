package cmd

import (
	"fmt"
	"os"

	"github.com/qorbani/mysql-toolkit/drivers"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [sql statement]",
	Short: "Execute SQL Query",
	Long:  `Execute SQL Query against the MySQL database provided`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		sql, err = drivers.NewMySQL()
		if err != nil {
			fmt.Printf("Error connection to db, %s\n", err.Error())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		result, err := sql.Query(args[0])
		if err != nil {
			fmt.Printf("Error executing query: %s\n", err.Error())
			os.Exit(1)
		} else {
			fmt.Print(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
