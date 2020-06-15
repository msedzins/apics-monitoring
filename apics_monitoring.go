package main

import (
	"apics-monitoring/configuration"
	"apics-monitoring/restapi"
	"flag"
	"fmt"
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

func main() {

	//LOAD CONFIGURATION
	configFile := handleInputParams()
	conf, err := configuration.LoadConfiguration(configFile)
	if err != nil {
		fmt.Println("Error loading configuration.", err)
		os.Exit(1)
	}

	//GET AUTH TOKEN
	auth := restapi.NewAuthentication()
	token, err := auth.GetToken(*conf)
	if err != nil {
		fmt.Println("Error getting IDCS token.", err)
		os.Exit(1)
	}

	fmt.Println("token", token)

}
