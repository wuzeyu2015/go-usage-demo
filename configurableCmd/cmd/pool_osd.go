package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type PoolOsdCmd struct {

}


type PoolOsd struct {
	PoolName string `json:"poolName"`
	NodeName string `json:"nodeName"`
	OsdName  string `json:"osdName"`
	Id       int    `json:"id"`
}

type PoolOsdList struct {
	RetCode string    `json:"retCode"`
	DataList []PoolOsd `json:"data"`
	Code	int 		`json:"code"`
}

type PoolOsdShow struct {
	RetCode string    `json:"retCode"`
	DataList []PoolOsd `json:"data"`
	Code	int 		`json:"code"`
}
//func PoolOsdBindBuffer(o *OperEntity, args []string) {
//
//	//and bind the parameter to buffer
//	if o.Name == "add" {
//		newPoolOsd := &PoolOsd{}
//		Buffer = newPoolOsd
//		newPoolOsd.PoolName = args[0]
//		newPoolOsd.NodeName = args[1]
//		newPoolOsd.OsdName = args[2]
//	}
//	if o.Name == "del" {
//		delPoolOsd := &PoolOsd{}
//		Buffer = delPoolOsd
//		delPoolOsd.PoolName = args[0]
//		delPoolOsd.NodeName = args[1]
//		delPoolOsd.OsdName = args[2]
//	}
//	if o.Name == "list" {
//		Buffer = &PoolOsdList{}
//		Display = ListPoolOsd
//	}
//	if o.Name == "show" {
//		showPoolOsd := &PoolOsd{}
//		showPoolOsd.PoolName = args[0]
//		Buffer = showPoolOsd
//		//Display = cmd.ShowPoolOsd
//	}
//}


func (p *PoolOsdCmd) ListPoolOsd(ret []byte){

	pl := PoolOsdList{}
	err := json.Unmarshal(ret, &pl)
	if err != nil {
		fmt.Printf("decode error, %v\n", err)
		os.Exit(1)
	}

	for _, n := range pl.DataList {
		fmt.Printf("%4d %-10s %15s %15s\n",
			n.Id, n.PoolName, n.NodeName, n.OsdName)
	}
}

func (p *PoolOsdCmd) ShowPoolOsd(ret []byte){

	ps := PoolOsdShow{}
	err := json.Unmarshal(ret, &ps)
	if err != nil {
		fmt.Printf("unknow error %v\n", err)
		os.Exit(1)
	}
	for _, n := range ps.DataList {
		fmt.Printf("%4d %-10s %15s %15s\n",
			n.Id, n.PoolName, n.NodeName, n.OsdName)
	}
}

func init() {
	CmdOperationMap["pool osd"] = make(map[string] interface{})
	CmdOperationMap["pool osd"]["-"] = &PoolOsdCmd{}
	CmdOperationMap["pool osd"]["add"] = &PoolOsd{}
	CmdOperationMap["pool osd"]["del"] = &PoolOsd{}
	CmdOperationMap["pool osd"]["show"] = &PoolOsd{}
	CmdOperationMap["pool osd"]["list"] = &PoolOsdList{}
}