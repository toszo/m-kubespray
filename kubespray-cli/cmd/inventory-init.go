package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toszo/kubespray-cli/pkg/inventory"
)

func init() {
	inventoryCmd.AddCommand(initInventoryCmd)

}

var initInventoryCmd = &cobra.Command{
	Use:   "init",
	Short: "Create init inventory",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("inventory init called")
		inventory.Init()
	},
}
