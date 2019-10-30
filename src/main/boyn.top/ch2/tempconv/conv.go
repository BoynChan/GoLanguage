package tempconv

func CToF(c Celsius) Fahrenheit {
	//创建一个新的Fahrenheit类型,这是类型转换的操作而不是调用一个函数
	//类型转换并不会改变值本身,但是会使其语义发生一定的变化
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
