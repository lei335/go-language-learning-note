1. 切片是对数组一个连续片段的引用（将该数组称为相关数组，通常是匿名的），所以切片是一个引用类型。多个切片如果表示同一个数组的片段，它们可以共享数据，即一个切片和相关数组的其他切片是共享存储的。

```go
s := make([]byte, 3) // 长度为3，容量默认也为3
```

2. make 初始化一个切片，并创建好相关数组。

3. 根据slice的类型大小去获取能够申请的最大容量大小
Go slice对应代码如下：
```go
func makeslice(et *_type, len, cap int) slice {
    maxElements := maxSliceCap(et.size)
    if len<0 || uintptr(len) > maxElements {
        ...
    }
    if cap < len || uintptr(cap) > maxElements {
        ...
    }

    p := mallocgc(et.size*uintptr(cap), et, true)
    return slice{p, len, cap}
}
```
可以看到，`maxSliceCap()`函数可以获取能够申请的最大容量，从而进行安全检查。
```go
func maxSliceCap(elemsize uintptr) uintptr {
    if elemsize < uintptr(len(maxElems)) {
        return maxElems[elemsize]
    }
    return maxAlloc/elesize
}
```
```go
var maxElems = [...]uintptr{
    ^uintptr(0),
    maxAlloc / 1, maxAlloc / 2, maxAlloc / 3, maxAlloc / 4,
    maxAlloc / 5, maxAlloc / 6, maxAlloc / 7, maxAlloc / 8,
    maxAlloc / 9, maxAlloc / 10, maxAlloc / 11, maxAlloc / 12,
    maxAlloc / 13, maxAlloc / 14, maxAlloc / 15, maxAlloc / 16,
    maxAlloc / 17, maxAlloc / 18, maxAlloc / 19, maxAlloc / 20,
    maxAlloc / 21, maxAlloc / 22, maxAlloc / 23, maxAlloc / 24,
    maxAlloc / 25, maxAlloc / 26, maxAlloc / 27, maxAlloc / 28,
    maxAlloc / 29, maxAlloc / 30, maxAlloc / 31, maxAlloc / 32,
}
```
`maxElems`是包含一些预定义的切片最大容量值的查找表，索引是切片元素的类型大小。值主要是包含以下三个核心点：
* ^uintptr(0)
* maxAlloc
* maxAlloc/typeSize
**^uintptr(0)**
通过对代码的分析，可得出结论：
* 在32位系统下，uintptr为uint32类型，占用大小为4字节
* 在64位系统下，uintptr为uint64类型，占用大小为8个字节
'^'位运算符的作用是按位异或
^uint64(0)
二进制：0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
按位取反：1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111
该数为无符号整数，得到十进制值为：18446744073709551615，刚好也是打印出来的^uintptr(0)的值，也印证了其底层数据类型为uint64的事实（本机为64位），同时它也代表如下：
* math.MaxUint64
* 2的64次方减1
**maxAlloc**
`maxAlloc`是允许用户分配的最大虚拟内存空间。在64位，理论上可分配最大`1<<headAddrBits`字节；在32位，最大可分配小于`1<<32`字节
**maxAlloc/typeSize**
在想得到该类型的最大容量大小时，会根据对应的类型大小去查找索引表，查找不到的情况下才会去手动计算它的值，最终计算得到的内存大小都为该类型大小的整数倍。查找表的设置，更像是一个优化逻辑，减少常用的计算开销。