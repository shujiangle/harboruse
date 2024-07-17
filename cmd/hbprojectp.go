/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	harborget "harboruse/pkg/api"
)

// hbprojectpCmd represents the hbprojectp command
var hbprojectpCmd = &cobra.Command{
	Use:   "hbprojectp",
	Short: "获取harbor project 列表",
	Long:  `获取harbor project 列表`,
	Run: func(cmd *cobra.Command, args []string) {
		parentCmd := cmd.Parent()
		url, _ := parentCmd.PersistentFlags().GetString("url")

		username, _ := parentCmd.PersistentFlags().GetString("username")
		password, _ := parentCmd.PersistentFlags().GetString("password")
		file, _ := parentCmd.PersistentFlags().GetString("file")
		//fmt.Println(url, username, password, file)
		fmt.Println("所有harbor项目列表:")
		harborget.Harborgetprojectmain(url, username, password, file)
	},
}

func init() {
	rootCmd.AddCommand(hbprojectpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hbprojectpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hbprojectpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//hbprojectpCmd.Flags().StringP("dst-password", "P", "Harbor12345", "输入目标地址的密码")

}
