/*
温度转换
*/
package tempconv

import "fmt"

//type就是c中的typedef,作为类型别名使用
//在这个程序来说,虽然摄氏温度和华氏温度的底层类型都是float64
//但是从实际意义上来说,他们是不能做比较的,所以我们需要定义两个类型
type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
