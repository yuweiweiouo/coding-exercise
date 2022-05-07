package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuweiweiouo/coding-exercise/internal/server"
)

var rootCmd = &cobra.Command{
	Use:   "coding-exercise",
	Short: "Start a Restful task list API server",
	Run: func(cmd *cobra.Command, args []string) {
		configName, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}
		server, cleanup, err := server.CreateServer(configName)
		if err != nil {
			log.Panic(err)
			os.Exit(1)
		}
		defer cleanup()
		server.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", "local", "想使用的配置檔名稱")
}
