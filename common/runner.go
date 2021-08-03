package common

import (
	"log"
	"os"
	"red-tldr/pkg"
	"red-tldr/utils"
)

func Runner(flagStruct * FlagStruct){
	switch os.Args[1] {
	case "search":
		log.Println("[Search Module is Running]...")
		defer log.Println("[Search Module is Done.]")
		if flagStruct.SearchFlag.Init {
			utils.GenerateConfig()
			break
		}

		if flagStruct.SearchFlag.Update{
			log.Println("[Search Module is Updating]...")
			pkg.UpdateDb()
		}else {
			fileList := pkg.SearchDB(pkg.SearchDbName,os.Args[2])
			if len(fileList) == 0{
				log.Println("[Search Module Not Found]...")
				break
			}
			pkg.SelectOneResult(fileList)
		}
		break
	default:
		log.Println("other...")
	}
}