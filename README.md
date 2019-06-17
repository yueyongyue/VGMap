# 可视化输出 Golang Map （GoLang 1.11.1）
## 为了便于理解 Golang Map， 基于 Golang Map 底层的实现，做的 Map 可视化输出

````
#########################
    1：输入key和value
    2：随机生成一个Map
    3：打印Map
    4：清空Map
    5：退出
#########################
请选择(1,2,3,4,5):
2
请输入范围最大值:
50
buckets[0]:
	|__tophash : [174 4 110 67 4 50 98 103]
	|__key     : [2 3 7 9 13 25 28 45]
	|__values  : [49 12 47 9 52 72 68 92]
	|__overflow: 0x0
buckets[1]:
	|__tophash : [105 12 135 0 0 0 0 0]
	|__key     : [0 10 40 0 0 0 0 0]
	|__values  : [31 54 53 0 0 0 0 0]
	|__overflow: 0x0
buckets[2]:
	|__tophash : [179 156 44 45 56 112 14 66]
	|__key     : [1 15 17 19 22 24 27 30]
	|__values  : [38 39 62 25 50 71 65 71]
	|__overflow: 0xc0000a21b0
		|__tophash : [118 149 248 40 0 0 0 0]
		|__key     : [34 36 42 49 0 0 0 0]
		|__values  : [63 73 86 52 0 0 0 0]
		|__overflow: 0x0
buckets[3]:
	|__tophash : [160 234 228 172 0 0 0 0]
	|__key     : [5 11 32 44 0 0 0 0]
	|__values  : [23 22 69 77 0 0 0 0]
	|__overflow: 0x0
buckets[4]:
	|__tophash : [37 68 118 17 166 0 0 0]
	|__key     : [4 8 16 37 46 0 0 0]
	|__values  : [35 14 27 68 74 0 0 0]
	|__overflow: 0x0
buckets[5]:
	|__tophash : [250 80 232 191 170 53 0 0]
	|__key     : [6 21 23 29 38 48 0 0]
	|__values  : [31 37 31 44 73 57 0 0]
	|__overflow: 0x0
buckets[6]:
	|__tophash : [162 122 251 203 247 180 110 189]
	|__key     : [12 14 18 31 33 35 39 43]
	|__values  : [24 42 55 39 64 41 65 56]
	|__overflow: 0xc0000a2240
		|__tophash : [31 0 0 0 0 0 0 0]
		|__key     : [47 0 0 0 0 0 0 0]
		|__values  : [71 0 0 0 0 0 0 0]
		|__overflow: 0x0
buckets[7]:
	|__tophash : [247 218 147 0 0 0 0 0]
	|__key     : [20 26 41 0 0 0 0 0]
	|__values  : [65 63 81 0 0 0 0 0]
	|__overflow: 0x0


````