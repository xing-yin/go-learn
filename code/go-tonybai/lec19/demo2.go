package main

/**
 * @Author: Alan Yin
 * @Date: 2022/5/4
 * @Desc:
 */

func main() {
	//var sl = []int{1, 2, 3, 4, 5}

	//for i := 0; i < len(sl); i++ {
	//	fmt.Printf("sl[%d] = %d\n", i, sl[i])
	//}

	//for i, v := range sl {
	//	fmt.Printf("sl[%d] = %d\n", i, v)
	//}

	//for i := range sl {
	//	fmt.Printf("sl[%d]\n", i)
	//}

	//for _, v := range sl {
	//	fmt.Printf("%d\n", v)
	//}

	//var s = "中国人"
	//for i, v := range s {
	//	fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	//}

	var m = map[string]int{
		"Alan":  21,
		"Bob":   22,
		"Cindy": 23,
	}

	for k, v := range m {
		println(k, v)
	}

	//var c = make(chan int)
	//for v := range c {
	//	fmt.Println(v)
	//}
}
