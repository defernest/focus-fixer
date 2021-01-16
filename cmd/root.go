package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Config string

var rootCmd = &cobra.Command{
	Use:                   "focusfix { bot | fix } -c config.csv [-t int]",
	Short:                 "FocusFix is an easy way to focus your cameras",
	Long:                  `A quick and easy way to adjust focus on AXIS cameras using the built-in autofocus feature`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactValidArgs(1),
	ValidArgs:             []string{"bot", "fix"},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(fixCmd, botCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
