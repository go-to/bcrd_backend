package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateShopsImageCmd = &cobra.Command{
	Use:   "updateShopsImage",
	Short: "get shops image url & update to db",
	Long:  "get shops image url & update to db",
	Run: func(cmd *cobra.Command, args []string) {
		startBatch()
	},
}

func init() {
	rootCmd.AddCommand(updateShopsImageCmd)
}

func startBatch() {
	loadEnv()

	u := setup()

	err := u.Shop.UpdateShopsImage()
	if err != nil {
		panic(err)
	}
	fmt.Println("updateShopsImage done")
}
