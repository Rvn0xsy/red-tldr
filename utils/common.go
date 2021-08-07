package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

const (
	 banner  = `
         /\_/\
    ____/ o o \   For Red Team [TL;DR]
  /~____  =ø= /   Github @Rvn0xsy
 (______)__m_m)   Blog: https://payloads.online

------------------------------------------------
Thank you for Use https://github.com/Rvn0xsy/red-tldr`
	configName = "config.toml"
	databaseName = string(os.PathSeparator) + "db" + string(os.PathSeparator) + "db.json"
	configDir = string(os.PathSeparator) + ".red-tldr" + string(os.PathSeparator)
	databaseDir = string(os.PathSeparator) + "red-tldr-db" + string(os.PathSeparator)
)


func GetPathSeparator()(pathSeparator string){
	pathSeparator = string(os.PathSeparator)
	return pathSeparator
}

func CheckDatabaseExist()(isExist bool){
	file := viper.GetString("red-tldr.path")
	if file == ""{
		file = GetDatabaseFilePath()
	}
	if _, err := os.Stat(GetDatabaseFilePath()); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func GetConfigPath()(Dir string){
	return path.Join(getHomeDir(), configDir)
}

func GetConfigFilePath()(configFilePath string){
	return path.Join(GetConfigPath(), configName)
}

func GetDatabasePath()(databasePath string){
	databasePath = viper.GetString("red-tldr.path")
	if databasePath == ""{
		databasePath = path.Join(getHomeDir(), databaseDir)
	}
	return databasePath
}

func GetDatabaseFilePath()(configFilePath string){
	return path.Join(GetDatabasePath(), databaseName)
}

func getHomeDir()(homeDir string){
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}


func GenerateConfig()  {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	configDir := GetConfigPath()
	_, err := os.Stat(configDir)
	if err == nil{
		_ = os.Remove(configDir)
	}
	CheckErrorOnExit(os.Mkdir(configDir,os.ModePerm))
	os.Chmod(configDir,os.ModePerm)

	viper.AddConfigPath(configDir)
	viper.SetDefault("red-tldr.path",GetDatabasePath())
	viper.SetDefault("red-tldr.index-update",false)
	viper.SetDefault("red-tldr.github-update",false)

	CheckErrorOnExit(viper.WriteConfigAs(GetConfigFilePath()))
	log.Println("[Generate Config Success, Config File Path : ",  GetConfigFilePath() ,"]")

}

func CheckErrorOnExit(err error)  {
	if err != nil{
		panic(err)
	}
	return
}
func CheckErrorOnPrint(err error)  {
	if err != nil{
		log.Println(err)
	}
	return
}

func ShowBanner()  {
	fmt.Println(banner)
}

func ShowHelp()  {
	ShowBanner()
	fmt.Println(`
Modules:
	search   [search Module]
	update   [update database from github https://github.com/Rvn0xsy/red-tldr-db]`)
}