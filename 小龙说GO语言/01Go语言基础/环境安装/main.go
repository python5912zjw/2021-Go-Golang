package main

import "fmt"

func main(){
	fmt.Println("Hello World!")
}

//这个面向对象，先不用管，主要是我用来讲解go mod模式
//小龙说GO语言文件夹其实就是一个module，
//大家可以在cmd切换到小龙说GO语言文件夹目录执行go mod why
//module是xiaolongtalkgo，这个是基准，以后所有的包引入都仰仗这个基准
type Test struct{

}

//这个是面向对象，先不用管，主要是我用来讲解go mod模式
func(test *Test) Test1(){
	fmt.Println("test")
}
