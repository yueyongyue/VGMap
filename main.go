package main

import (
	"fmt"
	"math/rand"
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
	bucketSize uintptr = 8
	maxKeySize uintptr = 128
	maxValSize uintptr = 128
	ptrSize            = 4 << (^uintptr(0) >> 63)
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

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	mp := make(map[int]int, 0)
	for i := 0; i < 50; i++ {
		mp[i] = i * 10
	}
	hMap := **(**hmap)(unsafe.Pointer(&mp))
	B := hMap.B
	fmt.Println("hmap.B:", B)
	fmt.Println("bmap array:", 1<<B) // 1 << B 相当于 2的B次幂
	fmt.Println("bmap oldbuckets:", hMap.oldbuckets)

	buckets := hMap.buckets
	fmt.Println("buckets指向的内存地址：", buckets)
	// 循环buckets数组
	for i := 0; i < 1<<B; i++ {
		t := "\t\t"
		b := *(*bmap)(add(buckets, uintptr(i)*uintptr(unsafe.Sizeof(bmap{}))))
		fmt.Printf("buckets[%d]:\n", i)
		fmt.Printf("%s|__tophash : %+v\n", t, b.tophash)
		fmt.Printf("%s|__key     : %+v\n", t, b.keys)
		fmt.Printf("%s|__values  : %+v\n", t, b.values)
		fmt.Printf("%s|__overflow: %p\n", t, b.overflow)
		// 循环bmap链表直到overflow为nil为止
		for ; b.overflow != nil; b = *(*bmap)(unsafe.Pointer(b.overflow)) {
			t += "\t\t"
			fmt.Printf("%s|__tophash : %+v\n", t, b.overflow.tophash)
			fmt.Printf("%s|__key     : %+v\n", t, b.overflow.keys)
			fmt.Printf("%s|__values  : %+v\n", t, b.overflow.values)
			fmt.Printf("%s|__overflow: %p\n", t, b.overflow.overflow)
		}

	}
}
