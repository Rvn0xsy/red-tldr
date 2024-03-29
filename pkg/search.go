package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"red-tldr/utils"
	"strings"
)

var (
	SearchDbDir = ""
	SearchDbName = ""
)

func init()  {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(utils.GetConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		utils.GenerateConfig()
	}
}

func SetDbDir()  {
	SearchDbDir = utils.GetDatabasePath() + string(os.PathSeparator)
	SearchDbName = utils.GetDatabaseFilePath()
}

type SearchDataStruct struct {
	Name string `yaml:"name"`
	Tags []string `yaml:"tags"`
	Data string `yaml:"data"`
}

type DataStruct struct {
	Name string `json:"name"`
	Tags []string `json:"tags"`
	File string `json:"file"`
}

type SearchDbStruct struct {
	Data []DataStruct`json:"data"`
}

type SearchResultStruct struct {
	Filename string
	Name string
}

func checkSliceContains(slice []string,value string) bool{
	for _,v := range slice{
		if value == v || strings.Contains(v,value) || strings.HasPrefix(v,value){
			return true
		}
	}
	return false
}

func Search(file string,keyword string)  {
	Data := new(SearchDataStruct)
	yamlFile, err := ioutil.ReadFile(file)
	utils.CheckErrorOnExit(err)
	err = yaml.Unmarshal(yamlFile, Data)
	utils.CheckErrorOnExit(err)
	if checkSliceContains(Data.Tags, keyword) {
		log.Println("Searching Success ...")
		fmt.Println(Data.Data)
	}else {
		log.Fatal("Not Found.")
	}
}

func ShowHighLightData(Data * SearchDataStruct){
	if viper.GetBool("red-tldr.color"){
		// Color
		color.Style{color.Red, color.OpBold}.Println(Data.Name)
		lines := strings.Split(Data.Data,"\n")
		printColor := false
		for _,line := range lines{
			if strings.HasPrefix(line,"```"){
				if printColor {
					printColor = false
					continue
				}else{
					fmt.Println()
					printColor = true
					continue
				}
			}
			if printColor {
				color.Style{color.Green, color.OpBold}.Println(line)
				// color.Green.Println(line)
				continue
			}
			fmt.Println(line)
		}
	}else{
		fmt.Println("=================")
		fmt.Println(Data.Name)
		fmt.Println("=================")
		fmt.Println(Data.Data)
		return
	}
}

func ShowDetails(file SearchResultStruct)  {
	Data := new(SearchDataStruct)
	yamlFile, err := ioutil.ReadFile(SearchDbDir+ utils.GetPathSeparator() + file.Filename)
	utils.CheckErrorOnExit(err)
	err = yaml.Unmarshal(yamlFile, Data)
	utils.CheckErrorOnExit(err)
	ShowHighLightData(Data)
}

func SelectOneResult(fileList []SearchResultStruct)  {
	var(
		i = 0
		count int
	)
	if len(fileList) == 1{
		ShowDetails(fileList[0])
		return
	}

	for n,f := range fileList{
		fmt.Println(n,") ",f.Name)
	}

	count = len(fileList)
	fmt.Print(fmt.Sprintf("[Count : %d] > Select Result Number : ",count))
	_, err := fmt.Scanf("%d", &i)
	// fmt.Println(i)
	if err != nil{
		i = count-1
	}
	if i >= len(fileList){
		i = count-1
	}
	ShowDetails(fileList[i])
}


// GetKeywordsTotal 获取关键字数量
func GetKeywordsTotal() int{
	Data := new(SearchDbStruct)
	DbFile, err := ioutil.ReadFile(SearchDbName)
	utils.CheckErrorOnExit(err)
	err = json.Unmarshal(DbFile, Data)
	utils.CheckErrorOnExit(err)
	return len(Data.Data)
}

func getDataStruct(file string) (Data * DataStruct ){
	Data = new(DataStruct)
	yamlFile, err := ioutil.ReadFile(file)
	utils.CheckErrorOnPrint(err)
	err = yaml.Unmarshal(yamlFile, Data)
	if err != nil{
		fmt.Println("ERROR >>>> ", file)
	}
	utils.CheckErrorOnPrint(err)
	return Data
}

func SearchDB(file string,keyword string) (yamlFile []SearchResultStruct) {
	Data := new(SearchDbStruct)
	DbFile, err := ioutil.ReadFile(file)
	utils.CheckErrorOnExit(err)
	err = json.Unmarshal(DbFile, Data)
	utils.CheckErrorOnExit(err)
	for _,o := range Data.Data{
		if checkSliceContains(o.Tags,keyword) {
			yamlFile = append(yamlFile, struct {
				Filename string
				Name     string
			}{Filename: o.File, Name: o.Name})
		}
	}
	return yamlFile
}

func UpdateDb()  {
	DbFile ,err := os.Create(SearchDbName)
	DbStruct := new(SearchDbStruct)
	utils.CheckErrorOnExit(err)
	defer DbFile.Close()
	fileList := utils.GetAllDataFile(SearchDbDir)
	if len(fileList) == 0 {
		return
	}
	for _,fileName := range fileList {
			Data := getDataStruct(fileName)
			dbFilename := strings.Replace(fileName, SearchDbDir, "", 1)
			DbStruct.Data = append(DbStruct.Data,DataStruct{
				Name: Data.Name,
				Tags: Data.Tags,
				File: dbFilename,
			})
	}
	DbJsonData , err := json.Marshal(DbStruct)
	utils.CheckErrorOnExit(err)
	_, err = DbFile.WriteString(string(DbJsonData))
	utils.CheckErrorOnExit(err)
}
