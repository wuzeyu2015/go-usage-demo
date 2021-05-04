package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const confFile = "zstorage.json"

type ZsNode struct {
	Name string	`json:"name"`
	Ip string	`json:"ip"`
}

type ZsConf struct {
	ClusterId string     `json:"cluster_id"`
	Controllers []ZsNode `json:"controllers"`
}

func getConf() *ZsConf {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Print("failed reading data from file: %s", err)
		return nil
	}
	var c ZsConf
	err2 := json.Unmarshal(data, &c)
	if err2!=nil {
		fmt.Print("Error: parse configuration file\n");
		return nil
	}
	return &c
}

