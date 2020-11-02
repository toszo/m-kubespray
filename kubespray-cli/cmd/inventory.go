package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(inventoryCmd)
}

var inventoryCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Inventory generation for kubespray",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inventory called")
	},
}
