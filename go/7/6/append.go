package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Println(a, b)
	//1 b... 将b中元素追加到a后面
	a = append(a, b...)
	fmt.Println(a, b)

	//2 将b中元素copy至新切片c
	c := make([]int, len(b))
	n := copy(c, b)
	fmt.Println(n, c, b)

	d := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(d)

	//3 删除索引i(3)的元素
	//d = append(d[:3],d[4:]...)

	//4 删除i-j(3,4)的元素
	//d = append(d[:3],d[5:]...)

	fmt.Println(d)

	//5 为切片 e 扩展 j(5) 个元素长度
	e := []int{0, 1, 2}
	fmt.Println(cap(e))
	e = append(e, make([]int, 5)...)
	fmt.Println(cap(e))

	//6 在索引i(4)的位置插入x(ee)元素
	f := []string{"aa", "bb", "cc", "dd", "ff", "gg"}
	fmt.Println(f)
	f = append(f[:4], append([]string{"ee"}, f[4:]...)...)
	fmt.Println(f)

	// 7 在索引为i的位置插入长度为j的新切片
	//在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	j := []int{400, 401, 402, 403, 404}
	fmt.Println(i, j)
	i = append(i[:5], append(j, i[5:]...)...)
	fmt.Println(i, j)

	//8 在索引为i(4)的位置 插入长度为j(3)的新切片
	g := []string{"aa", "bb", "cc", "dd", "hh", "ii"}
	h := []string{"ee", "ff", "gg"}
	fmt.Println(g, h)
	g = append(g[:4], append(h, g[4:]...)...)
	fmt.Println(g, h)

	//9 取出k末尾的元素
	k := []int{0, 1, 2, 3, 4, 5, 6}
	var l int
	fmt.Println(l, k)
	l, k = k[len(k)-1], k[:len(k)-1]
	fmt.Println(l, k)

	//10 将元素6追加到切片k
	k = append(k, 6)
	fmt.Println(k)

}
