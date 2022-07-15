/**
 * 装饰器模式?  怎么感觉和组合模式一样
 */
package main

import "fmt"

func main() {
	logger := Logger{}
	authorization := Authorization{}
	dt := Decorator{}
	dt.dt = append(dt.dt, logger)
	dt.dt = append(dt.dt, authorization)

	fmt.Println(dt.action())

	dt2 := Decorator{}
	dt2.dt = append(dt2.dt, logger)
	// dt 也实现了DecoratorInterface 接口方法，所以可以作为dt类型加入到切片中
	dt2.dt = append(dt2.dt, dt)

	fmt.Println(dt2.action())
}
