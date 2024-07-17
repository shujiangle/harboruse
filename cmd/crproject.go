/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// crprojectCmd represents the crproject command
var crprojectCmd = &cobra.Command{
	Use:   "crproject",
	Short: "创建项目",
	Long:  `创建项目`,
	Run: func(cmd *cobra.Command, args []string) {
		parentCmd := cmd.Parent()
		url, _ := parentCmd.PersistentFlags().GetString("url")

		username, _ := parentCmd.PersistentFlags().GetString("username")
		password, _ := parentCmd.PersistentFlags().GetString("password")
		dsturl, _ := cmd.Flags().GetString("dst-url")
		dstusername, _ := cmd.Flags().GetString("dst-username")
		dstpassword, _ := cmd.Flags().GetString("dst-password")
		fmt.Printf("源端url: %v, 源端username:%v, 源端password: %v\n", url, username, password)
		fmt.Printf("目标端url: %v, 目标端username:%v, 目标端password: %v", dsturl, dstusername, dstpassword)

	},
}

func init() {
	rootCmd.AddCommand(crprojectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crprojectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crprojectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	crprojectCmd.Flags().StringP("dst-url", "L", "https://192.168.153.10", "输入目标的URL")
	crprojectCmd.Flags().StringP("dst-username", "U", "admin", "输入目标地址的用户名")
	crprojectCmd.Flags().StringP("dst-password", "P", "Harbor12345", "输入目标地址的密码")
}
