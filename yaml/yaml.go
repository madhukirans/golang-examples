package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type EndpointConfig struct {
	ReqTimeout int  `yaml:"RequestTimeout"`
	UptimeRobotReqIntervel int  `yaml:"UptimeRobotRequestIntervalMinutes"`
	ExcludeEndPoints []string  `yaml:"ExcludeEndpoints"`
}


func (c *EndpointConfig) getEndpointConfig() *EndpointConfig {
	yamlFile, err := ioutil.ReadFile("/Users/madhuseelam/gopath/src/github.com/madhukirans/golang-examples/yaml/endpoints.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var C EndpointConfig
	C.getEndpointConfig()

	fmt.Println(C)
}
