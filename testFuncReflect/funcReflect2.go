package main

import "fmt"
import "reflect"
import "encoding/xml"

type ssssst struct{
}

func (this *ssssst)Echo(){
	fmt.Println("echo()")
}

func (this *ssssst)Echo2(str string){
	fmt.Println("Echo2 ", str)
}

var xmlstr string=`<root>
<func>Echo</func>
<func>Echo2</func>
</root>`

type st2 struct{
	E []string `xml:"func"`
}

func main() {
	s2 := st2{}
	xml.Unmarshal([]byte(xmlstr), &s2)

	s := &ssssst{}
	v := reflect.ValueOf(s)

	v.MethodByName(s2.E[1]).Call(nil)
}