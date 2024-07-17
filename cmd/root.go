/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"harboruse/pkg/api"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "harboruse",
	Short: "查看harbor镜像的地址",
	Long:  `查看harbor镜像的地址`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		file, _ := cmd.Flags().GetString("file")
		harborget.Harborgetmain(url, username, password, file)
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.harboruse.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("url", "l", "https://10.0.15.171", "输入你的URL")
	rootCmd.PersistentFlags().StringP("username", "u", "admin", "输入你的用户名")
	rootCmd.PersistentFlags().StringP("password", "p", "Harbor12345", "输入你的密码")
	rootCmd.PersistentFlags().StringP("file", "f", "harborget.txt", "输入你要保存到哪个文件")
}
