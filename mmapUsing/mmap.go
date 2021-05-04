package main


type Node struct {
	NodeName string `json:"nodeName"`
	IpAddr1  string `json:"ipAddr1"`
	IpAddr2  string `json:"ipAddr2"`
	Id       int    `json:"id"`
	Active   int    `json:"active"`
}

func main() {

	var CmdOperationMap = make(map[string] map[string] int)

	for i, _ := range CmdOperationMap {
		CmdOperationMap[i] = make(map[string]int)
	}
	//CmdOperationMap["a"]["b"] = 55

	m:=make(map[string]map[string]int)
	c:=make(map[string]int)
	c["b"]=1
	m["a"]=c
	m["a"]["b"] = 55

}

func init() {

	//for i, _ := range CmdOperationMap {
	//	CmdOperationMap[i] = make(map[string] interface{})
	//}
	//CmdOperationMap["node"]["add"] = &Node{}
}