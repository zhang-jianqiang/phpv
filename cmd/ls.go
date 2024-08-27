/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhang-jianqiang/phpv/pkg"
	"strings"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "php版本列表, 示例: phpv ls",
	Long:  `php版本列表, 示例: phpv ls`,
	Run: func(cmd *cobra.Command, args []string) {
		phpVersionItems, err := pkg.ListPhpVersion()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("可用php版本: \r\n" + strings.Join(phpVersionItems, "\r\n"))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
