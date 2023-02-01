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
type OpFunc func(int, int) int

// op函数接受两个init和一个函数抽象类型作为参数
func op(x, y int, f OpFunc) int {
	return f(x, y)
}

// OpFunc作为接收者，定义一个call方法
func (f OpFunc) call(x, y int) int {
	return f(x, y)
}

// 接收三个OpFunc作为参数,参数分别为total left和right，返回一个func(int, int, int, int) int
// 计算过程为total(left(left1, left2), right(right1, right2))
func (f OpFunc) binaryOpFunc(total OpFunc, left OpFunc, right OpFunc) func(left1 int, left2 int, right1 int, right2 int) int {
	return func(left1 int, left2 int, right1 int, right2 int) int {
		return total(left(left1, left2), right(right1, right2))
	}
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

	println(OpFunc(add).call(1, 2))
	println(OpFunc(sub).call(1, 2))
	println(OpFunc(mul).call(1, 2))
	println(OpFunc(div).call(1, 2))

	// 21
	println(OpFunc(nil).binaryOpFunc(mul, add, add)(1, 2, 3, 4))

	println(op2(1, 2, add))
	println(op2(1, 2, sub))
	println(op2(1, 2, mul))
	println(op2(1, 2, div))

}
