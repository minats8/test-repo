package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "hello",
	Short:   "first Command",
	Version: version,
	Long:    `a longer description for the first command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is th efirst cobra example")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
