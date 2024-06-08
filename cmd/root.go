/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/radareorg/r2pipe-go"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
		file := args[0]
		fmt.Println(file)
		pipe, err := r2pipe.NewPipe(file)
		if err != nil {
			panic(err)
		}
		pipe.Cmd("s main")
		s, err := pipe.Cmd("pdf")
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}
		ss := strings.Split(s, "\n")
		for i, v := range ss {
			fmt.Println(i, "\t", v)
		}
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
