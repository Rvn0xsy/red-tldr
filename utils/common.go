package utils


import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var banner = `

         /\_/\
    ____/ o o \   For Red Team [TL;DR]
  /~____  =ø= /   Github @Rvn0xsy
 (______)__m_m)   Blog: https://payloads.online

------------------------------------------------
Thank you for Use https://github.com/Rvn0xsy/red-tldr`

func GenerateConfig()  {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	homeDir := os.Getenv("HOME")
	configDir := homeDir+"/.red-tldr/"
	_, err := os.Stat(configDir)
	if err == nil{
		_ = os.Remove(configDir)
	}
	CheckErrorOnExit(os.Mkdir(configDir,os.ModePerm))
	os.Chmod(configDir,os.ModePerm)

	viper.AddConfigPath(configDir)
	viper.SetDefault("red-tldr.path","~/red-tldr-db/")
	viper.SetDefault("red-tldr.auto-update",true)

	CheckErrorOnExit(viper.WriteConfigAs(configDir+"config.toml"))
	log.Println("Generate Config Success, Config File Path : ", "~/.red-tldr/config.toml")

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
	search`)
}