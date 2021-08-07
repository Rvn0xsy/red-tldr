package main

import (
	"github.com/spf13/viper"
	"red-tldr/common"
	"red-tldr/pkg"
	"red-tldr/utils"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(utils.GetConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		utils.GenerateConfig()
	}

	pkg.SetDbDir()

	if utils.CheckDatabaseExist() == false || viper.GetBool("red-tldr.github-update"){
		pkg.GetLatestReleaseFromGithub()
	}
	if viper.GetBool("red-tldr.index-update") {
		pkg.UpdateDb()
	}
}

func main() {
	var flagArgs common.FlagStruct
	common.SetFlag(&flagArgs)
	common.Runner(&flagArgs)
}