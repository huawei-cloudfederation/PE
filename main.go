package main

import (
	//Standard packages
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"log"
	//Borrowed packages
	//Packages belong to our project
	"./Config"
	"./Poller"
)

func ProcessConfFile(filename string, conf *config.DCConfig) {

	file_content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Unable to read the config file %v", err)
	}

	err = json.Unmarshal(file_content, conf)

	if err != nil {
		log.Fatalf("unable to unmarshall the config file not a valid json err=%", err)
	}
}

func NewConfig() config.DCConfig {
	return config.DCConfig{}
}

func main() {

	log.Printf("The code just started")
	//Initalize parse config
	config := NewConfig()

	//Try to parse the config file
	conffile := flag.String("config", "./config.json", "Supply the location of MrRedis configuration file")
	dummyConfig := flag.Bool("printDummyConfig", false, "IF you want to print the default(false) config")
	flag.Parse()

	//
	if *dummyConfig == true {
		config_byte, err := json.MarshalIndent(config, " ", "  ")
		if err != nil {
			log.Printf("Error Marshalling the default config file %v", err)
			return
		}
		fmt.Printf("%s\n", string(config_byte))
		return
	}

	ProcessConfFile(*conffile, &config)
	log.Println(config.List)

	//Start Poller
	//go poller.Run(Array of Mesos Master/Zookeeper's endpoint)
	go poller.Run(config.List)

	//Start the PE Engine
	//go PE.Run(Policy Thresold information)

	//Start the HTTP Server
	//go httplib.Run(PortNumber, AvertizeIP)

	//wait for ever
	wait := make(chan struct{})
	<-wait
}
