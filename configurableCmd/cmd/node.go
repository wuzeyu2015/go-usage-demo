package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

var CmdOperationMap = make(map[string] map[string] interface{})

type NodeCmd struct {

}

type Node struct {
	NodeName string `json:"nodeName"`
	IpAddr1  string `json:"ipAddr1"`
	IpAddr2  string `json:"ipAddr2"`
	Id       int    `json:"id"`
	Active   int    `json:"active"`
}

type NodeShow struct {
	RetCode string `json:"retCode"`
	Data     Node   `json:"data"`
	Code	int 	`json:"code"`
}

type NodeList struct {
	RetCode string `json:"retCode"`
	Data 	[]Node 	`json:"data"`
	Code	int 	`json:"code"`
}

//func NodeBindBuffer(o *OperEntity, args []string) {
//
//	//and bind the parameter to buffer
//	if o.Name == "add" {
//		nodeInfo := &Node{}
//		Buffer = nodeInfo
//		nodeInfo.NodeName = args[0]
//		nodeInfo.IpAddr1 = args[1]
//		nodeInfo.Active = 1
//
//		if o.ArgsActualLen == 3 {
//			nodeInfo.IpAddr2 = args[2]
//		}
//	}
//	if o.Name == "del" {
//		nodeInfo := &Node{}
//		Buffer = nodeInfo
//		nodeInfo.NodeName = args[0]
//	}
//	if o.Name == "update" {
//		nodeInfo := &Node{}
//		Buffer = nodeInfo
//		nodeInfo.NodeName = args[0]
//		nodeInfo.IpAddr1 = args[1]
//
//		if o.ArgsActualLen == 3 {
//			nodeInfo.IpAddr2 = args[2]
//		}
//	}
//	if o.Name == "list" {
//		Buffer = &NodeList{}
//		Display = ListNode
//	}
//	if o.Name == "show" {
//		nodeInfo := &Node{}
//		nodeInfo.NodeName = args[0]
//		Buffer = nodeInfo
//		Display = ShowNode
//	}
//}

func (n *NodeCmd)ListNode(ret []byte){

	var nl NodeList
	err := json.Unmarshal(ret, &nl)
	if err != nil {
		fmt.Printf("decode error, %v\n", err)
		os.Exit(1)
	}
	for _, n := range nl.Data {
		fmt.Printf("%4d %-10s %15s %15s %d\n",
			n.Id, n.NodeName, n.IpAddr1,
			n.IpAddr2, n.Active)
	}
}

func (n *NodeCmd)ShowNode(ret []byte){

	var rn NodeShow
	err := json.Unmarshal(ret, &rn)
	if err != nil {
		fmt.Printf("unknow error %v\n", err)
		os.Exit(1)
	}
	p := &rn.Data
	if len(p.NodeName) != 0 {
		fmt.Printf("%-10s: %d\n", "Id", p.Id)
		fmt.Printf("%-10s: %s\n", "Name", p.NodeName)
		fmt.Printf("%-10s: %s\n", "IpAddr1", p.IpAddr1)
		fmt.Printf("%-10s: %s\n", "IpAddr2", p.IpAddr2)
		fmt.Printf("%-10s: %d\n", "Active", p.Active)
	}
}

func init() {

	CmdOperationMap["node"] = make(map[string] interface{})
	CmdOperationMap["node"]["-"] = &NodeCmd{}
	CmdOperationMap["node"]["add"] = &Node{}
	CmdOperationMap["node"]["del"] = &Node{}
	CmdOperationMap["node"]["update"] = &Node{}
	CmdOperationMap["node"]["show"] = &Node{}
	CmdOperationMap["node"]["list"] = &NodeList{}
}
