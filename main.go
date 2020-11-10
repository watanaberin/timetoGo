package main

import(
	"fmt"
)
func main()  {
	fmt.Println("Hello")
	var x int=20 //定义
	y :=100 //自动推断类型 用于函数内部
	if y>100{ 
		fmt.Println(y)
	}else if y<=100 {
		fmt.Println(x)
	}
	// switch x {
	// case x > 0:
	// 	fmt.Println(x)
	// case x < 0:
	// 	fmt.Println(-x)
	// default: 
	// 	fmt.Println(0)
	// }
	finger :=3
	switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
	default:
		fmt.Println("sorry")
	}

	x1 :=2
	for x1<5{
		fmt.Println(x1)
		x1++
	}
	for {
		fmt.Println(x1)
		x1--
		if x1==0{
			break
		}
	}


	x2:=[3]int{100,220,300}
	for i,n :=range x2{
		fmt.Println(i,":",n)
	}
}