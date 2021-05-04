package main


func main(){
	//一个空指针，无法被解引用
	var pMyInt *([]int)
	if pMyInt == nil {
		println("var pMyInt is nil")
	}
	//if *pMyInt == nil {
	//	println("var *pMyInt is nil")
	//}
	//一个空Int，被指针解引用
	var myInt []int
	if myInt == nil {
		println("var myInt is nil")
	}
	pMyInt = &myInt
	if *pMyInt == nil {
		println("pMyInt can point to a nil myInt")
	}
	return

}