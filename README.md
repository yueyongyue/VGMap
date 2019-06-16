# 可视化输出 Golang Map （GoLang 1.11.1）
## 基于 Golang Map 的实现，为了便于理解，做个 Map 的可视化输出

````
buckets[0]:
	|__tophash : [45 69 149 0 0 0 0 0]
	|__key     : [17 23 46 0 0 0 0 0]
	|__values  : [170 230 460 0 0 0 0 0]
	|__overflow: 0x0
buckets[1]:
	|__tophash : [164 156 135 179 41 72 0 0]
	|__key     : [10 13 20 24 43 45 0 0]
	|__values  : [100 130 200 240 430 450 0 0]
	|__overflow: 0x0
buckets[2]:
	|__tophash : [151 188 253 4 235 203 5 226]
	|__key     : [3 5 11 12 16 25 29 30]
	|__values  : [30 50 110 120 160 250 290 300]
	|__overflow: 0xc00008e090
		|__tophash : [211 0 0 0 0 0 0 0]
		|__key     : [35 0 0 0 0 0 0 0]
		|__values  : [350 0 0 0 0 0 0 0]
		|__overflow: 0x0
buckets[3]:
	|__tophash : [103 253 240 191 76 54 0 0]
	|__key     : [6 21 37 38 48 49 0 0]
	|__values  : [60 210 370 380 480 490 0 0]
	|__overflow: 0x0
buckets[4]:
	|__tophash : [134 67 148 0 0 0 0 0]
	|__key     : [7 19 42 0 0 0 0 0]
	|__values  : [70 190 420 0 0 0 0 0]
	|__overflow: 0x0
buckets[5]:
	|__tophash : [163 102 133 56 198 15 221 0]
	|__key     : [2 14 15 18 22 27 40 0]
	|__values  : [20 140 150 180 220 270 400 0]
	|__overflow: 0x0
buckets[6]:
	|__tophash : [75 85 145 248 242 0 0 0]
	|__key     : [0 1 26 28 32 0 0 0]
	|__values  : [0 10 260 280 320 0 0 0]
	|__overflow: 0x0
buckets[7]:
	|__tophash : [26 129 86 96 39 197 107 48]
	|__key     : [4 8 9 31 33 34 36 39]
	|__values  : [40 80 90 310 330 340 360 390]
	|__overflow: 0xc00008e120
		|__tophash : [174 243 141 0 0 0 0 0]
		|__key     : [41 44 47 0 0 0 0 0]
		|__values  : [410 440 470 0 0 0 0 0]
		|__overflow: 0x0
````