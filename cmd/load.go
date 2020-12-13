package cmd

import (
	"fmt"
	"os"

	"github.com/qorbani/mysql-toolkit/drivers"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [file.csv]",
	Short: "Load CSV File",
	Long:  `Load CSV File in to the specified MySQL table`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		sql, err = drivers.NewMySQL()
		if err != nil {
			fmt.Printf("Error connection to db, %s\n", err.Error())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// https://dev.mysql.com/doc/refman/8.0/en/load-data.html
		// Add Params:
		//	- Table Name
		//  - [Partition]
		// 	- [Character Set]
		//  - Fields
		// 	- Ignore X Rows
		//  - Set List=<Col/Val>

		// result, err := sql.Query(args[0])
		// if err != nil {
		// 	fmt.Printf("Error executing query: %s\n", err.Error())
		// 	os.Exit(1)
		// } else {
		// 	fmt.Print(result)
		// }
		fmt.Print(args)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
