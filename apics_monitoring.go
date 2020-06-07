package main

import (
	"apics-monitoring/configuration"
	"fmt"
)

func main() {
	/* This is my first sample program. */
	fmt.Println("Hello, World!")

	config, err := configuration.LoadConfiguration("configuration/configuration_test.json")

	fmt.Println("AAAA")
	fmt.Println("Error:", err)
	fmt.Printf("Config: %+v", config)
	fmt.Println("aaa:", config.GetAPIPlatformHost())

	//configuration.APIPlatformClientID
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//Load configuration
	//call modules
}
