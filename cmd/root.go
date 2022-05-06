package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuweiweiouo/coding-exercise/internal/app"
)

var rootCmd = &cobra.Command{
	Use:   "coding-exercise",
	Short: "Start a Restful task list API server",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		_, cleanup, err := app.CreateApp()
		if err != nil {
			log.Panic(err)
		}
		defer cleanup()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
