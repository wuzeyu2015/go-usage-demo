package cmd

import (
	"encoding/json"
	"fmt"
)

var redundancyToInt = map[string] int{
	"replicated": 0,
	"erasure": 1,
}

var operationToInt = map[string] int{
	"take": 0,
	"choose": 1,
	"emit": 2,
}

var classToInt = map[string] int{
	"hdd": 1,
	"ssd": 2,
	"ssd,hdd": 3,
}

var algToInt = map[string] int{
	"straw2": 0,
}

var hashToInt = map[string] int{
	"rjenkins1": 0,
}

var policyToInt = map[string] int{
	"firstn": 0,
	"indep": 1,
}

var failureDomainToInt = map[string] int{
	"osd":0,
	"host":1,
	"chassis":2,
	"rack":3,
	"row":4,
	"pdu":5,
	"pod":6,
	"room":7,
	"datacenter":8,
	"region":9,
	"root":10,
}

//just for user's reference and unmarlshal,do not edit this part in jason file
type FailureDomain struct {
	Osd 		int `json:"osd"`
	Host 		int `json:"host"`
	Chassis 	int `json:"chassis"`
	Rack 		int `json:"rack"`
	Row 		int `json:"row"`
	Pdu 		int `json:"pdu"`
	Pod 		int `json:"pod"`
	Room 		int `json:"room"`
	Datacenter 	int `json:"datacenter"`
	Region 		int `json:"region"`
	Root 		int `json:"root"`
}

type OSD struct {
	NodeName      string  `json:"node_name,omitempty"`
	OsdName       string  `json:"osd_name,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	Class         string  `json:"class,omitempty"`
}

type Bucket struct {
	FailureDomain string               `json:"failure_domain"`
	Alg           string                  `json:"alg,omitempty"`
	Hash          string                  `json:"hash,omitempty"`
	Bucket        map[string]*Bucket	`json:"bucket,omitempty"`
}

type Rule struct {
	Redundancy string       `json:"redundancy"`
	MinSize    int          `json:"min_size"`
	MaxSize    int          `json:"max_size"`
	Steps      []Step 		`json:"step"`
}

type Step struct {
	Operation     string    `json:"operation"`
	Root          string 	`json:"root"`
	Class         string    `json:"class"`
	FailureDomain string    `json:"failure_domain"`
	ChooseLeaf    bool   	`json:"choose_leaf"`
	Policy        string    `json:"policy"`
	NumRep        int    	`json:"num_rep"`
}

type Take struct {
	Operation     string    `json:"operation"`
	Root          string 	`json:"root"`
	Class         string    `json:"class"`
}

type Choose struct {
	Operation     string    `json:"operation"`
	FailureDomain string    `json:"failure_domain"`
	ChooseLeaf    bool   	`json:"choose_leaf"`
	Policy        string    `json:"policy"`
	NumRep        int    	`json:"num_rep"`
}

type Emit struct {
	Operation     string    `json:"operation"`
}

func (s Step) MarshalJSON() ([]byte, error) {
	if s.Operation == "take" {
		b, err := json.Marshal(Take{Operation: s.Operation, Root: s.Root, Class: s.Class})
		return b, err
	} else if s.Operation == "choose" {
		b, err := json.Marshal(Choose{Operation: 	s.Operation,
			FailureDomain: 	s.FailureDomain,
			ChooseLeaf: 	s.ChooseLeaf,
			Policy: 		s.Policy,
			NumRep: 		s.NumRep})
		return b, err
	} else if s.Operation == "emit" {
		b, err := json.Marshal(Emit{Operation: s.Operation})
		return b, err
	}
	return []byte{}, nil
}

func fillBucketString(bucket *map[string]*Bucket){
	for k, v := range *bucket {
		//只有叶子节点才没有hash和alg
		if v.FailureDomain == "osd"{
			(*bucket)[k].Hash = ""
			(*bucket)[k].Alg = ""
		}
		fillBucketString(&v.Bucket)
	}
}

func (C *Crush) crushParamsCheck() (bool, string){
	//max_retries, max_local_retries
	if C.MaxRetries == 0 || C.MaxLocalRetries == 0 {
		return false, "retries should not set 0."
	}
	//Osd
	for _, osd := range C.Osd {
		if osd.Class != "ssd" && osd.Class != "hdd" && osd.Class != "ssd,hdd" {
			return false, "osd class illegal."
		}
	}
	//bucket
	err, res := bucketParamCheck(&C.Bucket)
	if err == false {
		return false, res
	}
	//rules
	err, res = ruleParamCheck(&C.Rule)
	if err == false {
		return false, res
	}
	return true, ""
}

func bucketParamCheck(bucket *map[string]*Bucket) (bool, string){
	for k, v := range *bucket {

		if _, ok := failureDomainToInt[v.FailureDomain]; !ok {
			return false, fmt.Sprintf("bucket %s failure domain illegal\n", k)
		}
		if v.FailureDomain == "osd"{
			if _, ok := hashToInt[(*bucket)[k].Hash]; ok {
				return false, fmt.Sprintf("osd bucket %s too much parameters\n", k)
			}
			if _, ok := algToInt[(*bucket)[k].Alg]; ok {
				return false, fmt.Sprintf("osd bucket %s too much parameters\n", k)
			}
		} else {
			if _, ok := hashToInt[(*bucket)[k].Hash]; !ok {
				return false, fmt.Sprintf("osd bucket %s hash illegal\n", k)
			}
			if _, ok := algToInt[(*bucket)[k].Alg]; !ok {
				return false, fmt.Sprintf("osd bucket %s alg illegal\n", k)
			}
		}
		if err, res := bucketParamCheck(&v.Bucket); err == false {
			return false, res
		}
	}
	return true, ""
}

func ruleParamCheck(rule *map[string]*Rule) (bool, string){
	for _, v := range *rule {
		if v.Redundancy != "replicated" && v.Redundancy != "erasure"{
			return false, "rule illegal redundancy."
		}
		for _, oper := range v.Steps {
			if oper.Operation != "take" && oper.Operation != "choose" && oper.Operation != "emit" {
				return false, "step illegal Operation."
			}
			if oper.Operation == "take" && (oper.Class != "ssd" && oper.Class != "hdd" && oper.Class != "ssd,hdd") {
				return false, "class illegal for take operation."
			}
			if oper.Operation == "choose" {
				_, failureDomain := failureDomainToInt[oper.FailureDomain]
				if failureDomain == false {
					return false, "failure domain illegal for choose operation."
				}
				_, policy := policyToInt[oper.Policy]
				if policy == false {
					return false, "policy illegal for choose operation."
				}
			}
			if oper.Operation == "emit" {
				if oper.FailureDomain != "" || oper.Policy != "" || oper.NumRep != 0 {
					return false, "false parameters for emit operation."
				}
			}
		}
	}
	return true, ""
}