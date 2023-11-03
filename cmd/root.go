/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
    "strings"
    "time"
    
    "github.com/kelvinkuo/notion2jekyll/internal/convert"
    "github.com/spf13/cobra"
)

var (
    author      string
    createdTime string
    categoryStr string
    categories  []string
    tagStr      string
    tags        []string
    notionDir   string
    jkeyllDir   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "notion2jekyll",
    Short: "Convert markdown file exported by notion  into jekyll project dir.",
    Long: `examples:
  notion2j
`,
    
    Run: func(cmd *cobra.Command, args []string) {
        categories = strings.Split(categoryStr, ",")
        tags = strings.Split(tagStr, ",")
        
        convert.Notion2Jkeyll(author, createdTime, notionDir, jkeyllDir, categories, tags)
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    rootCmd.Flags().StringVarP(&author, "author", "a", "", "author of the post")
    rootCmd.Flags().StringVarP(&createdTime, "time", "m", time.Now().Format(time.DateTime), "created time format \"2006-01-02 15:04:05\"")
    rootCmd.Flags().StringVarP(&categoryStr, "categories", "c", "", "categories split by ','")
    rootCmd.Flags().StringVarP(&tagStr, "tags", "t", "", "tags split by ',' and should be lower-case")
    rootCmd.Flags().StringVarP(&notionDir, "notiondir", "n", "./", "the dir of exported notion post")
    rootCmd.Flags().StringVarP(&jkeyllDir, "jkeylldir", "j", "./", "the dir of created jkeyll post")
}
