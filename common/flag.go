package common

import (
	"flag"
	"os"
	"red-tldr/utils"
)

func SetFlag(flagStruct * FlagStruct)  {
	var err error
	flagSearch := flag.NewFlagSet("search", flag.ExitOnError)
	flagSearch.BoolVar(&flagStruct.SearchFlag.Update, "update",false,"Update Search DB.")
	flagSearch.BoolVar(&flagStruct.SearchFlag.Init, "init",false,"Init Search DB.")

	if len(os.Args) < 2{
		utils.ShowBanner()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "search":
		err = flagSearch.Parse(os.Args[2:])
		utils.CheckErrorOnExit(err)
	default:
		utils.ShowHelp()
	}

}