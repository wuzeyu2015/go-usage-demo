package cmd

import (
	"encoding/json"
	"fmt"
	"log"
)

type CrushCmd struct {

}


type Crush struct {
	MaxRetries      uint               `json:"max_retries"`
	MaxLocalRetries uint               `json:"max_local_retries"`
	FailureDomain   FailureDomain      `json:"failure_domain_root"`
	Osd             []OSD              `json:"osd"`
	Bucket          map[string]*Bucket `json:"bucket"`
	Rule            map[string]*Rule   `json:"rule"`
}

type CrushDump struct {
	RetCode string `json:"retCode"`
	Data     Crush   `json:"data"`
	Code	int 	`json:"code"`
}

func (c *CrushCmd)DumpCrash(ret []byte) {

	rsp := &CrushDump{}
	err := json.Unmarshal(ret, rsp)
	if err != nil {
		log.Fatal(err)
	}

	//fillStepString(rsp.Data.Rule)
	fillBucketString(&rsp.Data.Bucket)
	data, err := json.MarshalIndent(rsp.Data, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Print(string(data))
}


func init() {

	CmdOperationMap["crush"] = make(map[string] interface{})
	CmdOperationMap["crush"]["-"] = &CrushCmd{}
	CmdOperationMap["crush"]["dump"] = &CrushDump{}
	CmdOperationMap["crush"]["load"] = &Crush{}
}