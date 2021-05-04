package main

import (
	"encoding/json"
	"fmt"
)

type ClusterInfo struct {
	Nodes struct {
		Active   int `json:"active"`
		Inactive int `json:"inactive"`
		Total    int `json:"total"`
	}`json:"nodes"`
	Procs struct {
		Active   int `json:"active"`
		Inactive int `json:"inactive"`
		Total    int `json:"total"`
	} `json:"procs"`
}

type ClusterInfo2 struct {
	Nodes struct {
		Active   int `json:"active"`
		Inactive int `json:"inactive"`
		InactiveStr string `json:"inactive_str"`
		Total    int `json:"total"`
	}`json:"nodes"`
	Procs struct {
		Active   int `json:"active"`
		Inactive int `json:"inactive"`
		Total    int `json:"total"`
	} `json:"procs"`
}

func main() {

	fmt.Print("不同顶层tag之间的内部tag可以相同")
	clusterInfo := ClusterInfo{
		Nodes:	struct {
			Active   int `json:"active"`
			Inactive int `json:"inactive"`
			Total    int `json:"total"`
		} {
		Active:1,
		Inactive:2,
		Total:3,
	},
		Procs:	struct {
			Active   int `json:"active"`
			Inactive int `json:"inactive"`
			Total    int `json:"total"`
		} {
			Active:1,
			Inactive:2,
			Total:3,
		},
	}
	json_str, err := json.Marshal(clusterInfo);
	if err != nil {
		println(err)
	}
	println(string(json_str))

	um := &ClusterInfo2{}
	err = json.Unmarshal(json_str, um)
	if err != nil {
		println(err)
	}
	fmt.Print(*um)
	return
}