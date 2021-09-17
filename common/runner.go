package common

import (
	"log"
	"os"
	"red-tldr/pkg"
	"red-tldr/utils"
)

func Runner(){
	if len(os.Args) < 2{
		utils.ShowHelp()
		return
	}
	switch os.Args[1] {
		case "-h":
			utils.ShowHelp()
		case "update":
			pkg.UpdateDb()
		case "upgrade":
			pkg.GetLatestReleaseFromGithub()
			break
		default:
			fileList := pkg.SearchDB(pkg.SearchDbName,os.Args[1])
			if len(fileList) == 0{
				log.Println("[Search Module Not Found]...")
				break
			}
			pkg.SelectOneResult(fileList)
		}
}