package common

import (
	"fmt"
	"log"
	"os"
	"red-tldr/pkg"
)

const(
	banner  = `
         /\_/\
    ____/ o o \   For Red Team [TL;DR]
  /~____  =ø= /   Github @Rvn0xsy
 (______)__m_m)   Blog: https://payloads.online
                  Version: 0.4.2
------------------------------------------------
Thank you for Use https://github.com/Rvn0xsy/red-tldr`
)


func ShowBanner()  {
	fmt.Println(banner)
}

func ShowHelp()  {
	ShowBanner()
	fmt.Println(fmt.Sprintf("Keywords Total: %d ", pkg.GetKeywordsTotal()))
	fmt.Println(`
Command:
	<Keyword> [search keyword from database index]
	update    [update database index]
	upgrade   [update database from github https://github.com/Rvn0xsy/red-tldr-db]`)
}

func Runner(){
	if len(os.Args) < 2{
		ShowHelp()
		return
	}
	switch os.Args[1] {
		case "-h":
			ShowHelp()
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
