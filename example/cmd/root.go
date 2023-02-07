package cmd

import (
	"cobra/helper"
	"fmt"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:     "database",
	Short:   "database",
	Long:    `proceeds to database Storage`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proceeds to database storage...")
	},
}

var fileCmd = &cobra.Command{
	Use:     "file",
	Short:   "fileSystem",
	Long:    `proceeds to file storage`,
	Aliases: []string{"f"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proceeds to fileSystem storage...")
	},
}

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Short:   "Reverses a string",
	Aliases: []string{"rev"},
	Args:    cobra.ExactArgs(1),
	Run:     reverse,
}

func reverse(cmd *cobra.Command, args []string) {
	res := helper.Reverse(args[0])
	fmt.Println(res)
}

var uppercaseCmd = &cobra.Command{
	Use:     "uppercase",
	Short:   "Uppercase a string",
	Aliases: []string{"upper"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Uppcase(args[0])
		fmt.Println(res)
	},
}

var optBool bool
var modifyCmd = &cobra.Command{
	Use:     "modify",
	Short:   "modify a string",
	Aliases: []string{"mod"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := helper.Modify(args[0], optBool)
		fmt.Println(res)
	},
}

func init() {
	modifyCmd.Flags().BoolVarP(&optBool, "opt", "o", false, "Modify option")
	rootCmd.AddCommand(reverseCmd)
	rootCmd.AddCommand(uppercaseCmd)
	rootCmd.AddCommand(modifyCmd)
	rootCmd.AddCommand(databaseCmd)
	rootCmd.AddCommand(fileCmd)
}
