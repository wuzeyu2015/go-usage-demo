package main

import "unsafe"



type PgOsd struct {
	PoolId       int    	`json:"pool_id"`
	PgId       	 int    	`json:"pg_id"`
	NodeId       int    	`json:"node_id"`
	OsdId        int    	`json:"osd_id"`
	PgIndex      int    	`json:"pg_index"`
}


//sizeof返回大小固定？？
func main_sizeof(){
	//myMap := make(map[string]int, 0)
	var myMap map[string]int
	println(unsafe.Sizeof(myMap))
	myMap = make(map[string]int, 0)
	println(unsafe.Sizeof(myMap))
	myMap["sdf"] = 1
	println(unsafe.Sizeof(myMap))
	return

}

func main_map(){
	//myMap := make(map[string]int, 0)
	var myMap map[string]int
	if myMap == nil {
		println("myMap is nil")
	}
	println(len(myMap))
	myMap = make(map[string]int, 0)
	if myMap == nil {
		println("myMap is nil")
	}
	println(len(myMap))
	myMap["1"] = 1
	myMap["2"] = 2
	myMap["3"] = 3
	myMap["4"] = 4
	println(len(myMap))
	return

}

func main_nil(){
	//myMap := make(map[string]int, 0)
	var myInt []int
	//a := myInt[1]
	//myInt = append(myInt, a)
	if myInt == nil {
		println("myInt is nil")
	}
	pMyInt := &myInt
	println("pMyInt:", pMyInt)
	if *pMyInt == nil {
		println("*pmyInt is nil")
	}

	println(len(myInt))
	myInt = make([]int, 10)
	println("make a int:")
	println(len(myInt))
	for _, i := range myInt {
		print(i)
		println()
	}
	myInt = append(myInt, 1)
	myInt = append(myInt, 2)
	myInt = append(myInt, 3)

	println(len(myInt))
	return

}

//func main(){
//	pgOsd := PgOsd{
//		PoolId:  1,
//		PgId:    2,
//		NodeId:  3,
//		OsdId:   4,
//		PgIndex: 5,
//	}
//	pgOsdMap := make(map[int] []PgOsd)
//	pgOsdMap[0] = append(pgOsdMap[0], pgOsd)
//
//}