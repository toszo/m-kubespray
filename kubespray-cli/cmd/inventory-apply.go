package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toszo/kubespray-cli/pkg/inventory"
)

func init() {
	inventoryCmd.AddCommand(applyInventoryCmd)

}

var applyInventoryCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create execute cluster build",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("inventory apply called")
		inventory.Apply()
	},
}
