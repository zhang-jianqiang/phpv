/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhang-jianqiang/phpv/pkg"
	"path/filepath"
	"slices"
	"strings"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "切换php版本, 如: phpv use php7.3.4nts",
	Long:  `切换php版本, 如: phpv use php7.3.4nts`,
	Run: func(cmd *cobra.Command, args []string) {
		phpVersionItems, err := pkg.ListPhpVersion()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(args) < 1 {
			fmt.Println("版本为空")
			return
		}

		if !slices.Contains(phpVersionItems, args[0]) {
			fmt.Println("版本号输入错误, 支持的版本如下: " + strings.Join(args, "\r\n"))
			return
		}

		phpPath, err := pkg.ReadConfig()
		if err != nil {
			fmt.Println("php目录未配置, 请先配置php目录, 例如: phpv config D:/phpstudy_pro/Extensions/php")
			return
		}

		newPhpEnv := filepath.Join(phpPath, args[0])
		err = pkg.SetLink(newPhpEnv)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Success!\r\nVersion: " + args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
