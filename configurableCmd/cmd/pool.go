package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type PoolCmd struct {

}

type Pool struct {
	PoolName string `json:"poolName"`
	PgNum    int    `json:"pgNum"`
	Id       int    `json:"id"`
	Cap      uint64 `json:"cap"`
	Type     int    `json:"type"`
	Active   int    `json:"active"`
	CrushRule   string `json:"crush_rule"`
}

type PoolList struct {
	RetCode string `json:"retCode"`
	DataList []Pool `json:"data"`
	Code	int 	`json:"code"`
}

type PoolShow struct {
	RetCode string `json:"retCode"`
	Data     Pool   `json:"data"`
	Code	int 	`json:"code"`
}


func (p *PoolCmd)ListPool(ret []byte){

	var nl PoolList
	err := json.Unmarshal(ret, &nl)
	if err != nil {
		fmt.Printf("decode error, %v\n", err)
		os.Exit(1)
	}
	for _, n := range nl.DataList {
		fmt.Printf("%4d %-10s %5d %-4s %10d %d\n",
			n.Id, n.PoolName, n.PgNum, getType(n.Type),
			n.Cap, n.Active)
	}
}

//func PoolBindBuffer(o *OperEntity, args []string) {
//
//	//and bind the parameter to buffer
//	if o.Name == "add" {
//		newPool := &Pool{}
//		Buffer = newPool
//		newPool.PoolName = args[0]
//		newPool.PgNum, _ = strconv.Atoi(args[1])
//		newPool.Type, _ = TypeMap[args[2]]
//		cap, _ := strconv.Atoi(args[3])
//		newPool.Cap = uint64(cap)
//		newPool.Active = 1
//		if o.ArgsActualLen == 5 {
//			newPool.CrushRule = args[4]
//		}
//	}
//	if o.Name == "del" {
//		delPool := &Pool{}
//		Buffer = delPool
//		delPool.PoolName = args[0]
//	}
//	if o.Name == "update" {
//		updatePool := &Pool{}
//		Buffer = updatePool
//		updatePool.PoolName = args[0]
//		updatePool.PgNum, _ = strconv.Atoi(args[1])
//		updatePool.Type, _ = TypeMap[args[2]]
//		cap, _ := strconv.Atoi(args[3])
//		updatePool.Cap = uint64(cap)
//		updatePool.Active = 1
//		if o.ArgsActualLen == 5 {
//			updatePool.CrushRule = args[4]
//		}
//	}
//	if o.Name == "list" {
//		Buffer = &PoolList{}
//		Display = ListPool
//	}
//	if o.Name == "show" {
//		showPool := &Pool{}
//		showPool.PoolName = args[0]
//		Buffer = showPool
//		Display = ShowPool
//	}
//}

func (p *PoolCmd)ShowPool(ret []byte){

	var rn PoolShow
	err := json.Unmarshal(ret, &rn)
	if err != nil {
		fmt.Printf("unknow error %v\n", err)
		os.Exit(1)
	}
	pool := &rn.Data
	if len(pool.PoolName) != 0 {
		fmt.Printf("%-10s: %d\n", "Id", pool.Id)
		fmt.Printf("%-10s: %s\n", "Name", pool.PoolName)
		fmt.Printf("%-10s: %d\n", "PgNum", pool.PgNum)
		fmt.Printf("%-10s: %s\n", "Type", getType(pool.Type))
		fmt.Printf("%-10s: %d\n", "Cap", pool.Cap)
		fmt.Printf("%-10s: %d\n", "Active", pool.Active)
	}
}


func getType(tp int) string {
	if tp == 1 {
		return "3rep"
	}
	return "ec"
}

var TypeMap = map[string] int{"3rep": 1, "ec": 2}

func init() {
	CmdOperationMap["pool"] = make(map[string] interface{})
	CmdOperationMap["pool"]["-"] = &PoolCmd{}
	CmdOperationMap["pool"]["add"] = &Pool{}
	CmdOperationMap["pool"]["del"] = &Pool{}
	CmdOperationMap["pool"]["update"] = &Pool{}
	CmdOperationMap["pool"]["show"] = &Pool{}
	CmdOperationMap["pool"]["list"] = &PoolList{}
}