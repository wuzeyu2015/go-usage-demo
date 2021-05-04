package main

import (
	"configurableCmd/cmd"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)


func main() {
	file, err := os.OpenFile("cmdConfigure.json", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	byte, err := ioutil.ReadAll(file)
	if err != nil {
		print(err)
		return
	}

	if err = json.Unmarshal(byte, cmd.Rules); err != nil {
		print(err)
		return
	}
	//serialize, err := json.MarshalIndent(cmdRules, "", " ")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//println(string(serialize))


	//running cmd...
	cmd.Processer.CliHandler(os.Args[1:])
	return
}

