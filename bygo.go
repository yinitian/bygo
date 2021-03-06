package main

import (
	"os"
	"github.com/bysir-zl/bygo/artisan"
	"github.com/bysir-zl/bygo/log"
)

func main() {
	args := os.Args
	//if len(args) == 1 {
	//	args = []string{"create", "helloworld"}
	//}

	if len(args) == 1 {
		log.Error("Bygo", "error args")
		return
	}

	command := args[1]

	switch command {
	case "model":
		//table := args[2]
		//artisan.CreateModelFile(table)
	case "swagger":
		if len(args) != 4 {
			args = []string{"", "", "./", "./swagger.json"}
		}
		path := args[2]
		out := args[3]
		artisan.Swagger(path, out)
	case "json2go":
		err := artisan.Json2Go()
		if err != nil {
			log.ErrorT("bygo", err.Error())
		} else {
			log.InfoT("bygo", "gen go struct from clipboard success")
		}
	case "xml2go":
		err := artisan.Xml2Go()
		if err != nil {
			log.ErrorT("bygo", err.Error())
		} else {
			log.InfoT("bygo", "gen go struct from clipboard success")
		}
	case "go2pb":
		err := artisan.Go2Pb()
		if err != nil {
			log.ErrorT("bygo", err.Error())
		} else {
			log.InfoT("bygo", "gen pb struct from clipboard success")
		}
	}

}
