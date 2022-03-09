package main

import (
	"fmt"
	"sync"
)

func main() {
	// 注意map是无序的
	// 两种赋值操作
	//var countryCapitalMap map[string]string /*创建集合 */
	// 注意,最好使用下面这种来进行初始化,因为map属于引用类型,上面这种定义未分配空间
	var countryCapitalMap = make(map[string]string, 200) // 创建时可以选择初始容量来使其性能更高
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	/*使用键输出地图值 */
	// 使用key来遍历map
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	// range直接遍历
	for country, capital := range countryCapitalMap {
		fmt.Println(country, "首都是", capital)
	}
	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
	// 去除字典里的键值
	delete(countryCapitalMap, "France")
	fmt.Println(countryCapitalMap)

	// 清空map没有专门的函数 所以可以重新make一个
	//countryCapitalMap = make(map[string]string)

	/*线程不安全的map*/
	/*	m := make(map[int]int)
		// 不停的写
		go func() {
			for {
				m[100] = 100
			}

		}()

		// 不停的读取
		go func() {
			for {
				_ = m[100]
			}

		}()

		for {

		}*/

	/*线程安全的map*/
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	// map的引用传递
	test03(countryCapitalMap)
	fmt.Println("==========")
	fmt.Println(countryCapitalMap)

}

func test03(map1 map[string]string) {
	map1["France"] = "1233313"
}
