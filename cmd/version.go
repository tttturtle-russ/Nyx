/*
Copyright © 2024 Turtle Russ <tttturtleruss@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tttturtle-russ/Nyx/cfg"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current version of Nyx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
