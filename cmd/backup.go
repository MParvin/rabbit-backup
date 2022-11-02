package cmd

import (
	"fmt"

	tools "github.com/mparvin/rabbit-backup/tools"
	"github.com/spf13/viper"
)

func DoBackup(cfgFile string) {
	// Read config file
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	fmt.Println("RabbitMQ URL:", viper.GetString("rabbitmq.url"))

	// Delete git dir if exists
	git := tools.Git{
		GitPath: viper.GetString("git.path"),
	}
	err = tools.DeleteDir(viper.GetString("git.dir"))
	if err != nil {
		fmt.Printf("Error deleting git directory, %s", err)
	}

	// Git clone
	git.Clone(viper.GetString("git.url"), viper.GetString("git.dir"))

	// Change directory to git dir
	err = tools.ChangeDir(viper.GetString("git.dir"))
	if err != nil {
		fmt.Printf("Error changing directory, %s", err)
	}
	// Create backup
	err = tools.CreateBackup(viper.GetString("rabbitmq.url"), viper.GetString("rabbitmq.user"), viper.GetString("rabbitmq.password"))
	if err != nil {
		fmt.Printf("Error creating backup, %s", err)
	}

	// encrypt backup
	err = tools.EncryptBackup(viper.GetString("gpg.path"), viper.GetString("gpg.key"))
	if err != nil {
		fmt.Printf("Error encrypting backup, %s", err)
	}
	// if git status has changes, commit and push
	// else delete backup
	// Delete git dir

}
