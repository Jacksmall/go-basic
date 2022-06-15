package main

import "fmt"

func main() {
	// map是一个存储键值对的无序集合 （底层是散列表）
	// 创建和初始化
	// data := make(map[string]int)
	data := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(data)
	// 不能将切片作为map的键，因为切片｜函数这些具有引用语义
	// data2 := map[[]string]int{}
	// 可以将切片作为map的值
	data2 := map[string][]string{
		"a": {"a", "b"},
		"b": {"b", "c"},
		"c": {"c", "d"},
	}
	fmt.Println(data2)
	// 使用map
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	fmt.Println(colors)
	// 测试某个映射里面是否已经还有某个值
	value, exists := colors["Red"]
	if exists {
		fmt.Printf("value: %s exists\n", value)
	}
	// 数组、切片、映射 range
	for k, v := range data2 {
		fmt.Printf("data2 key: %s, value: %s\n", k, v)
	}
	// 映射的删除 delete
	delete(data2, "a")
	for k, v := range data2 {
		fmt.Printf("data2 key: %s, value: %s\n", k, v)
	}

	// 数据在不同函数间传递都是通过值传递的，数组的传递是制造副本传递，应该改为切片传递，因为切片是指向数组的指针
	// 长度&容量

	// 在函数间传递映射，并不会制造出一个副本，当某个函数对这个映射做了修改，所有对这个映射的引用都会察觉

	// int64 占 8 字节(64位)
	// float32 占 4 字节(32位)
	// bool 占 1 字节(8位)
}
