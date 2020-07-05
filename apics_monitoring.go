package main

import (
	"apics-monitoring/configuration"
	"apics-monitoring/modules"
	"apics-monitoring/modules/validatepolltime"
	"flag"
	"fmt"
	"log"
	"os"
)

func handleInputParams() string {
	help := flag.Bool("h", false, "Print help when set")
	config := flag.String("cf", "config.json", "Path to configuration file")
	flag.Parse()

	if *help == true {
		fmt.Printf("Program usage:%s [flags]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	return *config
}

//list of all supported modules to be run sequentially. Currently there is only one.
var listOfModules []modules.Module = []modules.Module{&validatepolltime.ValidatePollTime{}}

func main() {

	//LOAD CONFIGURATION
	configFile := handleInputParams()
	conf, err := configuration.LoadConfiguration(configFile)
	if err != nil {
		log.Fatalln("Error loading configuration.", err)
	}

	//GET AUTH TOKEN
	//TODO: Change it!
	//auth := restapi.NewAuthentication()
	token, err := "TOKEN", nil //auth.GetToken(*conf)
	if err != nil {
		log.Fatalln("Error getting IDCS token.", err)
	}

	//EXECUTE ALL MODULES
	for _, item := range listOfModules {
		fmt.Println("Calling module:,", item.GetName())
		alerts, err := item.Execute(token, *conf)
		if err != nil {
			log.Fatalln("Error calling the module.", err)
		}

		fmt.Printf("output:%+v", alerts)
	}

}
