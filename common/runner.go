package common

import (
	"log"
	"os"
	"red-tldr/pkg"
)

func Runner(flagStruct * FlagStruct){
	switch os.Args[1] {
	case "search":
		log.Println("[Search Module is Running]...")
		defer log.Println("[Search Module is Done.]")
		if flagStruct.SearchFlag.Update{
			log.Println("[Search Module is Updating]...")
			pkg.UpdateDb()
		}else {
			if len(os.Args) < 3 {
				log.Println("[Search keyword not found]...")
				break
			}
			fileList := pkg.SearchDB(pkg.SearchDbName,os.Args[2])
			if len(fileList) == 0{
				log.Println("[Search Module Not Found]...")
				break
			}
			pkg.SelectOneResult(fileList)
		}
		break
	case "update":
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