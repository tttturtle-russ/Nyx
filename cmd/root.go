/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tttturtle-russ/Nyx/internel/parser"
	"os"
)

var (
	File string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Nyx",
	Short: "Nyx is a command line reverse engineering tool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("debug") != nil {
			fmt.Println("Debug")
		}
		//text := []string{"test", "test", "test", "test", "test", "test", "test"}
		//funcs := []string{"func", "func", "func", "func", "func", "exit"}
		fmt.Println(parser.NewDisassembler(args[0], false).Functions())
		//display.InitScreen(text, funcs).Display()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&File, "file", "f", "", "specific the file that need to reverse")
	rootCmd.Flags().BoolP("debug", "d", false, "into debug mode")
}
