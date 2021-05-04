package main

import (
	"encoding/json"
	"log"
)

var Buffer interface{}

type Node struct {
	NodeName string `json:"nodeName"`
	IpAddr1  string `json:"ipAddr1"`
	IpAddr2  string `json:"ipAddr2"`
	Id       int    `json:"id"`
	Active   int    `json:"active"`
}


func main() {
	Buffer = &Node{NodeName: "sdf", IpAddr1: "sdf", IpAddr2: "fsd", Id: 3, Active: 4}
	buf, err := json.Marshal(Buffer)
	if err != nil {
		log.Fatal(err)
		return
	}
	print(string(buf))
	return
}
