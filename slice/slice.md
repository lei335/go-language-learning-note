切片是对数组一个连续片段的引用（将该数组称为相关数组，通常是匿名的），所以切片是一个引用类型。多个切片如果表示同一个数组的片段，它们可以共享数据，即一个切片和相关数组的其他切片是共享存储的。

```go
s := make([]byte, 3) // 长度为3，容量默认也为3
```

make 初始化一个切片，并创建好相关数组。