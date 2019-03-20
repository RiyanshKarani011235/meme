// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rhysd/abspath"
	"github.com/spf13/cobra"

	"github.com/riyanshkarani011235/meme/lexer"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "meme build builds the provided set of meme description files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		files := make([]string, len(args))
		for i, arg := range args {
			files[i] = getAbsoluteFilePath(arg)
		}

		build(files)
	},
}

func getAbsoluteFilePath(file string) string {
	absolutePath, err := abspath.ExpandFrom(file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	path := absolutePath.String()

	// check if the file actually exists
	if fi, err := os.Stat(path); err == nil {
		// path exists, make sure that this is a file
		// and not a directory
		switch mode := fi.Mode(); {
		case mode.IsRegular():
			// do nothing
		default:
			fmt.Printf("Error: `%s` is not a file\n", path)
			os.Exit(1)
		}

	} else if os.IsNotExist(err) {
		// path does not exist
		fmt.Printf("Error: file `%s` does not exits\n", path)
		os.Exit(1)
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if !strings.HasSuffix(path, ".meme") {
		fmt.Printf("`%s` is not a meme file\n", path)
		os.Exit(1)
	}

	return path
}

func build(files []string) {
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		l := lexer.NewLexer(string(data))

		tokens := l.Tokenize()

		for _, token := range tokens {
			fmt.Printf("%v\n", token)
		}
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
