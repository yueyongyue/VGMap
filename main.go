package main

import (
	"fmt"
	"math/rand"
	"os"
	"unsafe"
)

type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}

type mapextra struct {
	overflow     *[]*bmap
	oldoverflow  *[]*bmap
	nextOverflow *bmap
}

type bmap struct {
	tophash [8]uint8
	keys    [8]int
	values  [8]int
	//pad      uintptr
	overflow *bmap
}

const (
	empty        = 0 // cell is empty
	minTopHash   = 4 // minimum tophash for a normal filled cell.
	sameSizeGrow = 8 // the current map growth is to a new map of the same size
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func (h *hmap) sameSizeGrow() bool {
	return h.flags&sameSizeGrow != 0
}

// 排空条件： 0 < b.tophash[0] < 4， 表示排空
func evacuated(b *bmap) bool {
	h := b.tophash[0]
	return h > empty && h < minTopHash
}

func main() {
	mp := make(map[int]int, 0)
	var choice int
	for {
		var scope int
		var key int
		var value int
		fmt.Println("#########################")
		fmt.Println("    1：输入key和value")
		fmt.Println("    2：随机生成一个Map")
		fmt.Println("    3：打印Map")
		fmt.Println("    4：清空Map")
		fmt.Println("    5：退出")
		fmt.Println("#########################")
		fmt.Println("请选择(1,2,3,4,5):")
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Println("输入的值不正确请重新输入...")
		}
		switch choice {
		case 1:
			fmt.Println("请输入key:")
			if _, err := fmt.Scanln(&key); err != nil {
				fmt.Println("输入的key必须是整数")
				choice = 1
				continue
			}
			fmt.Println("请输入value:")
			if _, err := fmt.Scanln(&value); err != nil {
				fmt.Println("输入的value必须是整数")
				choice = 1
				continue
			}
			mp[key] = value
			vmap(mp)
		case 2:
			fmt.Println("请输入范围最大值:")
			if _, err := fmt.Scanln(&scope); err != nil {
				fmt.Println("输入的范围最大值必须是整数")
				choice = 2
				continue
			}
			mp = make(map[int]int, 0)
			for i := 0; i < scope; i++ {
				mp[i] = i + rand.Intn(scope)
			}
			vmap(mp)
		case 3:
			vmap(mp)
		case 4:
			for k := range mp {
				delete(mp, k)
			}
			fmt.Println("清空成功")
		case 5:
			os.Exit(0)
		default:
			fmt.Println("输入的值不正确请重新输入...")
		}
	}
}

func vmap(mp map[int]int) {
	h := **(**hmap)(unsafe.Pointer(&mp))
	bucketSize := unsafe.Sizeof(bmap{})
	m := 1 << h.B // 1 << B 相当于 2的B次幂
	//fmt.Println("hmap.B:", h.B)
	//fmt.Println("bmap array:", m)
	//fmt.Println("bmap oldbuckets:", h.oldbuckets)

	// 判断buckets是否有值
	if c := h.oldbuckets; c != nil {
		// 判断当前map是否增长到相同大小的新map
		if !h.sameSizeGrow() {
			// 不是的话，buckets数量减半
			m >>= 1
		}
		for i := 0; i < m; i++ {
			t := "\t"
			oldb := *(*bmap)(add(c, uintptr(i)*uintptr(bucketSize)))
			// 判断旧的bucket是否已经排空，没有排空的情况下
			if !evacuated(&oldb) {
				fmt.Printf("oldbuckets[%d]:\n", i)
				fmt.Printf("%s|__tophash : %+v\n", t, oldb.tophash)
				fmt.Printf("%s|__key     : %+v\n", t, oldb.keys)
				fmt.Printf("%s|__values  : %+v\n", t, oldb.values)
				fmt.Printf("%s|__overflow: %p\n", t, oldb.overflow)
				// 循环bmap链表直到overflow为nil为止
				for ; oldb.overflow != nil; oldb = *(*bmap)(unsafe.Pointer(oldb.overflow)) {
					t += "\t"
					fmt.Printf("%s|__tophash : %+v\n", t, oldb.overflow.tophash)
					fmt.Printf("%s|__key     : %+v\n", t, oldb.overflow.keys)
					fmt.Printf("%s|__values  : %+v\n", t, oldb.overflow.values)
					fmt.Printf("%s|__overflow: %p\n", t, oldb.overflow.overflow)
				}
			}
		}
	}
	// 循环buckets数组
	for i := 0; i < m; i++ {
		t := "\t"
		b := *(*bmap)(add(h.buckets, uintptr(i)*uintptr(bucketSize)))
		fmt.Printf("buckets[%d]:\n", i)
		fmt.Printf("%s|__tophash : %+v\n", t, b.tophash)
		fmt.Printf("%s|__key     : %+v\n", t, b.keys)
		fmt.Printf("%s|__values  : %+v\n", t, b.values)
		fmt.Printf("%s|__overflow: %p\n", t, b.overflow)
		// 循环bmap链表直到overflow为nil为止
		for ; b.overflow != nil; b = *(*bmap)(unsafe.Pointer(b.overflow)) {
			t += "\t"
			fmt.Printf("%s|__tophash : %+v\n", t, b.overflow.tophash)
			fmt.Printf("%s|__key     : %+v\n", t, b.overflow.keys)
			fmt.Printf("%s|__values  : %+v\n", t, b.overflow.values)
			fmt.Printf("%s|__overflow: %p\n", t, b.overflow.overflow)
		}
	}
}
