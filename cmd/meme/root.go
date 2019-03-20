package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meme",
	Short: "meme is a build tool for the meme language",
	Long:  "meme is an extensible data representation language",
	Run: func(cmd *cobra.Command, args []string) {
		// @todo do stuff here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
