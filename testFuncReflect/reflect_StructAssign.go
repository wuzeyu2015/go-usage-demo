package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Node struct {
	NodeName string `json:"nodeName"`
	IpAddr1  string `json:"ipAddr1"`
	IpAddr2  string `json:"ipAddr2"`
	Id       int    `json:"id"`
	Active   int    `json:"active"`
}


//1.通过tag反射
//将结构体里的成员按照json名字来赋值
func SetStructFieldByJsonName(ptr interface{}, fields map[string]interface{}) {
	println("fields:", fields)

	v := reflect.ValueOf(ptr).Elem() // the struct variable

	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")

		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]
		println("JSONnAME:", name)

		if value, ok := fields[name]; ok {

			println("fieldInfo.Name:", fieldInfo.Name)
			//给结构体赋值
			//保证赋值时数据类型一致
			println("类型1：", reflect.ValueOf(value).Type(), "类型2：", v.FieldByName(fieldInfo.Name).Type())
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}

	return
}
//2.通过结构体字段名称进行反射
// 通过反射，对user进行赋值
type user struct{
	name string
	age int
	feature map[string]interface{}
}

func main() {
	var u interface{}
	u=new(Node)

	fmt.Println(u)
}

//func main() {
//	var u interface{}
//	u=new(user)
//	value:=reflect.ValueOf(u)
//	if value.Kind()==reflect.Ptr{
//		elem:=value.Elem()
//		name:=elem.FieldByName("name")
//		if name.Kind()==reflect.String{
//			*(*string)(unsafe.Pointer(name.Addr().Pointer())) = "fangwendong"
//		}
//
//		age:=elem.FieldByName("age")
//		if age.Kind()==reflect.Int{
//			*(*int)(unsafe.Pointer(age.Addr().Pointer())) =24
//		}
//
//		feature:=elem.FieldByName("feature")
//		if feature.Kind()==reflect.Map{
//			*(*map[string]interface{})(unsafe.Pointer(feature.Addr().Pointer())) =map[string]interface{}{
//				"爱好":"篮球",
//				"体重":60,
//				"视力":5.2,
//			}
//		}
//
//	}
//
//	fmt.Println(u)
//}

