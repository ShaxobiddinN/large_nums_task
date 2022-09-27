package main

import (
	"fmt"
	"github.com/ShaxobiddinN/large_nums_task/bigint"
)

func main(){
	a,err :=bigint.NewInt("-048")
	if err != nil{
		panic(err)
	}
	b, err :=bigint.NewInt("-010")
	if err != nil{
		panic(err)
	}
	// err = b.Set("1") //b = "1"
	// if err != nil{
	// 	panic(err)
	// }

	c := bigint.Add(a,b)
	d := bigint.Sub(a,b)
	e := bigint.Multiply(a,b)
	f := bigint.Mod(a,b)
	l :=b.Abs()

	fmt.Println("Add two numbers",c)
	fmt.Println("Sub two numbers",d)
	fmt.Println("Multiply two numbers",e)
	fmt.Println("Mod two number",f)
	fmt.Println("Abs num",l)


}