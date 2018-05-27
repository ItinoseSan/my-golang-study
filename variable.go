package main

import "fmt"
// Define variable
var x int

func test(){
	return x
}

func DefineVariableInFunction(){
	a := 1
	b := "String"
}

const A = 1
const B = 1
//  const 宣言内で、定数ジェネレータ iota を使うと 0からの列挙値(int型)を生成できる。

const (
	_ = iota
	TOKEN1
	TOKEN2
	TOKEN3
)