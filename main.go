package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagAction     string
	flagUser       string
	flagRepo       string
	flagCreateRepo bool
	flagDeleteRepo bool
)

type Repo struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func init() {
	flag.StringVar(&flagAction, "a", "", "Specify the option\nremote: get html/ssh repo url")
	flag.StringVar(&flagUser, "u", "jflores-p", "User of github")
	flag.StringVar(&flagRepo, "r", "", "Repository name")
	flag.BoolVar(&flagCreateRepo, "c", false, "Triggers create repo")
	flag.BoolVar(&flagDeleteRepo, "d", false, "Triggers delete repo")

	flag.Parse()
}

func main() {
	if flag.NFlag() == 0 {
		fmt.Println("No flags called. Try using: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//var rep repo
	//
	//_ = json.Unmarshal([]byte(flagRepo), &rep)
	//
	//fmt.Println(rep)

	fmt.Println(flagCreateRepo)
	fmt.Println(flagDeleteRepo)
	fmt.Println(flagUser)
	fmt.Println(flagRepo)


	if flagCreateRepo || flagDeleteRepo {
		repo := Repo{
			Name:  flagRepo,
			Owner: flagUser,
		}
		fmt.Println(repo)
	}
}
