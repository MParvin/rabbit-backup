/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rabbit-backup",
	Short: "Download backup from RabbitMQ and push to git",
	Long: `This tool will download backup from RabbitMQ and push to git.
	It will also create a new branch for each backup and push it to git.
	It can also be used to restore a backup from git to RabbitMQ.
	Also it can be encrypted and decrypted using GPG.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
		DoBackup(cfgFile)
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file (default is $HOME/.rabbit-backup.yaml)")

	rootCmd.Flags().BoolP("debug", "d", false, "Show more messages for debugging")
}
