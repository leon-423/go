package main

//var a int = 100
//var b int = 200

//func main()  {
//	fmt.Print("测试导入包")
//	fmt.Print(a,b)
//}

/*匿名变量*/
/*func getData() (int,int)  {
	return 200,400
}

func get()(int,int)  {
	return 100,200
}

func main()  {
	a, _ := get()
	_, b := get()
	println(a,b)
}*/

/*变量的作用域*/
//全局变量 d
/*var d int = 10

//局部变量和全局变量名称可以相同，优先取局部变量的值
var a int = 222;

func main()  {
	//局部变量 a b
	var a int = 3
	var b int = 4

	c := a + b
    d := 2 * d
	fmt.Printf("a = %d, b = %d, c = %d",a,b,c)
    fmt.Println("=======分割线=======")
	fmt.Println(d)
}*/

/*浮点类型*/
//小数类型
/*var f float32 = 16777214;

func main()  {
	//fmt.Println(f)
	fmt.Println(math.Pi)
}
*/

/*复数类型*/
//Go语言拥有两种复数类型，分别是 complex64（32 位实数和虚数）类型和 complex128（64 位实数和虚数）类型。
