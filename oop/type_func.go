package main

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	z := x * y
	return z
}

func div(x, y int) int {
	q := x / y
	return q
}

// 用type定义一个函数抽象类型
type opFunc func(int, int) int

// op函数接受两个init和一个函数抽象类型作为参数
func op(x, y int, f opFunc) int {
	return f(x, y)
}

// op2函数接受两个init和一个函数作为参数
func op2(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func main() {
	println(add(1, 2))
	println(sub(1, 2))
	println(mul(1, 2))
	println(div(1, 2))

	println(op(1, 2, add))
	println(op(1, 2, sub))
	println(op(1, 2, mul))
	println(op(1, 2, div))

	println(op2(1, 2, add))
	println(op2(1, 2, sub))
	println(op2(1, 2, mul))
	println(op2(1, 2, div))
}
