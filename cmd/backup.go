package cmd

import (
	"fmt"

	"github.com/spf13/viper"
	// impor git from tools directory
	"github.com/username/rabbit-backup/tools/git"
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

	// Create git object
	git := git.NewGit()
	git.GitPath = viper.GetString("git.path")
	git.GitDir = viper.GetString("git.dir")

}
