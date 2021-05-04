package cmd

import (
	"configurableCmd/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type CmdRules struct {
	Version string   `json:"version"`
	Desc string      `json:"desc"`
	Copyright string `json:"copyright"`
	Exec string      `json:"exec"`
	Cmds []CmdEntity `json:"cmds"`
}

type CmdEntity struct {
	Name string             `json:"name"`
	Property int 			`json:"property"`
	Desc []string           `json:"desc"`
	Operations []OperEntity `json:"operations,omitempty"`
	Exec string             `json:"exec,omitempty"`
}

type OperEntity struct {
	Name string `json:"name"`
	Desc []string `json:"desc"`
	ArgsLen []int `json:"args_len"`
	ArgsActualLen int `json:"args_actual_len"`
	Args []struct {
		Arg string `json:"arg"`
		Desc []string `json:"desc"`
		Type []string `json:"type"`
		Limit string `json:"limit,omitempty"`
		Size int `json:"size,omitempty"`
		Optional bool `json:"optional,omitempty"`
	} `json:"args"`
	Exec string `json:"exec"`
	//contain PUT or POST data for a certain operation
	sndBuffer interface{}
	//contain PUT or POST data for a certain operation
	rcvBuffer interface{}
	//show the return information for a certain operation
	Display reflect.Value
	//Display func(ret []byte)
	Api struct{
		Method string `json:"method"`
		Path   string `json:"path"`
	} `json:"api"`
}

var Rules = &CmdRules{}

var Processer = CmdProcesser{
	rule: Rules,
}

type CmdProcesser struct {
	rule *CmdRules
}

//return a configured cmd entity.
func (p *CmdProcesser)cmdTraveler(args []string, subCmd bool) (*CmdEntity, []string) {

	for _, cmd := range p.rule.Cmds {
		if subCmd == true {
			if args[0] + " " + args[1] == cmd.Name {
				return &cmd, args[2:]
			}
		} else {
			if args[0] == cmd.Name {
				//find cmd, need check if a subcommand followed.
				if cmd.Property == 1 && len(args) > 1 {
					var updatedCmd *CmdEntity = nil
					updatedCmd, args = p.cmdTraveler(args, true)
						if updatedCmd != nil {
							return updatedCmd, args
						}
				}
				return &cmd, args[1:]
			}
		}
	}
	return nil, args
}

//return a configured operation entity.
func (c *CmdEntity)operTraveler(args []string) (*OperEntity, []string) {

	if len(args) == 0 {
		fmt.Println(c.Desc)
		return nil, []string{}
	}
	for _, oper := range c.Operations {
		if args[0] == oper.Name {
			return &oper, args[1:]
		}
	}
	fmt.Println(c.Desc)
	return nil, []string{}
}


//check exists of a cmd.
func (p *CmdProcesser)cmdCheck(name string) bool {

	for _, cmd := range p.rule.Cmds {
		if name == cmd.Name {
			return true
		}
	}
	return false
}

func (p *CmdProcesser)CliHandler(args []string) string {
	if len(args) == 0 {
		fmt.Println(p.rule.Desc)
		log.Fatal("no cmd found...")
		return ""
	}
	var cmdEntity *CmdEntity = nil
	cmdEntity, args = p.cmdTraveler(args, false)
	if cmdEntity == nil {
		return p.rule.Desc
	}
	var operEntity *OperEntity = nil
	operEntity, args = cmdEntity.operTraveler(args)
	if operEntity == nil {
		log.Fatal("terminate..")
	}
	operEntity.bindBuffer(args, cmdEntity.Name)
	operEntity.Excute()
	return ""
}

//excute operation.
func (o *OperEntity)Excute() {

	o.rcvBuffer = common.HttpSend(o.Api.Method, o.Api.Path, o.sndBuffer)

	//get method need to display return value.
	if o.Api.Method == "GET" {
		o.Display.Call([]reflect.Value{reflect.ValueOf(o.rcvBuffer)})
	}

	fmt.Printf("%s\n", common.GetRetMsg(o.rcvBuffer.([]byte)))
	return
}

func (o *OperEntity)ParamsLen(argslen int) bool{

	for _, len := range o.ArgsLen {
		if argslen == len {
			o.ArgsActualLen = argslen
			return true
		}
	}
	return false
}

func (o *OperEntity)ParamsConstraints(args []string) bool{

	for i, arg := range args {
		Match(o.Args[i].Type[0], arg)
	}
	return true
}


//read the parameters behind operation
func (o *OperEntity)bindBuffer(args []string, cmdName string) {

	//parameters valid check
	if err := o.ParamsLen(len(args)); err != true {
		log.Fatal(o.Desc)
		return
	}
	//parameters valid check
	if err := o.ParamsConstraints(args); err != true {
		log.Fatal(o.Desc)
		return
	}
	//bindBuffer
	o.sndBuffer = CmdOperationMap[cmdName][o.Name]
	//assign buffer
	o.setBufferWithRules(o.sndBuffer, args)

	//inputs:= make([]reflect.Value, len(args))
	//for i, _:= range args {
	//	inputs[i] = reflect.ValueOf(args[i])
	//}
	o.Display = reflect.ValueOf(CmdOperationMap[cmdName]["-"]).MethodByName(o.Exec)



	//switch cmdName {
	//case "node":
	//	NodeBindBuffer(o, args)
	//case "pool":
	//	PoolBindBuffer(o, args)
	//case "pool osd":
	//	PoolOsdBindBuffer(o, args)
	//	//todo:bind other command
	//default:
	//
	//}
}

func Match(typ string, arg string) {
	l := len(arg)
	regString := regexp.MustCompile(fmt.Sprintf("[0-9a-zA-Z_]{%d}", l))

	switch typ {
	case "string":
		if !regString.Match([]byte(arg)) {
			fmt.Printf("arg is invalid, %s\n", arg)
			os.Exit(1)
		}
	case "ipV4":
		common.CheckIP(arg)
	default:
	}
}


func (o *OperEntity) setBufferWithRules(ptr interface{}, args []string) {

	if o.Name == "load" {
		jsonFile, err := os.OpenFile(args[0], os.O_RDONLY, 0755)
		defer jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		fileStream, _ := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err)
		}
		//map structure auto erase duplicated key.
		//etcd transaction will prevent duplicated key falsely set in slice.
		err = json.Unmarshal(fileStream, ptr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		v := reflect.ValueOf(ptr).Elem() // the struct variable

		for i := 0; i < v.NumField() && i < o.ArgsActualLen; i++ {
			fieldInfo := v.Type().Field(i) // a reflect.StructField
			tag := fieldInfo.Tag           // a reflect.StructTag
			name := tag.Get("json")

			if name == "" {
				name = strings.ToLower(fieldInfo.Name)
			}
			//find json tag.
			name = strings.Split(name, ",")[0]
			println("JSONnAME:", name)

			//operation arg is matched with struct json tag.
			if o.Args[i].Arg == name /*&& o.Args[i].Type[0] ==*/ {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(args[i]))
			}
		}
	}

	return
}