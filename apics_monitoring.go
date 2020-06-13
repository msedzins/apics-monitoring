package main

import (
	"apics-monitoring/configuration"
	"apics-monitoring/restapi"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	auth := restapi.Authentication{}
	token, err := auth.GetToken(*conf)
	if err != nil {
		fmt.Println("Error getting IDCS token.", err)
		os.Exit(1)
	}

	fmt.Println("token", token)

}

func testRESTAPI() {

	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		log.Fatalln(err)
	}

	//req, err := http.NewRequest(http.MethodPost, "https://jsonplaceholder.typicode.com/todos/1", nil)
	//req.SetBasicAuth
	//client := http.Client{}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
}
