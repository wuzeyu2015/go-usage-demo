package common

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func IpValid(ip string) bool {
	return net.ParseIP(ip)!=nil
}

func CheckIP(ip string) {
	if !IpValid(ip) {
		fmt.Printf("Ip Address is invalid, %s\n", ip)
		os.Exit(1)
	}
}

var httpRetMsg = map[string]string {
	"RET_OK":"Operation successful\n",
	"RET_BAD_REQUEST":"Http request is wrong!\n",
	"RET_NO_EXIST":"The object doesn't exist!\n",
	"RET_EXIST":"The object already exist!\n",
	"RET_UNKNOWN_ERROR":"Internal error occurs.\n",
}

type ApiRetCode struct {
	RetCode string `json:"retCode"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Detail string `json:"detail"`
}

func GetRetMsg(rb []byte) string {
	var rc ApiRetCode
	err := json.Unmarshal(rb, &rc)
	if err!=nil {
		fmt.Printf("getRetMsg(): err = %v\n", err)
		return fmt.Sprintf("Unknow error code %s\n", string(rb))
	}

	if rc.Code != 0{
		if rc.Detail != ""{
			return fmt.Sprintf("%s (%s)", rc.Msg, rc.Detail)
		}else{
			return rc.Msg
		}
	}

	ret := rc.RetCode
	msg, ok := httpRetMsg[ret]
	if !ok {
		return fmt.Sprintf("Unknow error code %s\n", string(rb))
	}
	return msg
}