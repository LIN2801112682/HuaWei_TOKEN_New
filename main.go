package main

import (
	"build_VToken_index"
	"build_dictionary"
	"fmt"
	"precise_query"
	"runtime"
)

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func main() {
	fmt.Println("字典树D：===============================================================")
	fmt.Println("字典树D内存占用大小：")
	//traceMemStats()
	fmt.Println()
	root := build_dictionary.GererateTree("src/resources/500Dic.txt", 1, 3, 10) //
	fmt.Println()
	//traceMemStats()
	fmt.Println()

	fmt.Println("索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	traceMemStats()
	fmt.Println()
	_, indexTreeNode := build_VToken_index.GererateIndex("src/resources/100Index.txt", 1, 3, root) //
	fmt.Println()
	traceMemStats()
	fmt.Println()

	/*fmt.Println("新增索引后的索引项集：===============================================================")
	fmt.Println()
	fmt.Println("索引项集内存占用大小：")
	//traceMemStats()
	fmt.Println()
	go_dic.AddIndex("src/resources/add2000.txt", 1, 3, root, indexTree)
	fmt.Println()
	//traceMemStats()
	fmt.Println()*/

	resInt := precise_query.MatchSearch("english", root, indexTreeNode, 1, 3) //get english venues
	fmt.Println(resInt)
}
