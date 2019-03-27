package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type SnmptrapdDeviceDefaults struct {
	Host string `json:"host"`
}

type SnmptrapdTrapDefaults struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type SnmptrapdDefaults struct {
	Device SnmptrapdDeviceDefaults `json:"device"`
	Trap   SnmptrapdTrapDefaults   `json:"trap"`
}

type SnmptrapdSettings struct {
	Defaults SnmptrapdDefaults `json:"defaults"`
}

type SensuAgentApiSettings struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type SensuAgentSettings struct {
	API SensuAgentApiSettings `json:"api"`
}

type SensuCheckSettings struct {
	Namespace   string `json:"namespace"`
	LabelPrefix string `json:"label_prefix"`
	Status      int    `json:"status"`
}

type SensuSettings struct {
	Agent SensuAgentSettings `json:"agent"`
	Check SensuCheckSettings `json:"check"`
}

type Settings struct {
	Snmptrapd SnmptrapdSettings `json:"snmptrapd"`
	Sensu     SensuSettings     `json:"sensu"`
}

func LoadConfig(filename string) *Settings {
	var config *Settings
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	filebytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	json.Unmarshal(filebytes, &config)

	output, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	fmt.Println(string(output))
	return config
}
