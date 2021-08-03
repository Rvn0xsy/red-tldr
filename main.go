package main

import (
	"fmt"
	"github.com/spf13/viper"
	"red-tldr/common"
	"red-tldr/pkg"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("$HOME/.red-tldr/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	
	pkg.SetDbDir(viper.GetString("red-tldr.path"))
	if viper.GetBool("red-tldr.auto-update") {
		pkg.UpdateDb()
	}
}

func main() {
	var flagArgs common.FlagStruct
	common.SetFlag(&flagArgs)
	common.Runner(&flagArgs)
}